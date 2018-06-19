package main

import (
	"time"
)

type Failure struct {
	Id        int       `db:"id,omitempty"`
	Issue     string    `db:"issue"`
	Service   int64     `db:"service"`
	CreatedAt time.Time `db:"created_at"`
	Ago       string
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

func (s *Service) SelectAllFailures() ([]*Failure, error) {
	var fails []*Failure
	col := dbSession.Collection("failures").Find("session", s.Id)
	err := col.All(&fails)
	return fails, err
}

func (s *Service) LimitedFailures() ([]*Failure, error) {
	var fails []*Failure
	col := dbSession.Collection("failures").Find("session", s.Id).Limit(10)
	err := col.All(&fails)
	return fails, err
}

func CountFailures() (uint64, error) {
	col := dbSession.Collection("failures").Find()
	amount, err := col.Count()
	return amount, err
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
