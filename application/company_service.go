package application

import (
	"context"
	"errors"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain"
	logger "github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/logging"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/tracer"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyService struct {
	store    domain.CompanyStore
	jobStore domain.JobOfferStore
	logger   *logger.Logger
}

func NewCompanyService(store domain.CompanyStore, jobStore domain.JobOfferStore, logger *logger.Logger) *CompanyService {
	return &CompanyService{
		store:    store,
		jobStore: jobStore,
		logger:   logger,
	}
}

func (service *CompanyService) Get(ctx context.Context, id primitive.ObjectID) (*domain.Company, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE Get")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.GetActiveById(ctx, id)
}

func (service *CompanyService) GetAll(ctx context.Context) ([]*domain.Company, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetAll")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.GetAllActive(ctx)
}

func (service *CompanyService) CreateNewCompany(ctx context.Context, company *domain.Company) (*domain.Company, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE CreateNewCompany")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	dbCompany, _ := service.store.GetByUsername(ctx, (*company).Username)
	if dbCompany != nil {
		err := errors.New("username already exists")
		return nil, err
	}

	dbCompany, _ = service.store.GetByUsername(ctx, (*company).Username)
	if dbCompany != nil {
		err := errors.New("email already exists")
		return nil, err
	}

	(*company).Id = primitive.NewObjectID()
	(*company).IsActive = false
	err := service.store.Insert(ctx, company)
	if err != nil {
		err := errors.New("error while creating new company")
		return nil, err
	}

	return company, nil
}

func (service *CompanyService) InsertJobOffer(ctx context.Context, job *domain.JobOffer) error {
	span := tracer.StartSpanFromContext(ctx, "SERVICE InsertJobOffer")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.jobStore.Insert(ctx, job)
}

func (service *CompanyService) GetAllJobs(ctx context.Context) ([]*domain.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetAllJobs")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.jobStore.GetAllActive(ctx)
}

func (service *CompanyService) ActivateAccount(ctx context.Context, email string) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE ActivateAccount")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.store.UpdateIsActive(ctx, email)
	if err != nil {
		err := errors.New("error activating account")
		return "", err
	}

	return "Account successfully activated!", nil
}

func (service *CompanyService) FilterJobs(ctx context.Context, filter *domain.JobFilter) ([]*domain.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE FilterJobs")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.jobStore.FilterJobs(ctx, filter)
}
