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
	"encoding/json"
	"fmt"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"strconv"
	"time"
	"upper.io/db.v3"
)

type Service struct {
	s *types.Service
}

type Failure struct {
	F interface{}
}

func serviceCol() db.Collection {
	return DbSession.Collection("services")
}

func SelectService(id int64) *Service {
	for _, s := range CoreApp.Services {
		ser := s.ToService()
		if ser.Id == id {
			return &Service{ser}
		}
	}
	return nil
}

func SelectAllServices() ([]*Service, error) {
	var services []*types.Service
	var sers []*Service
	col := serviceCol().Find()
	err := col.All(&services)
	if err != nil {
		utils.Log(3, fmt.Sprintf("service error: %v", err))
		return nil, err
	}
	for _, s := range services {
		ser := NewService(s)
		sers = append(sers, ser)
		s.Checkins = SelectAllCheckins(s)
		s.Failures = SelectAllFailures(s)
	}
	CoreApp.Services = sers
	return sers, err
}

func (s *Service) AvgTime() float64 {
	total, _ := s.TotalHits()
	if total == 0 {
		return float64(0)
	}
	sum, _ := s.Sum()
	avg := sum / float64(total) * 100
	amount := fmt.Sprintf("%0.0f", avg*10)
	val, _ := strconv.ParseFloat(amount, 10)
	return val
}

func (ser *Service) Online24() float32 {
	s := ser.ToService()
	total, _ := ser.TotalHits()
	failed, _ := ser.TotalFailures24Hours()
	if failed == 0 {
		s.Online24Hours = 100.00
		return s.Online24Hours
	}
	if total == 0 {
		s.Online24Hours = 0
		return s.Online24Hours
	}
	avg := float64(failed) / float64(total) * 100
	avg = 100 - avg
	if avg < 0 {
		avg = 0
	}
	amount, _ := strconv.ParseFloat(fmt.Sprintf("%0.2f", avg), 10)
	s.Online24Hours = float32(amount)
	return s.Online24Hours
}

type DateScan struct {
	CreatedAt time.Time `json:"x"`
	Value     int64     `json:"y"`
}

func (s *Service) ToService() *types.Service {
	return s.s
}

func NewService(s *types.Service) *Service {
	return &Service{s}
}

func (ser *Service) SmallText() string {
	s := ser.ToService()
	last := ser.LimitedFailures()
	hits, _ := ser.LimitedHits()
	if !s.Online {
		if len(last) > 0 {
			lastFailure := MakeFailure(last[0].ToFailure())
			return fmt.Sprintf("%v on %v", lastFailure.ParseError(), last[0].ToFailure().CreatedAt.Format("Monday 3:04PM, Jan _2 2006"))
		} else {
			return fmt.Sprintf("%v is currently offline", s.Name)
		}
	} else {
		if len(last) == 0 {
			return fmt.Sprintf("Online since %v", s.CreatedAt.Format("Monday 3:04PM, Jan _2 2006"))
		} else {
			return fmt.Sprintf("Online, last failure was %v", hits[0].CreatedAt.Format("Monday 3:04PM, Jan _2 2006"))
		}
	}
	return fmt.Sprintf("No Failures in the last 24 hours! %v", hits[0])
}

func GroupDataBy(column string, id int64, tm time.Time, increment string) string {
	var sql string
	switch CoreApp.DbConnection {
	case "mysql":
		sql = fmt.Sprintf("SELECT CONCAT(date_format(created_at, '%%Y-%%m-%%dT%%H:%%i:00Z')) AS created_at, AVG(latency)*1000 AS value FROM %v WHERE service=%v AND DATE_FORMAT(created_at, '%%Y-%%m-%%dT%%TZ') BETWEEN DATE_FORMAT('%v', '%%Y-%%m-%%dT%%TZ') AND DATE_FORMAT(NOW(), '%%Y-%%m-%%dT%%TZ') GROUP BY 1 ORDER BY created_at ASC;", column, id, tm.Format(time.RFC3339))
	case "sqlite":
		sql = fmt.Sprintf("SELECT strftime('%%Y-%%m-%%dT%%H:%%M:00Z', created_at), AVG(latency)*1000 as value FROM %v WHERE service=%v AND created_at >= '%v' GROUP BY strftime('%%M:00', created_at) ORDER BY created_at ASC;", column, id, tm.Format(time.RFC3339))
	case "postgres":
		sql = fmt.Sprintf("SELECT date_trunc('%v', created_at), AVG(latency)*1000 AS value FROM %v WHERE service=%v AND created_at >= '%v' GROUP BY 1 ORDER BY date_trunc ASC;", increment, column, id, tm.Format(time.RFC3339))
	}
	return sql
}

func (ser *Service) GraphData() string {
	s := ser.ToService()
	var d []*DateScan
	since := time.Now().Add(time.Hour*-24 + time.Minute*0 + time.Second*0)

	sql := GroupDataBy("hits", s.Id, since, "minute")

	dated, err := DbSession.Query(db.Raw(sql))
	if err != nil {
		utils.Log(2, err)
		return ""
	}
	for dated.Next() {
		gd := new(DateScan)
		var tt string
		var ff float64
		err := dated.Scan(&tt, &ff)
		if err != nil {
			utils.Log(2, fmt.Sprintf("Issue loading chart data for service %v, %v", s.Name, err))
		}
		gd.CreatedAt, err = time.Parse(time.RFC3339, tt)
		if err != nil {
			utils.Log(2, fmt.Sprintf("Issue parsing time %v", err))
		}
		gd.Value = int64(ff)
		d = append(d, gd)
	}
	data, err := json.Marshal(d)
	if err != nil {
		utils.Log(2, err)
		return ""
	}
	return string(data)
}

func (ser *Service) AvgUptime() string {
	s := ser.ToService()
	failed, _ := ser.TotalFailures()
	total, _ := ser.TotalHits()
	if failed == 0 {
		s.TotalUptime = "100"
		return s.TotalUptime
	}
	if total == 0 {
		s.TotalUptime = "0"
		return s.TotalUptime
	}
	percent := float64(failed) / float64(total) * 100
	percent = 100 - percent
	if percent < 0 {
		percent = 0
	}
	s.TotalUptime = fmt.Sprintf("%0.2f", percent)
	if s.TotalUptime == "100.00" {
		s.TotalUptime = "100"
	}
	return s.TotalUptime
}

func RemoveArray(u *types.Service) []*Service {
	var srvcs []*Service
	for _, s := range CoreApp.Services {
		ser := s.ToService()
		if ser.Id != u.Id {
			srvcs = append(srvcs, s)
		}
	}
	CoreApp.Services = srvcs
	return srvcs
}

func DeleteService(u *types.Service) error {
	res := serviceCol().Find("id", u.Id)
	err := res.Delete()
	if err != nil {
		utils.Log(3, fmt.Sprintf("Failed to delete service %v. %v", u.Name, err))
		return err
	}
	utils.Log(1, fmt.Sprintf("Stopping %v Monitoring...", u.Name))
	if u.StopRoutine != nil {
		close(u.StopRoutine)
	}
	utils.Log(1, fmt.Sprintf("Stopped %v Monitoring Service", u.Name))
	RemoveArray(u)
	OnDeletedService(u)
	return err
}

func UpdateService(service *types.Service) *types.Service {
	service.CreatedAt = time.Now()
	res := serviceCol().Find("id", service.Id)
	err := res.Update(service)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Failed to update service %v. %v", service.Name, err))
		return service
	}
	CoreApp.Services, _ = SelectAllServices()
	OnUpdateService(service)
	return service
}

func updateService(u *types.Service) {
	var services []*Service
	for _, s := range CoreApp.Services {
		if s.s.Id == u.Id {
			s.s = u
		}
		services = append(services, s)
	}
	CoreApp.Services = services
}

func CreateService(u *types.Service) (int64, error) {
	u.CreatedAt = time.Now()
	uuid, err := serviceCol().Insert(u)
	if uuid == nil {
		utils.Log(3, fmt.Sprintf("Failed to create service %v. %v", u.Name, err))
		return 0, err
	}
	u.Id = uuid.(int64)
	u.StopRoutine = make(chan bool)
	CoreApp.Services = append(CoreApp.Services, &Service{u})
	return uuid.(int64), err
}

func CountOnline() int {
	amount := 0
	for _, s := range CoreApp.Services {
		ser := s.ToService()
		if ser.Online {
			amount++
		}
	}
	return amount
}
