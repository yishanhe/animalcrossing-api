package handlers

import (
	"context"
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/yishanhe/animalcrossing-api/models"
	"github.com/yishanhe/animalcrossing-api/pkg/database"
	"github.com/yishanhe/animalcrossing-api/pkg/entities"
	"github.com/yishanhe/animalcrossing-api/restapi/operations/fish"
	"go.mongodb.org/mongo-driver/bson"
)

type listFishes struct {
}

func NewListFishes() fish.ListFishesHandler {
	return &listFishes{}
}

func (d *listFishes) Handle(params fish.ListFishesParams) middleware.Responder {
	coll := database.NewMongoClient().Database("AnimalCrossingDB").Collection("fishes")

	cursor, err := coll.Find(context.Background(), bson.M{})

	var results []*models.Fish
	var count = 0

	for cursor.Next(context.TODO()) {
		count++
		var fishEntity entities.Critter
		if err = cursor.Decode(&fishEntity); err != nil {
			log.Fatal(err)
		}
		ele := &models.Fish{
			Price:  fishEntity.Sell,
			Shadow: fishEntity.Shadow,
			Availability: &models.Availability{
				Months: &models.Months{
					Northern: fishEntity.ActiveMonths.Northern,
					Southern: fishEntity.ActiveMonths.Southern,
				},
				Hours:    fishEntity.ActiveHours,
				Location: fishEntity.Location,
				Rarity:   models.Rarity(fishEntity.Rarity),
			},
		}
		if fishEntity.RainSnowCatchUp == true {
			ele.Availability.Weather = append(ele.Availability.Weather, models.WeatherSnow)
			ele.Availability.Weather = append(ele.Availability.Weather, models.WeatherRain)
		} else {
			ele.Availability.Weather = append(ele.Availability.Weather, models.WeatherAny)
		}

		ele.SetID(&fishEntity.ID)
		ele.SetName(&models.Name{
			NameEn: fishEntity.Name,
		})
		results = append(results, ele)
	}

	listResult := &models.FishListResult{
		ListResult: models.ListResult{
			PageCursor:  "a",
			ResultCount: int64(count),
		},
		Results: results,
	}

	return fish.NewListFishesOK().WithPayload(listResult)
}
