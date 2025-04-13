package controllers

import (
	svc "github.com/Jason2924/scanner/backend/services"
	"github.com/gin-gonic/gin"
)

type IReportController interface {
	ReadCurrent(ctxt *gin.Context)
}

type reportController struct {
	iReportService svc.IReportService
}

func NewReportController(reportSrvc svc.IReportService) IReportController {
	return &reportController{
		iReportService: reportSrvc,
	}
}

func (ctr *reportController) ReadCurrent(ctxt *gin.Context) {
	return
}
