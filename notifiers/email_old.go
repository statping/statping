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

var (
	mailer     *gomail.Dialer
	emailQueue []*types.Email
	EmailComm  *Notification
)

func EmailRoutine() {
	var sentAddresses []string
	for _, email := range emailQueue {
		if inArray(sentAddresses, email.To) || email.Sent {
			emailQueue = removeEmail(emailQueue, email)
			continue
		}
		e := email
		go func(email *types.Email) {
			err := dialSend(email)
			if err == nil {
				email.Sent = true
				sentAddresses = append(sentAddresses, email.To)
				utils.Log(1, fmt.Sprintf("Email '%v' sent to: %v using the %v template (size: %v)", email.Subject, email.To, email.Template, len([]byte(email.Source))))
				emailQueue = removeEmail(emailQueue, email)
			}
		}(e)
	}
	time.Sleep(60 * time.Second)
	if EmailComm.Enabled {
		EmailRoutine()
	}
}

func dialSend(email *types.Email) error {
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

func LoadEmailer(mail *Notification) {
	utils.Log(1, fmt.Sprintf("Loading SMTP Emailer using host: %v:%v", mail.Host, mail.Port))
	mailer = gomail.NewDialer(mail.Host, mail.Port, mail.Username, mail.Password)
	mailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
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
