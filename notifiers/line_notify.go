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

var _ notifier.Notifier = (*lineNotifier)(nil)

const (
	lineNotifyMethod = "line_notify"
)

type lineNotifier struct {
	*notifications.Notification
}

func (l *lineNotifier) Select() *notifications.Notification {
	return l.Notification
}

var LineNotify = &lineNotifier{&notifications.Notification{
	Method:      lineNotifyMethod,
	Title:       "LINE Notify",
	Description: "LINE Notify will send notifications to your LINE Notify account when services are offline or online. Based on the <a href=\"https://notify-bot.line.me/doc/en/\">LINE Notify API</a>.",
	Author:      "Kanin Peanviriyakulkit",
	AuthorUrl:   "https://github.com/dogrocker",
	Icon:        "far fa-bell",
	Limits:      60,
	Form: []notifications.NotificationForm{{
		Type:        "text",
		Title:       "Access Token",
		Placeholder: "Insert your Line Notify Access Token here.",
		DbField:     "api_secret",
	}}},
}

// Send will send a HTTP Post with the Authorization to the notify-api.line.me server. It accepts type: string
func (l *lineNotifier) sendMessage(message string) error {
	v := url.Values{}
	v.Set("message", message)
	headers := []string{fmt.Sprintf("Authorization=Bearer %v", l.ApiSecret)}
	_, _, err := utils.HttpRequest("https://notify-api.line.me/api/notify", "POST", "application/x-www-form-urlencoded", headers, strings.NewReader(v.Encode()), time.Duration(10*time.Second), true)
	return err
}

// OnFailure will trigger failing service
func (l *lineNotifier) OnFailure(s *services.Service, f *failures.Failure) error {
	msg := fmt.Sprintf("Your service '%v' is currently offline!", s.Name)
	return l.sendMessage(msg)
}

// OnSuccess will trigger successful service
func (l *lineNotifier) OnSuccess(s *services.Service) error {
	msg := fmt.Sprintf("Service %s is online!", s.Name)
	return l.sendMessage(msg)
}

// OnTest triggers when this notifier has been saved
func (l *lineNotifier) OnTest() error {
	msg := fmt.Sprintf("Testing if Line Notifier is working!")
	return l.sendMessage(msg)
}
