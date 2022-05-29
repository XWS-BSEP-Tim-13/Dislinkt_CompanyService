package api

import (
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain/enum"
	pb "github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/infrastructure/grpc/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapCompanyDomainToPb(company *domain.Company) *pb.Company {
	companyPb := &pb.Company{
		Id:          company.Id.Hex(),
		CompanyName: company.CompanyName,
		Username:    company.Username,
		Description: company.Description,
		Location:    company.Location,
		Website:     company.Website,
		CompanySize: company.CompanySize,
		Industry:    company.Industry,
	}
	return companyPb
}

func mapJobOfferDtoToDomain(jobOfferDto *pb.JobOfferDto) *domain.JobOffer {
	jobOffer := &domain.JobOffer{
		Id: primitive.NewObjectID(),
		Company: domain.Company{
			CompanyName: jobOfferDto.Company.CompanyName,
			Username:    jobOfferDto.Company.Username,
			Email:       jobOfferDto.Company.Email,
			PhoneNumber: jobOfferDto.Company.PhoneNumber,
			Description: jobOfferDto.Company.Description,
			Location:    jobOfferDto.Company.Location,
			Website:     jobOfferDto.Company.Website,
			CompanySize: jobOfferDto.Company.CompanySize,
			Industry:    jobOfferDto.Company.Industry,
		},
		JobDescription: jobOfferDto.JobDescription,
		Position:       jobOfferDto.Position,
		Prerequisites:  jobOfferDto.Prerequisites,
		EmploymentType: enum.EmploymentType(jobOfferDto.EmploymentType),
	}
	return jobOffer
}

func mapCompanyPbToDomain(companyPb *pb.Company) *domain.Company {
	company := &domain.Company{
		CompanyName: companyPb.CompanyName,
		Username:    companyPb.Username,
		Email:       companyPb.Email,
		PhoneNumber: companyPb.PhoneNumber,
		Description: companyPb.Description,
		Location:    companyPb.Location,
		Website:     companyPb.Website,
		CompanySize: companyPb.CompanySize,
		Industry:    companyPb.Industry,
	}
	return company
}
