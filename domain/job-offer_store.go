package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobOfferStore interface {
	GetActiveById(ctx context.Context, id primitive.ObjectID) (*JobOffer, error)
	GetAllActive(ctx context.Context) ([]*JobOffer, error)
	Insert(ctx context.Context, jobOffer *JobOffer) error
	DeleteAll(ctx context.Context)
	FilterJobs(ctx context.Context, filter *JobFilter) ([]*JobOffer, error)
}
