package main

import (
	"github.com/ararog/timeago"
	"github.com/hunterlong/statup/log"
	"time"
)

type Failure struct {
	Id        int       `db:"id,omitempty"`
	Issue     string    `db:"issue"`
	Method    string    `db:"method"`
	Service   int64     `db:"service"`
	CreatedAt time.Time `db:"created_at"`
}

func (s *Service) CreateFailure(data FailureData) (int64, error) {
	fail := &Failure{
		Issue:     data.Issue,
		Service:   s.Id,
		CreatedAt: time.Now(),
	}
	s.Failures = append(s.Failures, fail)
	col := dbSession.Collection("failures")
	uuid, err := col.Insert(fail)
	if uuid == nil {
		return 0, err
	}
	return uuid.(int64), err
}

func (s *Service) SelectAllFailures() []*Failure {
	var fails []*Failure
	col := dbSession.Collection("failures").Find("service", s.Id).OrderBy("-id")
	col.All(&fails)
	return fails
}

func (u *Service) DeleteFailures() {
	var fails []*Failure
	col := dbSession.Collection("failures")
	col.Find("service", u.Id).All(&fails)
	for _, fail := range fails {
		fail.Delete()
	}
}

func (s *Service) LimitedFailures() []*Failure {
	var fails []*Failure
	col := dbSession.Collection("failures").Find("service", s.Id).OrderBy("-id").Limit(10)
	col.All(&fails)
	return fails
}

func reverseFailures(input []*Failure) []*Failure {
	if len(input) == 0 {
		return input
	}
	return append(reverseFailures(input[1:]), input[0])
}

func (f *Failure) Ago() string {
	got, _ := timeago.TimeAgoWithTime(time.Now(), f.CreatedAt)
	return got
}

func (f *Failure) Delete() error {
	col := dbSession.Collection("failures").Find("id", f.Id)
	return col.Delete()
}

func CountFailures() uint64 {
	col := dbSession.Collection("failures").Find()
	amount, err := col.Count()
	if err != nil {
		log.Send(2, err)
		return 0
	}
	return amount
}

func (s *Service) TotalFailures() (uint64, error) {
	col := dbSession.Collection("failures").Find("service", s.Id)
	amount, err := col.Count()
	return amount, err
}

func (s *Service) TotalFailures24Hours() (uint64, error) {
	col := dbSession.Collection("failures").Find("service", s.Id)
	amount, err := col.Count()
	return amount, err
}
