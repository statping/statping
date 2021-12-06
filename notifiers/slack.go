package notifiers

import (
	"bytes"
	"errors"
	"strings"
	"time"

	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/notifier"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
)

var _ notifier.Notifier = (*slack)(nil)

const (
	slackMethod = "slack"
)

type slack struct {
	*notifications.Notification
}

func (s *slack) Select() *notifications.Notification {
	return s.Notification
}

var slacker = &slack{&notifications.Notification{
	Method:      slackMethod,
	Title:       "Slack",
	Description: "Send notifications to your slack channel when a service is offline. Insert your Incoming webhook URL for your channel to receive notifications. Based on the <a href=\"https://api.slack.com/incoming-webhooks\">Slack API</a>.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(10 * time.Second),
	Icon:        "fab fa-slack",
	SuccessData: null.NewNullString(`{ "blocks": [ { "type": "section", "text": { "type": "mrkdwn", "text": "The service {{.Service.Name}} is back online." } }, { "type": "actions", "elements": [ { "type": "button", "text": { "type": "plain_text", "text": "View Service", "emoji": true }, "style": "primary", "url": "{{.Core.Domain}}/service/{{.Service.Id}}" }, { "type": "button", "text": { "type": "plain_text", "text": "Go to Statping", "emoji": true }, "url": "{{.Core.Domain}}" } ] } ] }`),
	FailureData: null.NewNullString(`{ "blocks": [ { "type": "section", "text": { "type": "mrkdwn", "text": ":warning: The service {{.Service.Name}} is currently offline! :warning:" } }, { "type": "divider" }, { "type": "section", "fields": [ { "type": "mrkdwn", "text": "*Service:*\n{{.Service.Name}}" }, { "type": "mrkdwn", "text": "*URL:*\n{{.Service.Domain}}" }, { "type": "mrkdwn", "text": "*Status Code:*\n{{.Service.LastStatusCode}}" }, { "type": "mrkdwn", "text": "*When:*\n{{.Failure.CreatedAt}}" }, { "type": "mrkdwn", "text": "*Downtime:*\n{{.Service.Downtime.Human}}" }, { "type": "plain_text", "text": "*Error:*\n{{.Failure.Issue}}" } ] }, { "type": "divider" }, { "type": "actions", "elements": [ { "type": "button", "text": { "type": "plain_text", "text": "View Offline Service", "emoji": true }, "style": "danger", "url": "{{.Core.Domain}}/service/{{.Service.Id}}" }, { "type": "button", "text": { "type": "plain_text", "text": "Go to Statping", "emoji": true }, "url": "{{.Core.Domain}}" } ] } ] }`),
	DataType:    "json",
	RequestInfo: "Slack allows you to customize your own messages with many complex components. Checkout the <a target=\"_blank\" href=\"https://api.slack.com/reference/surfaces/formatting\">Slack Message API</a> to learn how you can create your own.",
	Limits:      60,
	Form: []notifications.NotificationForm{{
		Type:        "text",
		Title:       "Incoming Webhook Url",
		Placeholder: "https://hooks.slack.com/services/ETJ1B87WE/H76D6G8S30/H4d97R4EcZ40SpfyqPlAHr",
		SmallText:   "Incoming Webhook URL from <a href=\"https://api.slack.com/apps\" target=\"_blank\">Slack Apps</a>",
		DbField:     "Host",
		Required:    true,
	}}},
}

// Send will send a HTTP Post to the slack webhooker API. It accepts type: string
func (s *slack) sendSlack(msg string) (string, error) {
	resp, _, err := utils.HttpRequest(s.Host.String, "POST", "application/json", nil, strings.NewReader(msg), time.Duration(10*time.Second), true, nil)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func (s *slack) OnTest() (string, error) {
	example := services.Example(true)
	testMsg := ReplaceVars(s.SuccessData.String, example, failures.Failure{})
	contents, resp, err := utils.HttpRequest(s.Host.String, "POST", "application/json", nil, bytes.NewBuffer([]byte(testMsg)), time.Duration(10*time.Second), true, nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if string(contents) != "ok" {
		return string(contents), errors.New("the slack response was incorrect, check the URL")
	}
	return string(contents), nil
}

// OnFailure will trigger failing service
func (s *slack) OnFailure(srv services.Service, f failures.Failure) (string, error) {
	msg := ReplaceVars(s.FailureData.String, srv, f)
	out, err := s.sendSlack(msg)
	return out, err
}

// OnSuccess will trigger successful service
func (s *slack) OnSuccess(srv services.Service) (string, error) {
	msg := ReplaceVars(s.SuccessData.String, srv, failures.Failure{})
	out, err := s.sendSlack(msg)
	return out, err
}

// OnSave will trigger when this notifier is saved
func (s *slack) OnSave() (string, error) {
	return "", nil
}

func (s *slack) Valid(values notifications.Values) error {
	return nil
}
