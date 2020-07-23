package services

import (
	"github.com/statping/statping/types/incidents"
)

func (s *Service) Incidents() []*incidents.Incident {
	var i []*incidents.Incident
	db.Where("service = ?", s.Id).Find(&i)
	return i
}

func (s *Service) DeleteIncidents() error {
	for _, i := range s.Incidents() {
		if err := i.Delete(); err != nil {
			return err
		}
	}
	return nil
}
