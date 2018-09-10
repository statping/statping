// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
	emailQueue []*EmailOutgoing
	emailBox   *rice.Box
	mailer     *gomail.Dialer
)

type Email struct {
	*Notification
}

// DEFINE YOUR NOTIFICATION HERE.
func init() {

	emailer = &Email{
		Notification: &Notification{
			Id:     EMAIL_ID,
			Method: EMAIL_METHOD,
			Form: []NotificationForm{{
				Type:        "text",
				Title:       "SMTP Host",
				Placeholder: "Insert your SMTP Host here.",
				DbField:     "Host",
			}, {
				Type:        "text",
				Title:       "SMTP Username",
				Placeholder: "Insert your SMTP Username here.",
				DbField:     "Username",
			}, {
				Type:        "password",
				Title:       "SMTP Password",
				Placeholder: "Insert your SMTP Password here.",
				DbField:     "Password",
			}, {
				Type:        "number",
				Title:       "SMTP Port",
				Placeholder: "Insert your SMTP Port here.",
				DbField:     "Port",
			}, {
				Type:        "text",
				Title:       "Outgoing Email Address",
				Placeholder: "Insert your Outgoing Email Address",
				DbField:     "Var1",
			}, {
				Type:        "email",
				Title:       "Send Alerts To",
				Placeholder: "Email Address",
				DbField:     "Var2",
			}},
		}}

	err := AddNotifier(emailer)
	if err != nil {
		utils.Log(3, err)
	}
}

// WHEN NOTIFIER LOADS
func (u *Email) Init() error {
	emailBox = rice.MustFindBox("emails")
	err := u.Install()
	utils.Log(1, fmt.Sprintf("Creating Mailer: %v:%v", u.Notification.Host, u.Notification.Port))

	if err == nil {
		notifier, _ := SelectNotification(u.Id)
		forms := u.Form
		u.Notification = notifier
		u.Form = forms
		if u.Enabled {

			utils.Log(1, fmt.Sprintf("Loading SMTP Emailer using host: %v:%v", u.Notification.Host, u.Notification.Port))
			mailer = gomail.NewPlainDialer(u.Notification.Host, u.Notification.Port, u.Notification.Username, u.Notification.Password)
			mailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

			go u.Run()
		}
	}
	return nil
}

func (u *Email) Test() error {
	utils.Log(1, "Emailer notifier loaded")
	if u.Enabled {
		email := &EmailOutgoing{
			To:       emailer.Var2,
			Subject:  "Test Email",
			Template: "message.html",
			Data:     nil,
			From:     emailer.Var1,
		}
		SendEmail(emailBox, email)
	}
	return nil
}

type emailMessage struct {
	Service *types.Service
}

type EmailOutgoing struct {
	To       string
	Subject  string
	Template string
	From     string
	Data     interface{}
	Source   string
	Sent     bool
}

// AFTER NOTIFIER LOADS, IF ENABLED, START A QUEUE PROCESS
func (u *Email) Run() error {
	var sentAddresses []string
	for _, email := range emailQueue {
		if inArray(sentAddresses, email.To) || email.Sent {
			emailQueue = removeEmail(emailQueue, email)
			continue
		}
		if u.CanSend() {
			err := u.dialSend(email)
			if err == nil {
				email.Sent = true
				sentAddresses = append(sentAddresses, email.To)
				utils.Log(1, fmt.Sprintf("Email '%v' sent to: %v using the %v template (size: %v)", email.Subject, email.To, email.Template, len([]byte(email.Source))))
				emailQueue = removeEmail(emailQueue, email)
				u.Log(fmt.Sprintf("Subject: %v to %v", email.Subject, email.To))
			} else {
				utils.Log(3, fmt.Sprintf("Email Notifier could not send email: %v", err))
			}
		}
	}
	time.Sleep(60 * time.Second)
	if u.Enabled {
		return u.Run()
	}
	return nil
}

// ON SERVICE FAILURE, DO YOUR OWN FUNCTIONS
func (u *Email) OnFailure(s *types.Service, f *types.Failure) {
	if u.Enabled {
		msg := emailMessage{
			Service: s,
		}
		email := &EmailOutgoing{
			To:       emailer.Var2,
			Subject:  fmt.Sprintf("Service %v is Failing", s.Name),
			Template: "failure.html",
			Data:     msg,
			From:     emailer.Var1,
		}
		SendEmail(emailBox, email)

	}
}

// ON SERVICE SUCCESS, DO YOUR OWN FUNCTIONS
func (u *Email) OnSuccess(s *types.Service) {

}

// ON SAVE OR UPDATE OF THE NOTIFIER FORM
func (u *Email) OnSave() error {
	utils.Log(1, fmt.Sprintf("Notification %v is receiving updated information.", u.Method))
	// Do updating stuff here
	return nil
}

// ON SERVICE FAILURE, DO YOUR OWN FUNCTIONS
func (u *Email) Install() error {
	inDb := emailer.Notification.IsInDatabase()
	if !inDb {
		newNotifer, err := InsertDatabase(u.Notification)
		if err != nil {
			utils.Log(3, err)
			return err
		}
		utils.Log(1, fmt.Sprintf("new notifier #%v installed: %v", newNotifer, u.Method))
	}
	return nil
}

func (u *Email) dialSend(email *EmailOutgoing) error {
	m := gomail.NewMessage()
	m.SetHeader("From", email.From)
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/html", email.Source)
	if err := mailer.DialAndSend(m); err != nil {
		utils.Log(3, fmt.Sprintf("Email '%v' sent to: %v using the %v template (size: %v) %v", email.Subject, email.To, email.Template, len([]byte(email.Source)), err))
		return err
	}
	return nil
}

func SendEmail(box *rice.Box, email *EmailOutgoing) {
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

func removeEmail(emails []*EmailOutgoing, em *EmailOutgoing) []*EmailOutgoing {
	var newArr []*EmailOutgoing
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
