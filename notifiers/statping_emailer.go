package notifiers

import (
	"bytes"
	"encoding/json"
	"errors"
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
	Title:       "Statping Emailer",
	Description: "Send an email when a service becomes offline or back online using Statping's email service. You will need to verify your email address.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(10 * time.Second),
	Icon:        "fab fa-slack",
	RequestInfo: "Slack allows you to customize your own messages with many complex components. Checkout the <a target=\"_blank\" href=\"https://api.slack.com/reference/surfaces/formatting\">Slack Message API</a> to learn how you can create your own.",
	Limits:      60,
	Form: []notifications.NotificationForm{{
		Type:        "email",
		Title:       "Send to Email Address",
		Placeholder: "Insert your email address",
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
	example := services.Example(true)
	testMsg := ReplaceVars(s.SuccessData, example, nil)
	contents, resp, err := utils.HttpRequest(s.Host, "POST", "application/json", nil, bytes.NewBuffer([]byte(testMsg)), time.Duration(10*time.Second), true, nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if string(contents) != "ok" {
		return string(contents), errors.New("the slack response was incorrect, check the URL")
	}
	return string(contents), nil
}

type statpingMail struct {
	Email   string            `json:"email"`
	Core    *core.Core        `json:"core"`
	Service *services.Service `json:"service"`
	Failure *failures.Failure `json:"failure,omitempty"`
}

// OnFailure will trigger failing service
func (s *statpingEmailer) OnFailure(srv *services.Service, f *failures.Failure) (string, error) {
	ee := statpingMail{
		Email:   s.Host,
		Core:    core.App,
		Service: srv,
		Failure: f,
	}
	out, err := s.sendStatpingEmail(ee)
	return out, err
}

// OnSuccess will trigger successful service
func (s *statpingEmailer) OnSuccess(srv *services.Service) (string, error) {
	ee := statpingMail{
		Email:   s.Host,
		Core:    core.App,
		Service: srv,
		Failure: nil,
	}
	out, err := s.sendStatpingEmail(ee)
	return out, err
}
