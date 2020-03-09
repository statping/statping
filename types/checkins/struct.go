package checkins

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/utils"
	"time"
)

// Checkin struct will allow an application to send a recurring HTTP GET to confirm a service is online
type Checkin struct {
	Id          int64               `gorm:"primary_key;column:id" json:"id"`
	ServiceId   int64               `gorm:"index;column:service" json:"service_id"`
	Name        string              `gorm:"column:name" json:"name"`
	Interval    int64               `gorm:"column:check_interval" json:"interval"`
	GracePeriod int64               `gorm:"column:grace_period"  json:"grace"`
	ApiKey      string              `gorm:"column:api_key"  json:"api_key"`
	CreatedAt   time.Time           `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time           `gorm:"column:updated_at" json:"updated_at"`
	Running     chan bool           `gorm:"-" json:"-"`
	Failing     bool                `gorm:"-" json:"failing"`
	LastHitTime time.Time           `gorm:"-" json:"last_hit"`
	AllHits     []*CheckinHit       `gorm:"-" json:"hits"`
	AllFailures []*failures.Failure `gorm:"-" json:"failures"`
}

// CheckinHit is a successful response from a Checkin
type CheckinHit struct {
	Id        int64     `gorm:"primary_key;column:id" json:"id"`
	Checkin   int64     `gorm:"index;column:checkin" json:"-"`
	From      string    `gorm:"column:from_location" json:"from"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

// BeforeCreate for checkinHit will set CreatedAt to UTC
func (c *CheckinHit) BeforeCreate() (err error) {
	if c.CreatedAt.IsZero() {
		c.CreatedAt = time.Now().UTC()
	}
	return
}

func (c *Checkin) BeforeCreate() (err error) {
	c.ApiKey = utils.RandomString(7)
	if c.CreatedAt.IsZero() {
		c.CreatedAt = time.Now().UTC()
		c.UpdatedAt = time.Now().UTC()
	}
	return
}

func (c *Checkin) BeforeDelete(tx database.Database) (err error) {
	return tx.Where("id = ?", c.ServiceId).
		Update("group_id", 0).Error()
}
