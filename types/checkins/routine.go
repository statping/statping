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
	reCheck := c.Period()

CheckinLoop:
	for {
		select {
		case <-c.Running:
			log.Infoln(fmt.Sprintf("Stopping checkin routine: %s", c.Name))
			c.Failing = false
			break CheckinLoop
		case <-time.After(reCheck):
			lastHit := c.LastHit()
			ago := utils.Now().Sub(lastHit.CreatedAt)

			log.Infoln(fmt.Sprintf("Checkin '%s' expects a request every %s last request was %s ago", c.Name, c.Period(), utils.DurationReadable(ago)))

			if ago.Seconds() > c.Period().Seconds() {
				issue := fmt.Sprintf("Checkin expects a request every %d seconds", c.Interval)
				log.Warnln(issue)

				fail := &failures.Failure{
					Issue:    issue,
					Method:   "checkin",
					Service:  c.ServiceId,
					PingTime: ago.Milliseconds(),
				}

				if err := c.CreateFailure(fail); err != nil {
					log.Errorln(err)
				}
			}
			reCheck = c.Period()
		}
	}
}
