package notifiers

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/go-mail/mail"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/notifier"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"html/template"
)

var _ notifier.Notifier = (*emailer)(nil)

var (
	mailer *mail.Dialer
)

type emailer struct {
	*notifications.Notification
}

func (e *emailer) Select() *notifications.Notification {
	return e.Notification
}

var email = &emailer{&notifications.Notification{
	Method:      "email",
	Title:       "SMTP Mail",
	Description: "Send emails via SMTP when services are online or offline.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Icon:        "far fa-envelope",
	Limits:      30,
	Form: []notifications.NotificationForm{{
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
		Placeholder: "outgoing@email.com",
		DbField:     "Var1",
	}, {
		Type:        "email",
		Title:       "Send Alerts To",
		Placeholder: "sendto@email.com",
		DbField:     "Var2",
	}, {
		Type:        "switch",
		Title:       "Disable TLS/SSL",
		Placeholder: "",
		SmallText:   "Enabling this will set Insecure Skip Verify to true",
		DbField:     "api_key",
	}}},
}

type emailOutgoing struct {
	To       string
	Subject  string
	Template string
	From     string
	Data     replacer
	Source   string
	Sent     bool
}

// OnFailure will trigger failing service
func (e *emailer) OnFailure(s services.Service, f failures.Failure) (string, error) {
	subject := fmt.Sprintf("Service %s is Offline", s.Name)
	tmpl := renderEmail(s, f)
	email := &emailOutgoing{
		To:       e.Var2,
		Subject:  subject,
		Template: tmpl,
		From:     e.Var1,
	}
	return tmpl, e.dialSend(email)
}

// OnSuccess will trigger successful service
func (e *emailer) OnSuccess(s services.Service) (string, error) {
	subject := fmt.Sprintf("Service %s is Back Online", s.Name)
	tmpl := renderEmail(s, failures.Failure{})
	email := &emailOutgoing{
		To:       e.Var2,
		Subject:  subject,
		Template: tmpl,
		From:     e.Var1,
	}
	return tmpl, e.dialSend(email)
}

func renderEmail(s services.Service, f failures.Failure) string {
	wr := bytes.NewBuffer(nil)
	tmpl := template.New("email")
	tmpl, err := tmpl.Parse(emailBase)
	if err != nil {
		log.Errorln(err)
		return emailBase
	}

	data := replacer{
		Core:    *core.App,
		Service: s,
		Failure: f,
		Custom:  nil,
	}

	if err = tmpl.ExecuteTemplate(wr, "email", data); err != nil {
		log.Errorln(err)
		return emailBase
	}

	return wr.String()
}

// OnTest triggers when this notifier has been saved
func (e *emailer) OnTest() (string, error) {
	service := services.Example(true)
	subject := fmt.Sprintf("Service %v is Back Online", service.Name)
	email := &emailOutgoing{
		To:       e.Var2,
		Subject:  subject,
		Template: renderEmail(service, failures.Example()),
		From:     e.Var1,
	}
	return subject, e.dialSend(email)
}

// OnSave will trigger when this notifier is saved
func (e *emailer) OnSave() (string, error) {
	return "", nil
}

func (e *emailer) dialSend(email *emailOutgoing) error {
	mailer = mail.NewDialer(e.Host, e.Port, e.Username, e.Password)
	m := mail.NewMessage()
	// if email setting TLS is Disabled
	if e.ApiKey == "true" {
		mailer.SSL = false
	} else {
		mailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	m.SetAddressHeader("From", email.From, "Statping")
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/html", email.Template)

	if err := mailer.DialAndSend(m); err != nil {
		utils.Log.Errorln(fmt.Sprintf("email '%v' sent to: %v (size: %v) %v", email.Subject, email.To, len([]byte(email.Source)), err))
		return err
	}

	return nil
}
