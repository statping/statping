package notifiers

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/notifier"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
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

func (t *pushover) Valid(values notifications.Values) error {
	return nil
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
	SuccessData: null.NewNullString(`Your service '{{.Service.Name}}' is currently online!`),
	FailureData: null.NewNullString(`Your service '{{.Service.Name}}' is currently offline!`),
	DataType:    "text",
	Form: []notifications.NotificationForm{{
		Type:        "text",
		Title:       "User Token",
		Placeholder: "Insert your Pushover User Token",
		DbField:     "api_key",
		Required:    true,
	}, {
		Type:        "text",
		Title:       "Application API Key",
		Placeholder: "Create an Application and insert the API Key here",
		DbField:     "api_secret",
		Required:    true,
	}, {
		Type:        "list",
		Title:       "Priority",
		Placeholder: "Set the notification priority level",
		DbField:     "Var1",
		Required:    true,
		ListOptions: []string{"Lowest", "Low", "Normal", "High", "Emergency"},
	}, {
		Type:        "list",
		Title:       "Notification Sound",
		Placeholder: "Choose a sound for this Pushover notification",
		DbField:     "Var2",
		Required:    true,
		ListOptions: []string{"none", "pushover", "bike", "bugle", "cashregister", "classical", "cosmic", "falling", "gamelan", "incoming", "intermissioon", "magic", "mechanical", "painobar", "siren", "spacealarm", "tugboat", "alien", "climb", "persistent", "echo", "updown"},
	},
	}},
}

func priority(val string) string {
	switch strings.ToLower(val) {
	case "lowest":
		return "-2"
	case "low":
		return "-1"
	case "normal":
		return "0"
	case "high":
		return "1"
	case "emergency":
		return "2"
	default:
		return "0"
	}
}

// Send will send a HTTP Post to the Pushover API. It accepts type: string
func (t *pushover) sendMessage(message string) (string, error) {
	v := url.Values{}
	v.Set("token", t.ApiSecret.String)
	v.Set("user", t.ApiKey.String)
	v.Set("message", message)
	v.Set("priority", priority(t.Var1.String))
	if t.Var2.String != "" {
		v.Set("sound", t.Var2.String)
	}
	rb := strings.NewReader(v.Encode())

	content, _, err := utils.HttpRequest(pushoverUrl, "POST", "application/x-www-form-urlencoded", nil, rb, time.Duration(10*time.Second), true, nil)
	if err != nil {
		return "", err
	}
	return string(content), err
}

// OnFailure will trigger failing service
func (t *pushover) OnFailure(s services.Service, f failures.Failure) (string, error) {
	message := ReplaceVars(t.FailureData.String, s, f)
	out, err := t.sendMessage(message)
	return out, err
}

// OnSuccess will trigger successful service
func (t *pushover) OnSuccess(s services.Service) (string, error) {
	message := ReplaceVars(t.SuccessData.String, s, failures.Failure{})
	out, err := t.sendMessage(message)
	return out, err
}

// OnTest will test the Pushover SMS messaging
func (t *pushover) OnTest() (string, error) {
	example := services.Example(true)
	msg := fmt.Sprintf("Testing the Pushover Notifier, Your service '%s' is currently offline! Error: %s", example.Name, exampleFailure.Issue)
	content, err := t.sendMessage(msg)
	return content, err
}

// OnSave will trigger when this notifier is saved
func (t *pushover) OnSave() (string, error) {
	return "", nil
}
