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

// Routine for checking if the last Checkin was within its interval
func (c *Checkin) Routine() {
	if c.Last() == nil {
		return
	}
	reCheck := c.Period()
CheckinLoop:
	for {
		select {
		case <-c.Running:
			utils.Log(1, fmt.Sprintf("Stopping checkin routine: %v", c.Name))
			break CheckinLoop
		case <-time.After(reCheck):
			utils.Log(1, fmt.Sprintf("Checkin %v is expected at %v, checking every %v", c.Name, utils.FormatDuration(c.Expected()), utils.FormatDuration(c.Period())))
			if c.Expected().Seconds() <= 0 {
				issue := fmt.Sprintf("Checkin %v is failing, no request since %v", c.Name, c.Last().CreatedAt)
				utils.Log(3, issue)
				c.Service()
				c.CreateFailure()
			}
			reCheck = c.Period()
		}
		continue
	}
}

// String will return a Checkin API string
func (c *Checkin) String() string {
	return c.ApiKey
}

// ReturnCheckin converts *types.Checking to *core.Checkin
func ReturnCheckin(c *types.Checkin) *Checkin {
	return &Checkin{Checkin: c}
}

// ReturnCheckinHit converts *types.checkinHit to *core.checkinHit
func ReturnCheckinHit(c *types.CheckinHit) *CheckinHit {
	return &CheckinHit{CheckinHit: c}
}

func (c *Checkin) Service() *Service {
	service := SelectService(c.ServiceId)
	return service
}

func (c *Checkin) CreateFailure() (int64, error) {
	service := c.Service()
	fail := &types.Failure{
		Issue:    fmt.Sprintf("Checkin %v was not reported %v ago, it expects a request every %v", c.Name, utils.FormatDuration(c.Expected()), utils.FormatDuration(c.Period())),
		Method:   "checkin",
		MethodId: c.Id,
		Service:  service.Id,
		PingTime: c.Expected().Seconds() * 0.001,
	}
	row := failuresDB().Create(&fail)
	return fail.Id, row.Error
}

// AllCheckins returns all checkin in system
func AllCheckins() []*Checkin {
	var checkins []*Checkin
	checkinDB().Find(&checkins)
	return checkins
}

// SelectCheckin will find a Checkin based on the API supplied
func SelectCheckin(api string) *Checkin {
	var checkin Checkin
	checkinDB().Where("api_key = ?", api).First(&checkin)
	return &checkin
}

// SelectCheckin will find a Checkin based on the API supplied
func SelectCheckinId(id int64) *Checkin {
	var checkin Checkin
	checkinDB().Where("id = ?", id).First(&checkin)
	return &checkin
}

// Period will return the duration of the Checkin interval
func (c *Checkin) Period() time.Duration {
	duration, _ := time.ParseDuration(fmt.Sprintf("%vs", c.Interval))
	return duration
}

// Grace will return the duration of the Checkin Grace Period (after service hasn't responded, wait a bit for a response)
func (c *Checkin) Grace() time.Duration {
	duration, _ := time.ParseDuration(fmt.Sprintf("%vs", c.GracePeriod))
	return duration
}

// Expected returns the duration of when the serviec should receive a Checkin
func (c *Checkin) Expected() time.Duration {
	last := c.Last().CreatedAt
	now := time.Now()
	lastDir := now.Sub(last)
	sub := time.Duration(c.Period() - lastDir)
	return sub
}

// Last returns the last checkinHit for a Checkin
func (c *Checkin) Last() *CheckinHit {
	var hit CheckinHit
	checkinHitsDB().Where("checkin = ?", c.Id).Last(&hit)
	return &hit
}

func (c *Checkin) Link() string {
	return fmt.Sprintf("%v/checkin/%v", CoreApp.Domain, c.ApiKey)
}

// AllHits returns all of the CheckinHits for a given Checkin
func (c *Checkin) AllHits() []*types.CheckinHit {
	var checkins []*types.CheckinHit
	checkinHitsDB().Where("checkin = ?", c.Id).Order("id DESC").Find(&checkins)
	return checkins
}

// Hits returns all of the CheckinHits for a given Checkin
func (c *Checkin) AllFailures() []*types.Failure {
	var failures []*types.Failure
	col := failuresDB().Where("checkin = ?", c.Id).Where("method = 'checkin'").Order("id desc")
	col.Find(&failures)
	return failures
}

// Create will create a new Checkin
func (c *Checkin) Delete() error {
	c.Close()
	row := checkinDB().Delete(&c)
	return row.Error
}

// Create will create a new Checkin
func (c *Checkin) Create() (int64, error) {
	c.ApiKey = utils.RandomString(7)
	row := checkinDB().Create(&c)
	c.Start()
	go c.Routine()
	if row.Error != nil {
		utils.Log(2, row.Error)
		return 0, row.Error
	}
	return c.Id, row.Error
}

// Update will update a Checkin
func (c *Checkin) Update() (int64, error) {
	row := checkinDB().Update(c)
	if row.Error != nil {
		utils.Log(2, row.Error)
		return 0, row.Error
	}
	return c.Id, row.Error
}

// Create will create a new successful checkinHit
func (c *CheckinHit) Create() (int64, error) {
	if c.CreatedAt.IsZero() {
		c.CreatedAt = time.Now()
	}
	row := checkinHitsDB().Create(c)
	if row.Error != nil {
		utils.Log(2, row.Error)
		return 0, row.Error
	}
	return c.Id, row.Error
}

// Ago returns the duration of time between now and the last successful checkinHit
func (c *CheckinHit) Ago() string {
	got, _ := timeago.TimeAgoWithTime(time.Now(), c.CreatedAt)
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
