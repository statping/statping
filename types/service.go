package types

import "time"

type Service struct {
	Id             int64      `db:"id,omitempty" json:"id"`
	Name           string     `db:"name" json:"name"`
	Domain         string     `db:"domain" json:"domain"`
	Expected       string     `db:"expected" json:"expected"`
	ExpectedStatus int        `db:"expected_status" json:"expected_status"`
	Interval       int        `db:"check_interval" json:"check_interval"`
	Type           string     `db:"check_type" json:"type"`
	Method         string     `db:"method" json:"method"`
	PostData       string     `db:"post_data" json:"post_data"`
	Port           int        `db:"port" json:"port"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	Timeout        int        `db:"timeout" json:"timeout"`
	Order          int        `db:"order_id" json:"order_id"`
	Online         bool       `json:"online"`
	Latency        float64    `json:"latency"`
	Online24Hours  float32    `json:"24_hours_online"`
	AvgResponse    string     `json:"avg_response"`
	TotalUptime    string     `json:"uptime"`
	OrderId        int64      `json:"order_id"`
	Failures       []*Failure `json:"failures"`
	Checkins       []*Checkin `json:"checkins"`
	StopRoutine    chan bool  `json:"-"`
	LastResponse   string
	LastStatusCode int
	LastOnline     time.Time
	DnsLookup      float64 `json:"dns_lookup_time"`
}

func (s *Service) Start() {
	s.StopRoutine = make(chan bool)
}

func (s *Service) Close() {
	s.StopRoutine <- true
}
