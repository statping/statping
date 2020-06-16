package services

import (
	"fmt"
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/errors"
	"github.com/statping/statping/types/metrics"
	"github.com/statping/statping/utils"
	"sort"
)

var (
	db          database.Database
	log         = utils.Log.WithField("type", "service")
	allServices map[int64]*Service
)

func (s *Service) AfterFind() {
	metrics.Query("service", "find")
}

func (s *Service) AfterUpdate() {
	metrics.Query("service", "update")
}

func (s *Service) AfterDelete() {
	metrics.Query("service", "delete")
}

func init() {
	allServices = make(map[int64]*Service)
}

func Services() map[int64]*Service {
	return allServices
}

func SetDB(database database.Database) {
	db = database.Model(&Service{})
}

func Find(id int64) (*Service, error) {
	srv := allServices[id]
	if srv == nil {
		return nil, errors.Missing(&Service{}, id)
	}
	return srv, nil
}

func all() []*Service {
	var services []*Service
	db.Find(&services)
	return services
}

func All() map[int64]*Service {
	return allServices
}

func AllInOrder() []Service {
	var services []Service
	for _, service := range allServices {
		service.UpdateStats()
		services = append(services, *service)
	}
	sort.Sort(ServiceOrder(services))
	return services
}

func (s *Service) Create() error {
	err := db.Create(s)
	if err.Error() != nil {
		log.Errorln(fmt.Sprintf("Failed to create service %v #%v: %v", s.Name, s.Id, err))
		return err.Error()
	}
	return nil
}

func (s *Service) AfterCreate() error {
	allServices[s.Id] = s
	metrics.Query("service", "create")
	return nil
}

func (s *Service) Update() error {
	q := db.Update(s)
	allServices[s.Id] = s
	s.Close()
	s.SleepDuration = s.Duration()
	go ServiceCheckQueue(allServices[s.Id], true)
	return q.Error()
}

func (s *Service) Delete() error {
	s.Close()
	if err := s.DeleteFailures(); err != nil {
		return err
	}
	if err := s.DeleteHits(); err != nil {
		return err
	}
	delete(allServices, s.Id)
	q := db.Model(&Service{}).Delete(s)
	return q.Error()
}

func (s *Service) DeleteFailures() error {
	return s.AllFailures().DeleteAll()
}

func (s *Service) DeleteHits() error {
	return s.AllHits().DeleteAll()
}

func (s *Service) DeleteCheckins() error {
	for _, c := range s.Checkins() {
		if err := c.Delete(); err != nil {
			return err
		}
	}
	return nil
}
