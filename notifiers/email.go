package notifiers

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"gopkg.in/gomail.v2"
	"html/template"
	"time"
)

const (
	EMAIL_ID     int64 = 1
	EMAIL_METHOD       = "email"
)

var (
	emailer    *Email
	emailArray []string
	emailQueue []*types.Email
)

type Email struct {
	*Notification
	mailer *gomail.Dialer
}

// DEFINE YOUR NOTIFICATION HERE.
func init() {

	emailer = &Email{
		Notification: &Notification{
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
	err := u.Install()

	if err == nil {
		notifier, _ := SelectNotification(u.Id)
		forms := u.Form
		u.Notification = notifier
		u.Form = forms
		if u.Enabled {

			utils.Log(1, fmt.Sprintf("Loading SMTP Emailer using host: %v:%v", u.Notification.Host, u.Notification.Port))
			u.mailer = gomail.NewDialer(u.Notification.Host, u.Notification.Port, u.Notification.Username, u.Notification.Password)
			u.mailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

			go u.Run()
		}
	}

	//go u.Run()
	return nil
}

func (u *Email) Test() error {
	//email := &types.Email{
	//	To:       "info@socialeck.com",
	//	Subject:  "Test Email",
	//	Template: "message.html",
	//	Data:     nil,
	//	From:     emailer.Var1,
	//}
	//SendEmail(core.EmailBox, email)
	return nil
}

// AFTER NOTIFIER LOADS, IF ENABLED, START A QUEUE PROCESS
func (u *Email) Run() error {
	var sentAddresses []string
	for _, email := range emailQueue {
		if inArray(sentAddresses, email.To) || email.Sent {
			emailQueue = removeEmail(emailQueue, email)
			continue
		}
		e := email
		go func(email *types.Email) {
			err := u.dialSend(email)
			if err == nil {
				email.Sent = true
				sentAddresses = append(sentAddresses, email.To)
				utils.Log(1, fmt.Sprintf("Email '%v' sent to: %v using the %v template (size: %v)", email.Subject, email.To, email.Template, len([]byte(email.Source))))
				emailQueue = removeEmail(emailQueue, email)
			}
		}(e)
	}
	time.Sleep(60 * time.Second)
	if u.Enabled {
		return u.Run()
	}
	return nil
}

// ON SERVICE FAILURE, DO YOUR OWN FUNCTIONS
func (u *Email) OnFailure(data map[string]interface{}) error {
	if u.Enabled {
		utils.Log(1, fmt.Sprintf("Notification %v is receiving a failure notification.", u.Method))
		// Do failing stuff here!
	}
	return nil
}

// ON SERVICE SUCCESS, DO YOUR OWN FUNCTIONS
func (u *Email) OnSuccess(data map[string]interface{}) error {
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

func (u *Email) dialSend(email *types.Email) error {
	m := gomail.NewMessage()
	m.SetHeader("From", email.From)
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/html", email.Source)
	if err := u.mailer.DialAndSend(m); err != nil {
		utils.Log(3, fmt.Sprintf("Email '%v' sent to: %v using the %v template (size: %v) %v", email.Subject, email.To, email.Template, len([]byte(email.Source)), err))
		return err
	}
	return nil
}

func SendEmail(box *rice.Box, email *types.Email) {
	source := EmailTemplate(box, email.Template, email.Data)
	email.Source = source
	emailQueue = append(emailQueue, email)
}

func EmailTemplate(box *rice.Box, tmpl string, data interface{}) string {
	emailTpl, err := box.String(tmpl)
	if err != nil {
		utils.Log(3, err)
	}
	t := template.New("email")
	t, err = t.Parse(emailTpl)
	if err != nil {
		utils.Log(3, err)
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		utils.Log(2, err)
	}
	result := tpl.String()
	return result
}

func removeEmail(emails []*types.Email, em *types.Email) []*types.Email {
	var newArr []*types.Email
	for _, e := range emails {
		if e != em {
			newArr = append(newArr, e)
		}
	}
	return newArr
}

func inArray(a []string, v string) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}
