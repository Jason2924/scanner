package handlers

import (
	"net/http"
	"time"

	"github.com/Jason2924/scanner/backend/controllers"
	"github.com/Jason2924/scanner/backend/databases"
	"github.com/Jason2924/scanner/backend/middlewares"
	"github.com/Jason2924/scanner/backend/models"
	"github.com/Jason2924/scanner/backend/repositories"
	"github.com/Jason2924/scanner/backend/services"
	"github.com/gin-gonic/gin"
)

func Initialize(mysqlDtbs databases.IMysqlDatabase) *gin.Engine {
	// gin.SetMode(convertMode(conf.Mode))
	ngin := gin.New()

	reportRepo := repositories.NewReportRepository(mysqlDtbs)

	reportSrvc := services.NewReportService(reportRepo)

	reportCtrl := controllers.NewReportController(reportSrvc)

	apiGrup := ngin.Group("api/v1")
	apiGrup.Use(middlewares.TimeoutMiddleware(10 * time.Second))

	setCommonRoutes(apiGrup)
	setReportRoutes(apiGrup, reportCtrl)

	return ngin
}

// func convertMode(mode string) string {
// 	switch strings.ToLower(mode) {
// 	case "production":
// 		return gin.ReleaseMode
// 	case "test":
// 		return gin.TestMode
// 	default:
// 		return gin.DebugMode
// 	}
// }

func setCommonRoutes(apiGrup *gin.RouterGroup) {
	apiGrup.GET("ping", func(ctxt *gin.Context) {
		resp := models.NewResponse("pong", nil)
		ctxt.JSON(http.StatusOK, resp)
	})
}

func setReportRoutes(apiGrup *gin.RouterGroup, reportCtrl controllers.IReportController) {
	reportGrup := apiGrup.Group("reports")
	{
		reportGrup.GET("read-current", reportCtrl.ReadCurrent)
	}
}
