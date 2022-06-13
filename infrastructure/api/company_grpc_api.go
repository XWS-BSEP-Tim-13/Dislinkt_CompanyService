package api

import (
	"context"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/application"
	pb "github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/infrastructure/grpc/proto"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/jwt"
	logger "github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/logging"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/status"
)

type CompanyHandler struct {
	pb.UnimplementedCompanyServiceServer
	service     *application.CompanyService
	goValidator *util.GoValidator
	logger      *logger.Logger
}

func NewCompanyHandler(service *application.CompanyService, goValidator *util.GoValidator, logger *logger.Logger) *CompanyHandler {
	return &CompanyHandler{
		service:     service,
		goValidator: goValidator,
		logger:      logger,
	}
}

func (handler *CompanyHandler) CreateCompany(ctx context.Context, request *pb.NewCompany) (*pb.NewCompany, error) {
	company := mapCompanyPbToDomain(request.Company)

	err := handler.goValidator.Validator.Struct(company)
	if err != nil {
		handler.logger.WarningMessage("Action: Company registration with invalid data")
		return nil, status.Error(500, err.Error())
	}

	newCompany, err := handler.service.CreateNewCompany(company)
	if err != nil {
		handler.logger.ErrorMessage("Action: Company registration")
		return nil, status.Error(400, err.Error())
	}

	response := &pb.NewCompany{
		Company: mapCompanyDomainToPb(newCompany),
	}

	handler.logger.InfoMessage("Action: New company registered - " + company.CompanyName)
	return response, nil
}

func (handler *CompanyHandler) CreateJobOffer(ctx context.Context, request *pb.JobOfferRequest) (*pb.JobOfferResponse, error) {
	job := mapJobOfferDtoToDomain(request.Dto)
	err := handler.service.InsertJobOffer(job)
	if err != nil {
		handler.logger.ErrorMessage("Company: " + job.Company.Username + " | Action: Create job offer")
		return nil, status.Error(400, err.Error())
	}
	response := &pb.JobOfferResponse{
		Id: primitive.ObjectID.String(job.Id),
	}

	handler.logger.InfoMessage("Company: " + job.Company.Username + " | Action: Create job offer")
	return response, nil
}

func (handler *CompanyHandler) FilterJobOffers(ctx context.Context, request *pb.FilterJobsRequest) (*pb.GetAllJobsResponse, error) {

	filter := mapPbFilterToDomain(request.Filter)
	jobs, err := handler.service.FilterJobs(filter)
	if err != nil {
		return nil, status.Error(400, err.Error())
	}
	response := &pb.GetAllJobsResponse{
		Jobs: []*pb.JobOffer{},
	}

	for _, job := range jobs {
		current := mapJobDomainToPb(job)
		response.Jobs = append(response.Jobs, current)
	}
	return response, nil
}

func (handler *CompanyHandler) GetJobOffers(ctx context.Context, request *pb.EmptyMessage) (*pb.GetAllJobsResponse, error) {
	username, _ := jwt.ExtractUsernameFromToken(ctx)
	loggedUser := username
	if username == "" {
		loggedUser = "Anonymous"
	}
	resp, err := handler.service.GetAllJobs()
	if err != nil {
		handler.logger.ErrorMessage("User: " + loggedUser + " | Action: Get job offers")
		return nil, status.Error(500, err.Error())
	}
	response := &pb.GetAllJobsResponse{
		Jobs: []*pb.JobOffer{},
	}

	for _, job := range resp {
		current := mapJobDomainToPb(job)
		response.Jobs = append(response.Jobs, current)
	}

	handler.logger.InfoMessage("User: " + loggedUser + " | Action: Get job offers")
	return response, nil
}

func (handler *CompanyHandler) ActivateAccount(ctx context.Context, request *pb.ActivateAccountRequest) (*pb.ActivateAccountResponse, error) {
	email := request.Email

	resp, err := handler.service.ActivateAccount(email)
	if err != nil {
		handler.logger.ErrorMessage("Company: " + email + " | Action: Activate account")
		return nil, status.Error(500, err.Error())
	}

	response := &pb.ActivateAccountResponse{
		Message: resp,
	}

	handler.logger.InfoMessage("Company: " + email + " | Action: Activate account")
	return response, nil
}

func (handler *CompanyHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	username, _ := jwt.ExtractUsernameFromToken(ctx)
	loggedUser := username
	if username == "" {
		loggedUser = "Anonymous"
	}
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	company, err := handler.service.Get(objectId)
	if err != nil {
		handler.logger.ErrorMessage("User: " + loggedUser + " | Action: Get company by id")
		return nil, err
	}
	companyPb := mapCompanyDomainToPb(company)
	response := &pb.GetResponse{
		Company: companyPb,
	}

	handler.logger.InfoMessage("User: " + loggedUser + " | Action: Get company by id")
	return response, nil
}

func (handler *CompanyHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	username, _ := jwt.ExtractUsernameFromToken(ctx)
	loggedUser := username
	if username == "" {
		loggedUser = "Anonymous"
	}

	companies, err := handler.service.GetAll()
	if err != nil {
		handler.logger.ErrorMessage("User: " + loggedUser + " | Action: Get companies")
		return nil, err
	}
	response := &pb.GetAllResponse{
		Companies: []*pb.Company{},
	}
	for _, company := range companies {
		current := mapCompanyDomainToPb(company)
		response.Companies = append(response.Companies, current)
	}
	handler.logger.InfoMessage("User: " + loggedUser + " | Action: Get companies")
	return response, nil
}
