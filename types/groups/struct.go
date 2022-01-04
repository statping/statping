package groups

import (
	"github.com/statping/statping/types/null"
	"time"
)

// Group is the main struct for Groups
type Group struct {
	Id        int64         `gorm:"primary_key;column:id" json:"id"`
	Name      string        `gorm:"column:name" json:"name"`
	Public    null.NullBool `gorm:"default:true;column:public" json:"public"`
	Order     int           `gorm:"default:0;column:order_id" json:"order_id"`
	CreatedAt time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at" json:"updated_at"`
}

// GroupOrder will reorder the groups based on 'order_id' (Order)
type GroupOrder []*Group

// Sort interface for resorting the Groups in order
func (c GroupOrder) Len() int           { return len(c) }
func (c GroupOrder) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c GroupOrder) Less(i, j int) bool { return c[i].Order < c[j].Order }
