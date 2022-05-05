package api

import (
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain"
	pb "github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/infrastructure/grpc/proto"
)

func mapCompany(company *domain.Company) *pb.Company {
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
	for _, jobOffer := range company.JobOffers {
		companyPb.JobOffers = append(companyPb.JobOffers, &pb.JobOffer{
			Id:             jobOffer.Id.Hex(),
			Position:       jobOffer.Position,
			JobDescription: jobOffer.JobDescription,
			Prerequisites:  jobOffer.Prerequisites,
		})
	}
	return companyPb
}
