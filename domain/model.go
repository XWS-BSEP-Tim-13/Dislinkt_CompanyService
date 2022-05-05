package domain

import (
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	Id          primitive.ObjectID `bson:"_id"`
	CompanyName string             `bson:"company_name"`
	Username    string             `bson:"username"`
	Description string             `bson:"description"`
	Location    string             `bson:"location"`
	Website     string             `bson:"website"`
	CompanySize string             `bson:"company_size"`
	Industry    string             `bson:"industry"`
	JobOffers   []JobOffer         `bson:"job_offers"`
}

type JobOffer struct {
	Id             primitive.ObjectID  `bson:"_id"`
	Position       string              `bson:"position"`
	JobDescription string              `bson:"job_description"`
	Prerequisites  string              `bson:"prerequisites"`
	Company        Company             `bson:"company"`
	EmploymentType enum.EmploymentType `bson:"employment_type"`
}
