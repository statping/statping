package database

import "github.com/hunterlong/statping/types"

type Service struct {
	db      Database
	service *types.Service
}

type Servicer interface {
	Failures() Database
	Hits() Database
}

func (it *Db) GetService(id int64) (Servicer, error) {
	var service types.Service
	query := it.Model(&types.Service{}).Where("id = ?", id).Find(&service)
	return &Service{it, &service}, query.Error()
}

func (s *Service) Failures() Database {
	return s.db.Model(&types.Failure{}).Where("service = ?", s.service.Id)
}

func (s *Service) Hits() Database {
	return s.db.Model(&types.Hit{}).Where("service = ?", s.service.Id)
}
