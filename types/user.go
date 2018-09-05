package types

import (
	"time"
)

type User struct {
	Id            int64     `gorm:"primary_key;column:id" json:"id"`
	Username      string    `gorm:"type:varchar(100);unique;column:username;" json:"username"`
	Password      string    `gorm:"column:password" json:"-"`
	Email         string    `gorm:"type:varchar(100);unique;column:email" json:"-"`
	ApiKey        string    `gorm:"column:api_key" json:"api_key"`
	ApiSecret     string    `gorm:"column:api_secret" json:"-"`
	Admin         bool      `gorm:"column:administrator" json:"admin"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UserInterface `gorm:"-" json:"-"`
}

type UserInterface interface {
	// Database functions
	Create() (int64, error)
	Update() error
	Delete() error
}
