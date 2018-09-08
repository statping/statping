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
)

type Service struct {
	*types.Service
}

func ReturnService(s *types.Service) *Service {
	return &Service{s}
}

func SelectService(id int64) *Service {
	for _, s := range CoreApp.Services {
		if s.(*Service).Id == id {
			return s.(*Service)
		}
	}
	return nil
}

func (c *Core) SelectAllServices() ([]*Service, error) {
	var services []*Service
	db := servicesDB().Find(&services).Order("order_id desc")
	if db.Error != nil {
		utils.Log(3, fmt.Sprintf("service error: %v", db.Error))
		return nil, db.Error
	}
	for _, service := range services {
		service.Start()
		service.AllCheckins()
		service.AllFailures()
		CoreApp.Services = append(CoreApp.Services, service)
	}
	return services, db.Error
}

func (s *Service) ToJSON() string {
	data, _ := json.Marshal(s)
	return string(data)
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

func (s *Service) Online24() float32 {
	ago := time.Now().Add(-24 * time.Hour)
	return s.OnlineSince(ago)
}

func (s *Service) OnlineSince(ago time.Time) float32 {
	failed, _ := s.TotalFailuresSince(ago)
	if failed == 0 {
		s.Online24Hours = 100.00
		return s.Online24Hours
	}
	total, _ := s.TotalHitsSince(ago)
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

func (s *Service) lastFailure() *Failure {
	limited := s.LimitedFailures()
	last := limited[len(limited)-1]
	return last
}

func (s *Service) SmallText() string {
	last := s.LimitedFailures()
	hits, _ := s.LimitedHits()
	if s.Online {
		if len(last) == 0 {
			return fmt.Sprintf("Online since %v", s.CreatedAt.Format("Monday 3:04PM, Jan _2 2006"))
		} else {
			return fmt.Sprintf("Online, last failure was %v", hits[0].CreatedAt.Format("Monday 3:04PM, Jan _2 2006"))
		}
	}
	if len(last) > 0 {
		lastFailure := s.lastFailure()
		return fmt.Sprintf("%v on %v", lastFailure.ParseError(), last[0].CreatedAt.Format("Monday 3:04PM, Jan _2 2006"))
	} else {
		return fmt.Sprintf("%v is currently offline", s.Name)
	}
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

func (s *Service) GraphData() string {
	var d []*DateScan
	since := time.Now().Add(time.Hour*-24 + time.Minute*0 + time.Second*0)
	sql := GroupDataBy("hits", s.Id, since, "minute")
	rows, err := DbSession.Raw(sql).Rows()
	defer rows.Close()
	if err != nil {
		utils.Log(2, err)
		return ""
	}
	for rows.Next() {
		gd := new(DateScan)
		var tt string
		var ff float64
		err := rows.Scan(&tt, &ff)
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

func (s *Service) AvgUptime24() string {
	ago := time.Now().Add(-24 * time.Hour)
	return s.AvgUptime(ago)
}

func (s *Service) AvgUptime(ago time.Time) string {
	failed, _ := s.TotalFailuresSince(ago)
	if failed == 0 {
		return "100"
	}
	total, _ := s.TotalHitsSince(ago)
	if total == 0 {
		return "0.00"
	}
	percent := float64(failed) / float64(total) * 100
	percent = 100 - percent
	if percent < 0 {
		percent = 0
	}
	amount := fmt.Sprintf("%0.2f", percent)
	if amount == "100.00" {
		amount = "100"
	}
	return amount
}

func (s *Service) TotalUptime() string {
	hits, _ := s.TotalHits()
	failures, _ := s.TotalFailures()
	percent := float64(failures) / float64(hits) * 100
	percent = 100 - percent
	if percent < 0 {
		percent = 0
	}
	amount := fmt.Sprintf("%0.2f", percent)
	if amount == "100.00" {
		amount = "100"
	}
	return amount
}

func (s *Service) index() int {
	for k, service := range CoreApp.Services {
		if s.Id == service.(*Service).Id {
			return k
		}
	}
	return 0
}

func updateService(service *Service) {
	index := service.index()
	CoreApp.Services[index] = service
}

func (u *Service) Delete() error {
	i := u.index()
	err := servicesDB().Delete(u)
	if err.Error != nil {
		utils.Log(3, fmt.Sprintf("Failed to delete service %v. %v", u.Name, err.Error))
		return err.Error
	}
	u.Close()
	slice := CoreApp.Services
	CoreApp.Services = append(slice[:i], slice[i+1:]...)
	OnDeletedService(u)
	return err.Error
}

func (u *Service) UpdateSingle(attr ...interface{}) error {
	return servicesDB().Model(u).Update(attr).Error
}

func (u *Service) Update(restart bool) error {
	err := servicesDB().Update(u)
	if err.Error != nil {
		utils.Log(3, fmt.Sprintf("Failed to update service %v. %v", u.Name, err))
		return err.Error
	}
	if restart {
		u.Close()
		u.Start()
		u.SleepDuration = time.Duration(u.Interval) * time.Second
		go u.CheckQueue(true)
	}
	updateService(u)
	OnUpdateService(u)
	return err.Error
}

func (u *Service) Create() (int64, error) {
	u.CreatedAt = time.Now()
	db := servicesDB().Create(u)
	if db.Error != nil {
		utils.Log(3, fmt.Sprintf("Failed to create service %v #%v: %v", u.Name, u.Id, db.Error))
		return 0, db.Error
	}
	u.Start()
	go u.CheckQueue(true)
	CoreApp.Services = append(CoreApp.Services, u)
	return u.Id, nil
}

func (c *Core) ServicesCount() int {
	return len(c.Services)
}

func (c *Core) CountOnline() int {
	amount := 0
	for _, s := range CoreApp.Services {
		if s.(*Service).Online {
			amount++
		}
	}
	return amount
}
