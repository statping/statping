package notifiers

import (
	"bytes"
	"html/template"
	"time"

	"github.com/statping-ng/statping-ng/types/core"
	"github.com/statping-ng/statping-ng/types/failures"
	"github.com/statping-ng/statping-ng/types/services"
	"github.com/statping-ng/statping-ng/utils"
)

//go:generate go run generate.go

var log = utils.Log.WithField("type", "notifier")

type replacer struct {
	Core    core.Core
	Service services.Service
	Failure failures.Failure
	Email   string
	Custom  map[string]string
}

func InitNotifiers() {
	Add(
		slacker,
		mattermoster,
		Command,
		Discorder,
		email,
		LineNotify,
		Telegram,
		Twilio,
		Webhook,
		Mobile,
		Pushover,
		statpingMailer,
		Gotify,
		AmazonSNS,
	)

	services.UpdateNotifiers()
}

func ReplaceTemplate(tmpl string, data replacer) string {
	buf := new(bytes.Buffer)
	tmp, err := template.New("replacement").Parse(tmpl)
	if err != nil {
		log.Error(err)
		return err.Error()
	}
	err = tmp.Execute(buf, data)
	if err != nil {
		log.Error(err)
		return err.Error()
	}
	return buf.String()
}

func Add(notifs ...services.ServiceNotifier) {
	for _, n := range notifs {
		notif := n.Select()
		if err := notif.Create(); err != nil {
			log.Error(err)
		}
		services.AddNotifier(n)
	}
}

func ReplaceVars(input string, s services.Service, f failures.Failure) string {
	return ReplaceTemplate(input, replacer{Service: s, Failure: f, Core: *core.App})
}

var exampleFailure = &failures.Failure{
	Id:        1,
	Issue:     "HTTP returned a 500 status code",
	ErrorCode: 500,
	Service:   1,
	PingTime:  43203,
	CreatedAt: utils.Now().Add(-10 * time.Minute),
}
