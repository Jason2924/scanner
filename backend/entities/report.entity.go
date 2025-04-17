package entities

import (
	"time"

	"github.com/google/uuid"
)

type ReportSchema struct {
	ID          uuid.UUID `gorm:"column:id;type:varchar(255);primaryKey"`
	Latitude    float64   `gorm:"column:latitude;index"`
	Longitude   float64   `gorm:"column:longitude;index"`
	Location    string    `gorm:"column:location;type:varchar(100)"`
	Timestamp   int64     `gorm:"column:timestamp;index"`
	Timezone    int       `gorm:"column:timezone;type:smallint"`
	Unit        string    `gorm:"column:unit;type:varchar(50)"`
	Temperature float64   `gorm:"column:temperature"`
	Pressure    int       `gorm:"column:pressure;type:smallint"`
	Humidity    int       `gorm:"column:humidity;type:smallint"`
	CloudCover  int       `gorm:"column:cloud_cover;type:smallint"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}

func (ntt *ReportSchema) TableName() string {
	return "reports"
}
