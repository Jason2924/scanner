package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Jason2924/scanner/backend/databases"
	"github.com/Jason2924/scanner/backend/mappers"
	"github.com/Jason2924/scanner/backend/models"
	"github.com/Jason2924/scanner/backend/repositories"
	"gorm.io/gorm"
)

type IReportService interface {
	ReadCurrent(ctxt context.Context, reqt *models.ReportReadCurrentReqt) (*models.ReportReadResp, error)
	ReadMany(ctxt context.Context, reqt *models.ReportReadManyReqt) (*models.ReportReadManyResp, error)
	CompareByIds(ctxt context.Context, reqt *models.ReportCompareByIdsReqt) (*models.ReportCompareByIdsResp, error)
	CountMany(ctxt context.Context, reqt *models.ReportCountManyReqt) (*models.ReportCountManyResp, error)
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
	cacheName := fmt.Sprintf("weather:list:%d:%d", reqt.Limit, reqt.Page)
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
		resp.FromEntities(items)
		if erro = svc.iRedisCache.Store(ctxt, cacheName, resp, 15*time.Minute); erro != nil {
			return nil, erro
		}
	}
	return resp, nil
}

func (svc *reportService) CompareByIds(ctxt context.Context, reqt *models.ReportCompareByIdsReqt) (*models.ReportCompareByIdsResp, error) {
	resp := &models.ReportCompareByIdsResp{}
	items, erro := svc.iReportReporitory.CompareByIds(ctxt, reqt)
	if erro != nil {
		return nil, erro
	}
	if len(items) != len(reqt.Ids) {
		return nil, fmt.Errorf("Error the length of request and response items not equal")
	}
	resp.FromEntities(items)
	return resp, nil
}

func (svc *reportService) CountMany(ctxt context.Context, reqt *models.ReportCountManyReqt) (*models.ReportCountManyResp, error) {
	cacheName := "weather:count"
	resp := &models.ReportCountManyResp{Total: 0}
	hasCache, erro := svc.iRedisCache.Retrieve(ctxt, cacheName, resp)
	if erro != nil {
		return nil, erro
	}
	if !hasCache {
		total, erro := svc.iReportReporitory.CountMany(ctxt, reqt)
		if erro != nil {
			return nil, erro
		}
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
	reportEntity := mappers.MapOpenWeatherToReport(openWeatherModel, reqt.Unit)
	conditionReqt := &models.ReportReadByConditionReqt{
		Latitude:  reportEntity.Latitude,
		Longitude: reportEntity.Longitude,
		Unit:      reportEntity.Unit,
		Timestamp: reportEntity.Timestamp,
	}
	_, erro = svc.iReportReporitory.ReadByCondition(ctxt, conditionReqt)
	if errors.Is(erro, gorm.ErrRecordNotFound) {
		erro = svc.iReportReporitory.InsertCurrent(ctxt, reportEntity)
		if erro != nil {
			return erro
		}
	} else if erro != nil {
		return erro
	}
	return nil
}
