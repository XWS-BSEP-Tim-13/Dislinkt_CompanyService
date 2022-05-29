package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type JobOfferStore interface {
	GetActiveById(id primitive.ObjectID) (*JobOffer, error)
	GetAllActive() ([]*JobOffer, error)
	Insert(jobOffer *JobOffer) error
	DeleteAll()
}
