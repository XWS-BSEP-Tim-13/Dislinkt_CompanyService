package application

import (
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
