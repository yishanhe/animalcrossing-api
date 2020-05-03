package handlers

import (
	"context"
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/yishanhe/animalcrossing-api/models"
	"github.com/yishanhe/animalcrossing-api/pkg/entities"
	"github.com/yishanhe/animalcrossing-api/restapi/operations/bug"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type listBugs struct {
	db *mongo.Client
}

func NewListBugs(db *mongo.Client) bug.ListBugsHandler {
	return &listBugs{
		db: db,
	}
}

func (d *listBugs) Handle(params bug.ListBugsParams) middleware.Responder {
	coll := d.db.Database("AnimalCrossingDB").Collection("bugs")

	cursor, err := coll.Find(context.Background(), bson.M{})

	var results []*models.Bug
	var count = 0

	for cursor.Next(context.TODO()) {
		count++
		var entity entities.Critter
		if err = cursor.Decode(&entity); err != nil {
			log.Fatal(err)
		}
		ele := &models.Bug{
			Price:  entity.Sell,
			Shadow: entity.Shadow,
			Availability: &models.Availability{
				Months: &models.Months{
					Northern: entity.ActiveMonths.Northern,
					Southern: entity.ActiveMonths.Southern,
				},
				Hours:    entity.ActiveHours,
				Location: entity.Location,
				Rarity:   models.Rarity(entity.Rarity),
			},
		}
		if entity.RainSnowCatchUp == true {
			ele.Availability.Weather = append(ele.Availability.Weather, models.WeatherSnow)
			ele.Availability.Weather = append(ele.Availability.Weather, models.WeatherRain)
		} else {
			ele.Availability.Weather = append(ele.Availability.Weather, models.WeatherAny)
		}
		ele.SetID(&entity.ID)
		ele.SetName(&models.Name{
			NameEn: entity.Name,
		})
		results = append(results, ele)
	}
	listResult := &models.BugListResult{
		ListResult: models.ListResult{
			PageCursor:  "a",
			ResultCount: int64(count),
		},
		Results: results,
	}
	return bug.NewListBugsOK().WithPayload(listResult)
}

