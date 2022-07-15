package domain

import (
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Company struct {
	Id          primitive.ObjectID `bson:"_id"`
	CompanyName string             `bson:"company_name" validate:"required,companyName"`
	Username    string             `bson:"username" validate:"required,username"`
	Email       string             `bson:"email" validate:"required,email"`
	PhoneNumber string             `bson:"phone_number" validate:"required,numeric,min=9,max=10"`
	Description string             `bson:"description"`
	Location    string             `bson:"location" validate:"required,max=256"`
	Website     string             `bson:"website" validate:"required,website"`
	CompanySize string             `bson:"company_size" validate:"required,companyName"`
	Industry    string             `bson:"industry" validate:"required,max=256"`
	IsActive    bool               `bson:"is_active"`
}

type JobOffer struct {
	Id             primitive.ObjectID  `bson:"_id"`
	Position       string              `bson:"position"`
	JobDescription string              `bson:"job_description"`
	Prerequisites  string              `bson:"prerequisites"`
	Company        Company             `bson:"company"`
	EmploymentType enum.EmploymentType `bson:"employment_type"`
	Published      time.Time           `bson:"published" validate:"required"`
}
