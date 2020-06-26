package notifiers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/notifier"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"strings"
	"time"
)

var _ notifier.Notifier = (*discord)(nil)

type discord struct {
	*notifications.Notification
}

var Discorder = &discord{&notifications.Notification{
	Method:      "discord",
	Title:       "discord",
	Description: "Send notifications to your discord channel using discord webhooks. Insert your discord channel Webhook URL to receive notifications. Based on the <a href=\"https://discordapp.com/developers/docs/resources/Webhook\">discord webhooker API</a>.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(5 * time.Second),
	Host:        "https://discordapp.com/api/webhooks/****/*****",
	Icon:        "fab fa-discord",
	SuccessData: `{"content": "Your service '{{.Service.Name}}' is currently online!"}`,
	FailureData: `{"content": "Your service '{{.Service.Name}}' is currently failing! Reason: {{.Failure.Issue}}"}`,
	DataType:    "json",
	Limits:      60,
	Form: []notifications.NotificationForm{{
		Type:        "text",
		Title:       "discord webhooker URL",
		Placeholder: "Insert your Webhook URL here",
		DbField:     "host",
	}}},
}

// Send will send a HTTP Post to the discord API. It accepts type: []byte
func (d *discord) sendRequest(msg string) (string, error) {
	out, _, err := utils.HttpRequest(Discorder.GetValue("host"), "POST", "application/json", nil, strings.NewReader(msg), time.Duration(10*time.Second), true, nil)
	return string(out), err
}

func (d *discord) Select() *notifications.Notification {
	return d.Notification
}

// OnFailure will trigger failing service
func (d *discord) OnFailure(s services.Service, f failures.Failure) (string, error) {
	out, err := d.sendRequest(ReplaceVars(d.FailureData, s, f))
	return out, err
}

// OnSuccess will trigger successful service
func (d *discord) OnSuccess(s services.Service) (string, error) {
	out, err := d.sendRequest(ReplaceVars(d.SuccessData, s, failures.Failure{}))
	return out, err
}

// OnSave triggers when this notifier has been saved
func (d *discord) OnTest() (string, error) {
	outError := errors.New("Incorrect discord URL, please confirm URL is correct")
	message := `{"content": "Testing the discord notifier"}`
	contents, _, err := utils.HttpRequest(Discorder.Host, "POST", "application/json", nil, bytes.NewBuffer([]byte(message)), time.Duration(10*time.Second), true, nil)
	if string(contents) == "" {
		return "", nil
	}
	var dtt discordTestJson
	err = json.Unmarshal(contents, &dtt)
	if err != nil {
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
