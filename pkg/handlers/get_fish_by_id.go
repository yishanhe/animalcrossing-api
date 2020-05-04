package handlers

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/yishanhe/animalcrossing-api/pkg/converters"
	"github.com/yishanhe/animalcrossing-api/pkg/database"
	"github.com/yishanhe/animalcrossing-api/pkg/utils"
	"github.com/yishanhe/animalcrossing-api/restapi/operations/fish"
	"go.mongodb.org/mongo-driver/mongo"
)

type getFishById struct {
}

func NewGetFishById() fish.GetFishHandler {
	return &getFishById{}
}

func (f *getFishById) Handle(params fish.GetFishParams) middleware.Responder {
	id := int(params.ID)
	fmt.Println("query id ", id)
	found, err := database.NewDatabaseClient().FindCritterByID(id, "fishes")
	if err == mongo.ErrNoDocuments {
		return fish.NewGetFishNotFound()
	}
	utils.PrettyPrint(found)
	return fish.NewGetFishOK().WithPayload(converters.ToFishDTO(found))
}
