package persistence

import (
	"context"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_CompanyService/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
)

const (
	DatabaseJobs   = "jobs"
	CollectionJobs = "job"
)

type JobOfferMongoDBStore struct {
	jobs *mongo.Collection
}

func (store JobOfferMongoDBStore) FilterJobs(filter *domain.JobFilter) ([]*domain.JobOffer, error) {
	empType := bson.M{}
	position := bson.M{"position": primitive.Regex{Pattern: filter.Position, Options: "i"}}
	company := bson.M{}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"likes", -1}, {"date", -1}})
	if filter.SortDate == 0 {
		findOptions.SetSort(bson.D{{"published", -1}})
	}
	if filter.SortDate == 1 {
		findOptions.SetSort(bson.D{{"published", 1}})
	}
	if filter.EmploymentType != 3 {
		empType = bson.M{"employment_type": filter.EmploymentType}
	}
	compId, _ := strconv.Atoi(filter.Company)
	if compId != -1 {
		company = bson.M{"company._id": filter.Company}
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
	jobs, _ := decodeJobs(cursor)
	return jobs, nil
}

func NewJobOfferMongoDBStore(client *mongo.Client) domain.JobOfferStore {
	jobs := client.Database(DatabaseJobs).Collection(CollectionJobs)
	return &JobOfferMongoDBStore{
		jobs: jobs,
	}
}

func (store JobOfferMongoDBStore) GetActiveById(id primitive.ObjectID) (*domain.JobOffer, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store JobOfferMongoDBStore) GetAllActive() ([]*domain.JobOffer, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store JobOfferMongoDBStore) Insert(jobOffer *domain.JobOffer) error {
	result, err := store.jobs.InsertOne(context.TODO(), jobOffer)
	if err != nil {
		return err
	}
	jobOffer.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *JobOfferMongoDBStore) filterOne(filter interface{}) (job *domain.JobOffer, err error) {
	result := store.jobs.FindOne(context.TODO(), filter)
	err = result.Decode(&job)
	return
}

func (store *JobOfferMongoDBStore) filter(filter interface{}) ([]*domain.JobOffer, error) {
	cursor, err := store.jobs.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeJobs(cursor)
}

func (store JobOfferMongoDBStore) DeleteAll() {
	store.jobs.DeleteMany(context.TODO(), bson.D{{}})
}

func decodeJobs(cursor *mongo.Cursor) (jobs []*domain.JobOffer, err error) {
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
