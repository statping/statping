package types

import "time"

// Group is the main struct for Groups
type Group struct {
	Id        int64     `gorm:"primary_key;column:id" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Public    NullBool  `gorm:"default:true;column:public" json:"public"`
	Order     int       `gorm:"default:0;column:order_id" json:"order_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
