package services

import (
	"context"
	"time"

	"github.com/Jason2924/scanner/backend/entities"
	"github.com/Jason2924/scanner/backend/repositories"
	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
)

type IFakeDataService interface {
	InsertReports(ctxt context.Context, amount int) error
}

type fakeDataService struct {
	iReportReporitory repositories.IReportRepository
}

func NewFakeDataService(reportRepo repositories.IReportRepository) IFakeDataService {
	return &fakeDataService{
		iReportReporitory: reportRepo,
	}
}

func (svc *fakeDataService) InsertReports(ctxt context.Context, amount int) error {
	reports := make([]entities.ReportSchema, 0, amount)
	now := time.Now()
	for i := amount - 1; i >= 0; i-- {
		timestamp := now.Add(-time.Duration(i) * time.Hour)
		report := svc.generateReport(timestamp.Unix())
		reports = append(reports, *report)
	}
	if erro := svc.iReportReporitory.InsertMany(ctxt, reports); erro != nil {
		return erro
	}
	return nil
}

func (svc *fakeDataService) generateReport(timestamp int64) *entities.ReportSchema {
	gofakeit.Seed(time.Now().UnixNano())
	return &entities.ReportSchema{
		ID:          uuid.New(),
		Latitude:    1.3586,
		Longitude:   103.9899,
		Location:    "Tampines Estate, SG",
		Timestamp:   timestamp,
		Timezone:    28800,
		Unit:        "metric",
		Temperature: float32(gofakeit.Float64Range(10.0, 35.0)),
		Pressure:    gofakeit.Number(980, 1050),
		Humidity:    gofakeit.Number(40, 100),
		CloudCover:  gofakeit.Number(0, 100),
		CreatedAt:   time.Now(),
	}
}
