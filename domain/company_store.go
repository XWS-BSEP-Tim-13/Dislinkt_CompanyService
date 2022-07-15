package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyStore interface {
	GetActiveById(ctx context.Context, id primitive.ObjectID) (*Company, error)
	GetAllActive(ctx context.Context) ([]*Company, error)
	GetActiveByUsername(ctx context.Context, username string) (*Company, error)
	GetByUsername(ctx context.Context, username string) (*Company, error)
	GetByEmail(ctx context.Context, email string) (*Company, error)
	Insert(ctx context.Context, company *Company) error
	DeleteAll(ctx context.Context)
	UpdateIsActive(ctx context.Context, email string) error
}
