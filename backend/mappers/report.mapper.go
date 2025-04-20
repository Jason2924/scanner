package mappers

import (
	"fmt"
	"time"

	"github.com/Jason2924/scanner/backend/entities"
	"github.com/Jason2924/scanner/backend/models"
	"github.com/google/uuid"
)

func MapOpenWeatherToReport(model *models.OpenWeatherCurrentResp, unit string) *entities.ReportSchema {
	return &entities.ReportSchema{
		ID:          uuid.New(),
		Latitude:    model.Coordinate.Latitude,
		Longitude:   model.Coordinate.Longitude,
		Location:    fmt.Sprintf("%s, %s", model.CityName, model.System.Country),
		Unit:        unit,
		Timestamp:   model.TimeOfData,
		Timezone:    model.Timezone,
		Temperature: model.Main.Temperature,
		Pressure:    model.Main.Pressure,
		Humidity:    model.Main.Humidity,
		CloudCover:  model.Clouds.All,
		CreatedAt:   time.Now(),
	}
}
