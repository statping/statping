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
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"time"
)

type Checkin struct {
	*database.CheckinObj
}

type CheckinHit struct {
	*types.CheckinHit
}

// Select returns a *types.Checkin
func (c *Checkin) Select() *types.Checkin {
	return c.Checkin
}

// Routine for checking if the last Checkin was within its interval
func CheckinRoutine(checkin database.Checkiner) {
	c := checkin.Object()
	if c.Hits().Last() == nil {
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
				issue := fmt.Sprintf("Checkin %v is failing, no request since %v", c.Name, c.Hits().Last().CreatedAt)
				log.Errorln(issue)

				CreateCheckinFailure(c)
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

func CreateCheckinFailure(checkin database.Checkiner) (int64, error) {
	c := checkin.Object()
	service := c.Service()
	c.Failing = true
	fail := &types.Failure{
		Issue:    fmt.Sprintf("Checkin %v was not reported %v ago, it expects a request every %v", c.Name, utils.FormatDuration(c.Expected()), utils.FormatDuration(c.Period())),
		Method:   "checkin",
		MethodId: c.Id,
		Service:  service.Id,
		Checkin:  c.Id,
		PingTime: c.Expected().Seconds(),
	}
	_, err := database.Create(fail)
	if err != nil {
		return 0, err
	}
	//sort.Sort(types.FailSort(c.Failures()))
	return fail.Id, err
}

// AllCheckins returns all checkin in system
func AllCheckins() []*database.CheckinObj {
	checkins := database.AllCheckins()
	return checkins
}

// SelectCheckin will find a Checkin based on the API supplied
func SelectCheckin(api string) *Checkin {
	for _, s := range Services() {
		for _, c := range s.AllCheckins() {
			if c.ApiKey == api {
				return &Checkin{c}
			}
		}
	}
	return nil
}

// AllHits returns all of the CheckinHits for a given Checkin
func (c *Checkin) AllHits() []*types.CheckinHit {
	var checkins []*types.CheckinHit
	Database(&types.CheckinHit{}).Where("checkin = ?", c.Id).Order("id DESC").Find(&checkins)
	return checkins
}

// Hits returns all of the CheckinHits for a given Checkin
func (c *Checkin) AllFailures() []*types.Failure {
	var failures []*types.Failure
	Database(&types.Failure{}).
		Where("checkin = ?", c.Id).
		Where("method = 'checkin'").
		Order("id desc").
		Find(&failures)

	return failures
}

func (c *Checkin) GetFailures(count int) []*types.Failure {
	var failures []*types.Failure
	Database(&types.Failure{}).
		Where("checkin = ?", c.Id).
		Where("method = 'checkin'").
		Limit(count).
		Order("id desc").
		Find(&failures)

	return failures
}

// Create will create a new Checkin
func (c *Checkin) Delete() error {
	c.Close()
	i := c.index()
	service := c.Service()
	slice := service.Checkins
	service.Checkins = append(slice[:i], slice[i+1:]...)
	row := Database(c).Delete(&c)
	return row.Error()
}

// index returns a checkin index int for updating the *checkin.Service slice
func (c *Checkin) index() int {
	for k, checkin := range c.Service().Checkins {
		if c.Id == checkin.Model().Id {
			return k
		}
	}
	return 0
}

// Create will create a new Checkin
func (c *Checkin) Create() (int64, error) {
	c.ApiKey = utils.RandomString(7)
	_, err := database.Create(c)
	if err != nil {
		log.Warnln(err)
		return 0, err
	}
	service := SelectService(c.ServiceId)
	service.Checkins = append(service.Checkins, c)
	c.Start()
	go CheckinRoutine(c)
	return c.Id, err
}

// Update will update a Checkin
func (c *Checkin) Update() (int64, error) {
	row := Database(c).Update(&c)
	if row.Error() != nil {
		log.Warnln(row.Error())
		return 0, row.Error()
	}
	return c.Id, row.Error()
}

// Create will create a new successful checkinHit
func (c *CheckinHit) Create() (int64, error) {
	if c.CreatedAt.IsZero() {
		c.CreatedAt = utils.Now()
	}
	row := Database(c).Create(&c)
	if row.Error() != nil {
		log.Warnln(row.Error())
		return 0, row.Error()
	}
	return c.Id, row.Error()
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
