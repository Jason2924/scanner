package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jason2924/scanner/backend/models"
	svc "github.com/Jason2924/scanner/backend/services"
	"github.com/Jason2924/scanner/backend/ultilities"
	"github.com/gin-gonic/gin"
)

type IReportController interface {
	ReadCurrent(ctxt *gin.Context)
	ReadMany(ctxt *gin.Context)
	InsertCurrent(ctxt *gin.Context)
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
	var (
		reqt = &models.ReportReadCurrentReqt{}
		resp = &models.ReportReadResp{}
		erro error
	)
	if erro = ultilities.BindRequest(ctxt, ultilities.BindTypeQuery, reqt); erro != nil {
		log.Println("Error:", erro.Error())
		result := models.NewResponse("error occured while getting data", nil)
		ctxt.AbortWithStatusJSON(http.StatusConflict, result)
		return
	}
	resp, erro = ctr.iReportService.ReadCurrent(ctxt, reqt)
	if erro != nil {
		log.Println("Error:", erro.Error())
		result := models.NewResponse("error occured while getting data", nil)
		ctxt.AbortWithStatusJSON(http.StatusConflict, result)
		return
	}
	result := models.NewResponse("success", resp)
	ctxt.JSON(http.StatusOK, result)
}

func (ctr *reportController) ReadMany(ctxt *gin.Context) {
	var (
		reqt = &models.ReportReadManyReqt{}
		resp = &models.ReportReadManyResp{}
		erro error
	)
	if erro = ultilities.BindRequest(ctxt, ultilities.BindTypeQuery, reqt); erro != nil {
		log.Println("Error:", erro.Error())
		result := models.NewResponse("error occured while getting data", nil)
		ctxt.AbortWithStatusJSON(http.StatusConflict, result)
		return
	}
	fmt.Println(reqt)
	if resp, erro = ctr.iReportService.ReadMany(ctxt, reqt); erro != nil {
		log.Println("Error:", erro.Error())
		result := models.NewResponse("error occured while getting data", nil)
		ctxt.AbortWithStatusJSON(http.StatusConflict, result)
		return
	}
	result := models.NewResponse("success", resp)
	ctxt.JSON(http.StatusOK, result)
}

func (ctr *reportController) InsertCurrent(ctxt *gin.Context) {
	var (
		reqt = &models.ReportInsertCurrentReqt{}
		erro error
	)
	if erro = ultilities.BindRequest(ctxt, ultilities.BindTypeJson, reqt); erro != nil {
		log.Println("Error:", erro.Error())
		result := models.NewResponse("error occured while writing data", nil)
		ctxt.AbortWithStatusJSON(http.StatusConflict, result)
		return
	}
	if erro = ctr.iReportService.InsertCurrent(ctxt, reqt); erro != nil {
		log.Println("Error:", erro.Error())
		result := models.NewResponse("error occured while writing data", nil)
		ctxt.AbortWithStatusJSON(http.StatusConflict, result)
		return
	}
	result := models.NewResponse("success", nil)
	ctxt.JSON(http.StatusOK, result)
}
