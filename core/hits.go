package core

import (
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"time"
	"upper.io/db.v3"
)

type Hit types.Hit

func hitCol() db.Collection {
	return DbSession.Collection("hits")
}

func CreateServiceHit(s *types.Service, d HitData) (int64, error) {
	h := Hit{
		Service:   s.Id,
		Latency:   d.Latency,
		CreatedAt: time.Now(),
	}
	uuid, err := hitCol().Insert(h)
	if uuid == nil {
		utils.Log(2, err)
		return 0, err
	}
	return uuid.(int64), err
}

func (ser *Service) Hits() ([]Hit, error) {
	s := ser.ToService()
	var hits []Hit
	col := hitCol().Find("service", s.Id).OrderBy("-id")
	err := col.All(&hits)
	return hits, err
}

func (ser *Service) LimitedHits() ([]*Hit, error) {
	s := ser.ToService()
	var hits []*Hit
	col := hitCol().Find("service", s.Id).OrderBy("-id").Limit(1024)
	err := col.All(&hits)
	return reverseHits(hits), err
}

func reverseHits(input []*Hit) []*Hit {
	if len(input) == 0 {
		return input
	}
	return append(reverseHits(input[1:]), input[0])
}

func (ser *Service) SelectHitsGroupBy(group string) ([]Hit, error) {
	s := ser.ToService()
	var hits []Hit
	col := hitCol().Find("service", s.Id)
	err := col.All(&hits)
	return hits, err
}

func (ser *Service) TotalHits() (uint64, error) {
	s := ser.ToService()
	col := hitCol().Find("service", s.Id)
	amount, err := col.Count()
	return amount, err
}

func (s *Service) Sum() (float64, error) {
	var amount float64
	hits, err := s.Hits()
	if err != nil {
		utils.Log(2, err)
	}
	for _, h := range hits {
		amount += h.Latency
	}
	return amount, err
}
