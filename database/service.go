package database

import "github.com/hunterlong/statping/types"

type ServiceObj struct {
	*types.Service
	failures
}

func (o *Object) AsService() *types.Service {
	return o.model.(*types.Service)
}

func Service(id int64) (HitsFailures, error) {
	var service types.Service
	query := database.Model(&types.Service{}).Where("id = ?", id).Find(&service)
	return &ServiceObj{Service: &service}, query.Error()
}

func (s *ServiceObj) Hits() *hits {
	return &hits{
		database.Model(&types.Hit{}).Where("service = ?", s.Id),
	}
}

func (s *ServiceObj) Failures() *failures {
	return &failures{
		database.Model(&types.Failure{}).
			Where("method != 'checkin' AND service = ?", s.Id).Order("id desc"),
	}
}
