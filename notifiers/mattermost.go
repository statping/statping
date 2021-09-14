package notifiers

import (
	"bytes"
	"errors"
	"strings"
	"time"

	"github.com/statping-ng/statping-ng/types/failures"
	"github.com/statping-ng/statping-ng/types/notifications"
	"github.com/statping-ng/statping-ng/types/notifier"
	"github.com/statping-ng/statping-ng/types/null"
	"github.com/statping-ng/statping-ng/types/services"
	"github.com/statping-ng/statping-ng/utils"
)

var _ notifier.Notifier = (*mattermost)(nil)

const (
	mattermostMethod = "mattermost"
)

type mattermost struct {
	*notifications.Notification
}

func (s *mattermost) Select() *notifications.Notification {
	return s.Notification
}

var mattermoster = &mattermost{&notifications.Notification{
	Method:      mattermostMethod,
	Title:       "Mattermost",
	Description: "Send notifications to your mattermost channel when a service is offline. Insert your Incoming webhook URL for your channel to receive notifications. Based on the <a href=\"https://docs.mattermost.com/developer/webhooks-incoming.html\">Mattermost API</a>.",
	Author:      "Adam Boutcher",
	AuthorUrl:   "https://github.com/adamboutcher",
	Delay:       time.Duration(10 * time.Second),
	Icon:        "far fa-comments",
	SuccessData: null.NewNullString(`{"icon_emoji":":white_check_mark:", "text": "The service {{.Service.Name}} is back online."}`),
	FailureData: null.NewNullString(`{"icon_emoji":":x:", "text": "The service {{.Service.Name}} has gone offline."}`),
	DataType:    "json",
	RequestInfo: "Mattermost allows you to customize your own messages with many complex components.",
	Limits:      60,
	Form: []notifications.NotificationForm{{
		Type:        "text",
		Title:       "Incoming Webhook Url",
		Placeholder: "https://mattermost.example.com/hooks/abcd1234efgh5678ijkl9012xy",
		SmallText:   "Incoming Webhook URL generated on your mattermost server.",
		DbField:     "Host",
		Required:    true,
	}}},
}

// Send will send a HTTP Post to the mattermost webhooker API. It accepts type: string
func (s *mattermost) sendMattermost(msg string) (string, error) {
	resp, _, err := utils.HttpRequest(s.Host.String, "POST", "application/json", nil, strings.NewReader(msg), time.Duration(10*time.Second), true, nil)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func (s *mattermost) OnTest() (string, error) {
	example := services.Example(true)
	testMsg := ReplaceVars(s.SuccessData.String, example, failures.Failure{})
	contents, resp, err := utils.HttpRequest(s.Host.String, "POST", "application/json", nil, bytes.NewBuffer([]byte(testMsg)), time.Duration(10*time.Second), true, nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if string(contents) != "ok" {
		return string(contents), errors.New("the mattermost response was incorrect, check the URL")
	}
	return string(contents), nil
}

// OnFailure will trigger failing service
func (s *mattermost) OnFailure(srv services.Service, f failures.Failure) (string, error) {
	msg := ReplaceVars(s.FailureData.String, srv, f)
	out, err := s.sendMattermost(msg)
	return out, err
}

// OnSuccess will trigger successful service
func (s *mattermost) OnSuccess(srv services.Service) (string, error) {
	msg := ReplaceVars(s.SuccessData.String, srv, failures.Failure{})
	out, err := s.sendMattermost(msg)
	return out, err
}

// OnSave will trigger when this notifier is saved
func (s *mattermost) OnSave() (string, error) {
	return "", nil
}

func (s *mattermost) Valid(values notifications.Values) error {
	return nil
}
