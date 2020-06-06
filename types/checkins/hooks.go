package checkins

import "github.com/statping/statping/utils"

func (c *Checkin) BeforeCreate() error {
	if c.ApiKey == "" {
		c.ApiKey = utils.RandomString(32)
	}
	return nil
}

func (c *Checkin) AfterCreate() error {
	c.Start()
	go c.CheckinRoutine()
	return nil
}

func (c *Checkin) BeforeDelete() error {
	c.Close()
	q := dbHits.Where("checkin = ?", c.Id).Delete(&CheckinHit{})
	if err := q.Error(); err != nil {
		return err
	}
	return nil
}
