package checkins

func (c *Checkin) LastHit() *CheckinHit {
	var hit CheckinHit
	dbHits.Where("checkin = ?", c.Id).Last(&hit)
	return &hit
}

func (c *Checkin) Hits() []*CheckinHit {
	var hits []*CheckinHit
	dbHits.Where("checkin = ?", c.Id).Order("id DESC").Find(&hits)
	c.AllHits = hits
	return hits
}

func (c *CheckinHit) Create() error {
	q := dbHits.Create(c)
	return q.Error()
}

func (c *CheckinHit) Update() error {
	q := dbHits.Update(c)
	return q.Error()
}

func (c *CheckinHit) Delete() error {
	q := dbHits.Delete(c)
	return q.Error()
}
