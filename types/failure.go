package types

import (
	"time"
)

type Failure struct {
	Id               int64     `gorm:"primary_key;column:id" json:"id"`
	Issue            string    `gorm:"column:issue" json:"issue"`
	Method           string    `gorm:"column:method" json:"method,omitempty"`
	Service          int64     `gorm:"index;column:service" json:"-"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	FailureInterface `gorm:"-" json:"-"`
}

type FailureInterface interface {
	Ago() string        // Ago returns a human readble timestamp
	ParseError() string // ParseError returns a human readable error for a service failure
}
