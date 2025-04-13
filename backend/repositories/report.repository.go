package repositories

import (
	"context"

	dtb "github.com/Jason2924/scanner/backend/databases"
)

type IReportRepository interface {
	ReadCurrent(ctxt context.Context)
}

type reportRepository struct {
	iMysqlDatabase dtb.IMysqlDatabase
}

func NewReportRepository(mysqlDtbs dtb.IMysqlDatabase) IReportRepository {
	return &reportRepository{
		iMysqlDatabase: mysqlDtbs,
	}
}

func (rep *reportRepository) ReadCurrent(ctxt context.Context) {

}
