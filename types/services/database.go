package services

import (
	"fmt"
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/utils"
)

var log = utils.Log

func DB() database.Database {
	return database.DB().Model(&Service{})
}

func Find(id int64) (*Service, error) {
	var service *Service
	db := DB().Where("id = ?", id).Find(&service)
	return service, db.Error()
}

func All() []*Service {
	var services []*Service
	DB().Find(&services)
	return services
}

func (s *Service) Create() error {

	err := s.Create()
	if err != nil {
		log.Errorln(fmt.Sprintf("Failed to create service %v #%v: %v", s.Name, s.Id, err))
		return err
	}
	allServices[s.Id] = s

	go ServiceCheckQueue(s, true)
	reorderServices()
	//notifications.OnNewService(s)

	return nil
}

func (s *Service) Update() error {
	db := DB().Update(&s)

	allServices[s.Id] = s

	if !s.AllowNotifications.Bool {
		//for _, n := range CoreApp.Notifications {
		//	notif := n.(notifier.Notifier).Select()
		//	notif.ResetUniqueQueue(fmt.Sprintf("service_%v", s.Id))
		//}
	}
	s.Close()
	s.Start()
	s.SleepDuration = s.Duration()
	go ServiceCheckQueue(s, true)

	reorderServices()
	//notifier.OnUpdatedService(s.Service)

	return db.Error()
}

func (s *Service) Delete() error {
	db := database.DB().Delete(&s)

	s.Close()
	delete(allServices, s.Id)
	reorderServices()
	//notifier.OnDeletedService(s.Service)

	return db.Error()
}

func (s *Service) DeleteFailures() error {
	query := database.DB().Exec(`DELETE FROM failures WHERE service = ?`, s.Id)
	return query.Error()
}
