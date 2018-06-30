package core

import (
	"fmt"
	"github.com/ararog/timeago"
	"github.com/hunterlong/statup/utils"
	"strings"
	"time"
)

func (s *Service) CreateFailure(data FailureData) (int64, error) {
	fail := &Failure{
		Issue:     data.Issue,
		Service:   s.Id,
		CreatedAt: time.Now(),
	}
	s.Failures = append(s.Failures, fail)
	col := DbSession.Collection("failures")
	uuid, err := col.Insert(fail)
	if uuid == nil {
		return 0, err
	}
	return uuid.(int64), err
}

func (s *Service) SelectAllFailures() []*Failure {
	var fails []*Failure
	col := DbSession.Collection("failures").Find("service", s.Id).OrderBy("-id")
	col.All(&fails)
	return fails
}

func (u *Service) DeleteFailures() {
	var fails []*Failure
	col := DbSession.Collection("failures")
	col.Find("service", u.Id).All(&fails)
	for _, fail := range fails {
		fail.Delete()
	}
}

func (s *Service) LimitedFailures() []*Failure {
	var fails []*Failure
	col := DbSession.Collection("failures").Find("service", s.Id).OrderBy("-id").Limit(10)
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
	col := DbSession.Collection("failures").Find("id", f.Id)
	return col.Delete()
}

func CountFailures() uint64 {
	col := DbSession.Collection("failures").Find()
	amount, err := col.Count()
	if err != nil {
		utils.Log(2, err)
		return 0
	}
	return amount
}

func (s *Service) TotalFailures() (uint64, error) {
	col := DbSession.Collection("failures").Find("service", s.Id)
	amount, err := col.Count()
	return amount, err
}

func (s *Service) TotalFailures24Hours() (uint64, error) {
	col := DbSession.Collection("failures").Find("service", s.Id)
	amount, err := col.Count()
	return amount, err
}

func (f *Failure) ParseError() string {
	err := strings.Contains(f.Issue, "operation timed out")
	if err {
		return fmt.Sprintf("HTTP Request Timed Out")
	}
	err = strings.Contains(f.Issue, "x509: certificate is valid")
	if err {
		return fmt.Sprintf("SSL Certificate invalid")
	}
	err = strings.Contains(f.Issue, "no such host")
	if err {
		return fmt.Sprintf("Domain is offline or not found")
	}
	err = strings.Contains(f.Issue, "HTTP Status Code")
	if err {
		return fmt.Sprintf("Incorrect HTTP Status Code")
	}
	err = strings.Contains(f.Issue, "connection refused")
	if err {
		return fmt.Sprintf("Connection Failed")
	}
	return f.Issue
}
