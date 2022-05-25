package persistence

import (
	"context"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "companies"
	COLLECTION = "company"
)

type CompanyMongoDBStore struct {
	companies *mongo.Collection
}

func NewCompanyMongoDBStore(client *mongo.Client) domain.CompanyStore {
	companies := client.Database(DATABASE).Collection(COLLECTION)
	return &CompanyMongoDBStore{
		companies: companies,
	}
}

func (store *CompanyMongoDBStore) GetActiveById(id primitive.ObjectID) (*domain.Company, error) {
	filter := bson.M{"_id": id, "is_active": true}
	return store.filterOne(filter)
}

func (store *CompanyMongoDBStore) GetAllActive() ([]*domain.Company, error) {
	filter := bson.D{{"is_active", "true"}}
	return store.filter(filter)
}

func (store *CompanyMongoDBStore) GetActiveByUsername(username string) (*domain.Company, error) {
	filter := bson.M{"username": username, "is_active": true}
	return store.filterOne(filter)
}

func (store *CompanyMongoDBStore) GetByUsername(username string) (*domain.Company, error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *CompanyMongoDBStore) GetByEmail(email string) (*domain.Company, error) {
	filter := bson.M{"email": email}
	return store.filterOne(filter)
}

func (store *CompanyMongoDBStore) Insert(company *domain.Company) error {
	result, err := store.companies.InsertOne(context.TODO(), company)
	if err != nil {
		return err
	}
	company.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *CompanyMongoDBStore) UpdateIsActive(email string) error {
	_, err := store.companies.UpdateOne(
		context.TODO(),
		bson.M{"email": email},
		bson.D{{"$set", bson.D{{"is_active", true}}}},
	)
	return err
}

func (store *CompanyMongoDBStore) DeleteAll() {
	store.companies.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *CompanyMongoDBStore) filter(filter interface{}) ([]*domain.Company, error) {
	cursor, err := store.companies.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *CompanyMongoDBStore) filterOne(filter interface{}) (company *domain.Company, err error) {
	result := store.companies.FindOne(context.TODO(), filter)
	err = result.Decode(&company)
	return
}

func decode(cursor *mongo.Cursor) (companies []*domain.Company, err error) {
	for cursor.Next(context.TODO()) {
		var company domain.Company
		err = cursor.Decode(&company)
		if err != nil {
			return
		}
		companies = append(companies, &company)
	}
	err = cursor.Err()
	return
}
