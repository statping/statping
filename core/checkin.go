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
	"fmt"
	"github.com/ararog/timeago"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"time"
)

type checkin struct {
	*types.Checkin
}

type checkinHit struct {
	*types.CheckinHit
}

// String will return a checkin API string
func (c *checkin) String() string {
	return c.ApiKey
}

// ReturnCheckin converts *types.Checking to *core.checkin
func ReturnCheckin(s *types.Checkin) *checkin {
	return &checkin{Checkin: s}
}

// ReturnCheckinHit converts *types.checkinHit to *core.checkinHit
func ReturnCheckinHit(h *types.CheckinHit) *checkinHit {
	return &checkinHit{CheckinHit: h}
}

// SelectCheckin will find a checkin based on the API supplied
func SelectCheckin(api string) *checkin {
	var checkin checkin
	checkinDB().Where("api_key = ?", api).First(&checkin)
	return &checkin
}

// Period will return the duration of the checkin interval
func (c *checkin) Period() time.Duration {
	duration, _ := time.ParseDuration(fmt.Sprintf("%vs", c.Interval))
	return duration
}

// Grace will return the duration of the checkin Grace Period (after service hasn't responded, wait a bit for a response)
func (c *checkin) Grace() time.Duration {
	duration, _ := time.ParseDuration(fmt.Sprintf("%vs", c.GracePeriod))
	return duration
}

// Expected returns the duration of when the serviec should receive a checkin
func (c *checkin) Expected() time.Duration {
	last := c.Last().CreatedAt
	now := time.Now()
	lastDir := now.Sub(last)
	sub := time.Duration(c.Period() - lastDir)
	return sub
}

// Last returns the last checkinHit for a checkin
func (c *checkin) Last() checkinHit {
	var hit checkinHit
	checkinHitsDB().Where("checkin = ?", c.Id).Last(&hit)
	return hit
}

// Hits returns all of the CheckinHits for a given checkin
func (c *checkin) Hits() []checkinHit {
	var checkins []checkinHit
	checkinHitsDB().Where("checkin = ?", c.Id).Order("id DESC").Find(&checkins)
	return checkins
}

// Create will create a new checkin
func (c *checkin) Create() (int64, error) {
	c.ApiKey = utils.RandomString(7)
	row := checkinDB().Create(&c)
	if row.Error != nil {
		utils.Log(2, row.Error)
		return 0, row.Error
	}
	return c.Id, row.Error
}

// Update will update a checkin
func (u *checkin) Update() (int64, error) {
	row := checkinDB().Update(&u)
	if row.Error != nil {
		utils.Log(2, row.Error)
		return 0, row.Error
	}
	return u.Id, row.Error
}

// Create will create a new successful checkinHit
func (u *checkinHit) Create() (int64, error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	row := checkinHitsDB().Create(u)
	if row.Error != nil {
		utils.Log(2, row.Error)
		return 0, row.Error
	}
	return u.Id, row.Error
}

// Ago returns the duration of time between now and the last successful checkinHit
func (f *checkinHit) Ago() string {
	got, _ := timeago.TimeAgoWithTime(time.Now(), f.CreatedAt)
	return got
}

// RecheckCheckinFailure will check if a Service checkin has been reported yet
func (c *checkin) RecheckCheckinFailure(guard chan struct{}) {
	between := time.Now().Sub(time.Now()).Seconds()
	if between > float64(c.Interval) {
		fmt.Println("rechecking every 15 seconds!")
		time.Sleep(15 * time.Second)
		guard <- struct{}{}
		c.RecheckCheckinFailure(guard)
	} else {
		fmt.Println("i recovered!!")
	}
	<-guard
}
