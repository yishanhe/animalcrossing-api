package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/yishanhe/animalcrossing-api/pkg/database"
	"github.com/yishanhe/animalcrossing-api/restapi/operations/fish"
)

type getFishById struct {
}

func NewGetFishById() fish.GetFishHandler {
	return &getFishById{}
}

func (f *getFishById) Handle(params fish.GetFishParams) middleware.Responder {
	id := int(params.ID)
	found := database.NewDatabaseClient().FindCritterByID(id, "fishes")

	
	return fish.NewGetFishNotFound()
}
