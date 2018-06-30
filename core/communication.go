package core

import (
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"time"
)

type Communication types.Communication

func LoadDefaultCommunications() {
	emailer := SelectCommunication(1)
	if emailer.Enabled {
		//LoadMailer(emailer)
		//go EmailerQueue()
	}
}

func LoadComms() {
	for _, c := range CoreApp.Communications {
		if c.Enabled {

		}
	}
}

func Run(c *Communication) {

	//sample := &Email{
	//	To:      "info@socialeck.com",
	//	Subject: "Test Email from Statup",
	//}

	//AddEmail(sample)
}

func SelectAllCommunications() ([]*Communication, error) {
	var c []*Communication
	col := DbSession.Collection("communication").Find()
	err := col.All(&c)
	CoreApp.Communications = c
	return c, err
}

func Create(c *Communication) (int64, error) {
	c.CreatedAt = time.Now()
	uuid, err := DbSession.Collection("communication").Insert(c)
	if err != nil {
		utils.Log(3, err)
	}
	if uuid == nil {
		utils.Log(2, err)
		return 0, err
	}
	c.Id = uuid.(int64)
	if CoreApp != nil {
		CoreApp.Communications = append(CoreApp.Communications, c)
	}
	return uuid.(int64), err
}

func Disable(c *Communication) {
	c.Enabled = false
	Update(c)
}

func Enable(c *Communication) {
	c.Enabled = true
	Update(c)
}

func Update(c *Communication) *Communication {
	col := DbSession.Collection("communication").Find("id", c.Id)
	col.Update(c)
	return c
}

func SelectCommunication(id int64) *Communication {
	for _, c := range CoreApp.Communications {
		if c.Id == id {
			return c
		}
	}
	return nil
}
