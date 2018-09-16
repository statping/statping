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
)

type Hit struct {
	*types.Hit
}

// CreateHit will create a new 'hit' record in the database for a successful/online service
func (s *Service) CreateHit(h *types.Hit) (int64, error) {
	h.CreatedAt = time.Now().UTC()
	db := hitsDB().Create(h)
	if db.Error != nil {
		utils.Log(2, db.Error)
		return 0, db.Error
	}
	return h.Id, db.Error
}

// Hits returns all successful hits for a service
func (s *Service) Hits() ([]*types.Hit, error) {
	var hits []*types.Hit
	col := hitsDB().Where("service = ?", s.Id).Order("id desc")
	err := col.Find(&hits)
	return hits, err.Error
}

// LimitedHits returns the last 1024 successful/online 'hit' records for a service
func (s *Service) LimitedHits() ([]*types.Hit, error) {
	var hits []*types.Hit
	col := hitsDB().Where("service = ?", s.Id).Order("id desc").Limit(1024)
	err := col.Find(&hits)
	return reverseHits(hits), err.Error
}

// reverseHits will reverse the service's hit slice
func reverseHits(input []*types.Hit) []*types.Hit {
	if len(input) == 0 {
		return input
	}
	return append(reverseHits(input[1:]), input[0])
}

// SelectHitsGroupBy returns all hits from the group by function
func (s *Service) SelectHitsGroupBy(group string) ([]*types.Hit, error) {
	var hits []*types.Hit
	col := hitsDB().Where("service = ?", s.Id)
	err := col.Find(&hits)
	return hits, err.Error
}

// TotalHits returns the total amount of successful hits a service has
func (s *Service) TotalHits() (uint64, error) {
	var count uint64
	col := hitsDB().Where("service = ?", s.Id)
	err := col.Count(&count)
	return count, err.Error
}

// TotalHitsSince returns the total amount of hits based on a specific time/date
func (s *Service) TotalHitsSince(ago time.Time) (uint64, error) {
	var count uint64
	rows := hitsDB().Where("service = ? AND created_at > ?", s.Id, ago.Format("2006-01-02 15:04:05"))
	err := rows.Count(&count)
	return count, err.Error
}

// Sum returns the added value Latency for all of the services successful hits.
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
