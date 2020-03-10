package services

import (
	"github.com/statping/statping/types/hits"
	"time"
)

func (s *Service) HitsColumnID() (string, int64) {
	return "service", s.Id
}

func (s *Service) FirstHit() *hits.Hit {
	return hits.AllHits(s).First()
}

func (s *Service) LastHit() *hits.Hit {
	return hits.AllHits(s).Last()
}

func (s *Service) AllHits() hits.Hitters {
	return hits.AllHits(s)
}

func (s *Service) HitsSince(t time.Time) hits.Hitters {
	return hits.Since(t, s)
}
