// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"bytes"
	"fmt"
	"github.com/hunterlong/statup/notifiers"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func CheckServices() {
	utils.Log(1, fmt.Sprintf("Starting monitoring process for %v Services", len(CoreApp.Services)))
	for _, ser := range CoreApp.Services {
		//go obj.StartCheckins()
		go ser.CheckQueue(true)
	}
}

func (s *Service) CheckQueue(record bool) {
	s.Checkpoint = time.Now()
	s.SleepDuration = time.Duration((time.Duration(s.Id) * 100) * time.Millisecond)
CheckLoop:
	for {
		select {
		case <-s.Running:
			utils.Log(1, fmt.Sprintf("Stopping service: %v", s.Name))
			break CheckLoop
		case <-time.After(s.SleepDuration):
			s.Check(record)
			s.Checkpoint = s.Checkpoint.Add(s.duration())
			sleep := s.Checkpoint.Sub(time.Now())
			if !s.Online {
				s.SleepDuration = s.duration()
			} else {
				s.SleepDuration = sleep
			}
		}
		continue
	}
}

func (s *Service) duration() time.Duration {
	var amount time.Duration
	if s.Interval >= 10000 {
		amount = time.Duration(s.Interval) * time.Microsecond
	} else {
		amount = time.Duration(s.Interval) * time.Second
	}
	return amount
}

func (s *Service) dnsCheck() (float64, error) {
	t1 := time.Now()
	domain := s.Domain
	hasPort, _ := regexp.MatchString(`\:([0-9]+)`, domain)
	if hasPort {
		splitDomain := strings.Split(s.Domain, ":")
		domain = splitDomain[len(splitDomain)-2]
	}
	url, err := url.Parse(domain)
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

func (s *Service) checkTcp(record bool) *Service {
	t1 := time.Now()
	domain := fmt.Sprintf("%v", s.Domain)
	if s.Port != 0 {
		domain = fmt.Sprintf("%v:%v", s.Domain, s.Port)
	}
	conn, err := net.DialTimeout("tcp", domain, time.Duration(s.Timeout)*time.Second)
	if err != nil {
		if record {
			RecordFailure(s, fmt.Sprintf("TCP Dial Error %v", err))
		}
		return s
	}
	if err := conn.Close(); err != nil {
		if record {
			RecordFailure(s, fmt.Sprintf("TCP Socket Close Error %v", err))
		}
		return s
	}
	t2 := time.Now()
	s.Latency = t2.Sub(t1).Seconds()
	s.LastResponse = ""
	if record {
		RecordSuccess(s)
	}
	return s
}

func (s *Service) Check(record bool) {
	switch s.Type {
	case "http":
		s.checkHttp(record)
	case "tcp":
		s.checkTcp(record)
	}
}

func (s *Service) checkHttp(record bool) *Service {
	dnsLookup, err := s.dnsCheck()
	if err != nil {
		if record {
			RecordFailure(s, fmt.Sprintf("Could not get IP address for domain %v, %v", s.Domain, err))
		}
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
		if record {
			RecordFailure(s, fmt.Sprintf("HTTP Error %v", err))
		}
		return s
	}
	response.Header.Set("Connection", "close")
	response.Header.Set("User-Agent", "StatupMonitor")
	t2 := time.Now()
	s.Latency = t2.Sub(t1).Seconds()
	if err != nil {
		if record {
			RecordFailure(s, fmt.Sprintf("HTTP Error %v", err))
		}
		return s
	}
	defer response.Body.Close()

	if s.Expected != "" {
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			utils.Log(2, err)
		}
		match, err := regexp.MatchString(s.Expected, string(contents))
		if err != nil {
			utils.Log(2, err)
		}
		if !match {
			s.LastResponse = string(contents)
			s.LastStatusCode = response.StatusCode
			if record {
				RecordFailure(s, fmt.Sprintf("HTTP Response Body did not match '%v'", s.Expected))
			}
			return s
		}
	}
	if s.ExpectedStatus != response.StatusCode {
		//s.LastResponse = string(contents)
		s.LastStatusCode = response.StatusCode
		if record {
			RecordFailure(s, fmt.Sprintf("HTTP Status Code %v did not match %v", response.StatusCode, s.ExpectedStatus))
		}
		return s
	}
	s.LastStatusCode = response.StatusCode
	s.Online = true
	if record {
		RecordSuccess(s)
	}
	return s
}

type HitData struct {
	Latency float64
}

func RecordSuccess(s *Service) {
	s.Online = true
	s.LastOnline = time.Now()
	hit := &types.Hit{
		Service:   s.Id,
		Latency:   s.Latency,
		CreatedAt: time.Now(),
	}
	utils.Log(1, fmt.Sprintf("Service %v Successful: %0.2f ms", s.Name, hit.Latency*1000))
	s.CreateHit(hit)
	notifiers.OnSuccess(s.Service)
}

func RecordFailure(s *Service, issue string) {
	s.Online = false
	fail := &types.Failure{
		Service:   s.Id,
		Issue:     issue,
		CreatedAt: time.Now(),
	}
	utils.Log(2, fmt.Sprintf("Service %v Failing: %v", s.Name, issue))
	s.CreateFailure(fail)
	notifiers.OnFailure(s.Service, fail)
}
