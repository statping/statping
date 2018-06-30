package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"gopkg.in/gomail.v2"
	"html/template"
	"time"
)

var (
	emailQue *Que
)

type Email types.Email

type Que struct {
	Mailer       *gomail.Dialer
	Outgoing     []*Email
	LastSent     int
	LastSentTime time.Time
}

func AddEmail(email *Email) {
	if emailQue == nil {
		return
	}
	emailQue.Outgoing = append(emailQue.Outgoing, email)
}

func EmailerQueue() {
	if emailQue == nil {
		return
	}
	uniques := []*Email{}
	for _, out := range emailQue.Outgoing {
		if isUnique(uniques, out) {
			msg := fmt.Sprintf("sending email to: %v \n", out.To)
			Send(out)
			utils.Log(0, msg)
			uniques = append(uniques, out)
		}
	}
	emailQue.Outgoing = nil
	fmt.Println("running emailer queue")
	time.Sleep(60 * time.Second)
	EmailerQueue()
}

func isUnique(arr []*Email, obj *Email) bool {
	for _, v := range arr {
		if v.To == obj.To && v.Subject == obj.Subject {
			return false
		}
	}
	return true
}

func Send(em *Email) {
	source := EmailTemplate(em.Template, em.Data)
	m := gomail.NewMessage()
	m.SetHeader("From", "info@betatude.com")
	m.SetHeader("To", em.To)
	m.SetHeader("Subject", em.Subject)
	m.SetBody("text/html", source)
	if err := emailQue.Mailer.DialAndSend(m); err != nil {
		utils.Log(2, err)
	}
	emailQue.LastSent++
	emailQue.LastSentTime = time.Now()
}

func SendFailureEmail(service *core.Service) {
	email := &Email{
		To:       "info@socialeck.com",
		Subject:  fmt.Sprintf("Service %v is Failing", service.Name),
		Template: "failure.html",
		Data:     service,
	}
	AddEmail(email)
}

func LoadMailer(config *core.Communication) *gomail.Dialer {
	if config.Host == "" || config.Username == "" || config.Password == "" {
		return nil
	}
	emailQue = new(Que)
	emailQue.Outgoing = []*Email{}
	emailQue.Mailer = gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	emailQue.Mailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return emailQue.Mailer
}

func EmailTemplate(tmpl string, data interface{}) string {
	emailTpl, err := core.EmailBox.String(tmpl)
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
