package api

import (
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain"
	pb "github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/infrastructure/grpc/proto"
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
