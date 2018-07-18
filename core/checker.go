package core

import (
	"bytes"
	"fmt"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

type FailureData types.FailureData

func CheckServices() {
	CoreApp.Services, _ = SelectAllServices()
	utils.Log(1, fmt.Sprintf("Starting monitoring process for %v Services", len(CoreApp.Services)))
	for _, ser := range CoreApp.Services {
		s := ser.ToService()
		//go obj.StartCheckins()
		s.StopRoutine = make(chan struct{})
		go CheckQueue(s)
	}
}

func CheckQueue(s *types.Service) {
	for {
		select {
		case <-s.StopRoutine:
			return
		default:
			ServiceCheck(s)
		}
		time.Sleep(time.Duration(s.Interval) * time.Second)
	}
}

func DNSCheck(s *types.Service) (float64, error) {
	t1 := time.Now()
	url, err := url.Parse(s.Domain)
	if err != nil {
		return 0, err
	}
	_, err = net.LookupIP(url.Host)
	if err != nil {
		return 0, err
	}
	t2 := time.Now()
	subTime := t2.Sub(t1).Seconds()
	return subTime, err
}

func ServiceTCPCheck(s *types.Service) *types.Service {
	t1 := time.Now()
	domain := fmt.Sprintf("%v", s.Domain)
	if s.Port != 0 {
		domain = fmt.Sprintf("%v:%v", s.Domain, s.Port)
	}
	conn, err := net.DialTimeout("tcp", domain, time.Duration(s.Timeout)*time.Second)
	if err != nil {
		RecordFailure(s, fmt.Sprintf("TCP Dial Error %v", err))
		return s
	}
	if err := conn.Close(); err != nil {
		RecordFailure(s, fmt.Sprintf("TCP Socket Close Error %v", err))
		return s
	}
	t2 := time.Now()
	s.Latency = t2.Sub(t1).Seconds()
	s.LastResponse = ""
	RecordSuccess(s)
	return s
}

func ServiceCheck(s *types.Service) *types.Service {
	switch s.Type {
	case "http":
		ServiceHTTPCheck(s)
	case "tcp":
		ServiceTCPCheck(s)
	}
	return s
}

func ServiceHTTPCheck(s *types.Service) *types.Service {
	dnsLookup, err := DNSCheck(s)
	if err != nil {
		RecordFailure(s, fmt.Sprintf("Could not get IP address for domain %v, %v", s.Domain, err))
		return s
	}
	s.DnsLookup = dnsLookup
	t1 := time.Now()
	timeout := time.Duration(s.Timeout)
	client := http.Client{
		Timeout: timeout * time.Second,
	}

	var response *http.Response
	if s.Method == "POST" {
		response, err = client.Post(s.Domain, "application/json", bytes.NewBuffer([]byte(s.PostData)))
	} else {
		response, err = client.Get(s.Domain)
	}
	if err != nil {
		RecordFailure(s, fmt.Sprintf("HTTP Error %v", err))
		return s
	}
	response.Header.Set("User-Agent", "StatupMonitor")
	t2 := time.Now()
	s.Latency = t2.Sub(t1).Seconds()
	if err != nil {
		RecordFailure(s, fmt.Sprintf("HTTP Error %v", err))
		return s
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		utils.Log(2, err)
	}
	if s.Expected != "" {
		match, err := regexp.MatchString(s.Expected, string(contents))
		if err != nil {
			utils.Log(2, err)
		}
		if !match {
			s.LastResponse = string(contents)
			s.LastStatusCode = response.StatusCode
			RecordFailure(s, fmt.Sprintf("HTTP Response Body did not match '%v'", s.Expected))
			return s
		}
	}
	if s.ExpectedStatus != response.StatusCode {
		s.LastResponse = string(contents)
		s.LastStatusCode = response.StatusCode
		RecordFailure(s, fmt.Sprintf("HTTP Status Code %v did not match %v", response.StatusCode, s.ExpectedStatus))
		return s
	}
	s.LastResponse = string(contents)
	s.LastStatusCode = response.StatusCode
	s.Online = true
	RecordSuccess(s)
	return s
}

type HitData struct {
	Latency float64
}

func RecordSuccess(s *types.Service) {
	s.Online = true
	s.LastOnline = time.Now()
	data := HitData{
		Latency: s.Latency,
	}
	utils.Log(1, fmt.Sprintf("Service %v Successful: %0.2f ms", s.Name, data.Latency*1000))
	CreateServiceHit(s, data)
	OnSuccess(s)
}

func RecordFailure(s *types.Service, issue string) {
	s.Online = false
	data := FailureData{
		Issue: issue,
	}
	utils.Log(2, fmt.Sprintf("Service %v Failing: %v", s.Name, issue))
	CreateServiceFailure(s, data)
	//SendFailureEmail(s)
	OnFailure(s, data)
}
