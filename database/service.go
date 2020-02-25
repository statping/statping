package database

import (
	"fmt"
	"github.com/hunterlong/statping/types"
	"strconv"
	"strings"
	"time"
)

type ServiceObj struct {
	*types.Service
	o *Object

	Servicer
}

type Servicer interface {
	Hits() *HitObj
	Failures() *FailureObj
	AllCheckins() []*CheckinObj
	Model() *types.Service
	Interval() time.Duration
	DowntimeText() string

	Hittable
}

type Hittable interface {
	CreateHit(*types.Hit) (int64, error)
}

func Service(id int64) (*ServiceObj, error) {
	var service types.Service
	query := database.Services().Where("id = ?", id)
	finer := query.Find(&service)
	return &ServiceObj{Service: &service, o: wrapObject(id, &service, query)}, finer.Error()
}

func wrapServices(all []*types.Service, db Database) []*ServiceObj {
	var arr []*ServiceObj
	for _, v := range all {
		arr = append(arr, &ServiceObj{Service: v, o: wrapObject(v.Id, v, db)})
	}
	return arr
}

func Services() []*ServiceObj {
	var services []*types.Service
	db := database.Services().Order("order_id desc")
	db.Find(&services)
	return wrapServices(services, db)
}

func (s *ServiceObj) AllCheckins() []*CheckinObj {
	var checkins []*types.Checkin
	query := database.Checkins().Where("service = ?", s.Id)
	query.Find(&checkins)
	return wrapCheckins(checkins, query)
}

func (s *ServiceObj) DowntimeText() string {
	last := s.Failures().Last(1)
	return parseError(last)
}

// ParseError returns a human readable error for a Failure
func parseError(f *types.Failure) string {
	if f.Method == "checkin" {
		return fmt.Sprintf("Checkin is Offline")
	}
	err := strings.Contains(f.Issue, "connection reset by peer")
	if err {
		return fmt.Sprintf("Connection Reset")
	}
	err = strings.Contains(f.Issue, "operation timed out")
	if err {
		return fmt.Sprintf("HTTP Request Timed Out")
	}
	err = strings.Contains(f.Issue, "x509: certificate is valid")
	if err {
		return fmt.Sprintf("SSL Certificate invalid")
	}
	err = strings.Contains(f.Issue, "Client.Timeout exceeded while awaiting headers")
	if err {
		return fmt.Sprintf("Connection Timed Out")
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
	err = strings.Contains(f.Issue, "can't assign requested address")
	if err {
		return fmt.Sprintf("Unable to Request Address")
	}
	err = strings.Contains(f.Issue, "no route to host")
	if err {
		return fmt.Sprintf("Domain is offline or not found")
	}
	err = strings.Contains(f.Issue, "i/o timeout")
	if err {
		return fmt.Sprintf("Connection Timed Out")
	}
	err = strings.Contains(f.Issue, "Client.Timeout exceeded while reading body")
	if err {
		return fmt.Sprintf("Timed Out on Response Body")
	}
	return f.Issue
}

func (s *ServiceObj) Interval() time.Duration {
	return time.Duration(s.Service.Interval) * time.Second
}

func (s *ServiceObj) Model() *types.Service {
	return s.Service
}

func (s *ServiceObj) Hits() *HitObj {
	fmt.Println("hits")
	query := database.Hits().Where("service = ?", s.Id)
	return &HitObj{wrapObject(s.Id, nil, query)}
}

func (s *ServiceObj) Failures() *FailureObj {
	q := database.Failures().Where("method != 'checkin' AND service = ?", s.Id)
	return &FailureObj{wrapObject(s.Id, nil, q)}
}

func (s *ServiceObj) Group() *GroupObj {
	var group types.Group
	q := database.Groups().Where("id = ?", s.GroupId)
	finder := q.Find(&group)
	if finder.Error() != nil {
		return nil
	}
	return &GroupObj{Group: &group, o: wrapObject(group.Id, &group, q)}
}

func (s *ServiceObj) object() *Object {
	return s.o
}

func (s *ServiceObj) UpdateStats() *types.Stats {
	s.Online24Hours = s.OnlineDaysPercent(1)
	s.Online7Days = s.OnlineDaysPercent(7)
	s.AvgResponse = s.AvgTime()
	s.FailuresLast24Hours = len(s.Failures().Since(time.Now().Add(-time.Hour * 24)))
	return s.Stats
}

// AvgTime will return the average amount of time for a service to response back successfully
func (s *ServiceObj) AvgTime() float64 {
	var sum float64
	database.Hits().
		Select("AVG(latency) as amount").
		Where("service = ?", s.Id).Pluck("amount", &sum).Debug()

	total := s.Hits().Count()

	if total == 0 {
		return 0
	}
	avg := sum / float64(total) * 100
	f, _ := strconv.ParseFloat(fmt.Sprintf("%0.0f", avg*10), 32)
	return f
}

// OnlineDaysPercent returns the service's uptime percent within last 24 hours
func (s *ServiceObj) OnlineDaysPercent(days int) float32 {
	ago := time.Now().UTC().Add((-24 * time.Duration(days)) * time.Hour)
	return s.OnlineSince(ago)
}

// OnlineSince accepts a time since parameter to return the percent of a service's uptime.
func (s *ServiceObj) OnlineSince(ago time.Time) float32 {
	failed := s.Failures().Since(ago)
	if len(failed) == 0 {
		s.Online24Hours = 100.00
		return s.Online24Hours
	}
	total := s.Hits().Since(ago)
	if len(total) == 0 {
		s.Online24Hours = 0
		return s.Online24Hours
	}
	avg := float64(len(failed)) / float64(len(total)) * 100
	avg = 100 - avg
	if avg < 0 {
		avg = 0
	}
	amount, _ := strconv.ParseFloat(fmt.Sprintf("%0.2f", avg), 10)
	s.Online24Hours = float32(amount)
	return s.Online24Hours
}

// Downtime returns the amount of time of a offline service
func (s *ServiceObj) Downtime() time.Duration {
	hits := s.Hits().Last(1)
	fail := s.Failures().Last(1)
	if fail == nil {
		return time.Duration(0)
	}
	if hits == nil {
		return time.Now().UTC().Sub(fail.CreatedAt.UTC())
	}
	since := fail.CreatedAt.UTC().Sub(hits.CreatedAt.UTC())
	return since
}
