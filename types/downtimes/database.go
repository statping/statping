package downtimes

import (
	"fmt"
	"github.com/statping/statping/database"
	"time"
)

var db database.Database
var dbHits database.Database

func SetDB(database database.Database) {
	db = database.Model(&Downtime{})
}

func Find(id int64) (*Downtime, error) {
	var downtime Downtime
	q := db.Where("id = ?", id).Find(&downtime)
	if q.Error() != nil {
		return nil, q.Error()
	}
	if q.RecordNotFound() {
		return nil, fmt.Errorf(" Downtime record not found : %s", id)
	}
	return &downtime, q.Error()
}

func FindByService(service int64, start time.Time, end time.Time) (*[]Downtime, error) {
	var downtime []Downtime
	q := db.Where("service = ? and start BETWEEN ? AND ? ", service, start, end)
	q = q.Order("id ASC ").Find(&downtime)
	return &downtime, q.Error()
}

func (c *Downtime) Create() error {
	q := db.Create(c)
	return q.Error()
}

func (c *Downtime) Update() error {
	q := db.Where(" id = ? ", c.Id).Updates(c)
	return q.Error()
}

func (c *Downtime) Delete() error {
	q := dbHits.Where("id = ?", c.Id).Delete(&Downtime{})
	if err := q.Error(); err != nil {
		return err
	}
	q = db.Model(&Downtime{}).Delete(c)
	return q.Error()
}
