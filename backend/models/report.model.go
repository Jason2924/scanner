package models

import (
	"fmt"
	"time"

	"github.com/Jason2924/scanner/backend/entities"
	"github.com/google/uuid"
)

type ReportReadResp struct {
	ID          uuid.UUID `json:"id"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"dateTime"`
	Timestamp   int64     `json:"timestamp"`
	Timezone    int       `json:"timezone"`
	Temperature float32   `json:"temperature"`
	Pressure    int       `json:"pressure"`
	Humidity    int       `json:"humidity"`
	CloudCover  int       `json:"cloudCover"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (mod *ReportReadResp) FromEntity(item *entities.ReportSchema) {
	utcTime := time.Unix(item.Timestamp, 0).UTC()
	zoneTime := item.Timestamp / 3600
	zoneName := "UTC"
	if zoneTime < 0 {
		zoneName += fmt.Sprintf("%s%d", zoneName, zoneTime)
	} else {
		zoneName += fmt.Sprintf("%s-%d", zoneName, zoneTime)
	}
	location := time.FixedZone(zoneName, item.Timezone)
	localTime := utcTime.In(location)

	mod.ID = item.ID
	mod.Latitude = item.Latitude
	mod.Longitude = item.Longitude
	mod.Location = item.Location
	mod.DateTime = localTime
	mod.Timestamp = item.Timestamp
	mod.Timezone = item.Timezone
	mod.Temperature = item.Temperature
	mod.Pressure = item.Pressure
	mod.Humidity = item.Humidity
	mod.CloudCover = item.CloudCover
	mod.CreatedAt = item.CreatedAt

}

type ReportReadCurrentReqt struct {
	Latitude  float64 `form:"latitude" binding:"required"`
	Longitude float64 `form:"longitude" binding:"required"`
}

type ReportReadByConditionReqt struct {
	Latitude  float64
	Longitude float64
	Unit      string
	Timestamp int64
}

type ReportReadManyReqt struct {
	Latitude  float64 `form:"latitude" binding:"required"`
	Longitude float64 `form:"longitude" binding:"required"`
	Limit     int     `form:"limit" binding:"required"`
	Page      int     `form:"page" binding:"required"`
}

type ReportReadManyResp struct {
	List *[]ReportReadResp `json:"list"`
}

func (mod *ReportReadManyResp) FromEntities(items []entities.ReportSchema) {
	list := make([]ReportReadResp, 0, len(items))
	for _, item := range items {
		parsedItem := ReportReadResp{}
		parsedItem.FromEntity(&item)
		list = append(list, parsedItem)
	}
	mod.List = &list
}

type ReportCompareByIdsReqt struct {
	Ids []string `form:"ids[]" binding:"required"`
}

type ReportCompareByIdsResp struct {
	List *[]ReportReadResp `json:"list"`
}

func (mod *ReportCompareByIdsResp) FromEntities(items []entities.ReportSchema) {
	list := make([]ReportReadResp, 0, len(items))
	for _, item := range items {
		parsedItem := ReportReadResp{}
		parsedItem.FromEntity(&item)
		list = append(list, parsedItem)
	}
	mod.List = &list
}

type ReportCountManyReqt struct {
	Latitude  float64 `form:"latitude" binding:"required"`
	Longitude float64 `form:"longitude" binding:"required"`
}

type ReportCountManyResp struct {
	Total int64 `json:"total"`
}

type ReportInsertCurrentReqt struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Unit      string  `json:"unit"`
}
