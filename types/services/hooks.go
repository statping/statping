package services

import "github.com/statping/statping/utils"

// BeforeCreate for Service will set CreatedAt to UTC
func (s *Service) BeforeCreate() (err error) {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = utils.Now()
		s.UpdatedAt = utils.Now()
	}
	return
}

func (s *Service) AfterCreate() error {
	allServices[s.Id] = s
	return nil
}

func (s *Service) AfterUpdate() error {
	allServices[s.Id] = s
	s.Close()
	s.SleepDuration = s.Duration()
	go ServiceCheckQueue(allServices[s.Id], true)
	return nil
}

func (s *Service) BeforeDelete() error {
	s.Close()
	if err := s.DeleteFailures(); err != nil {
		return err
	}
	if err := s.DeleteHits(); err != nil {
		return err
	}
	delete(allServices, s.Id)
	return nil
}
