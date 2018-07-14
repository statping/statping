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
		obj := s
		//go obj.StartCheckins()
		obj.StopRoutine = make(chan struct{})
		go CheckQueue(obj)
	}
}

func CheckQueue(s *types.Service) {
	for {
		select {
		case <-s.StopRoutine:
			return
		default:
			ServiceCheck(s)
			if s.Interval < 1 {
				s.Interval = 1
			}
			msg := fmt.Sprintf("Service: %v | Online: %v | Latency: %0.0fms", s.Name, s.Online, (s.Latency * 1000))
			utils.Log(1, msg)
			time.Sleep(time.Duration(s.Interval) * time.Second)
		}
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

func ServiceCheck(s *types.Service) *types.Service {
	dnsLookup, err := DNSCheck(s)
	if err != nil {
		RecordFailure(s, fmt.Sprintf("Could not get IP address for domain %v, %v", s.Domain, err))
		return s
	}
	s.DnsLookup = dnsLookup
	t1 := time.Now()
	client := http.Client{
		Timeout: 30 * time.Second,
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
	RecordSuccess(s, response)
	return s
}

type HitData struct {
	Latency float64
}

func RecordSuccess(s *types.Service, response *http.Response) {
	s.Online = true
	s.LastOnline = time.Now()
	data := HitData{
		Latency: s.Latency,
	}
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
