package checkins

import (
	"fmt"
	"github.com/statping/statping/utils"
	"time"
)

func (c *Checkin) Expected() time.Duration {
	last := c.LastHit()
	now := utils.Now()
	lastDir := now.Sub(last.CreatedAt)
	return c.Period() - lastDir
}

func (c *Checkin) Period() time.Duration {
	return time.Duration(c.Interval) * time.Minute
}

// Start will create a channel for the checkin checking go routine
func (c *Checkin) Start() {
	log.Infoln(fmt.Sprintf("Starting checkin routine: %s", c.Name))
	c.Running = make(chan bool)
	go c.checkinRoutine()
}

// Close will stop the checkin routine
func (c *Checkin) Close() {
	if c.Running != nil {
		close(c.Running)
	}
}

// IsRunning returns true if the checkin go routine is running
func (c *Checkin) IsRunning() bool {
	if c.Running == nil {
		return false
	}
	select {
	case <-c.Running:
		return false
	default:
		return true
	}
}
