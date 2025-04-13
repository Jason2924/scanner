package entities

import (
	"time"

	"github.com/google/uuid"
)

type ReportSchema struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey"`
	Latitude    string    `gorm:"column:latitude"`
	Longitude   string    `gorm:"column:longitude"`
	Location    string    `gorm:"column:location"`
	Time        time.Time `gorm:"column:time"`
	Temperature float64   `gorm:"column:temperature"`
	Pressure    int       `gorm:"column:pressure"`
	Humidity    int       `gorm:"column:humidity"`
	CloudCover  int       `gorm:"column:cloud_cover"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}

func (ntt *ReportSchema) TableName() string {
	return "reports"
}
