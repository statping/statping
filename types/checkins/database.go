package checkins

import (
	"github.com/statping/statping/database"
)

var db database.Database
var dbHits database.Database

func SetDB(database database.Database) {
	db = database.Model(&Checkin{})
	dbHits = database.Model(&CheckinHit{})
}

func Find(id int64) (*Checkin, error) {
	var checkin Checkin
	q := db.Where("id = ?", id).Find(&checkin)
	return &checkin, q.Error()
}

func FindByAPI(key string) (*Checkin, error) {
	var checkin Checkin
	q := db.Where("api_key = ?", key).Find(&checkin)
	return &checkin, q.Error()
}

func All() []*Checkin {
	var checkins []*Checkin
	db.Find(&checkins)
	return checkins
}

func (c *Checkin) Create() error {
	q := db.Create(c)
	return q.Error()
}

func (c *Checkin) Update() error {
	q := db.Update(c)
	return q.Error()
}

func (c *Checkin) Delete() error {
	q := db.Delete(c)
	return q.Error()
}
