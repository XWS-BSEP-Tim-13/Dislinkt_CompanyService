package startup

import (
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var companies = []*domain.Company{
	{
		Id:          getObjectId("623b0cc3a34d25d8567f9f82"),
		CompanyName: "Levi9",
		Username:    "levi9",
		Location:    "ns",
		Description: "Technology services",
		Website:     "www.levi9.com",
		CompanySize: "1000",
		Industry:    "IT",
		IsActive:    true,
	},
	{
		Id:          getObjectId("623b0cc3a34d25d8567f9f83"),
		CompanyName: "VegaIT",
		Username:    "VegaIT",
		Location:    "ns",
		Description: "Technology services",
		Website:     "www.vegait.com",
		CompanySize: "1000",
		Industry:    "IT",
		IsActive:    true,
	},
}

var jobs = []*domain.JobOffer{
	{
		Id:             getObjectId("623b0cc3a34d25d8567f9f82"),
		EmploymentType: 0,
		Position:       "Softver developer",
		Prerequisites:  "2 years of expirience.",
		Company:        *companies[0],
		JobDescription: "Great expirience for self development and work with experts",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
