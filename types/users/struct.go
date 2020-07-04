package users

import (
	"github.com/statping/statping/types/null"
	"time"
)

// User is the main struct for Users
type User struct {
	Id        int64         `gorm:"primary_key;column:id" json:"id"`
	Username  string        `gorm:"type:varchar(100);unique;column:username;" json:"username,omitempty"`
	Password  string        `gorm:"column:password" json:"password,omitempty"`
	Email     string        `gorm:"type:varchar(100);column:email" json:"email,omitempty"`
	ApiKey    string        `gorm:"column:api_key" json:"api_key,omitempty"`
	Scopes    string        `gorm:"column:scopes" json:"scopes,omitempty"`
	Admin     null.NullBool `gorm:"column:administrator" json:"admin,omitempty"`
	CreatedAt time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at" json:"updated_at"`
	Token     string        `gorm:"-" json:"token"`
}
