package types

import (
	"time"
)

type Checkin struct {
	Id               int64     `gorm:"primary_key;column:id"`
	Service          int64     `gorm:"index;column:service"`
	Interval         int64     `gorm:"column:check_interval"`
	Api              string    `gorm:"column:api"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	Hits             int64     `json:"hits"`
	Last             time.Time `json:"last"`
	CheckinInterface `json:"-"`
}

type CheckinInterface interface {
	// Database functions
	Create() (int64, error)
	//Update() error
	//Delete() error
	Ago() string
	Receivehit()
}
