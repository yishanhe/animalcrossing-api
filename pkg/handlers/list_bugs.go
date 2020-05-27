package handlers

import (
	"context"
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/yishanhe/animalcrossing-api/models"
	"github.com/yishanhe/animalcrossing-api/pkg/database"
	"github.com/yishanhe/animalcrossing-api/restapi/operations/bug"
	"go.mongodb.org/mongo-driver/bson"
)

type listBugs struct {
}

func NewListBugs() bug.ListBugsHandler {
	return &listBugs{}
}

func (d *listBugs) Handle(params bug.ListBugsParams) middleware.Responder {

	coll := database.NewMongoClient().Database("AnimalCrossingDevDB").Collection("bug")

	cursor, err := coll.Find(context.Background(), bson.M{})

	var results []*models.Bug
	var count = 0

	for cursor.Next(context.TODO()) {
		count++
		var entity *models.Bug
		if err = cursor.Decode(&entity); err != nil {
			log.Fatal(err)
		}
		results = append(results, entity)
	}
	log.Println(count)
	listResult := &models.BugListResult{
		ListResult: models.ListResult{
			PageCursor:  "a",
			ResultCount: int64(count),
		},
		Results: results,
	}
	return bug.NewListBugsOK().WithPayload(listResult)
}
