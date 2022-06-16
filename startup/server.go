package startup

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/application"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/infrastructure/api"
	company "github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/infrastructure/grpc/proto"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/infrastructure/persistence"
	logger "github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/logging"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/startup/config"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/util"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
)

type Server struct {
	config *config.Config
}

const (
	serverCertFile = "cert/cert.pem"
	serverKeyFile  = "cert/key.pem"
	clientCertFile = "cert/client-cert.pem"
)

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	logger := logger.InitLogger("company-service", context.TODO())

	mongoClient := server.initMongoClient()
	companyStore := server.initCompanyStore(mongoClient)
	jobsStore := server.initJobsStore(mongoClient)
	companyService := server.initCompanyService(companyStore, jobsStore, logger)
	goValidator := server.initGoValidator()
	companyHandler := server.initCompanyHandler(companyService, goValidator, logger)

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

func (server *Server) initCompanyService(store domain.CompanyStore, jobStore domain.JobOfferStore, logger *logger.Logger) *application.CompanyService {
	return application.NewCompanyService(store, jobStore, logger)
}

func (server *Server) initGoValidator() *util.GoValidator {
	return util.NewGoValidator()
}

func (server *Server) initCompanyHandler(service *application.CompanyService, goValidator *util.GoValidator, logger *logger.Logger) *api.CompanyHandler {
	return api.NewCompanyHandler(service, goValidator, logger)
}

func (server *Server) startGrpcServer(productHandler *api.CompanyHandler) {
	cert, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	pemClientCA, err := ioutil.ReadFile(clientCertFile)
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		log.Fatal(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.NoClientCert,
		//ClientCAs:    certPool,
	}

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewTLS(config)),
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(opts...)
	company.RegisterCompanyServiceServer(grpcServer, productHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
