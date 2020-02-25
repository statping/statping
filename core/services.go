// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
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
	"fmt"
	"github.com/hunterlong/statping/core/notifier"
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types"
	"sort"
	"time"
)

type Service struct {
	*types.Service
}

type Servicer interface{}

func Services() []database.Servicer {
	return CoreApp.services
}

// SelectService returns a *core.Service from in memory
func SelectService(id int64) *types.Service {
	for _, s := range Services() {
		if s.Model().Id == id {
			fmt.Println("service: ", s.Model())
			return s.Model()
		}
	}
	return nil
}

// CheckinProcess runs the checkin routine for each checkin attached to service
func CheckinProcess(s database.Servicer) {
	for _, c := range s.AllCheckins() {
		c.Start()
		go CheckinRoutine(c)
	}
}

// SelectAllServices returns a slice of *core.Service to be store on []*core.Services
// should only be called once on startup.
func SelectAllServices(start bool) ([]*database.ServiceObj, error) {
	srvs := database.Services()
	for _, s := range srvs {
		fmt.Println("services: ", s.Id, s.Name)
	}

	for _, s := range srvs {
		if start {
			service := s.Model()
			service.Start()
			CheckinProcess(s)
		}
		//fails := service.Service (limitedFailures)
		//for _, f := range fails {
		//	service.Failures = append(service.Failures, f)
		//}
		for _, c := range s.AllCheckins() {
			s.Checkins = append(s.Checkins, c)
		}
		// collect initial service stats
		s.Service.Stats = s.UpdateStats()
		CoreApp.services = append(CoreApp.services, s)
	}
	reorderServices()
	return srvs, nil
}

// reorderServices will sort the services based on 'order_id'
func reorderServices() {
	sort.Sort(ServiceOrder(CoreApp.services))
}

// GraphData will return all hits or failures
func GraphData(q *database.GroupQuery, dbType interface{}, by database.By) []*database.TimeValue {
	dbQuery, err := q.Database().GroupQuery(q, by).ToTimeValue(dbType)

	if err != nil {
		log.Error(err)
		return nil
	}
	if q.FillEmpty {
		return dbQuery.FillMissing(q.Start, q.End)
	}
	return dbQuery.ToValues()
}

// index returns a services index int for updating the []*core.Services slice
func index(s database.Servicer) int {
	for k, service := range CoreApp.services {
		if s.Model().Id == service.Model().Id {
			return k
		}
	}
	return 0
}

// updateService will update a service in the []*core.Services slice
func updateService(s database.Servicer) {
	CoreApp.services[index(s)] = s
}

// Delete will remove a service from the database, it will also end the service checking go routine
func Delete(srv database.Servicer) error {
	i := index(srv)
	s := srv.Model()
	err := database.Delete(s)
	if err != nil {
		log.Errorln(fmt.Sprintf("Failed to delete service %v. %v", s.Name, err))
		return err
	}
	s.Close()
	slice := CoreApp.services
	CoreApp.services = append(slice[:i], slice[i+1:]...)
	reorderServices()
	notifier.OnDeletedService(s)
	return err
}

// Update will update a service in the database, the service's checking routine can be restarted by passing true
func Update(srv database.Servicer, restart bool) error {
	s := srv.Model()
	err := database.Update(s)
	if err != nil {
		log.Errorln(fmt.Sprintf("Failed to update service %v. %v", s.Name, err))
		return err
	}
	// clear the notification queue for a service
	if !s.AllowNotifications.Bool {
		for _, n := range CoreApp.Notifications {
			notif := n.(notifier.Notifier).Select()
			notif.ResetUniqueQueue(fmt.Sprintf("service_%v", s.Id))
		}
	}
	if restart {
		s.Close()
		s.Start()
		s.SleepDuration = time.Duration(s.Interval) * time.Second
		go ServiceCheckQueue(srv, true)
	}
	reorderServices()
	updateService(srv)
	notifier.OnUpdatedService(s)
	return err
}

// Create will create a service and insert it into the database
func Create(srv database.Servicer, check bool) (int64, error) {
	s := srv.Model()
	s.CreatedAt = time.Now().UTC()
	_, err := database.Create(s)
	if err != nil {
		log.Errorln(fmt.Sprintf("Failed to create service %v #%v: %v", s.Name, s.Id, err))
		return 0, err
	}
	s.Start()
	go ServiceCheckQueue(srv, check)
	CoreApp.services = append(CoreApp.services, srv)
	reorderServices()
	notifier.OnNewService(s)
	return s.Id, nil
}
