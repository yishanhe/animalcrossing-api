package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Critter struct {
	ObjectID        primitive.ObjectID `bson:"_id"`
	ID              int64              `bson:"internalId"`
	EntryID         string             `bson:"uniqueEntryId"`
	Name            string             `bson:"name"`
	ActiveMonths    Months             `bson:"activeMonths"`
	Sell            int64              `bson:"sell"`
	Location        string             `bson:"whereHow"`
	Shadow          string             `bson:"shadow"`
	Colors          []string           `bson:"colors"`
	ActiveHours     [][]string         `bson:"activeHours"`
	RainSnowCatchUp bool               `bson:"rainSnowCatchup"`
	Size            string             `bson:"size"`
	SpecialSell     int64              `bson:"specialSell"`
	Rarity          string             `bson:"rarity"`
}

type Months struct {
	Northern []int64 `bson: "northern"`
	Southern []int64 `bson: "northern"`
}
