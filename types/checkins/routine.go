package checkins

import (
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/utils"
	"time"
)

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

// Routine for checking if the last Checkin was within its interval
func (c *Checkin) CheckinRoutine() {
	lastHit := c.LastHit()
	if lastHit == nil {
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
			if c.Expected() <= 0 {
				issue := fmt.Sprintf("Checkin %v is failing, no request since %v", c.Name, lastHit.CreatedAt)
				log.Errorln(issue)

				fail := &failures.Failure{
					Issue:     issue,
					Method:    "checkin",
					Service:   c.ServiceId,
					Checkin:   c.Id,
					PingTime:  c.Expected().Milliseconds(),
					CreatedAt: time.Time{},
				}

				c.CreateFailure(fail)
			}
			reCheck = c.Period()
		}
		continue
	}
}
