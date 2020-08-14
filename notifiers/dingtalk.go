package notifiers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/notifier"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"strings"
	"time"
)

var _ notifier.Notifier = (*dingtalk)(nil)

type dingtalk struct {
	*notifications.Notification
}

var Dingtalker = &dingtalk{&notifications.Notification{
	Method:      "dingtalk",
	Title:       "dingtalk",
	Description: "Send notifications to your dingtalk group using dingtalk webhooks. Insert your dingtalk bot Webhook URL to receive notifications. Based on the <a href=\"https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq\">dingtalk bot API</a>.",
	Author:      "Huangchun Zhang",
	AuthorUrl:   "https://github.com/hoilc",
	Delay:       time.Duration(5 * time.Second),
	Host:        null.NewNullString("https://oapi.dingtalk.com/robot/send?access_token=*****"),
	Icon:        "dingtalk",
	SuccessData: null.NewNullString(`{"msgtype": "text","text": {"content": "Your service '{{.Service.Name}}' is currently online!"}}`),
	FailureData: null.NewNullString(`{"msgtype": "text","text": {"content": "Your service '{{.Service.Name}}' is currently failing!\nReason: {{.Failure.Issue}}"}}`),
	DataType:    "json",
	Limits:      20,
	Form: []notifications.NotificationForm{{
		Type:        "text",
		Title:       "dingtalk bot webhook URL",
		Placeholder: "Insert your Webhook URL here",
		DbField:     "host",
	}}},
}

// Send will send a HTTP Post to the dingtalk API. It accepts type: []byte
func (d *dingtalk) sendRequest(msg string) (string, error) {
	out, _, err := utils.HttpRequest(d.Host.String, "POST", "application/json", nil, strings.NewReader(msg), time.Duration(10*time.Second), true, nil)
	return string(out), err
}

func (d *dingtalk) Select() *notifications.Notification {
	return d.Notification
}

func (d *dingtalk) Valid(values notifications.Values) error {
	return nil
}

// OnFailure will trigger failing service
func (d *dingtalk) OnFailure(s services.Service, f failures.Failure) (string, error) {
	out, err := d.sendRequest(ReplaceVars(d.FailureData.String, s, f))
	return out, err
}

// OnSuccess will trigger successful service
func (d *dingtalk) OnSuccess(s services.Service) (string, error) {
	out, err := d.sendRequest(ReplaceVars(d.SuccessData.String, s, failures.Failure{}))
	return out, err
}

// OnSave triggers when test dingtalk notify
func (d *dingtalk) OnTest() (string, error) {
	outError := errors.New("incorrect dingtalk URL, please confirm URL is correct")
	message := `{"msgtype": "text","text": {"content": "Testing dingtalk notifier from statping"}}`
	contents, _, err := utils.HttpRequest(Dingtalker.Host.String, "POST", "application/json", nil, bytes.NewBuffer([]byte(message)), time.Duration(10*time.Second), true, nil)
	if string(contents) == "" {
		return "", nil
	}
	var dtt dingtalkTestJson
	if err != nil {
		return "", err
	}
	if err = json.Unmarshal(contents, &dtt); err != nil {
		return string(contents), outError
	}
	if dtt.Code != 0 {
		return string(contents), errors.New(dtt.Message)
	}
	return string(contents), nil
}

// OnSave will trigger when this notifier is saved
func (d *dingtalk) OnSave() (string, error) {
	return "", nil
}

type dingtalkTestJson struct {
	Code    int    `json:"errcode"`
	Message string `json:"errmsg"`
}
