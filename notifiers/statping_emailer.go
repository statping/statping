package notifiers

import (
	"bytes"
	"encoding/json"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/notifier"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"time"
)

var _ notifier.Notifier = (*statpingEmailer)(nil)

const (
	statpingEmailerName = "statping_emailer"
	statpingEmailerHost = "https://news.statping.com"
)

type statpingEmailer struct {
	*notifications.Notification
}

func (s *statpingEmailer) Select() *notifications.Notification {
	return s.Notification
}

var statpingMailer = &statpingEmailer{&notifications.Notification{
	Method:      statpingEmailerName,
	Title:       "Email",
	Description: "Send an email when a service becomes offline or back online using Statping's email service. You will need to verify your email address.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(10 * time.Second),
	Icon:        "fas envelope-square",
	Limits:      60,
	Form: []notifications.NotificationForm{{
		Type:        "email",
		Title:       "Send to Email Address",
		Placeholder: "info@statping.com",
		DbField:     "Host",
		Required:    true,
	}}},
}

// Send will send a HTTP Post to the slack webhooker API. It accepts type: string
func (s *statpingEmailer) sendStatpingEmail(msg statpingMail) (string, error) {
	data, _ := json.Marshal(msg)
	resp, _, err := utils.HttpRequest(statpingEmailerHost+"/notifier", "POST", "application/json", nil, bytes.NewBuffer(data), time.Duration(10*time.Second), true, nil)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func (s *statpingEmailer) OnTest() (string, error) {
	return "", nil
}

type statpingMail struct {
	Email   string           `json:"email"`
	Core    core.Core        `json:"core,omitempty"`
	Service services.Service `json:"service,omitempty"`
	Failure failures.Failure `json:"failure,omitempty"`
}

// OnFailure will trigger failing service
func (s *statpingEmailer) OnFailure(srv services.Service, f failures.Failure) (string, error) {
	ee := statpingMail{
		Email:   s.Host,
		Core:    *core.App,
		Service: srv,
		Failure: f,
	}
	return s.sendStatpingEmail(ee)
}

// OnSuccess will trigger successful service
func (s *statpingEmailer) OnSuccess(srv services.Service) (string, error) {
	ee := statpingMail{
		Email:   s.Host,
		Core:    *core.App,
		Service: srv,
		Failure: failures.Failure{},
	}
	return s.sendStatpingEmail(ee)
}

// OnSave will trigger when this notifier is saved
func (s *statpingEmailer) OnSave() (string, error) {
	ee := statpingMail{
		Email:   s.Host,
		Core:    *core.App,
		Service: services.Service{},
		Failure: failures.Failure{},
	}
	out, err := s.sendStatpingEmail(ee)
	log.Println("statping emailer response", out)
	return out, err
}
