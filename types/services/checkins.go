package services

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types/checkins"
)

// CheckinProcess runs the checkin routine for each checkin attached to service
func CheckinProcess(s *Service) {
	for _, c := range s.Checkins() {
		c.Start()
		go c.CheckinRoutine()
	}
}

func (s *Service) Checkins() []*checkins.Checkin {
	var chks []*checkins.Checkin
	database.DB().Where("service = ?", s.Id).Find(&chks)
	return chks
}
