package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"math/rand"
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
	Port           int        `db:"port" json:"port"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	Online         bool       `json:"online"`
	Latency        float64    `json:"latency"`
	Online24Hours  float32    `json:"24_hours_online"`
	AvgResponse    string     `json:"avg_response"`
	TotalUptime    string     `json:"uptime"`
	Failures       []*Failure `json:"failures"`
	Checkins       []*Checkin `json:"checkins"`
}

func serviceCol() db.Collection {
	return dbSession.Collection("services")
}

func SelectService(id int64) (*Service, error) {
	var service *Service
	res := serviceCol().Find("id", id)
	err := res.One(&service)
	service.Checkins = service.SelectAllCheckins()
	return service, err
}

func SelectAllServices() ([]*Service, error) {
	var services []*Service
	col := serviceCol().Find()
	err := col.All(&services)
	for _, s := range services {
		s.Checkins = s.SelectAllCheckins()
	}
	return services, err
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

func (s *Service) GraphData() string {
	var d []GraphJson
	hits, _ := s.LimitedHits()
	for _, h := range hits {
		val := h.CreatedAt
		o := GraphJson{
			X: val.String(),
			Y: h.Latency * 1000,
		}
		d = append(d, o)
	}
	data, _ := json.Marshal(d)
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

func (u *Service) Delete() error {
	res := serviceCol().Find("id", u.Id)
	err := res.Delete()
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

func NewSHA1Hash(n ...int) string {
	noRandomCharacters := 32
	if len(n) > 0 {
		noRandomCharacters = n[0]
	}
	randString := RandomString(noRandomCharacters)
	hash := sha1.New()
	hash.Write([]byte(randString))
	bs := hash.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString generates a random string of n length
func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	return string(b)
}
