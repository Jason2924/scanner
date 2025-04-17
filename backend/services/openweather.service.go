package services

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/Jason2924/scanner/backend/models"
	"github.com/Jason2924/scanner/backend/ultilities"
)

type IOpenWeatherService interface {
	GetCurrentReport(ctxt context.Context, longitude float64, latitude float64, unit string) (*models.OpenWeatherCurrentResp, error)
}

type openWeatherService struct {
	apiKey string
}

func NewOpenWeatherService(apiKey string) IOpenWeatherService {
	return &openWeatherService{
		apiKey: apiKey,
	}
}

func (svc *openWeatherService) GetCurrentReport(ctxt context.Context, longitude float64, latitude float64, unit string) (*models.OpenWeatherCurrentResp, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=%s", latitude, longitude, svc.apiKey, unit)
	model := &models.OpenWeatherCurrentResp{}
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	erro := ultilities.HttpGet(ctxt, url, headers, model)
	if erro != nil {
		return nil, erro
	}
	return model, nil
}

func (svc *openWeatherService) GetLocalCurrentReport(longitude float64, latitude float64, unit string) (*models.OpenWeatherCurrentResp, error) {
	rootPath, _ := os.Getwd()
	file, erro := os.Open(rootPath + "/data/openweather/openweather-current.data.json")
	if erro != nil {
		return nil, erro
	}
	defer file.Close()
	data, erro := io.ReadAll(file)
	if erro != nil {
		return nil, erro
	}
	model := &models.OpenWeatherCurrentResp{}
	erro = ultilities.ParseObjectFromJson(data, model)
	if erro != nil {
		return nil, erro
	}
	return model, nil
}
