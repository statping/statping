// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
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
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"sort"
	"time"
)

type Checkin struct {
	*types.Checkin
}

type CheckinHit struct {
	*types.CheckinHit
}

// Select returns a *types.Checkin
func (c *Checkin) Select() *types.Checkin {
	return c.Checkin
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
			log.Infoln(fmt.Sprintf("Stopping checkin routine: %v", c.Name))
			c.Failing = false
			break CheckinLoop
		case <-time.After(reCheck):
			log.Infoln(fmt.Sprintf("Checkin %v is expected at %v, checking every %v", c.Name, utils.FormatDuration(c.Expected()), utils.FormatDuration(c.Period())))
			if c.Expected().Seconds() <= 0 {
				issue := fmt.Sprintf("Checkin %v is failing, no request since %v", c.Name, c.Last().CreatedAt)
				log.Errorln(issue)
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
	return SelectService(c.ServiceId)
}

func (c *Checkin) CreateFailure() (int64, error) {
	service := c.Service()
	c.Failing = true
	fail := &Failure{&types.Failure{
		Issue:    fmt.Sprintf("Checkin %v was not reported %v ago, it expects a request every %v", c.Name, utils.FormatDuration(c.Expected()), utils.FormatDuration(c.Period())),
		Method:   "checkin",
		MethodId: c.Id,
		Service:  service.Id,
		Checkin:  c.Id,
		PingTime: c.Expected().Seconds(),
	}}
	row := failuresDB().Create(&fail)
	sort.Sort(types.FailSort(c.Failures))
	c.Failures = append(c.Failures, fail)
	if len(c.Failures) > limitedFailures {
		c.Failures = c.Failures[1:]
	}
	return fail.Id, row.Error
}

// LimitedHits will return the last amount of successful hits from a checkin
func (c *Checkin) LimitedHits(amount int64) []*types.CheckinHit {
	var hits []*types.CheckinHit
	checkinHitsDB().Where("checkin = ?", c.Id).Order("id desc").Limit(amount).Find(&hits)
	return hits
}

// AllCheckins returns all checkin in system
func AllCheckins() []*Checkin {
	var checkins []*Checkin
	checkinDB().Find(&checkins)
	return checkins
}

// SelectCheckin will find a Checkin based on the API supplied
func SelectCheckin(api string) *Checkin {
	for _, s := range Services() {
		for _, c := range s.Select().Checkins {
			if c.Select().ApiKey == api {
				return c.(*Checkin)
			}
		}
	}
	return nil
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
	now := utils.Now()
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
func (c *Checkin) LimitedFailures(amount int64) []types.FailureInterface {
	var failures []*Failure
	var failInterfaces []types.FailureInterface
	col := failuresDB().Where("checkin = ?", c.Id).Where("method = 'checkin'").Limit(amount).Order("id desc")
	col.Find(&failures)
	for _, f := range failures {
		failInterfaces = append(failInterfaces, f)
	}
	return failInterfaces
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
	i := c.index()
	service := c.Service()
	slice := service.Checkins
	service.Checkins = append(slice[:i], slice[i+1:]...)
	row := checkinDB().Delete(&c)
	return row.Error
}

// index returns a checkin index int for updating the *checkin.Service slice
func (c *Checkin) index() int {
	for k, checkin := range c.Service().Checkins {
		if c.Id == checkin.Select().Id {
			return k
		}
	}
	return 0
}

// Create will create a new Checkin
func (c *Checkin) Create() (int64, error) {
	c.ApiKey = utils.RandomString(7)
	row := checkinDB().Create(&c)
	if row.Error != nil {
		log.Warnln(row.Error)
		return 0, row.Error
	}
	service := SelectService(c.ServiceId)
	service.Checkins = append(service.Checkins, c)
	c.Start()
	go c.Routine()
	return c.Id, row.Error
}

// Update will update a Checkin
func (c *Checkin) Update() (int64, error) {
	row := checkinDB().Update(&c)
	if row.Error != nil {
		log.Warnln(row.Error)
		return 0, row.Error
	}
	return c.Id, row.Error
}

// Create will create a new successful checkinHit
func (c *CheckinHit) Create() (int64, error) {
	if c.CreatedAt.IsZero() {
		c.CreatedAt = utils.Now()
	}
	row := checkinHitsDB().Create(&c)
	if row.Error != nil {
		log.Warnln(row.Error)
		return 0, row.Error
	}
	return c.Id, row.Error
}

// Ago returns the duration of time between now and the last successful checkinHit
func (c *CheckinHit) Ago() string {
	got, _ := timeago.TimeAgoWithTime(utils.Now(), c.CreatedAt)
	return got
}

// RecheckCheckinFailure will check if a Service Checkin has been reported yet
func (c *Checkin) RecheckCheckinFailure(guard chan struct{}) {
	between := utils.Now().Sub(utils.Now()).Seconds()
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
