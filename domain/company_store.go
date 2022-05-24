package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyStore interface {
	GetActiveById(id primitive.ObjectID) (*Company, error)
	GetAllActive() ([]*Company, error)
	GetActiveByUsername(username string) (*Company, error)
	GetByUsername(username string) (*Company, error)
	GetByEmail(email string) (*Company, error)
	Insert(company *Company) error
	DeleteAll()
}
