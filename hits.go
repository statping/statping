package main

import (
	"time"
	"upper.io/db.v3"
)

type Hit struct {
	Id        int       `db:"id,omitempty"`
	Service   int64     `db:"service"`
	Latency   float64   `db:"latency"`
	CreatedAt time.Time `db:"created_at"`
}

func hitCol() db.Collection {
	return dbSession.Collection("hits")
}

func (s *Service) CreateHit(d HitData) (int64, error) {
	h := Hit{
		Service:   s.Id,
		Latency:   d.Latency,
		CreatedAt: time.Now(),
	}
	uuid, err := hitCol().Insert(h)
	if uuid == nil {
		return 0, err
	}
	return uuid.(int64), err
}

func (s *Service) Hits() ([]Hit, error) {
	var hits []Hit
	col := hitCol().Find("service", s.Id).OrderBy("-id")
	err := col.All(&hits)
	return hits, err
}

func (s *Service) LimitedHits() ([]Hit, error) {
	var hits []Hit
	col := hitCol().Find("service", s.Id).Limit(1056).OrderBy("-id")
	err := col.All(&hits)
	return hits, err
}

func (s *Service) SelectHitsGroupBy(group string) ([]Hit, error) {
	var hits []Hit
	col := hitCol().Find("service", s.Id)
	err := col.All(&hits)
	return hits, err
}

func (s *Service) TotalHits() (uint64, error) {
	col := hitCol().Find("service", s.Id)
	amount, err := col.Count()
	return amount, err
}

func (s *Service) Sum() (float64, error) {
	var amount float64
	hits, err := s.Hits()
	for _, h := range hits {
		amount += h.Latency
	}
	return amount, err
}
