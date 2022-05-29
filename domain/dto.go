package domain

import (
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain/enum"
)

type JobFilter struct {
	Position       string              `bson:"position"`
	Company        string              `bson:"company"`
	EmploymentType enum.EmploymentType `bson:"employment_type"`
	SortDate       enum.DateSort       `bson:"sort-date"`
}
