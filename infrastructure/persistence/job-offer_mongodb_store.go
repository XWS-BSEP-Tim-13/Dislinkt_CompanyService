package persistence

import (
	"context"
	"fmt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/tracer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DatabaseJobs   = "jobs"
	CollectionJobs = "job"
)

type JobOfferMongoDBStore struct {
	jobs *mongo.Collection
}

func (store JobOfferMongoDBStore) FilterJobs(ctx context.Context, filter *domain.JobFilter) ([]*domain.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "DB FilterJobs")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	empType := bson.M{}
	position := bson.M{"position": primitive.Regex{Pattern: filter.Position, Options: "i"}}
	company := bson.M{}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"likes", -1}, {"date", -1}})
	if filter.SortDate == 0 {
		findOptions.SetSort(bson.D{{"published", 1}})
	}
	if filter.SortDate == 1 {
		findOptions.SetSort(bson.D{{"published", -1}})
	}
	fmt.Printf("Employment type : %s\n", filter.Position)
	if filter.EmploymentType != 3 {
		empType = bson.M{"employment_type": filter.EmploymentType}
	}
	compId, _ := primitive.ObjectIDFromHex(filter.Company)
	if filter.Company != "-1" {
		fmt.Printf("Company id: %s\n", compId)
		company = bson.M{"company._id": compId}
	}

	filterr := bson.M{
		"$and": []bson.M{
			empType,
			position,
			company,
		},
	}
	cursor, err := store.jobs.Find(context.TODO(), filterr, findOptions)
	if err != nil {
		return nil, err
	}
	jobs, _ := decodeJobs(ctx, cursor)
	fmt.Printf("Broj poslova: %d\n", len(jobs))
	return jobs, nil
}

func NewJobOfferMongoDBStore(client *mongo.Client) domain.JobOfferStore {
	jobs := client.Database(DatabaseJobs).Collection(CollectionJobs)
	return &JobOfferMongoDBStore{
		jobs: jobs,
	}
}

func (store JobOfferMongoDBStore) GetActiveById(ctx context.Context, id primitive.ObjectID) (*domain.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "DB GetActiveById")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"_id": id}
	return store.filterOne(ctx, filter)
}

func (store JobOfferMongoDBStore) GetAllActive(ctx context.Context) ([]*domain.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "DB GetAllActive")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.D{{}}
	return store.filter(ctx, filter)
}

func (store JobOfferMongoDBStore) Insert(ctx context.Context, jobOffer *domain.JobOffer) error {
	span := tracer.StartSpanFromContext(ctx, "DB Insert")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	jobOffer.Id = primitive.NewObjectID()
	result, err := store.jobs.InsertOne(context.TODO(), jobOffer)
	if err != nil {
		return err
	}
	jobOffer.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *JobOfferMongoDBStore) filterOne(ctx context.Context, filter interface{}) (job *domain.JobOffer, err error) {
	span := tracer.StartSpanFromContext(ctx, "DB Insert")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := store.jobs.FindOne(context.TODO(), filter)
	err = result.Decode(&job)
	return
}

func (store *JobOfferMongoDBStore) filter(ctx context.Context, filter interface{}) ([]*domain.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "DB filter")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	cursor, err := store.jobs.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeJobs(ctx, cursor)
}

func (store JobOfferMongoDBStore) DeleteAll(ctx context.Context) {
	span := tracer.StartSpanFromContext(ctx, "DB DeleteAll")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	store.jobs.DeleteMany(context.TODO(), bson.D{{}})
}

func decodeJobs(ctx context.Context, cursor *mongo.Cursor) (jobs []*domain.JobOffer, err error) {
	span := tracer.StartSpanFromContext(ctx, "DB decodeJobs")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	for cursor.Next(context.TODO()) {
		var job domain.JobOffer
		err = cursor.Decode(&job)
		if err != nil {
			return
		}
		jobs = append(jobs, &job)
	}
	err = cursor.Err()
	return
}
