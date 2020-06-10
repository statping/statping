package services

import (
	"github.com/statping/statping/types/checkins"
)

// CheckinProcess runs the checkin routine for each checkin attached to service
func CheckinProcess(s *Service) {
	for _, c := range s.Checkins() {
		c.Start()
	}
}

func (s *Service) Checkins() []*checkins.Checkin {
	var chks []*checkins.Checkin
	db.Where("service = ?", s.Id).Find(&chks)
	return chks
}
