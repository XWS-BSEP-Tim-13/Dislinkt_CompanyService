package api

import (
	"context"
	"fmt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/application"
	pb "github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/infrastructure/grpc/proto"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/status"
)

type CompanyHandler struct {
	pb.UnimplementedCompanyServiceServer
	service     *application.CompanyService
	goValidator *util.GoValidator
}

func NewCompanyHandler(service *application.CompanyService, goValidator *util.GoValidator) *CompanyHandler {
	return &CompanyHandler{
		service:     service,
		goValidator: goValidator,
	}
}

func (handler *CompanyHandler) CreateCompany(ctx context.Context, request *pb.NewCompany) (*pb.NewCompany, error) {
	company := mapCompanyPbToDomain(request.Company)
	fmt.Println(company)

	err := handler.goValidator.Validator.Struct(company)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newCompany, err := handler.service.CreateNewCompany(company)
	if err != nil {
		return nil, status.Error(400, err.Error())
	}

	response := &pb.NewCompany{
		Company: mapCompanyDomainToPb(newCompany),
	}

	return response, nil
}

func (handler *CompanyHandler) ActivateAccount(ctx context.Context, request *pb.ActivateAccountRequest) (*pb.ActivateAccountResponse, error) {
	email := request.Email

	resp, err := handler.service.ActivateAccount(email)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	response := &pb.ActivateAccountResponse{
		Message: resp,
	}

	return response, nil
}

func (handler *CompanyHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	company, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	companyPb := mapCompanyDomainToPb(company)
	response := &pb.GetResponse{
		Company: companyPb,
	}
	return response, nil
}

func (handler *CompanyHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	companies, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Companies: []*pb.Company{},
	}
	for _, company := range companies {
		current := mapCompanyDomainToPb(company)
		response.Companies = append(response.Companies, current)
	}
	return response, nil
}
