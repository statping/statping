package core

import (
	"github.com/hunterlong/statup/notifiers"
	"github.com/hunterlong/statup/utils"
)

func LoadDefaultCommunications() {
	//communications.EmailComm = SelectCommunication(1)
	//emailer := communications.EmailComm
	//if emailer.Enabled {
	//	admin, _ := SelectUser(1)
	//	communications.LoadEmailer(emailer)
	//	email := &types.Email{
	//		To:       admin.Email,
	//		Subject:  "Test Email",
	//		Template: "message.html",
	//		Data:     nil,
	//		From:     emailer.Var1,
	//	}
	//	communications.SendEmail(EmailBox, email)
	//	go communications.EmailRoutine()
	//}
	//communications.SlackComm = SelectCommunication(2)
	//slack := communications.SlackComm
	//if slack.Enabled {
	//	communications.LoadSlack(slack.Host)
	//	msg := fmt.Sprintf("Slack loaded on your Statup Status Page!")
	//	communications.SendSlack(msg)
	//	go communications.SlackRoutine()
	//}
}

func LoadComms() {
	for _, c := range CoreApp.Communications {
		if c.Enabled {

		}
	}
}

func SelectAllCommunications() ([]*notifiers.Notification, error) {
	var c []*notifiers.Notification
	col := DbSession.Collection("communication").Find()
	err := col.OrderBy("id").All(&c)
	//CoreApp.Communications = c
	//communications.LoadComms(c)
	return c, err
}

func Create(c *notifiers.Notification) (int64, error) {
	//c.CreatedAt = time.Now()
	//uuid, err := DbSession.Collection("communication").Insert(c)
	//if err != nil {
	//	utils.Log(3, err)
	//}
	//if uuid == nil {
	//	utils.Log(2, err)
	//	return 0, err
	//}
	//c.Id = uuid.(int64)
	//c.Routine = make(chan struct{})
	//if CoreApp != nil {
	//	CoreApp.Communications = append(CoreApp.Communications, c.Communicator)
	//}
	//return uuid.(int64), err
	return 0, nil
}

func Disable(c *notifiers.Notification) {
	c.Enabled = false
	Update(c)
}

func Enable(c *notifiers.Notification) {
	c.Enabled = true
	Update(c)
}

func Update(c *notifiers.Notification) *notifiers.Notification {
	col := DbSession.Collection("communication").Find("id", c.Id)
	col.Update(c)
	SelectAllCommunications()
	return c
}

func SelectCommunication(id int64) *notifiers.Notification {
	var comm *notifiers.Notification
	col := DbSession.Collection("communication").Find("id", id)
	err := col.One(&comm)
	if err != nil {
		utils.Log(2, err)
		return nil
	}
	return comm
}
