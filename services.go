package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

var (
	services []*Service
)

type Service struct {
	Id             int64
	Name           string
	Domain         string
	Expected       string
	ExpectedStatus int
	Interval       int
	Method         string
	Port           int
	CreatedAt      time.Time
	Data           string
	Online         bool
	Latency        float64
	Online24Hours  float64
	AvgResponse    string
	TotalUptime    float64
}

func SelectService(id string) Service {
	var tk Service
	rows, err := db.Query("SELECT * FROM services WHERE id=$1", id)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&tk.Id, &tk.Name, &tk.Domain, &tk.Method, &tk.Port, &tk.Expected, &tk.ExpectedStatus, &tk.Interval, &tk.CreatedAt)
		if err != nil {
			panic(err)
		}
	}
	return tk
}

func SelectAllServices() []*Service {
	var tks []*Service
	rows, err := db.Query("SELECT * FROM services ORDER BY id ASC")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var tk Service
		err = rows.Scan(&tk.Id, &tk.Name, &tk.Domain, &tk.Method, &tk.Port, &tk.Expected, &tk.ExpectedStatus, &tk.Interval, &tk.CreatedAt)
		if err != nil {
			panic(err)
		}
		tk.FormatData()
		tks = append(tks, &tk)
	}
	return tks
}

func (s *Service) FormatData() *Service {
	s.GraphData()
	s.AvgUptime()
	s.Online24()
	s.AvgTime()
	return s
}

func (s *Service) AvgTime() float64 {
	total := s.TotalHits()
	sum := s.Sum()
	avg := sum / float64(total) * 100
	s.AvgResponse = fmt.Sprintf("%0.0f", avg*10)
	return avg
}

func (s *Service) Online24() float64 {
	total := s.TotalHits()
	failed := s.TotalFailures24Hours()
	if failed == 0 {
		s.Online24Hours = 100.00
		return s.Online24Hours
	}
	if total == 0 {
		s.Online24Hours = 0
		return s.Online24Hours
	}
	avg := float64(failed) / float64(total) * 100
	s.Online24Hours = avg
	return avg
}

type GraphJson struct {
	X string  `json:"x"`
	Y float64 `json:"y"`
}

func (s *Service) GraphData() string {
	hits := SelectAllHits(s.Id)
	var d []*GraphJson
	for _, h := range hits {
		val := h.CreatedAt
		o := &GraphJson{
			X: val.String(),
			Y: h.Value * 1000,
		}
		d = append(d, o)
	}
	data, _ := json.Marshal(d)
	s.Data = string(data)
	return s.Data
}

func (s *Service) AvgUptime() float64 {
	failed := s.TotalFailures()
	total := s.TotalHits()
	if failed == 0 {
		s.TotalUptime = 100.00
		return s.TotalUptime
	}
	if total == 0 {
		s.TotalUptime = 0
		return s.TotalUptime
	}
	percent := float64(failed) / float64(total) * 100
	s.TotalUptime = percent
	return percent
}

func (u *Service) Create() int {
	var lastInsertId int
	err := db.QueryRow("INSERT INTO services(name, domain, method, port, expected, expected_status, interval, created_at) VALUES($1,$2,$3,$4,$5,$6,$7,NOW()) returning id;", u.Name, u.Domain, u.Method, u.Port, u.Expected, u.ExpectedStatus, u.Interval).Scan(&lastInsertId)
	if err != nil {
		panic(err)
	}
	services = SelectAllServices()
	return lastInsertId
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
