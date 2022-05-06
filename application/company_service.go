package application

import (
	"errors"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyService struct {
	store domain.CompanyStore
}

func NewCompanyService(store domain.CompanyStore) *CompanyService {
	return &CompanyService{
		store: store,
	}
}

func (service *CompanyService) Get(id primitive.ObjectID) (*domain.Company, error) {
	return service.store.Get(id)
}

func (service *CompanyService) GetAll() ([]*domain.Company, error) {
	return service.store.GetAll()
}

func (service *CompanyService) CreateNewCompany(company *domain.Company) (*domain.Company, error) {
	dbCompany, _ := service.store.GetByUsername((*company).Username)
	if dbCompany != nil {
		err := errors.New("username already exists")
		return nil, err
	}
	(*company).Id = primitive.NewObjectID()
	err := service.store.Insert(company)
	if err != nil {
		err := errors.New("error while creating new company")
		return nil, err
	}

	return company, nil
}
