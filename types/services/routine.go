package services

import (
	"bytes"
	"fmt"
	"github.com/hunterlong/statping/types/failures"
	"github.com/hunterlong/statping/types/hits"
	"github.com/hunterlong/statping/utils"
	"github.com/tatsushid/go-fastping"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// checkServices will start the checking go routine for each service
func CheckServices() {
	log.Infoln(fmt.Sprintf("Starting monitoring process for %v Services", len(allServices)))
	for _, s := range allServices {
		//go CheckinRoutine()
		time.Sleep(250 * time.Millisecond)
		go ServiceCheckQueue(s, true)
	}
}

// CheckQueue is the main go routine for checking a service
func ServiceCheckQueue(s *Service, record bool) {
	s.Start()
	s.Checkpoint = time.Now()
	s.SleepDuration = (time.Duration(s.Id) * 100) * time.Millisecond

CheckLoop:
	for {
		select {
		case <-s.Running:
			log.Infoln(fmt.Sprintf("Stopping service: %v", s.Name))
			break CheckLoop
		case <-time.After(s.SleepDuration):
			s.CheckService(record)
			s.Checkpoint = s.Checkpoint.Add(s.Duration())
			sleep := s.Checkpoint.Sub(time.Now())
			if !s.Online {
				s.SleepDuration = s.Duration()
			} else {
				s.SleepDuration = sleep
			}
		}
		continue
	}
}

func parseHost(s *Service) string {
	if s.Type == "tcp" || s.Type == "udp" {
		return s.Domain
	} else {
		u, err := url.Parse(s.Domain)
		if err != nil {
			return s.Domain
		}
		return strings.Split(u.Host, ":")[0]
	}
}

// dnsCheck will check the domain name and return a float64 for the amount of time the DNS check took
func dnsCheck(s *Service) (float64, error) {
	var err error
	t1 := time.Now()
	host := parseHost(s)
	if s.Type == "tcp" {
		_, err = net.LookupHost(host)
	} else {
		_, err = net.LookupIP(host)
	}
	if err != nil {
		return 0, err
	}
	t2 := time.Now()
	subTime := t2.Sub(t1).Seconds()
	return subTime, err
}

func isIPv6(address string) bool {
	return strings.Count(address, ":") >= 2
}

// checkIcmp will send a ICMP ping packet to the service
func CheckIcmp(s *Service, record bool) *Service {
	defer s.updateLastCheck()

	p := fastping.NewPinger()
	resolveIP := "ip4:icmp"
	if isIPv6(s.Domain) {
		resolveIP = "ip6:icmp"
	}
	ra, err := net.ResolveIPAddr(resolveIP, s.Domain)
	if err != nil {
		recordFailure(s, fmt.Sprintf("Could not send ICMP to service %v, %v", s.Domain, err))
		return s
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		s.Latency = rtt.Seconds()
		recordSuccess(s)
	}
	err = p.Run()
	if err != nil {
		recordFailure(s, fmt.Sprintf("Issue running ICMP to service %v, %v", s.Domain, err))
		return s
	}
	s.LastResponse = ""
	return s
}

// checkTcp will check a TCP service
func CheckTcp(s *Service, record bool) *Service {
	defer s.updateLastCheck()

	dnsLookup, err := dnsCheck(s)
	if err != nil {
		if record {
			recordFailure(s, fmt.Sprintf("Could not get IP address for TCP service %v, %v", s.Domain, err))
		}
		return s
	}
	s.PingTime = dnsLookup
	t1 := time.Now()
	domain := fmt.Sprintf("%v", s.Domain)
	if s.Port != 0 {
		domain = fmt.Sprintf("%v:%v", s.Domain, s.Port)
		if isIPv6(s.Domain) {
			domain = fmt.Sprintf("[%v]:%v", s.Domain, s.Port)
		}
	}
	conn, err := net.DialTimeout(s.Type, domain, time.Duration(s.Timeout)*time.Second)
	if err != nil {
		if record {
			recordFailure(s, fmt.Sprintf("Dial Error %v", err))
		}
		return s
	}
	if err := conn.Close(); err != nil {
		if record {
			recordFailure(s, fmt.Sprintf("%v Socket Close Error %v", strings.ToUpper(s.Type), err))
		}
		return s
	}
	t2 := time.Now()
	s.Latency = t2.Sub(t1).Seconds()
	s.LastResponse = ""
	if record {
		recordSuccess(s)
	}
	return s
}

func (s *Service) updateLastCheck() {
	s.LastCheck = time.Now()
}

// checkHttp will check a HTTP service
func CheckHttp(s *Service, record bool) *Service {
	defer s.updateLastCheck()

	dnsLookup, err := dnsCheck(s)
	if err != nil {
		if record {
			recordFailure(s, fmt.Sprintf("Could not get IP address for domain %v, %v", s.Domain, err))
		}
		return s
	}
	s.PingTime = dnsLookup
	t1 := time.Now()

	timeout := time.Duration(s.Timeout) * time.Second
	var content []byte
	var res *http.Response

	var headers []string
	if s.Headers.Valid {
		headers = strings.Split(s.Headers.String, ",")
	} else {
		headers = nil
	}

	if s.Method == "POST" {
		content, res, err = utils.HttpRequest(s.Domain, s.Method, "application/json", headers, bytes.NewBuffer([]byte(s.PostData.String)), timeout, s.VerifySSL.Bool)
	} else {
		content, res, err = utils.HttpRequest(s.Domain, s.Method, nil, headers, nil, timeout, s.VerifySSL.Bool)
	}
	if err != nil {
		if record {
			recordFailure(s, fmt.Sprintf("HTTP Error %v", err))
		}
		return s
	}
	t2 := time.Now()
	s.Latency = t2.Sub(t1).Seconds()
	s.LastResponse = string(content)
	s.LastStatusCode = res.StatusCode

	if s.Expected.String != "" {
		match, err := regexp.MatchString(s.Expected.String, string(content))
		if err != nil {
			log.Warnln(fmt.Sprintf("Service %v expected: %v to match %v", s.Name, string(content), s.Expected.String))
		}
		if !match {
			if record {
				recordFailure(s, fmt.Sprintf("HTTP Response Body did not match '%v'", s.Expected))
			}
			return s
		}
	}
	if s.ExpectedStatus != res.StatusCode {
		if record {
			recordFailure(s, fmt.Sprintf("HTTP Status Code %v did not match %v", res.StatusCode, s.ExpectedStatus))
		}
		return s
	}
	if record {
		recordSuccess(s)
	}
	return s
}

// recordSuccess will create a new 'hit' record in the database for a successful/online service
func recordSuccess(s *Service) {
	s.LastOnline = time.Now().UTC()
	s.Online = true
	hit := &hits.Hit{
		Service:   s.Id,
		Latency:   s.Latency,
		PingTime:  s.PingTime,
		CreatedAt: time.Now().UTC(),
	}
	if err := hit.Create(); err != nil {
		log.Error(err)
	}
	log.WithFields(utils.ToFields(hit, s)).Infoln(
		fmt.Sprintf("Service #%d '%v' Successful Response: %0.2f ms | Lookup in: %0.2f ms | Online: %v | Interval: %d seconds", s.Id, s.Name, hit.Latency*1000, hit.PingTime*1000, s.Online, s.Interval))
	s.LastLookupTime = int64(hit.PingTime * 1000)
	s.LastLatency = int64(hit.Latency * 1000)
	//notifier.OnSuccess(s)
	s.SuccessNotified = true
}

// recordFailure will create a new 'Failure' record in the database for a offline service
func recordFailure(s *Service, issue string) {
	s.LastOffline = time.Now().UTC()

	fail := &failures.Failure{
		Service:   s.Id,
		Issue:     issue,
		PingTime:  s.PingTime,
		CreatedAt: time.Now().UTC(),
		ErrorCode: s.LastStatusCode,
	}
	log.WithFields(utils.ToFields(fail, s)).
		Warnln(fmt.Sprintf("Service %v Failing: %v | Lookup in: %0.2f ms", s.Name, issue, fail.PingTime*1000))

	if err := fail.Create(); err != nil {
		log.Error(err)
	}
	s.Online = false
	s.SuccessNotified = false
	s.DownText = s.DowntimeText()
	//notifier.OnFailure(s, fail)
}

// Check will run checkHttp for HTTP services and checkTcp for TCP services
// if record param is set to true, it will add a record into the database.
func (s *Service) CheckService(record bool) {
	switch s.Type {
	case "http":
		CheckHttp(s, record)
	case "tcp", "udp":
		CheckTcp(s, record)
	case "icmp":
		CheckIcmp(s, record)
	}
}
