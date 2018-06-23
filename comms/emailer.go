package comms

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
	Emailer  *gomail.Dialer
	Outgoing []*types.Email
)

func AddEmail(email *types.Email) {
	Outgoing = append(Outgoing, email)
}

func EmailerQueue() {
	defer EmailerQueue()
	for _, out := range Outgoing {
		fmt.Printf("sending email to: %v \n", out.To)
		Send(out)
	}
	Outgoing = nil
	fmt.Println("running emailer queue")
	time.Sleep(10 * time.Second)
}

func Send(em *types.Email) {
	source := EmailTemplate("comms/templates/error.html", nil)
	m := gomail.NewMessage()
	m.SetHeader("From", "info@betatude.com")
	m.SetHeader("To", em.To)
	m.SetHeader("Subject", em.Subject)
	m.SetBody("text/html", source)
	if err := Emailer.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}

func SendSample(em *types.Email) {
	source := EmailTemplate("comms/templates/error.html", nil)
	m := gomail.NewMessage()
	m.SetHeader("From", "info@betatude.com")
	m.SetHeader("To", em.To)
	m.SetHeader("Subject", em.Subject)
	m.SetBody("text/html", source)
	if err := Emailer.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}

func LoadMailer(config *types.Communication) *gomail.Dialer {
	Emailer = gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	Emailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return Emailer
}

func EmailTemplate(tmpl string, data interface{}) string {
	t := template.New("error.html")
	var err error
	t, err = t.ParseFiles(tmpl)
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
