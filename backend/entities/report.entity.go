package entities

import (
	"time"

	"github.com/google/uuid"
)

type ReportSchema struct {
	ID          uuid.UUID `gorm:"column:id;type:varchar(255);primaryKey"`
	Latitude    float64   `gorm:"column:latitude;type:decimal(10,6);index:idx_lat_lon_time,priority:1"`
	Longitude   float64   `gorm:"column:longitude;type:decimal(10,6);index:idx_lat_lon_time,priority:2"`
	Location    string    `gorm:"column:location;type:varchar(100)"`
	Timestamp   int64     `gorm:"column:timestamp;idx_lat_lon_time,priority:3"`
	Timezone    int       `gorm:"column:timezone;type:smallint"`
	Unit        string    `gorm:"column:unit;type:varchar(50)"`
	Temperature float32   `gorm:"column:temperature;type:decimal(10,2)"`
	Pressure    int       `gorm:"column:pressure;type:smallint"`
	Humidity    int       `gorm:"column:humidity;type:smallint"`
	CloudCover  int       `gorm:"column:cloud_cover;type:smallint"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}

func (ntt *ReportSchema) TableName() string {
	return "reports"
}
