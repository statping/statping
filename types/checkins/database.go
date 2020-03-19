package checkins

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/utils"
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
	c.ApiKey = utils.RandomString(7)
	q := db.Create(c)

	c.Start()
	go c.CheckinRoutine()
	return q.Error()
}

func (c *Checkin) Update() error {
	q := db.Update(c)
	return q.Error()
}

func (c *Checkin) Delete() error {
	c.Close()
	q := dbHits.Where("checkin = ?", c.Id).Delete(&CheckinHit{})
	if err := q.Error(); err != nil {
		return err
	}
	q = db.Model(&Checkin{}).Delete(c)
	return q.Error()
}

//func (c *Checkin) AfterDelete() error {
//	//q := dbHits.Where("checkin = ?", c.Id).Delete(&CheckinHit{})
//	return q.Error()
//}
