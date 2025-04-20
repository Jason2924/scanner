package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/Jason2924/scanner/backend/controllers"
	"github.com/Jason2924/scanner/backend/databases"
	"github.com/Jason2924/scanner/backend/middlewares"
	"github.com/Jason2924/scanner/backend/models"
	"github.com/Jason2924/scanner/backend/repositories"
	"github.com/Jason2924/scanner/backend/server"
	"github.com/Jason2924/scanner/backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Initialize(openWeatherKey string, mysqlDtbs databases.IMysqlDatabase, redisCache databases.IRedisCache, scheduler server.IScheduler) *gin.Engine {
	gin.SetMode(convertMode(""))
	ngin := gin.New()
	ngin.SetTrustedProxies(nil)
	ngin.Use(cors.New(setCORS()))

	reportRepo := repositories.NewReportRepository(mysqlDtbs)

	openWeatherSrvc := services.NewOpenWeatherService(openWeatherKey)
	reportSrvc := services.NewReportService(openWeatherSrvc, redisCache, reportRepo)

	reportCtrl := controllers.NewReportController(reportSrvc)

	scheduler.AddOpenWeatherJob(reportSrvc)
	scheduler.Start()

	apiGrup := ngin.Group("api/v1")
	apiGrup.Use(middlewares.TimeoutMiddleware(10 * time.Second))

	setCommonRoutes(apiGrup)
	setReportRoutes(apiGrup, reportCtrl)

	return ngin
}

func convertMode(mode string) string {
	switch strings.ToLower(mode) {
	case "production":
		return gin.ReleaseMode
	case "test":
		return gin.TestMode
	default:
		return gin.DebugMode
	}
}

func setCORS() cors.Config {
	conf := cors.DefaultConfig()
	conf.AllowOrigins = []string{"http://localhost:5173"}
	conf.AllowCredentials = true
	return conf
}

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
		reportGrup.GET("read-many", reportCtrl.ReadMany)
		reportGrup.GET("compare-ids", reportCtrl.CompareByIds)
		reportGrup.GET("count-many", reportCtrl.CountMany)
		reportGrup.POST("insert-current", reportCtrl.InsertCurrent)
	}
}
