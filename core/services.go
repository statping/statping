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
	"github.com/ararog/timeago"
	"github.com/hunterlong/statup/core/notifier"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"sort"
	"strconv"
	"time"
)

type Service struct {
	*types.Service
}

func (s *Service) Select() *types.Service {
	return s.Service
}

func ReturnService(s *types.Service) *Service {
	return &Service{s}
}

// SelectService returns a *core.Service from in memory
func SelectService(id int64) *Service {
	for _, s := range CoreApp.Services {
		if s.Select().Id == id {
			return s.(*Service)
		}
	}
	return nil
}

// SelectAllServices returns a slice of *core.Service to be store on []*core.Services, should only be called once on startup.
func (c *Core) SelectAllServices() ([]*Service, error) {
	var services []*Service
	db := servicesDB().Find(&services).Order("order_id desc")
	if db.Error != nil {
		utils.Log(3, fmt.Sprintf("service error: %v", db.Error))
		return nil, db.Error
	}
	CoreApp.Services = nil
	for _, service := range services {
		service.Start()
		service.AllCheckins()
		service.AllFailures()
		CoreApp.Services = append(CoreApp.Services, service)
	}
	sort.Sort(ServiceOrder(CoreApp.Services))
	return services, db.Error
}

// reorderServices will sort the services based on 'order_id'
func reorderServices() {
	sort.Sort(ServiceOrder(CoreApp.Services))
}

// ToJSON will convert a service to a JSON string
func (s *Service) ToJSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}

// AvgTime will return the average amount of time for a service to response back successfully
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

// Online24 returns the service's uptime percent within last 24 hours
func (s *Service) Online24() float32 {
	ago := time.Now().Add(-24 * time.Hour)
	return s.OnlineSince(ago)
}

// OnlineSince accepts a time since parameter to return the percent of a service's uptime.
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

// DateScan struct is for creating the charts.js graph JSON array
type DateScan struct {
	CreatedAt time.Time `json:"x"`
	Value     int64     `json:"y"`
}

// DateScanObj struct is for creating the charts.js graph JSON array
type DateScanObj struct {
	Array []DateScan
}

// lastFailure returns the last failure a service had
func (s *Service) lastFailure() *Failure {
	limited := s.LimitedFailures()
	if len(limited) == 0 {
		return nil
	}
	last := limited[len(limited)-1]
	return last
}

// SmallText returns a short description about a services status
func (s *Service) SmallText() string {
	last := s.LimitedFailures()
	hits, _ := s.LimitedHits()
	zone := CoreApp.Timezone
	if s.Online {
		if len(last) == 0 {
			return fmt.Sprintf("Online since %v", utils.Timezoner(s.CreatedAt, zone).Format("Monday 3:04:05PM, Jan _2 2006"))
		} else {
			return fmt.Sprintf("Online, last failure was %v", utils.Timezoner(hits[0].CreatedAt, zone).Format("Monday 3:04:05PM, Jan _2 2006"))
		}
	}
	if len(last) > 0 {
		lastFailure := s.lastFailure()
		got, _ := timeago.TimeAgoWithTime(time.Now().Add(s.Downtime()), time.Now())
		return fmt.Sprintf("Reported offline %v, %v", got, lastFailure.ParseError())
	} else {
		return fmt.Sprintf("%v is currently offline", s.Name)
	}
}

func (s *Service) DowntimeText() string {
	lastFailure := s.lastFailure()
	if lastFailure == nil {
		return ""
	}
	got, _ := timeago.TimeAgoWithTime(time.Now().UTC().Add(s.Downtime()), time.Now().UTC())
	return fmt.Sprintf("Reported offline %v, %v", got, lastFailure.ParseError())
}

// GroupDataBy returns a SQL query as a string to group a column by a time
func GroupDataBy(column string, id int64, start, end time.Time, increment string) string {
	var sql string
	switch CoreApp.DbConnection {
	case "mysql":
		sql = fmt.Sprintf("SELECT CONCAT(date_format(created_at, '%%Y-%%m-%%dT%%H:%%i:00Z')) AS created_at, AVG(latency)*1000 AS value FROM %v WHERE service=%v AND DATE_FORMAT(created_at, '%%Y-%%m-%%dT%%TZ') BETWEEN DATE_FORMAT('%v', '%%Y-%%m-%%dT%%TZ') AND DATE_FORMAT('%v', '%%Y-%%m-%%dT%%TZ') GROUP BY 1 ORDER BY created_at ASC;", column, id, start.UTC().Format(time.RFC3339), end.UTC().Format(time.RFC3339))
	case "sqlite":
		sql = fmt.Sprintf("SELECT strftime('%%Y-%%m-%%dT%%H:%%M:00Z', created_at), AVG(latency)*1000 as value FROM %v WHERE service=%v AND created_at >= '%v' AND created_at <= '%v' GROUP BY strftime('%%M:00', created_at) ORDER BY created_at ASC;", column, id, start.UTC().Format(time.RFC3339), end.UTC().Format(time.RFC3339))
	case "postgres":
		sql = fmt.Sprintf("SELECT date_trunc('%v', created_at), AVG(latency)*1000 AS value FROM %v WHERE service=%v AND created_at >= '%v' AND created_at <= '%v' GROUP BY 1 ORDER BY date_trunc ASC;", increment, column, id, start.UTC().Format(time.RFC3339), end.UTC().Format(time.RFC3339))
	}
	return sql
}

// Downtime returns the amount of time of a offline service
func (s *Service) Downtime() time.Duration {
	hits, _ := s.Hits()
	if len(hits) == 0 {
		return time.Duration(0)
	}
	fails := s.LimitedFailures()
	if len(fails) == 0 {
		return time.Duration(0)
	}
	since := fails[0].CreatedAt.UTC().Sub(hits[0].CreatedAt.UTC())
	return since
}

func GraphDataRaw(service types.ServiceInterface, start, end time.Time) *DateScanObj {
	var d []DateScan
	s := service.Select()
	sql := GroupDataBy("hits", s.Id, start, end, "minute")
	rows, err := DbSession.Raw(sql).Rows()
	if err != nil {
		utils.Log(2, err)
		return nil
	}
	for rows.Next() {
		var gd DateScan
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
		gd.CreatedAt = utils.Timezoner(gd.CreatedAt, CoreApp.Timezone)
		gd.Value = int64(ff)
		d = append(d, gd)
	}
	return &DateScanObj{d}
}

func (d *DateScanObj) ToString() string {
	data, err := json.Marshal(d.Array)
	if err != nil {
		utils.Log(2, err)
		return "{}"
	}
	return string(data)
}

// GraphData returns the JSON object used by Charts.js to render the chart
func (s *Service) GraphData() string {
	start := time.Now().Add(time.Hour*-24 + time.Minute*0 + time.Second*0)
	end := time.Now()
	obj := GraphDataRaw(s, start, end)
	data, err := json.Marshal(obj)
	if err != nil {
		utils.Log(2, err)
		return ""
	}
	return string(data)
}

// AvgUptime24 returns a service's average online status for last 24 hours
func (s *Service) AvgUptime24() string {
	ago := time.Now().Add(-24 * time.Hour)
	return s.AvgUptime(ago)
}

// AvgUptime returns average online status for last 24 hours
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

// TotalUptime returns the total uptime percent of a service
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

// index returns a services index int for updating the []*core.Services slice
func (s *Service) index() int {
	for k, service := range CoreApp.Services {
		if s.Id == service.(*Service).Id {
			return k
		}
	}
	return 0
}

// updateService will update a service in the []*core.Services slice
func updateService(service *Service) {
	index := service.index()
	CoreApp.Services[index] = service
}

// Delete will remove a service from the database, it will also end the service checking go routine
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
	reorderServices()
	notifier.OnDeletedService(u.Service)
	return err.Error
}

// UpdateSingle will update a single column for a service
func (u *Service) UpdateSingle(attr ...interface{}) error {
	return servicesDB().Model(u).Update(attr).Error
}

// Update will update a service in the database, the service's checking routine can be restarted by passing true
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
	reorderServices()
	updateService(u)
	notifier.OnUpdatedService(u.Service)
	return err.Error
}

// Create will create a service and insert it into the database
func (u *Service) Create(check bool) (int64, error) {
	u.CreatedAt = time.Now()
	db := servicesDB().Create(u)
	if db.Error != nil {
		utils.Log(3, fmt.Sprintf("Failed to create service %v #%v: %v", u.Name, u.Id, db.Error))
		return 0, db.Error
	}
	u.Start()
	go u.CheckQueue(check)
	CoreApp.Services = append(CoreApp.Services, u)
	reorderServices()
	notifier.OnNewService(u.Service)
	return u.Id, nil
}

// ServicesCount returns the amount of services inside the []*core.Services slice
func (c *Core) ServicesCount() int {
	return len(c.Services)
}

// CountOnline
func (c *Core) CountOnline() int {
	amount := 0
	for _, s := range CoreApp.Services {
		if s.Select().Online {
			amount++
		}
	}
	return amount
}
