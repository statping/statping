package downtimes

import (
	"fmt"
	"github.com/statping/statping/database"
	"strconv"
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

func ConvertToUnixTime(str string) (time.Time, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	var t time.Time
	if err != nil {
		return t, err
	}
	tm := time.Unix(i, 0)
	return tm, nil
}

func FindAll(vars map[string]string) (*[]Downtime, error) {
	var downtime []Downtime
	var start time.Time
	var end time.Time
	st, err1 := vars["start"]
	en, err2 := vars["end"]
	startInt, err := strconv.ParseInt(st, 10, 64)
	endInt, err := strconv.ParseInt(en, 10, 64)
	if err1 && err2 && (endInt > startInt) {
		start, err = ConvertToUnixTime(vars["start"])
		if err != nil {
			return &downtime, err
		}
		end, err = ConvertToUnixTime(vars["end"])
		if err != nil {
			return &downtime, err
		}
	} else {
		ninetyDaysAgo := time.Now().Add(time.Duration(-90*24) * time.Hour)
		start = ninetyDaysAgo
		end = time.Now()
	}
	q := db.Where("start BETWEEN ? AND ?", start, end)
	if subStatusVar, subStatusErr := vars["sub_status"]; subStatusErr {
		q = q.Where("sub_status = ?", subStatusVar)
	}
	if serviceIdVar, serviceIdErr := vars["service_id"]; serviceIdErr {
		q = q.Where("service = ?", serviceIdVar)
	}
	if typeVar, typeErr := vars["type"]; typeErr {
		q = q.Where("type = ?", typeVar)
	}
	var count int64
	if countVar, countErr := vars["count"]; countErr {
		count, err = strconv.ParseInt(countVar, 10, 64)
		if count > 100 {
			count = 100
		}
	} else {
		count = 20
	}
	var skip int64
	if skipVar, err6 := vars["skip"]; err6 {
		skip, err = strconv.ParseInt(skipVar, 10, 64)
	} else {
		skip = 0
	}
	q = q.Order("start DESC")
	q = q.Limit((int)(count)).Offset((int)(skip)).Find(&downtime)
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
	q := db.Model(&Downtime{}).Delete(c)
	return q.Error()
}
