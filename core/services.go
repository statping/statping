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
	*database.ServiceObj
}

func Services() map[int64]*Service {
	return CoreApp.services
}

// SelectService returns a *core.Service from in memory
func SelectService(id int64) *Service {
	service := CoreApp.services[id]
	if service != nil {
		return service
	}
	return nil
}

func (s *Service) AfterCreate(obj interface{}, err error) {

}

// CheckinProcess runs the checkin routine for each checkin attached to service
func CheckinProcess(s database.Servicer) {
	for _, c := range s.Checkins() {
		c.Start()
		go CheckinRoutine(c)
	}
}

// SelectAllServices returns a slice of *core.Service to be store on []*core.Services
// should only be called once on startup.
func SelectAllServices(start bool) (map[int64]*Service, error) {
	services := make(map[int64]*Service)
	if len(CoreApp.services) > 0 {
		return CoreApp.services, nil
	}

	for _, s := range database.Services() {
		if start {
			s.Start()
			CheckinProcess(s)
		}

		fails := s.Failures().Last(limitedFailures)
		s.Service.Failures = fails

		for _, c := range s.Checkins() {
			s.Service.Checkins = append(s.Service.Checkins, c.Checkin)
		}

		// collect initial service stats
		s.UpdateStats()

		services[s.Id] = &Service{s}
	}

	CoreApp.services = services
	reorderServices()

	return services, nil
}

func wrapFailures(f []*types.Failure) []*Failure {
	var fails []*Failure
	for _, v := range f {
		fails = append(fails, &Failure{v})
	}
	return fails
}

// reorderServices will sort the services based on 'order_id'
func reorderServices() {
	sort.Sort(ServiceOrder(CoreApp.services))
}

// updateService will update a service in the []*core.Services slice
func updateService(s *Service) {
	CoreApp.services[s.Id] = s
}

// Delete will remove a service from the database, it will also end the service checking go routine
func (s *Service) Delete() error {
	err := database.Delete(s)
	if err != nil {
		log.Errorln(fmt.Sprintf("Failed to delete service %v. %v", s.Name, err))
		return err
	}
	s.Close()
	CoreApp.services[s.Id] = nil
	reorderServices()
	notifier.OnDeletedService(s.Service)
	return err
}

// Update will update a service in the database, the service's checking routine can be restarted by passing true
func Update(s *Service, restart bool) error {
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
		s.SleepDuration = s.Duration()
		go ServiceCheckQueue(s, true)
	}
	reorderServices()
	updateService(s)
	notifier.OnUpdatedService(s.Service)
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
	service := &Service{s}
	s.Start()
	CoreApp.services[service.Id] = service
	go ServiceCheckQueue(service, check)
	reorderServices()
	notifier.OnNewService(s.Service)
	return s.Id, nil
}
