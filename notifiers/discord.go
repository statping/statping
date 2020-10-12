package notifiers

import (
	"bytes"
	"encoding/json"
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

var _ notifier.Notifier = (*discord)(nil)

type discord struct {
	*notifications.Notification
}

var Discorder = &discord{&notifications.Notification{
	Method:      "discord",
	Title:       "Discord",
	Description: "Send notifications to your discord channel using discord webhooks. Insert your discord channel Webhook URL to receive notifications. Based on the <a href=\"https://discordapp.com/developers/docs/resources/Webhook\">discord webhooker API</a>.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(5 * time.Second),
	Icon:        "fab fa-discord",
	SuccessData: null.NewNullString(`{"content": "Your service '{{.Service.Name}}' is currently back online and was down for {{.Service.Downtime.Human}}."}`),
	FailureData: null.NewNullString(`{"content": "Your service '{{.Service.Name}}' is has been failing for {{.Service.Downtime.Human}}! Reason: {{.Failure.Issue}}"}`),
	DataType:    "json",
	Limits:      60,
	Form: []notifications.NotificationForm{{
		Type:        "text",
		Title:       "discord webhooker URL",
		Placeholder: "https://discordapp.com/api/webhooks/****/*****",
		DbField:     "host",
	}}},
}

// Send will send a HTTP Post to the discord API. It accepts type: []byte
func (d *discord) sendRequest(msg string) (string, error) {
	out, _, err := utils.HttpRequest(d.Host.String, "POST", "application/json", nil, strings.NewReader(msg), time.Duration(10*time.Second), true, nil)
	return string(out), err
}

func (d *discord) Select() *notifications.Notification {
	return d.Notification
}

func (d *discord) Valid(values notifications.Values) error {
	return nil
}

// OnFailure will trigger failing service
func (d *discord) OnFailure(s services.Service, f failures.Failure) (string, error) {
	out, err := d.sendRequest(ReplaceVars(d.FailureData.String, s, f))
	return out, err
}

// OnSuccess will trigger successful service
func (d *discord) OnSuccess(s services.Service) (string, error) {
	out, err := d.sendRequest(ReplaceVars(d.SuccessData.String, s, failures.Failure{}))
	return out, err
}

// OnSave triggers when this notifier has been saved
func (d *discord) OnTest() (string, error) {
	outError := errors.New("incorrect discord URL, please confirm URL is correct")
	message := `{"content": "Testing the discord notifier"}`
	contents, _, err := utils.HttpRequest(Discorder.Host.String, "POST", "application/json", nil, bytes.NewBuffer([]byte(message)), time.Duration(10*time.Second), true, nil)
	if string(contents) == "" {
		return "", nil
	}
	var dtt discordTestJson
	if err != nil {
		return "", err
	}
	if err = json.Unmarshal(contents, &dtt); err != nil {
		return string(contents), outError
	}
	if dtt.Code == 0 {
		return string(contents), outError
	}
	return string(contents), nil
}

// OnSave will trigger when this notifier is saved
func (d *discord) OnSave() (string, error) {
	return "", nil
}

type discordTestJson struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
