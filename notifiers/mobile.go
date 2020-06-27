package notifiers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/notifier"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"time"
)

var _ notifier.Notifier = (*mobilePush)(nil)

type mobilePush struct {
	*notifications.Notification
}

func (m *mobilePush) Select() *notifications.Notification {
	return m.Notification
}

var Mobile = &mobilePush{&notifications.Notification{
	Method: "mobile",
	Title:  "Mobile",
	Description: `Receive push notifications on your Mobile device using the Statping App. You can scan the Authentication QR Code found in Settings to get the Mobile app setup in seconds.
				 <p align="center"><a href="https://play.google.com/store/apps/details?id=com.statping"><img src="https://img.cjx.io/google-play.svg"></a><a href="https://itunes.apple.com/us/app/apple-store/id1445513219"><img src="https://img.cjx.io/app-store-badge.svg"></a></p>`,
	Author:    "Hunter Long",
	AuthorUrl: "https://github.com/hunterlong",
	Delay:     time.Duration(5 * time.Second),
	Icon:      "fas fa-mobile-alt",
	Limits:    30,
	Form: []notifications.NotificationForm{{
		Type:        "text",
		Title:       "Device Identifiers",
		Placeholder: "A list of your Mobile device push notification ID's.",
		DbField:     "var1",
		IsHidden:    true,
	}, {
		Type:        "number",
		Title:       "Array of device numbers",
		Placeholder: "1 for iphone 2 for android",
		DbField:     "var2",
		IsHidden:    true,
	}}},
}

func dataJson(s services.Service, f failures.Failure) map[string]interface{} {
	serviceId := "0"
	serviceId = utils.ToString(s.Id)
	online := "online"
	if !s.Online {
		online = "offline"
	}
	issue := f.Issue
	link := fmt.Sprintf("statping://service?id=%v", serviceId)
	out := map[string]interface{}{
		"status": online,
		"id":     serviceId,
		"issue":  issue,
		"link":   link,
	}
	return out
}

// OnFailure will trigger failing service
func (m *mobilePush) OnFailure(s services.Service, f failures.Failure) (string, error) {
	data := dataJson(s, f)
	msg := &pushArray{
		Message: fmt.Sprintf("Your service '%v' is currently failing! Reason: %v", s.Name, f.Issue),
		Title:   "Service Offline",
		Data:    data,
	}
	return "notification sent", m.Send(msg)
}

// OnSuccess will trigger successful service
func (m *mobilePush) OnSuccess(s services.Service) (string, error) {
	data := dataJson(s, failures.Failure{})
	msg := &pushArray{
		Message:  "Service is Online!",
		Title:    "Service Online",
		Data:     data,
		Platform: 2,
	}
	return "notification sent", m.Send(msg)
}

// OnTest triggers when this notifier has been saved
func (m *mobilePush) OnTest() (string, error) {
	msg := &pushArray{
		Message:  "Testing the Mobile Notifier",
		Title:    "Testing Notifications",
		Tokens:   []string{m.Var1},
		Platform: 2,
	}
	body, err := pushRequest(msg)
	if err != nil {
		return "", err
	}
	var output mobileResponse
	err = json.Unmarshal(body, &output)
	if err != nil {
		return string(body), err
	}
	if len(output.Logs) == 0 {
		return string(body), err
	} else {
		firstLog := output.Logs[0].Error
		return string(body), fmt.Errorf("Mobile Notification error: %v", firstLog)
	}
}

// Send will send message to Statping push notifications endpoint
func (m *mobilePush) Send(pushMessage *pushArray) error {
	pushMessage.Tokens = []string{m.Var1}
	pushMessage.Platform = utils.ToInt(m.Var2)
	_, err := pushRequest(pushMessage)
	if err != nil {
		return err
	}
	return nil
}

// OnSave will trigger when this notifier is saved
func (m *mobilePush) OnSave() (string, error) {
	return "", nil
}

func pushRequest(msg *pushArray) ([]byte, error) {
	body, err := json.Marshal(&PushNotification{[]*pushArray{msg}})
	if err != nil {
		return nil, err
	}
	url := "https://push.statping.com/api/push"
	body, _, err = utils.HttpRequest(url, "POST", "application/json", nil, bytes.NewBuffer(body), time.Duration(20*time.Second), false, nil)
	return body, err
}

type PushNotification struct {
	Array []*pushArray `json:"notifications"`
}

type pushArray struct {
	Tokens   []string               `json:"tokens"`
	Platform int64                  `json:"platform"`
	Message  string                 `json:"message"`
	Topic    string                 `json:"topic"`
	Title    string                 `json:"title,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty"`
}

type mobileResponse struct {
	Counts  int                   `json:"counts"`
	Logs    []*mobileResponseLogs `json:"logs"`
	Success string                `json:"success"`
}

type mobileResponseLogs struct {
	Type     string `json:"type"`
	Platform string `json:"platform"`
	Token    string `json:"token"`
	Message  string `json:"message"`
	Error    string `json:"error"`
}
