package persistence

import (
	"context"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/tracer"
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

func (store *CompanyMongoDBStore) GetActiveById(ctx context.Context, id primitive.ObjectID) (*domain.Company, error) {
	span := tracer.StartSpanFromContext(ctx, "DB GetActiveById")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"_id": id}
	return store.filterOne(ctx, filter)
}

func (store *CompanyMongoDBStore) GetAllActive(ctx context.Context) ([]*domain.Company, error) {
	span := tracer.StartSpanFromContext(ctx, "DB GetAllActive")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.D{{}}
	return store.filter(ctx, filter)
}

func (store *CompanyMongoDBStore) GetActiveByUsername(ctx context.Context, username string) (*domain.Company, error) {
	span := tracer.StartSpanFromContext(ctx, "DB GetActiveByUsername")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"username": username}
	return store.filterOne(ctx, filter)
}

func (store *CompanyMongoDBStore) GetByUsername(ctx context.Context, username string) (*domain.Company, error) {
	span := tracer.StartSpanFromContext(ctx, "DB GetByUsername")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"username": username}
	return store.filterOne(ctx, filter)
}

func (store *CompanyMongoDBStore) GetByEmail(ctx context.Context, email string) (*domain.Company, error) {
	span := tracer.StartSpanFromContext(ctx, "DB GetByEmail")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"email": email}
	return store.filterOne(ctx, filter)
}

func (store *CompanyMongoDBStore) Insert(ctx context.Context, company *domain.Company) error {
	span := tracer.StartSpanFromContext(ctx, "DB Insert")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	result, err := store.companies.InsertOne(context.TODO(), company)
	if err != nil {
		return err
	}
	company.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *CompanyMongoDBStore) UpdateIsActive(ctx context.Context, email string) error {
	span := tracer.StartSpanFromContext(ctx, "DB UpdateIsActive")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := store.companies.UpdateOne(
		context.TODO(),
		bson.M{"email": email},
		bson.D{{"$set", bson.D{{"is_active", true}}}},
	)
	return err
}

func (store *CompanyMongoDBStore) DeleteAll(ctx context.Context) {
	span := tracer.StartSpanFromContext(ctx, "DB DeleteAll")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	store.companies.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *CompanyMongoDBStore) filter(ctx context.Context, filter interface{}) ([]*domain.Company, error) {
	span := tracer.StartSpanFromContext(ctx, "DB filter")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	cursor, err := store.companies.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *CompanyMongoDBStore) filterOne(ctx context.Context, filter interface{}) (company *domain.Company, err error) {
	span := tracer.StartSpanFromContext(ctx, "DB filterOne")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

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
