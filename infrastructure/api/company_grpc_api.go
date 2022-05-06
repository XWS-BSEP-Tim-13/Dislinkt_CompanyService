package api

import (
	"context"
	"fmt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/application"
	pb "github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/infrastructure/grpc/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/status"
)

type CompanyHandler struct {
	pb.UnimplementedCompanyServiceServer
	service *application.CompanyService
}

func NewCompanyHandler(service *application.CompanyService) *CompanyHandler {
	return &CompanyHandler{
		service: service,
	}
}

func (handler *CompanyHandler) CreateCompany(ctx context.Context, request *pb.NewCompany) (*pb.NewCompany, error) {
	fmt.Println((*request).Company)
	company := mapCompanyPbToDomain(request.Company)
	fmt.Println(company)

	newCompany, err := handler.service.CreateNewCompany(company)
	if err != nil {
		return nil, status.Error(400, err.Error())
	}

	response := &pb.NewCompany{
		Company: mapCompanyDomainToPb(newCompany),
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
