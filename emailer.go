package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/hunterlong/statup/types"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"time"
)

var (
	emailQue *Que
)

type Que struct {
	Mailer       *gomail.Dialer
	Outgoing     []*types.Email
	LastSent     int
	LastSentTime time.Time
}

func AddEmail(email *types.Email) {
	if emailQue == nil {
		return
	}
	emailQue.Outgoing = append(emailQue.Outgoing, email)
}

func EmailerQueue() {
	defer EmailerQueue()
	uniques := []*types.Email{}
	for _, out := range emailQue.Outgoing {
		if isUnique(uniques, out) {
			fmt.Printf("sending email to: %v \n", out.To)
			Send(out)
			uniques = append(uniques, out)
		}
	}
	emailQue.Outgoing = nil
	fmt.Println("running emailer queue")
	time.Sleep(60 * time.Second)
}

func isUnique(arr []*types.Email, obj *types.Email) bool {
	for _, v := range arr {
		if v.To == obj.To && v.Subject == obj.Subject {
			return false
		}
	}
	return true
}

func Send(em *types.Email) {
	source := EmailTemplate(em.Template, em.Data)
	m := gomail.NewMessage()
	m.SetHeader("From", "info@betatude.com")
	m.SetHeader("To", em.To)
	m.SetHeader("Subject", em.Subject)
	m.SetBody("text/html", source)
	if err := emailQue.Mailer.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
	emailQue.LastSent++
	emailQue.LastSentTime = time.Now()
}

func SendFailureEmail(service *Service) {
	email := &types.Email{
		To:       "info@socialeck.com",
		Subject:  fmt.Sprintf("Service %v is Failing", service.Name),
		Template: "failure.html",
		Data:     service,
	}
	AddEmail(email)
}

func LoadMailer(config *types.Communication) *gomail.Dialer {
	emailQue = &Que{}
	emailQue.Outgoing = []*types.Email{}
	emailQue.Mailer = gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	emailQue.Mailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return emailQue.Mailer
}

func EmailTemplate(tmpl string, data interface{}) string {
	emailTpl, err := emailBox.String(tmpl)
	if err != nil {
		panic(err)
	}
	t := template.New("email")
	t, err = t.Parse(emailTpl)
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		log.Println(err)
	}
	result := tpl.String()
	return result
}
