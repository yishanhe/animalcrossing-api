package database

import (
	"context"
	"log"

	"github.com/yishanhe/animalcrossing-api/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbClient struct {
	db *mongo.Client
}

type AnimalCrossingDB interface {
	FindCritterByID(id int, resourceType string) (*entities.Critter, error)
}

func NewMongoClient() *mongo.Client {

	var err error
	var client *mongo.Client
	opts := options.Client()
	opts.ApplyURI("mongodb://localhost:27017")
	opts.SetMaxPoolSize(5)
	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		log.Println(err.Error())
	}
	return client
}

func NewDatabaseClient() AnimalCrossingDB {

	var err error
	var client *mongo.Client
	opts := options.Client()
	opts.ApplyURI("mongodb://localhost:27017")
	opts.SetMaxPoolSize(5)
	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		log.Println(err.Error())
	}
	return &dbClient{
		db: client,
	}
}

func (d dbClient) FindCritterByID(id int, resourceType string) (*entities.Critter, error) {
	coll := d.db.Database("AnimalCrossingDB").Collection(resourceType)
	filter := bson.M{
		"internalId": id,
	}
	var found *entities.Critter
	err := coll.FindOne(context.Background(), filter).Decode(&found)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Panicln(err)
		panic(err)
	}
	if err == mongo.ErrNoDocuments {
		return nil, err
	}
	return found, err
}
