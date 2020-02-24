package database

import "github.com/hunterlong/statping/types"

type ServiceObj struct {
	db      Database
	service *types.Service
}

type Servicer interface {
	Failures() Database
	Hits() Database
}

func Service(id int64) (Servicer, error) {
	var service types.Service
	query := database.Model(&types.Service{}).Where("id = ?", id).Find(&service)
	return &ServiceObj{query, &service}, query.Error()
}

func (s *ServiceObj) Failures() Database {
	return database.Model(&types.Failure{}).Where("service = ?", s.service.Id)
}

func (s *ServiceObj) Hits() Database {
	return database.Model(&types.Hit{}).Where("service = ?", s.service.Id)
}
