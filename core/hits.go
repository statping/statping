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

func CreateServiceHit(s *Service, d HitData) (int64, error) {
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

func (s *Service) Hits() ([]Hit, error) {
	var hits []Hit
	col := hitCol().Find("service", s.Id).OrderBy("-id")
	err := col.All(&hits)
	return hits, err
}

func (s *Service) LimitedHits() ([]*Hit, error) {
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
	if err != nil {
		utils.Log(2, err)
	}
	for _, h := range hits {
		amount += h.Latency
	}
	return amount, err
}
