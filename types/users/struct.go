package users

import (
	"github.com/hunterlong/statping/types/null"
	"github.com/hunterlong/statping/utils"
	"time"
)

// User is the main struct for Users
type User struct {
	Id        int64         `gorm:"primary_key;column:id" json:"id"`
	Username  string        `gorm:"type:varchar(100);unique;column:username;" json:"username,omitempty"`
	Password  string        `gorm:"column:password" json:"password,omitempty"`
	Email     string        `gorm:"type:varchar(100);unique;column:email" json:"email,omitempty"`
	ApiKey    string        `gorm:"column:api_key" json:"api_key,omitempty"`
	ApiSecret string        `gorm:"column:api_secret" json:"api_secret,omitempty"`
	Admin     null.NullBool `gorm:"column:administrator" json:"admin,omitempty"`
	CreatedAt time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at" json:"updated_at"`
}

// BeforeCreate for User will set CreatedAt to UTC
func (u *User) BeforeCreate() (err error) {
	u.ApiKey = utils.RandomString(16)
	u.ApiSecret = utils.RandomString(16)
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now().UTC()
		u.UpdatedAt = time.Now().UTC()
	}
	return
}
