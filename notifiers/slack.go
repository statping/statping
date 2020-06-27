package notifiers

import (
	"bytes"
	"errors"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/notifier"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"strings"
	"time"
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
	Title:       "slack",
	Description: "Send notifications to your slack channel when a service is offline. Insert your Incoming webhook URL for your channel to receive notifications. Based on the <a href=\"https://api.slack.com/incoming-webhooks\">Slack API</a>.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(10 * time.Second),
	Host:        "https://webhooksurl.slack.com/***",
	Icon:        "fab fa-slack",
	SuccessData: `{ "blocks": [ { "type": "section", "text": { "type": "mrkdwn", "text": "The service {{.Service.Name}} is back online." } }, { "type": "actions", "elements": [ { "type": "button", "text": { "type": "plain_text", "text": "View Service", "emoji": true }, "style": "primary", "url": "{{.Core.Domain}}/service/{{.Service.Id}}" }, { "type": "button", "text": { "type": "plain_text", "text": "Go to Statping", "emoji": true }, "url": "{{.Core.Domain}}" } ] } ] }`,
	FailureData: `{ "blocks": [ { "type": "section", "text": { "type": "mrkdwn", "text": ":warning: The service {{.Service.Name}} is currently offline! :warning:" } }, { "type": "divider" }, { "type": "section", "fields": [ { "type": "mrkdwn", "text": "*Service:*\n{{.Service.Name}}" }, { "type": "mrkdwn", "text": "*URL:*\n{{.Service.Domain}}" }, { "type": "mrkdwn", "text": "*Status Code:*\n{{.Service.LastStatusCode}}" }, { "type": "mrkdwn", "text": "*When:*\n{{.Failure.CreatedAt}}" }, { "type": "mrkdwn", "text": "*Downtime:*\n{{.Service.DowntimeAgo}}" }, { "type": "plain_text", "text": "*Error:*\n{{.Failure.Issue}}" } ] }, { "type": "divider" }, { "type": "actions", "elements": [ { "type": "button", "text": { "type": "plain_text", "text": "View Offline Service", "emoji": true }, "style": "danger", "url": "{{.Core.Domain}}/service/{{.Service.Id}}" }, { "type": "button", "text": { "type": "plain_text", "text": "Go to Statping", "emoji": true }, "url": "{{.Core.Domain}}" } ] } ] }`,
	DataType:    "json",
	RequestInfo: "Slack allows you to customize your own messages with many complex components. Checkout the <a target=\"_blank\" href=\"https://api.slack.com/reference/surfaces/formatting\">Slack Message API</a> to learn how you can create your own.",
	Limits:      60,
	Form: []notifications.NotificationForm{{
		Type:        "text",
		Title:       "Incoming Webhook Url",
		Placeholder: "Insert your Slack Webhook URL here.",
		SmallText:   "Incoming Webhook URL from <a href=\"https://api.slack.com/apps\" target=\"_blank\">Slack Apps</a>",
		DbField:     "Host",
		Required:    true,
	}}},
}

// Send will send a HTTP Post to the slack webhooker API. It accepts type: string
func (s *slack) sendSlack(msg string) (string, error) {
	resp, _, err := utils.HttpRequest(s.Host, "POST", "application/json", nil, strings.NewReader(msg), time.Duration(10*time.Second), true, nil)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func (s *slack) OnTest() (string, error) {
	example := services.Example(true)
	testMsg := ReplaceVars(s.SuccessData, example, failures.Failure{})
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

// OnFailure will trigger failing service
func (s *slack) OnFailure(srv services.Service, f failures.Failure) (string, error) {
	msg := ReplaceVars(s.FailureData, srv, f)
	out, err := s.sendSlack(msg)
	return out, err
}

// OnSuccess will trigger successful service
func (s *slack) OnSuccess(srv services.Service) (string, error) {
	msg := ReplaceVars(s.SuccessData, srv, failures.Failure{})
	out, err := s.sendSlack(msg)
	return out, err
}

// OnSave will trigger when this notifier is saved
func (s *slack) OnSave() (string, error) {
	return "", nil
}
