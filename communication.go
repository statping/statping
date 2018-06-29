package main

import (
	"github.com/hunterlong/statup/log"
	"github.com/hunterlong/statup/types"
	"time"
)

func LoadDefaultCommunications() {
	emailer := SelectCommunication(1)
	if emailer.Enabled {
		LoadMailer(emailer)
		go EmailerQueue()
	}
}

func LoadComms() {
	for _, c := range core.Communications {
		if c.Enabled {

		}
	}
}

func Run(c *types.Communication) {

	sample := &types.Email{
		To:      "info@socialeck.com",
		Subject: "Test Email from Statup",
	}

	AddEmail(sample)
}

func SelectAllCommunications() ([]*types.Communication, error) {
	var c []*types.Communication
	col := dbSession.Collection("communication").Find()
	err := col.All(&c)
	core.Communications = c
	return c, err
}

func Create(c *types.Communication) (int64, error) {
	c.CreatedAt = time.Now()
	uuid, err := dbSession.Collection("communication").Insert(c)
	if err != nil {
		log.Send(3, err)
	}
	if uuid == nil {
		log.Send(2, err)
		return 0, err
	}
	c.Id = uuid.(int64)
	if core != nil {
		core.Communications = append(core.Communications, c)
	}
	return uuid.(int64), err
}

func Disable(c *types.Communication) {
	c.Enabled = false
	Update(c)
}

func Enable(c *types.Communication) {
	c.Enabled = true
	Update(c)
}

func Update(c *types.Communication) *types.Communication {
	col := dbSession.Collection("communication").Find("id", c.Id)
	col.Update(c)
	return c
}

func SelectCommunication(id int64) *types.Communication {
	for _, c := range core.Communications {
		if c.Id == id {
			return c
		}
	}
	return nil
}
