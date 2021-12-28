package downtimes

import (
	"fmt"
	"github.com/statping/statping/database"
	"time"
)

var (
	zeroInt64 int64
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

func (c *Downtime) Validate() error {
	if c.Type == "manual" {
		if c.End != nil && c.End.After(time.Now()) || c.Start.After(time.Now()) {
			return fmt.Errorf("Downtime cannot be in future")
		}
		if c.ServiceId == zeroInt64 {
			return fmt.Errorf("Service ID cannot be null")
		}
		if c.SubStatus != "down" && c.SubStatus != "degraded" {
			return fmt.Errorf("SubStatus can only be 'down' or 'degraded'")
		}
	}
	return nil
}

func (c *Downtime) BeforeCreate() error {
	return c.Validate()
}

func (c *Downtime) BeforeUpdate() error {
	return c.Validate()
}

func FindByService(service int64, start time.Time, end time.Time) (*[]Downtime, error) {
	var downtime []Downtime
	q := db.Where("service = ? and start BETWEEN ? AND ? ", service, start, end)
	q = q.Order("id ASC ").Find(&downtime)
	return &downtime, q.Error()
}

func FindDowntime(service int64, timeVar time.Time) *Downtime {
	var downtime []Downtime
	q := db.Where("service = $1 and start BETWEEN $2 AND $3 ", service, time.Time{}, timeVar)
	q = q.Order("id ASC ").Find(&downtime)
	downtimeList := *(&downtime)
	for _,dtime := range downtimeList{
		if (*(dtime.End)).Unix() > timeVar.Unix(){
			return &dtime
		}
	}
	return nil
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
	q := db.Model(&Downtime{}).Delete(c)
	return q.Error()
}
