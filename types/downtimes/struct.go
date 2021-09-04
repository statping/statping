package downtimes

import (
	"time"
)

// Checkin struct will allow an application to send a recurring HTTP GET to confirm a service is online
type Downtime struct {
	Id        int64     `gorm:"primary_key;column:id" json:"id"`
	ServiceId int64     `gorm:"index;column:service" json:"service_id"`
	SubStatus string    `gorm:"column:sub_status" json:"sub_status"`
	Failures  int       `gorm:"column:failures" json:"failures"`
	Start     time.Time `gorm:"index;column:start" json:"start"`
	End       time.Time `gorm:"column:end" json:"end"`
}
