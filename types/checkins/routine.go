package checkins

import (
	"fmt"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/utils"
	"time"
)

var log = utils.Log.WithField("type", "checkin")

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

// checkinRoutine for checking if the last Checkin was within its interval
func (c *Checkin) checkinRoutine() {
	lastHit := c.LastHit()
	if lastHit == nil {
		return
	}
	reCheck := c.Period()
CheckinLoop:
	for {
		select {
		case <-c.Running:
			log.Infoln(fmt.Sprintf("Stopping checkin routine: %s", c.Name))
			c.Failing = false
			break CheckinLoop
		case <-time.After(reCheck):
			log.Infoln(fmt.Sprintf("Checkin '%s' expects a request every %v", c.Name, utils.FormatDuration(c.Period())))
			if c.Expected() <= 0 {
				issue := fmt.Sprintf("Checkin '%s' is failing, no request since %v", c.Name, lastHit.CreatedAt)
				//log.Errorln(issue)

				fail := &failures.Failure{
					Issue:    issue,
					Method:   "checkin",
					Service:  c.ServiceId,
					Checkin:  c.Id,
					PingTime: c.Expected().Milliseconds(),
				}

				c.CreateFailure(fail)
			}
			reCheck = c.Period()
		}
	}
}
