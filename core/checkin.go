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

type Checkin struct {
	*types.Checkin
}

type CheckinHit struct {
	*types.CheckinHit
}

// String will return a Checkin API string
func (c *Checkin) String() string {
	return c.ApiKey
}

// ReturnCheckin converts *types.Checking to *core.Checkin
func ReturnCheckin(s *types.Checkin) *Checkin {
	return &Checkin{Checkin: s}
}

// ReturnCheckinHit converts *types.CheckinHit to *core.CheckinHit
func ReturnCheckinHit(h *types.CheckinHit) *CheckinHit {
	return &CheckinHit{CheckinHit: h}
}

// SelectCheckin will find a Checkin based on the API supplied
func SelectCheckin(api string) *Checkin {
	var checkin Checkin
	checkinDB().Where("api_key = ?", api).First(&checkin)
	return &checkin
}

// Period will return the duration of the Checkin interval
func (u *Checkin) Period() time.Duration {
	duration, _ := time.ParseDuration(fmt.Sprintf("%vs", u.Interval))
	return duration
}

// Grace will return the duration of the Checkin Grace Period (after service hasn't responded, wait a bit for a response)
func (u *Checkin) Grace() time.Duration {
	duration, _ := time.ParseDuration(fmt.Sprintf("%vs", u.GracePeriod))
	return duration
}

// Expected returns the duration of when the serviec should receive a checkin
func (u *Checkin) Expected() time.Duration {
	last := u.Last().CreatedAt
	now := time.Now()
	lastDir := now.Sub(last)
	sub := time.Duration(u.Period() - lastDir)
	return sub
}

// Last returns the last CheckinHit for a Checkin
func (u *Checkin) Last() CheckinHit {
	var hit CheckinHit
	checkinHitsDB().Where("checkin = ?", u.Id).Last(&hit)
	return hit
}

// Hits returns all of the CheckinHits for a given Checkin
func (u *Checkin) Hits() []CheckinHit {
	var checkins []CheckinHit
	checkinHitsDB().Where("checkin = ?", u.Id).Order("id DESC").Find(&checkins)
	return checkins
}

// Create will create a new Checkin
func (u *Checkin) Create() (int64, error) {
	u.ApiKey = utils.RandomString(7)
	row := checkinDB().Create(&u)
	if row.Error != nil {
		utils.Log(2, row.Error)
		return 0, row.Error
	}
	return u.Id, row.Error
}

// Update will update a Checkin
func (u *Checkin) Update() (int64, error) {
	row := checkinDB().Update(&u)
	if row.Error != nil {
		utils.Log(2, row.Error)
		return 0, row.Error
	}
	return u.Id, row.Error
}

// Create will create a new successful CheckinHit
func (u *CheckinHit) Create() (int64, error) {
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

// Ago returns the duration of time between now and the last successful CheckinHit
func (f *CheckinHit) Ago() string {
	got, _ := timeago.TimeAgoWithTime(time.Now(), f.CreatedAt)
	return got
}

// RecheckCheckinFailure will check if a Service Checkin has been reported yet
func (c *Checkin) RecheckCheckinFailure(guard chan struct{}) {
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
