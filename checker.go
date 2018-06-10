package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func CheckServices() {
	services := SelectAllServices()
	for _, v := range services {
		obj := v
		go obj.CheckQueue()
	}
}

func (s *Service) CheckQueue() {
	time.Sleep(time.Duration(s.Interval) * time.Second)
	s.Check()
	fmt.Printf("Service: %v | Online: %v | Latency: %v\n", s.Name, s.Online, s.Latency)
	s.CheckQueue()
}

func (s *Service) Check() {
	t1 := time.Now()
	response, err := http.Get(s.Domain)
	t2 := time.Now()
	s.Latency = t2.Sub(t1).Seconds()
	if err != nil {
		s.Failure(response, fmt.Sprintf("HTTP Error %v", err))
		return
	}

	if s.Expected != "" {
		contents, _ := ioutil.ReadAll(response.Body)
		match, _ := regexp.MatchString(s.Expected, string(contents))
		if !match {
			s.Failure(response, fmt.Sprintf("HTTP Response Body did not match '%v'", s.Expected))
			return
		}
	}
	if s.ExpectedStatus != response.StatusCode {
		s.Failure(response, fmt.Sprintf("HTTP Status Code %v did not match %v", response.StatusCode, s.ExpectedStatus))
		return
	}
	s.Record(response)
}

func (s *Service) Record(response *http.Response) {
	defer response.Body.Close()
	s.Online = true
	db.QueryRow("INSERT INTO hits(service,latency,created_at) VALUES($1,$2,NOW()) returning id;", s.Id, s.Latency).Scan()
}

func (s *Service) Failure(response *http.Response, issue string) {
	db.QueryRow("INSERT INTO failures(issue,service,created_at) VALUES($1,$2,NOW()) returning id;", issue, s.Id).Scan()
	s.Record(response)
}
