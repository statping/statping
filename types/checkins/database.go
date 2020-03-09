package checkins

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/utils"
)

func DB() database.Database {
	return database.DB().Model(&Checkin{})
}

func DBhits() database.Database {
	return database.DB().Model(&CheckinHit{})
}

func Find(id int64) (*Checkin, error) {
	var checkin Checkin
	db := DB().Where("id = ?", id).Find(&checkin)
	return &checkin, db.Error()
}

func FindByAPI(key string) (*Checkin, error) {
	var checkin Checkin
	db := DB().Where("api = ?", key).Find(&checkin)
	return &checkin, db.Error()
}

func All() []*Checkin {
	var checkins []*Checkin
	DB().Find(&checkins)
	return checkins
}

func (c *Checkin) Create() error {
	c.ApiKey = utils.RandomString(7)
	db := DB().Create(c)

	c.Start()
	go c.CheckinRoutine()
	return db.Error()
}

func (c *Checkin) Update() error {
	db := DB().Update(c)
	return db.Error()
}

func (c *Checkin) Delete() error {
	c.Close()
	db := DB().Delete(c)
	return db.Error()
}
