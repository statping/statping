package comms

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"os"
)

var (
	mailer *gomail.Dialer
)

func NewMailer() {
	mailer = gomail.NewDialer(os.Getenv("HOST"), 587, os.Getenv("USER"), os.Getenv("PASS"))
	mailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	source := EmailTemplate("comms/templates/error.html", "this is coooool")

	fmt.Println("source: ", source)
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

func SendEmail(to, subject, body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "info@email.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if err := mailer.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}
