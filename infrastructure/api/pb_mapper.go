package api

import (
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain/enum"
	pb "github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/infrastructure/grpc/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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
		Published:      time.Now(),
	}
	return jobOffer
}

func mapPbFilterToDomain(filter *pb.FilterBody) *domain.JobFilter {
	filterDomain := &domain.JobFilter{
		Company:        filter.CompanyId,
		EmploymentType: enum.EmploymentType(filter.Type),
		SortDate:       enum.DateSort(filter.Date),
		Position:       filter.Position,
	}
	return filterDomain
}

func mapJobDomainToPb(job *domain.JobOffer) *pb.JobOffer {
	jobPb := &pb.JobOffer{
		Id:             job.Id.Hex(),
		JobDescription: job.JobDescription,
		Position:       job.Position,
		Prerequisites:  job.Prerequisites,
		EmploymentType: pb.EmploymentType(job.EmploymentType),
		Company: &pb.Company{
			Id:          job.Company.Id.Hex(),
			CompanyName: job.Company.CompanyName,
			Username:    job.Company.Username,
			Description: job.Company.Description,
			Location:    job.Company.Location,
			Website:     job.Company.Website,
			CompanySize: job.Company.CompanySize,
			Industry:    job.Company.Industry,
		},
	}
	return jobPb
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
