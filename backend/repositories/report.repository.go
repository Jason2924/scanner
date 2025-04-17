package repositories

import (
	"context"

	"github.com/Jason2924/scanner/backend/databases"
	"github.com/Jason2924/scanner/backend/entities"
	"github.com/Jason2924/scanner/backend/models"
)

type IReportRepository interface {
	ReadCurrent(ctxt context.Context, reqt *models.ReportReadCurrentReqt) (*entities.ReportSchema, error)
	ReadByCondition(ctxt context.Context, reqt *models.ReportReadByConditionReqt) (*entities.ReportSchema, error)
	ReadMany(ctxt context.Context, reqt *models.ReportReadManyReqt) ([]entities.ReportSchema, error)
	CountAll(ctxt context.Context) (int64, error)
	InsertCurrent(ctxt context.Context, entity *entities.ReportSchema) error
}

type reportRepository struct {
	iMysqlDatabase databases.IMysqlDatabase
}

func NewReportRepository(mysqlDtbs databases.IMysqlDatabase) IReportRepository {
	return &reportRepository{
		iMysqlDatabase: mysqlDtbs,
	}
}

func (rep *reportRepository) ReadCurrent(ctxt context.Context, reqt *models.ReportReadCurrentReqt) (*entities.ReportSchema, error) {
	item := &entities.ReportSchema{}
	conn := rep.iMysqlDatabase.Connect().WithContext(ctxt)
	result := conn.Where("latitude = ? AND longitude = ?", reqt.Latitude, reqt.Longitude).Order("timestamp desc").Take(item)
	if result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}

func (rep *reportRepository) ReadByCondition(ctxt context.Context, reqt *models.ReportReadByConditionReqt) (*entities.ReportSchema, error) {
	item := &entities.ReportSchema{}
	conn := rep.iMysqlDatabase.Connect().WithContext(ctxt)
	result := conn.Where("latitude = ? AND longitude = ? AND timestamp = ?", reqt.Latitude, reqt.Longitude, reqt.Timestamp).Take(item)
	if result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}

func (rep *reportRepository) ReadMany(ctxt context.Context, reqt *models.ReportReadManyReqt) ([]entities.ReportSchema, error) {
	items := []entities.ReportSchema{}
	conn := rep.iMysqlDatabase.Connect().WithContext(ctxt)
	result := conn.Where("latitude = ? AND longitude = ?", reqt.Latitude, reqt.Longitude).Order("timestamp desc").Offset(reqt.Limit * (reqt.Page - 1)).Limit(reqt.Limit).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

func (rep *reportRepository) CountAll(ctxt context.Context) (int64, error) {
	total := int64(0)
	conn := rep.iMysqlDatabase.Connect().WithContext(ctxt)
	result := conn.Model(entities.ReportSchema{}).Count(&total)
	if result.Error != nil {
		return 0, result.Error
	}
	return total, nil
}

func (rep *reportRepository) InsertCurrent(ctxt context.Context, entity *entities.ReportSchema) error {
	conn := rep.iMysqlDatabase.Connect().WithContext(ctxt)
	result := conn.Model(&entities.ReportSchema{}).Create(entity)
	return result.Error
}
