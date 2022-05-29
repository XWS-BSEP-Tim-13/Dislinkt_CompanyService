package startup

import (
	"fmt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/application"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/infrastructure/api"
	company "github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/infrastructure/grpc/proto"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/infrastructure/persistence"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/startup/config"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/util"
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
	jobsStore := server.initJobsStore(mongoClient)
	companyService := server.initCompanyService(companyStore, jobsStore)
	goValidator := server.initGoValidator()

	companyHandler := server.initCompanyHandler(companyService, goValidator)

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
	for _, company := range companies {
		err := store.Insert(company)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initJobsStore(client *mongo.Client) domain.JobOfferStore {
	store := persistence.NewJobOfferMongoDBStore(client)
	store.DeleteAll()
	for _, job := range jobs {
		err := store.Insert(job)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initCompanyService(store domain.CompanyStore, jobStore domain.JobOfferStore) *application.CompanyService {
	return application.NewCompanyService(store, jobStore)
}

func (server *Server) initGoValidator() *util.GoValidator {
	return util.NewGoValidator()
}

func (server *Server) initCompanyHandler(service *application.CompanyService, goValidator *util.GoValidator) *api.CompanyHandler {
	return api.NewCompanyHandler(service, goValidator)
}

func (server *Server) startGrpcServer(productHandler *api.CompanyHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	company.RegisterCompanyServiceServer(grpcServer, productHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
