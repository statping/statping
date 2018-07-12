package notifiers

import (
	"fmt"
	"github.com/hunterlong/statup/utils"
	"time"
)

const (
	EMAIL_ID     int64 = 1
	EMAIL_METHOD       = "email"
)

var (
	emailer    *Email
	emailArray []string
)

type Email struct {
	*Notification
}

// DEFINE YOUR NOTIFICATION HERE.
func init() {

	emailer = &Email{&Notification{
		Id:     EMAIL_ID,
		Method: EMAIL_METHOD,
		Form: []NotificationForm{{
			id:          1,
			Type:        "text",
			Title:       "SMTP Host",
			Placeholder: "Insert your SMTP Host here.",
			DbField:     "Host",
		}, {
			id:          1,
			Type:        "text",
			Title:       "SMTP Username",
			Placeholder: "Insert your SMTP Username here.",
			DbField:     "Username",
		}, {
			id:          1,
			Type:        "password",
			Title:       "SMTP Password",
			Placeholder: "Insert your SMTP Password here.",
			DbField:     "Password",
		}, {
			id:          1,
			Type:        "number",
			Title:       "SMTP Port",
			Placeholder: "Insert your SMTP Port here.",
			DbField:     "Port",
		}, {
			id:          1,
			Type:        "text",
			Title:       "Outgoing Email Address",
			Placeholder: "Insert your Outgoing Email Address",
			DbField:     "Var1",
		}, {
			id:          1,
			Type:        "number",
			Title:       "Limits per Hour",
			Placeholder: "How many emails can it send per hour",
			DbField:     "Limits",
		}},
	}}

	add(emailer)
}

// Select Obj
func (u *Email) Select() *Notification {
	return u.Notification
}

// WHEN NOTIFIER LOADS
func (u *Email) Init() error {
	//err := SendSlack("its online")

	u.Install()

	//go u.Run()
	return nil
}

// AFTER NOTIFIER LOADS, IF ENABLED, START A QUEUE PROCESS
func (u *Email) Run() error {
	//var sentAddresses []string
	//for _, email := range emailArray {
	//	if inArray(sentAddresses, email.To) || email.Sent {
	//		emailQueue = removeEmail(emailQueue, email)
	//		continue
	//	}
	//	e := email
	//	go func(email *types.Email) {
	//		err := dialSend(email)
	//		if err == nil {
	//			email.Sent = true
	//			sentAddresses = append(sentAddresses, email.To)
	//			utils.Log(1, fmt.Sprintf("Email '%v' sent to: %v using the %v template (size: %v)", email.Subject, email.To, email.Template, len([]byte(email.Source))))
	//			emailQueue = removeEmail(emailQueue, email)
	//		}
	//	}(e)
	//}
	time.Sleep(60 * time.Second)
	//if EmailComm.Enabled {
	//EmailRoutine()
	//}
	return nil
}

// ON SERVICE FAILURE, DO YOUR OWN FUNCTIONS
func (u *Email) OnFailure() error {
	if u.Enabled {
		utils.Log(1, fmt.Sprintf("Notification %v is receiving a failure notification.", u.Method))
		// Do failing stuff here!
	}
	return nil
}

// ON SERVICE SUCCESS, DO YOUR OWN FUNCTIONS
func (u *Email) OnSuccess() error {
	if u.Enabled {
		utils.Log(1, fmt.Sprintf("Notification %v is receiving a failure notification.", u.Method))
		// Do failing stuff here!
	}
	return nil
}

// ON SAVE OR UPDATE OF THE NOTIFIER FORM
func (u *Email) OnSave() error {
	utils.Log(1, fmt.Sprintf("Notification %v is receiving updated information.", u.Method))

	// Do updating stuff here

	return nil
}

// ON SERVICE FAILURE, DO YOUR OWN FUNCTIONS
func (u *Email) Install() error {
	inDb, err := emailer.Notification.isInDatabase()
	if !inDb {
		newNotifer, err := InsertDatabase(u.Notification)
		if err != nil {
			utils.Log(3, err)
			return err
		}
		utils.Log(1, fmt.Sprintf("new notifier #%v installed: %v", newNotifer, u.Method))
	}
	return err
}
