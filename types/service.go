// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package types

import (
	"time"
)

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
	Failures       []*Failure `json:"failures"`
	Checkins       []*Checkin `json:"checkins"`
	Running        chan bool  `json:"-"`
	Checkpoint     time.Time  `json:"-"`
	LastResponse   string
	LastStatusCode int
	LastOnline     time.Time
	DnsLookup      float64 `json:"dns_lookup_time"`
	ServiceInterface
}

type ServiceInterface interface {
	// Database functions
	Create() (int64, error)
	Update() error
	Delete() error
	// Basic Method functions
	AvgTime() float64
	Online24() float32
	SmallText() string
	GraphData() string
	AvgUptime() string
	ToJSON() string
	// Failure functions
	CreateFailure(*Failure) (int64, error)
	LimitedFailures() []*Failure
	AllFailures() []*Failure
	TotalFailures() (uint64, error)
	TotalFailures24Hours() (uint64, error)
	// Hits functions (successful responses)
	CreateHit(*Hit) (int64, error)
	Hits() ([]*Hit, error)
	TotalHits() (uint64, error)
	Sum() (float64, error)
	LimitedHits() ([]*Hit, error)
	SelectHitsGroupBy(string) ([]*Hit, error)
	// Go Routines
	CheckQueue(bool)
	Check(bool) *Service
	checkHttp(bool) *Service
	checkTcp(bool) *Service
	// Checkin functions
	AllCheckins() []*Checkin
}

func (s *Service) Start() {
	if s.Running == nil {
		s.Running = make(chan bool)
	}
}

func (s *Service) Close() {
	if s.IsRunning() {
		close(s.Running)
	}
}

func (s *Service) IsRunning() bool {
	if s.Running == nil {
		return false
	}
	select {
	case <-s.Running:
		return false
	default:
		return true
	}
}
