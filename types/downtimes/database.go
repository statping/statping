package downtimes

import (
	"fmt"
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/services"
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

func ConvertToUnixTime(str string) (time.Time,error){
	i, err := strconv.ParseInt(str, 10, 64)
	var t time.Time
	if err != nil {
		return t,err
	}
	tm := time.Unix(i, 0)
	return tm,nil
}
type invalidTimeDurationError struct{}

func (m *invalidTimeDurationError) Error() string {
	return "invalid time duration"
}
func FindAll(vars map[string]string ,start time.Time, end time.Time) (*[]Downtime, error) {
	var downtime []Downtime
	q := db.Where("start BETWEEN ? AND ? ", start, end)
	for key,val :=range vars{
		switch key{
		case "start":
			_,ok := vars["end"]
			if ok && (vars["end"]>vars["start"]) {
				start,err := ConvertToUnixTime(vars["start"])
				if err!=nil {
					return &downtime,err
				}
				end,err := ConvertToUnixTime(vars["end"])
				if err!=nil {
					return &downtime,err
				}
				q = q.Where("start BETWEEN ? AND ? ", start, end)
			}else {
				return &downtime,&invalidTimeDurationError{}
			}
		case "sub_status":
			q = q.Where(" sub_status = ?",val)
		case "service":
			allServices := services.All()
			for k,v := range allServices{
			if v.Name == val {
				q = q.Where(" service = ?",k)
				}
			}
		case "type":
			q = q.Where(" type = ?",val)
		}
	}
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
	q := db.Model(&Downtime{}).Delete(c)
	return q.Error()
}
