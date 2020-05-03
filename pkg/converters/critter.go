package converters

import "github.com/yishanhe/animalcrossing-api/pkg/entities"

import "github.com/yishanhe/animalcrossing-api/models"

func ToFishDTO(entity entities.Critter) *models.Fish {
	ele := &models.Fish{
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
	return ele
}

func ToBugDTO(entity entities.Critter) *models.Fish {
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
	return ele
}
