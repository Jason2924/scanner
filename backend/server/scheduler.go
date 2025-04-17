package server

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/Jason2924/scanner/backend/models"
	"github.com/Jason2924/scanner/backend/services"
	"github.com/robfig/cron"
)

var (
	schedulerInstance IScheduler
	schedulerOnece    sync.Once
)

type IScheduler interface {
	Start()
	Stop()
	AddOpenWeatherJob(reportSrvc services.IReportService)
}

type scheduler struct {
	cronjob *cron.Cron
}

func GetScheduler() IScheduler {
	schedulerOnece.Do(func() {
		schedulerInstance = &scheduler{
			cronjob: cron.New(),
		}
	})
	return schedulerInstance
}

func (sch *scheduler) Start() {
	sch.cronjob.Start()
}

func (sch *scheduler) Stop() {
	sch.cronjob.Stop()
}

func (sch *scheduler) AddOpenWeatherJob(reportSrvc services.IReportService) {
	sch.cronjob.AddFunc("0 */30 * * * *", func() {
		log.Println("Open Weather job is running...")
		ctxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		reqt := &models.ReportInsertCurrentReqt{
			Latitude:  1.3586,
			Longitude: 103.9899,
			Unit:      "metric",
		}
		if erro := reportSrvc.InsertCurrent(ctxt, reqt); erro != nil {
			log.Printf("Error occurred while running Open Weather job: %v\n", erro.Error())
		}
		log.Println("Open Weather job is completed")
	})
}
