package notifiers

import (
	"fmt"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/notifier"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"net/url"
	"strings"
	"time"
)

const (
	pushoverUrl = "https://api.pushover.net/1/messages.json"
)

var _ notifier.Notifier = (*pushover)(nil)

type pushover struct {
	*notifications.Notification
}

func (t *pushover) Select() *notifications.Notification {
	return t.Notification
}

var Pushover = &pushover{&notifications.Notification{
	Method:      "pushover",
	Title:       "Pushover",
	Description: "Use Pushover to receive push notifications. You will need to create a <a href=\"https://pushover.net/apps/build\">New Application</a> on Pushover before using this notifier.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Icon:        "fa dot-circle",
	Delay:       time.Duration(10 * time.Second),
	Limits:      60,
	Form: []notifications.NotificationForm{{
		Type:        "text",
		Title:       "User Token",
		Placeholder: "Insert your device's Pushover Token",
		DbField:     "api_key",
		Required:    true,
	}, {
		Type:        "text",
		Title:       "Application API Key",
		Placeholder: "Create an Application and insert the API Key here",
		DbField:     "api_secret",
		Required:    true,
	},
	}},
}

// Send will send a HTTP Post to the Pushover API. It accepts type: string
func (t *pushover) sendMessage(message string) (string, error) {
	v := url.Values{}
	v.Set("token", t.ApiSecret)
	v.Set("user", t.ApiKey)
	v.Set("message", message)
	rb := strings.NewReader(v.Encode())

	content, _, err := utils.HttpRequest(pushoverUrl, "POST", "application/x-www-form-urlencoded", nil, rb, time.Duration(10*time.Second), true)
	if err != nil {
		return "", err
	}
	return string(content), err
}

// OnFailure will trigger failing service
func (t *pushover) OnFailure(s *services.Service, f *failures.Failure) error {
	msg := fmt.Sprintf("Your service '%s' is currently offline!", s.Name)
	_, err := t.sendMessage(msg)
	return err
}

// OnSuccess will trigger successful service
func (t *pushover) OnSuccess(s *services.Service) error {
	msg := fmt.Sprintf("Your service '%s' is currently online!", s.Name)
	_, err := t.sendMessage(msg)
	return err
}

// OnTest will test the Pushover SMS messaging
func (t *pushover) OnTest() (string, error) {
	msg := fmt.Sprintf("Testing the Pushover Notifier, Your service '%s' is currently offline! Error: %s", exampleService.Name, exampleFailure.Issue)
	content, err := t.sendMessage(msg)
	return content, err
}
