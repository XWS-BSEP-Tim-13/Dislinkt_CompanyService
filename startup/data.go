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
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
