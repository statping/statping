package core

import (
	"fmt"
	"github.com/hunterlong/statup/notifications"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"time"
)

func LoadDefaultCommunications() {
	notifications.EmailComm = SelectCommunication(1)
	emailer := notifications.EmailComm
	if emailer.Enabled {
		admin, _ := SelectUser(1)
		notifications.LoadEmailer(emailer)
		email := &types.Email{
			To:       admin.Email,
			Subject:  "Test Email",
			Template: "message.html",
			Data:     nil,
			From:     emailer.Var1,
		}
		notifications.SendEmail(EmailBox, email)
		go notifications.EmailRoutine()
	}
	notifications.SlackComm = SelectCommunication(2)
	slack := notifications.SlackComm
	if slack.Enabled {
		notifications.LoadSlack(slack.Host)
		msg := fmt.Sprintf("Slack loaded on your Statup Status Page!")
		notifications.SendSlack(msg)
		go notifications.SlackRoutine()
	}
}

func LoadComms() {
	for _, c := range CoreApp.Communications {
		if c.Enabled {

		}
	}
}

func SelectAllCommunications() ([]*types.Communication, error) {
	var c []*types.Communication
	col := DbSession.Collection("communication").Find()
	err := col.OrderBy("id").All(&c)
	CoreApp.Communications = c
	return c, err
}

func Create(c *types.Communication) (int64, error) {
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

func Disable(c *types.Communication) {
	c.Enabled = false
	Update(c)
}

func Enable(c *types.Communication) {
	c.Enabled = true
	Update(c)
}

func Update(c *types.Communication) *types.Communication {
	col := DbSession.Collection("communication").Find("id", c.Id)
	col.Update(c)
	SelectAllCommunications()
	return c
}

func SelectCommunication(id int64) *types.Communication {
	var comm *types.Communication
	col := DbSession.Collection("communication").Find("id", id)
	err := col.One(&comm)
	if err != nil {
		utils.Log(2, err)
		return nil
	}
	return comm
}
