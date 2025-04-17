package services

import (
	"context"
	"fmt"
	"time"

	"github.com/Jason2924/scanner/backend/databases"
	"github.com/Jason2924/scanner/backend/mappers"
	"github.com/Jason2924/scanner/backend/models"
	"github.com/Jason2924/scanner/backend/repositories"
)

type IReportService interface {
	ReadCurrent(ctxt context.Context, reqt *models.ReportReadCurrentReqt) (*models.ReportReadResp, error)
	ReadMany(ctxt context.Context, reqt *models.ReportReadManyReqt) (*models.ReportReadManyResp, error)
	InsertCurrent(ctxt context.Context, reqt *models.ReportInsertCurrentReqt) error
}

type reportService struct {
	iOpenWeatherService IOpenWeatherService
	iRedisCache         databases.IRedisCache
	iReportReporitory   repositories.IReportRepository
}

func NewReportService(openWeatherSrvc IOpenWeatherService, redisCache databases.IRedisCache, reportRepo repositories.IReportRepository) IReportService {
	return &reportService{
		iOpenWeatherService: openWeatherSrvc,
		iRedisCache:         redisCache,
		iReportReporitory:   reportRepo,
	}
}

func (svc *reportService) ReadCurrent(ctxt context.Context, reqt *models.ReportReadCurrentReqt) (*models.ReportReadResp, error) {
	cacheName := "weather:current"
	resp := &models.ReportReadResp{}
	hasCache, erro := svc.iRedisCache.Retrieve(ctxt, cacheName, resp)
	if erro != nil {
		return nil, erro
	}
	if !hasCache {
		item, erro := svc.iReportReporitory.ReadCurrent(ctxt, reqt)
		if erro != nil {
			return nil, erro
		}
		resp.FromEntity(item)
		if erro = svc.iRedisCache.Store(ctxt, cacheName, resp, 15*time.Minute); erro != nil {
			return nil, erro
		}
	}
	return resp, nil
}

func (svc *reportService) ReadMany(ctxt context.Context, reqt *models.ReportReadManyReqt) (*models.ReportReadManyResp, error) {
	cacheName := fmt.Sprintf("weather:history:%d:%d", reqt.Limit, reqt.Page)
	resp := &models.ReportReadManyResp{}
	hasCache, erro := svc.iRedisCache.Retrieve(ctxt, cacheName, resp)
	if erro != nil {
		return nil, erro
	}
	if !hasCache {
		items, erro := svc.iReportReporitory.ReadMany(ctxt, reqt)
		if erro != nil {
			return nil, erro
		}
		total, erro := svc.iReportReporitory.CountAll(ctxt)
		if erro != nil {
			return nil, erro
		}
		resp.FromEntities(items)
		resp.Total = total
		if erro = svc.iRedisCache.Store(ctxt, cacheName, resp, 15*time.Minute); erro != nil {
			return nil, erro
		}
	}
	return resp, nil
}

func (svc *reportService) InsertCurrent(ctxt context.Context, reqt *models.ReportInsertCurrentReqt) error {
	openWeatherModel, erro := svc.iOpenWeatherService.GetCurrentReport(ctxt, reqt.Longitude, reqt.Latitude, reqt.Unit)
	if erro != nil {
		return erro
	}
	reportEntity := mappers.MapOpenWeatherToReport(openWeatherModel)
	erro = svc.iReportReporitory.InsertCurrent(ctxt, reportEntity)
	if erro != nil {
		return erro
	}
	return nil
}
