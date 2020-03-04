package checkins

func (c *Checkin) LastHit() *CheckinHit {
	var hit *CheckinHit
	DBhits().Where("checkin = ?", c.Id).Last(&hit)
	return hit
}

func (c *Checkin) Hits() []*CheckinHit {
	var hits []*CheckinHit
	DBhits().Where("checkin = ?", c.Id).Find(&hits)
	c.AllHits = hits
	return hits
}

func (c *CheckinHit) Create() error {
	db := DBhits().Create(&c)
	return db.Error()
}

func (c *CheckinHit) Update() error {
	db := DBhits().Update(&c)
	return db.Error()
}

func (c *CheckinHit) Delete() error {
	db := DBhits().Delete(&c)
	return db.Error()
}
