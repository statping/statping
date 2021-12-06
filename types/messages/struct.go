package messages

import (
	"github.com/statping/statping/types/null"
	"time"
)

// Message is for creating Announcements, Alerts and other messages for the end users
type Message struct {
	Id                int64          `gorm:"primary_key;column:id" json:"id"`
	Title             string         `gorm:"column:title" json:"title"`
	Description       string         `gorm:"column:description" json:"description"`
	StartOn           time.Time      `gorm:"column:start_on" json:"start_on"`
	EndOn             time.Time      `gorm:"column:end_on" json:"end_on"`
	ServiceId         int64          `gorm:"index;column:service" json:"service"`
	NotifyUsers       null.NullBool  `gorm:"column:notify_users" json:"notify_users" scope:"user,admin"`
	NotifyMethod      string         `gorm:"column:notify_method" json:"notify_method" scope:"user,admin"`
	NotifyBefore      null.NullInt64 `gorm:"column:notify_before" json:"notify_before" scope:"user,admin"`
	NotifyBeforeScale string         `gorm:"column:notify_before_scale" json:"notify_before_scale" scope:"user,admin"`
	CreatedAt         time.Time      `gorm:"column:created_at" json:"created_at" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at" json:"updated_at" json:"updated_at"`
}
