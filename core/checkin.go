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

func (c *Checkin) String() string {
	return c.ApiKey
}

func ReturnCheckin(s *types.Checkin) *Checkin {
	return &Checkin{Checkin: s}
}

func ReturnCheckinHit(h *types.CheckinHit) *CheckinHit {
	return &CheckinHit{CheckinHit: h}
}

func SelectCheckin(api string) *Checkin {
	var checkin Checkin
	checkinDB().Where("api_key = ?", api).First(&checkin)
	return &checkin
}

func (u Checkin) Period() time.Duration {
	duration, _ := time.ParseDuration(fmt.Sprintf("%vs", u.Interval))
	return duration
}

func (u Checkin) Grace() time.Duration {
	duration, _ := time.ParseDuration(fmt.Sprintf("%vs", u.GracePeriod))
	return duration
}

func (u Checkin) Expected() time.Duration {
	last := u.Last().CreatedAt
	now := time.Now()
	lastDir := now.Sub(last)
	sub := time.Duration(u.Period() - lastDir)
	return sub
}

func (u Checkin) Last() CheckinHit {
	var hit CheckinHit
	checkinHitsDB().Where("checkin = ?", u.Id).Last(&hit)
	return hit
}

func (u *Checkin) Hits() []CheckinHit {
	var checkins []CheckinHit
	checkinHitsDB().Where("checkin = ?", u.Id).Order("id DESC").Find(&checkins)
	return checkins
}

func (u *Checkin) Create() (int64, error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	u.ApiKey = utils.NewSHA1Hash(7)
	row := checkinDB().Create(u)
	if row.Error == nil {
		utils.Log(2, row.Error)
		return 0, row.Error
	}
	return u.Id, row.Error
}

func (u *CheckinHit) Create() (int64, error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	row := checkinHitsDB().Create(u)
	if row.Error == nil {
		utils.Log(2, row.Error)
		return 0, row.Error
	}
	return u.Id, row.Error
}

func SelectCheckinApi(api string) *Checkin {
	var checkin *Checkin
	checkinDB().Where("api = ?", api).Find(&checkin)
	return checkin
}

func (c *Checkin) CreateHit() (int64, error) {
	c.CreatedAt = time.Now()
	row := checkinDB().Create(c)
	if row.Error == nil {
		utils.Log(2, row.Error)
		return 0, row.Error
	}
	return c.Id, row.Error
}

func (c *Checkin) RecheckCheckinFailure(guard chan struct{}) {
	between := time.Now().Sub(time.Now()).Seconds()
	if between > float64(c.Interval) {
		fmt.Println("rechecking every 15 seconds!")
		c.CreateFailure()
		time.Sleep(15 * time.Second)
		guard <- struct{}{}
		c.RecheckCheckinFailure(guard)
	} else {
		fmt.Println("i recovered!!")
	}
	<-guard
}

func (f *Checkin) CreateFailure() {

}

func (f *Checkin) Ago() string {
	got, _ := timeago.TimeAgoWithTime(time.Now(), time.Now())
	return got
}

func (f *CheckinHit) Ago() string {
	got, _ := timeago.TimeAgoWithTime(time.Now(), time.Now())
	return got
}
