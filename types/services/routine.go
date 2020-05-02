package services

import (
	"bytes"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/hits"
	"github.com/statping/statping/utils"
	"github.com/tatsushid/go-fastping"
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
	s.Checkpoint = utils.Now()
	s.SleepDuration = (time.Duration(s.Id) * 100) * time.Millisecond

CheckLoop:
	for {
		select {
		case <-s.Running:
			log.Infoln(fmt.Sprintf("Stopping service: %v", s.Name))
			break CheckLoop
		case <-time.After(s.SleepDuration):
			s.CheckService(record)
			s.UpdateStats()
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
	if s.Type == "tcp" || s.Type == "udp" || s.Type == "grpc" {
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
func dnsCheck(s *Service) (int64, error) {
	var err error
	t1 := utils.Now()
	host := parseHost(s)
	if s.Type == "tcp" || s.Type == "udp" || s.Type == "grpc" {
		_, err = net.LookupHost(host)
	} else {
		_, err = net.LookupIP(host)
	}
	if err != nil {
		return 0, err
	}
	t2 := utils.Now()
	subTime := t2.Sub(t1).Microseconds()
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
		s.Latency = rtt.Microseconds()
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

// CheckGrpc will check a gRPC service
func CheckGrpc(s *Service, record bool) *Service {
	defer s.updateLastCheck()

	dnsLookup, err := dnsCheck(s)
	if err != nil {
		if record {
			recordFailure(s, fmt.Sprintf("Could not get IP address for GRPC service %v, %v", s.Domain, err))
		}
		return s
	}
	s.PingTime = dnsLookup
	t1 := utils.Now()
	domain := fmt.Sprintf("%v", s.Domain)
	if s.Port != 0 {
		domain = fmt.Sprintf("%v:%v", s.Domain, s.Port)
		if isIPv6(s.Domain) {
			domain = fmt.Sprintf("[%v]:%v", s.Domain, s.Port)
		}
	}
	conn, err := grpc.Dial(domain, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
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
	t2 := utils.Now()
	s.Latency = t2.Sub(t1).Microseconds()
	s.LastResponse = ""
	if record {
		recordSuccess(s)
	}
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
	t1 := utils.Now()
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
	t2 := utils.Now()
	s.Latency = t2.Sub(t1).Microseconds()
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
	t1 := utils.Now()

	timeout := time.Duration(s.Timeout) * time.Second
	var content []byte
	var res *http.Response
	var data *bytes.Buffer
	var headers []string
	contentType := "application/json" // default Content-Type

	if s.Headers.Valid {
		headers = strings.Split(s.Headers.String, ",")
	} else {
		headers = nil
	}

	// check if 'Content-Type' header was defined
	for _, header := range headers {
		if strings.Split(header, "=")[0] == "Content-Type" {
			contentType = strings.Split(header, "=")[1]
			break
		}
	}

	if s.Redirect.Bool {
		headers = append(headers, "Redirect=true")
	}

	if s.PostData.String != "" {
		data = bytes.NewBuffer([]byte(s.PostData.String))
	} else {
		data = bytes.NewBuffer(nil)
	}

	// force set Content-Type to 'application/json' if requests are made
	// with POST method
	if s.Method == "POST" && contentType != "application/json" {
		contentType = "application/json"
	}

	content, res, err = utils.HttpRequest(s.Domain, s.Method, contentType,
		headers, data, timeout, s.VerifySSL.Bool)
	if err != nil {
		if record {
			recordFailure(s, fmt.Sprintf("HTTP Error %v", err))
		}
		return s
	}
	t2 := utils.Now()
	s.Latency = t2.Sub(t1).Microseconds()
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
	s.LastOnline = utils.Now()
	s.Online = true
	hit := &hits.Hit{
		Service:   s.Id,
		Latency:   s.Latency,
		PingTime:  s.PingTime,
		CreatedAt: utils.Now(),
	}
	if err := hit.Create(); err != nil {
		log.Error(err)
	}
	log.WithFields(utils.ToFields(hit, s)).Infoln(
		fmt.Sprintf("Service #%d '%v' Successful Response: %s | Lookup in: %s | Online: %v | Interval: %d seconds", s.Id, s.Name, humanMicro(hit.Latency), humanMicro(hit.PingTime), s.Online, s.Interval))
	s.LastLookupTime = hit.PingTime
	s.LastLatency = hit.Latency
	sendSuccess(s)
	s.SuccessNotified = true
}

func AddNotifier(n ServiceNotifier) {
	notif := n.Select()
	allNotifiers[notif.Method] = n
}

func sendSuccess(s *Service) {
	if !s.AllowNotifications.Bool {
		return
	}
	// dont send notification if server was already previous online
	if s.SuccessNotified {
		return
	}

	for _, n := range allNotifiers {
		notif := n.Select()
		if notif.CanSend() {
			log.Infof("Sending notification to: %s!", notif.Method)
			if err := n.OnSuccess(s); err != nil {
				notif.Logger().Errorln(err)
			}
			s.UserNotified = true
			s.SuccessNotified = true
			//s.UpdateNotify.Bool
		}
	}
	s.notifyAfterCount = 0
}

// recordFailure will create a new 'Failure' record in the database for a offline service
func recordFailure(s *Service, issue string) {
	s.LastOffline = utils.Now()

	fail := &failures.Failure{
		Service:   s.Id,
		Issue:     issue,
		PingTime:  s.PingTime,
		CreatedAt: utils.Now(),
		ErrorCode: s.LastStatusCode,
	}
	log.WithFields(utils.ToFields(fail, s)).
		Warnln(fmt.Sprintf("Service %v Failing: %v | Lookup in: %v", s.Name, issue, humanMicro(fail.PingTime)))

	if err := fail.Create(); err != nil {
		log.Error(err)
	}
	s.Online = false
	s.SuccessNotified = false
	s.DownText = s.DowntimeText()
	sendFailure(s, fail)
}

func sendFailure(s *Service, f *failures.Failure) {
	if !s.AllowNotifications.Bool {
		return
	}

	// ignore failure if user was already notified and
	// they have "continuous notifications" switched off.
	if s.UserNotified && !s.UpdateNotify.Bool {
		return
	}

	if s.notifyAfterCount > s.NotifyAfter {
		for _, n := range allNotifiers {
			notif := n.Select()
			if notif.CanSend() {
				log.Infof("Sending Failure notification to: %s!", notif.Method)
				if err := n.OnFailure(s, f); err != nil {
					notif.Logger().WithField("failure", f.Issue).Errorln(err)
				}
				s.UserNotified = true
				s.SuccessNotified = true
				//s.UpdateNotify.Bool
			}
		}
	}

	s.notifyAfterCount++
}

// Check will run checkHttp for HTTP services and checkTcp for TCP services
// if record param is set to true, it will add a record into the database.
func (s *Service) CheckService(record bool) {
	switch s.Type {
	case "http":
		CheckHttp(s, record)
	case "tcp", "udp":
		CheckTcp(s, record)
	case "grpc":
		CheckGrpc(s, record)
	case "icmp":
		CheckIcmp(s, record)
	}
}
