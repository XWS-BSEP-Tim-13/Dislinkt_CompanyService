package startup

import (
	"fmt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/application"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/infrastructure/api"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/infrastructure/persistence"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	companyStore := server.initCompanyStore(mongoClient)

	companyService := server.initCompanyService(companyStore)

	companyHandler := server.initCompanyHandler(companyService)

	server.startGrpcServer(companyHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.CompanyDBHost, server.config.CompanyDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initCompanyStore(client *mongo.Client) domain.CompanyStore {
	store := persistence.NewCompanyMongoDBStore(client)
	store.DeleteAll()
	for _, product := range products {
		err := store.Insert(product)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initCompanyService(store domain.CompanyStore) *application.CompanyService {
	return application.NewCompanyService(store)
}

func (server *Server) initCompanyHandler(service *application.CompanyService) *api.CompanyHandler {
	return api.NewCompanyHandler(service)
}

func (server *Server) startGrpcServer(productHandler *api.CompanyHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	//post.RegisterCompanyServiceServer(grpcServer, productHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
