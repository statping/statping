package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"upper.io/db.v3"
)

var (
	services []*Service
)

type Service struct {
	Id             int64      `db:"id,omitempty" json:"id"`
	Name           string     `db:"name" json:"name"`
	Domain         string     `db:"domain" json:"domain"`
	Expected       string     `db:"expected" json:"expected"`
	ExpectedStatus int        `db:"expected_status" json:"expected_status"`
	Interval       int        `db:"check_interval" json:"check_interval"`
	Type           string     `db:"check_type" json:"type"`
	Method         string     `db:"method" json:"method"`
	PostData       string     `db:"post_data" json:"post_data"`
	Port           int        `db:"port" json:"port"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	Online         bool       `json:"online"`
	Latency        float64    `json:"latency"`
	Online24Hours  float32    `json:"24_hours_online"`
	AvgResponse    string     `json:"avg_response"`
	TotalUptime    string     `json:"uptime"`
	OrderId        int64      `json:"order_id"`
	Failures       []*Failure `json:"failures"`
	Checkins       []*Checkin `json:"checkins"`
	runRoutine     bool
	LastResponse   string
	LastStatusCode int
	LastOnline     time.Time
}

func serviceCol() db.Collection {
	return dbSession.Collection("services")
}

func SelectService(id int64) *Service {
	for _, s := range services {
		if s.Id == id {
			return s
		}
	}
	return nil
}

func SelectAllServices() ([]*Service, error) {
	var srvcs []*Service
	col := serviceCol().Find()
	err := col.All(&srvcs)
	for _, s := range srvcs {
		s.Checkins = s.SelectAllCheckins()
		s.Failures = s.SelectAllFailures()
	}
	return srvcs, err
}

func (s *Service) AvgTime() float64 {
	total, _ := s.TotalHits()
	if total == 0 {
		return float64(0)
	}
	sum, _ := s.Sum()
	avg := sum / float64(total) * 100
	amount := fmt.Sprintf("%0.0f", avg*10)
	val, _ := strconv.ParseFloat(amount, 10)
	return val
}

func (s *Service) Online24() float32 {
	total, _ := s.TotalHits()
	failed, _ := s.TotalFailures24Hours()
	if failed == 0 {
		s.Online24Hours = 100.00
		return s.Online24Hours
	}
	if total == 0 {
		s.Online24Hours = 0
		return s.Online24Hours
	}
	avg := float64(failed) / float64(total) * 100
	avg = 100 - avg
	if avg < 0 {
		avg = 0
	}
	amount, _ := strconv.ParseFloat(fmt.Sprintf("%0.2f", avg), 10)
	s.Online24Hours = float32(amount)
	return s.Online24Hours
}

type GraphJson struct {
	X string  `json:"x"`
	Y float64 `json:"y"`
}

type DateScan struct {
	CreatedAt time.Time `json:"x"`
	Value     int64     `json:"y"`
}

func (s *Service) GraphData() string {
	var d []DateScan
	since := time.Now().Add(time.Hour*-24 + time.Minute*0 + time.Second*0)
	sql := fmt.Sprintf("SELECT date_trunc('minute', created_at), AVG(latency)*1000 AS value FROM hits WHERE service=%v AND created_at > '%v' GROUP BY 1 ORDER BY date_trunc ASC;", s.Id, since.Format(time.RFC3339))
	dated, err := dbSession.Query(db.Raw(sql))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for dated.Next() {
		var gd DateScan
		var ff float64
		dated.Scan(&gd.CreatedAt, &ff)
		gd.Value = int64(ff)
		d = append(d, gd)
	}
	data, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(data)
}

func (s *Service) AvgUptime() string {
	failed, _ := s.TotalFailures()
	total, _ := s.TotalHits()
	if failed == 0 {
		s.TotalUptime = "100"
		return s.TotalUptime
	}
	if total == 0 {
		s.TotalUptime = "0"
		return s.TotalUptime
	}
	percent := float64(failed) / float64(total) * 100
	percent = 100 - percent
	if percent < 0 {
		percent = 0
	}
	s.TotalUptime = fmt.Sprintf("%0.2f", percent)
	if s.TotalUptime == "100.00" {
		s.TotalUptime = "100"
	}
	return s.TotalUptime
}

func (u *Service) RemoveArray() []*Service {
	var srvcs []*Service
	for _, s := range services {
		if s.Id != u.Id {
			srvcs = append(srvcs, s)
		}
	}
	services = srvcs
	return srvcs
}

func (u *Service) Delete() error {
	res := serviceCol().Find("id", u.Id)
	err := res.Delete()
	u.RemoveArray()
	OnDeletedService(u)
	return err
}

func (u *Service) Update() {
	OnUpdateService(u)
}

func (u *Service) Create() (int64, error) {
	u.CreatedAt = time.Now()
	uuid, err := serviceCol().Insert(u)
	if uuid == nil {
		return 0, err
	}
	u.Id = uuid.(int64)
	services = append(services, u)
	go u.CheckQueue()
	OnNewService(u)
	return uuid.(int64), err
}

func CountOnline() int {
	amount := 0
	for _, v := range services {
		if v.Online {
			amount++
		}
	}
	return amount
}
