package services

import (
	"context"

	rep "github.com/Jason2924/scanner/backend/repositories"
)

type IReportService interface {
	ReadCurrent(ctxt context.Context)
}

type reportService struct {
	iReportReporitory rep.IReportRepository
}

func NewReportService(reportRepo rep.IReportRepository) IReportService {
	return &reportService{
		iReportReporitory: reportRepo,
	}
}

func (svc *reportService) ReadCurrent(ctxt context.Context) {

}
