package services

import (
	"fmt"
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/errors"
	"github.com/statping/statping/utils"
	"sort"
)

var (
	db          database.Database
	log         = utils.Log.WithField("type", "service")
	allServices map[int64]*Service
)

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

func (s *Service) Update() error {
	q := db.Update(s)
	return q.Error()
}

func (s *Service) Delete() error {
	q := db.Delete(s)
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
