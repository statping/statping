package services

import (
	"errors"
	"fmt"
	"github.com/statping/statping/database"
	"github.com/statping/statping/utils"
	"sort"
)

var log = utils.Log

var db database.Database

func SetDB(database database.Database) {
	db = database.Model(&Service{})
}

func Find(id int64) (*Service, error) {
	srv := allServices[id]
	if srv == nil {
		return nil, errors.New("service not found")
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
	return nil
}

func (s *Service) Update() error {
	q := db.Update(s)

	allServices[s.Id] = s

	if !s.AllowNotifications.Bool {
		//for _, n := range CoreApp.Notifications {
		//	notif := n.(notifier.Notifier).Select()
		//	notif.ResetUniqueQueue(fmt.Sprintf("service_%v", s.Id))
		//}
	}
	s.Close()
	s.SleepDuration = s.Duration()
	go ServiceCheckQueue(allServices[s.Id], true)

	//notifier.OnUpdatedService(s.Service)

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
	//notifier.OnDeletedService(s.Service)
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

//func (s *Service) AfterDelete() error {
//
//	return nil
//}

func (s *Service) AfterFind() error {
	return nil
}
