package application

import (
	"errors"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyService struct {
	store    domain.CompanyStore
	jobStore domain.JobOfferStore
}

func NewCompanyService(store domain.CompanyStore, jobStore domain.JobOfferStore) *CompanyService {
	return &CompanyService{
		store:    store,
		jobStore: jobStore,
	}
}

func (service *CompanyService) Get(id primitive.ObjectID) (*domain.Company, error) {
	return service.store.GetActiveById(id)
}

func (service *CompanyService) GetAll() ([]*domain.Company, error) {
	return service.store.GetAllActive()
}

func (service *CompanyService) CreateNewCompany(company *domain.Company) (*domain.Company, error) {
	dbCompany, _ := service.store.GetByUsername((*company).Username)
	if dbCompany != nil {
		err := errors.New("username already exists")
		return nil, err
	}

	dbCompany, _ = service.store.GetByUsername((*company).Username)
	if dbCompany != nil {
		err := errors.New("email already exists")
		return nil, err
	}

	(*company).Id = primitive.NewObjectID()
	(*company).IsActive = false
	err := service.store.Insert(company)
	if err != nil {
		err := errors.New("error while creating new company")
		return nil, err
	}

	return company, nil
}

func (service *CompanyService) InsertJobOffer(job *domain.JobOffer) error {
	return service.jobStore.Insert(job)
}

func (service *CompanyService) GetAllJobs() ([]*domain.JobOffer, error) {
	return service.jobStore.GetAllActive()
}

func (service *CompanyService) ActivateAccount(email string) (string, error) {
	err := service.store.UpdateIsActive(email)
	if err != nil {
		err := errors.New("error activating account")
		return "", err
	}

	return "Account successfully activated!", nil
}
