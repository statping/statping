// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package notifiers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hunterlong/statping/core/notifier"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"time"
)

const mobileIdentifier = "com.statping"

type mobilePush struct {
	*notifier.Notification
}

var Mobile = &mobilePush{&notifier.Notification{
	Method: "mobile",
	Title:  "Mobile Notifications",
	Description: `Receive push notifications on your Mobile device using the Statping App. You can scan the Authentication QR Code found in Settings to get the Mobile app setup in seconds.
				 <p align="center"><a href="https://play.google.com/store/apps/details?id=com.statping"><img src="https://img.cjx.io/google-play.svg"></a><a href="https://itunes.apple.com/us/app/apple-store/id1445513219"><img src="https://img.cjx.io/app-store-badge.svg"></a></p>`,
	Author:    "Hunter Long",
	AuthorUrl: "https://github.com/hunterlong",
	Delay:     time.Duration(5 * time.Second),
	Icon:      "fas fa-mobile-alt",
	Form: []notifier.NotificationForm{{
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

func (u *mobilePush) Select() *notifier.Notification {
	return u.Notification
}

func dataJson(s *types.Service, f *types.Failure) map[string]interface{} {
	serviceId := "0"
	if s != nil {
		serviceId = utils.ToString(s.Id)
	}
	online := "online"
	if !s.Online {
		online = "offline"
	}
	issue := ""
	if f != nil {
		issue = f.Issue
	}
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
func (u *mobilePush) OnFailure(s *types.Service, f *types.Failure) {
	data := dataJson(s, f)
	msg := &pushArray{
		Message: fmt.Sprintf("Your service '%v' is currently failing! Reason: %v", s.Name, f.Issue),
		Title:   "Service Offline",
		Topic:   mobileIdentifier,
		Data:    data,
	}
	u.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnSuccess will trigger successful service
func (u *mobilePush) OnSuccess(s *types.Service) {
	data := dataJson(s, nil)
	if !s.Online || !s.SuccessNotified {
		var msgStr string
		if s.UpdateNotify {
			s.UpdateNotify = false
		}
		msgStr = s.DownText

		u.ResetUniqueQueue(fmt.Sprintf("service_%v", s.Id))
		msg := &pushArray{
			Message: msgStr,
			Title:   "Service Online",
			Topic:   mobileIdentifier,
			Data:    data,
		}
		u.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
	}
}

// OnSave triggers when this notifier has been saved
func (u *mobilePush) OnSave() error {
	return nil
}

// OnTest triggers when this notifier has been saved
func (u *mobilePush) OnTest() error {
	msg := &pushArray{
		Message:  "Testing the Mobile Notifier",
		Title:    "Testing Notifications",
		Topic:    mobileIdentifier,
		Tokens:   []string{u.Var1},
		Platform: utils.ToInt(u.Var2),
	}
	body, err := pushRequest(msg)
	if err != nil {
		return err
	}
	var output mobileResponse
	err = json.Unmarshal(body, &output)
	if err != nil {
		return err
	}
	if len(output.Logs) == 0 {
		return nil
	} else {
		firstLog := output.Logs[0].Error
		return fmt.Errorf("Mobile Notification error: %v", firstLog)
	}
	return err
}

// Send will send message to Statping push notifications endpoint
func (u *mobilePush) Send(msg interface{}) error {
	pushMessage := msg.(*pushArray)
	pushMessage.Tokens = []string{u.Var1}
	pushMessage.Platform = utils.ToInt(u.Var2)
	_, err := pushRequest(pushMessage)
	if err != nil {
		return err
	}
	return nil
}

func pushRequest(msg *pushArray) ([]byte, error) {
	if msg.Platform == 1 {
		msg.Title = ""
	}
	body, err := json.Marshal(&PushNotification{[]*pushArray{msg}})
	if err != nil {
		return nil, err
	}
	url := "https://push.statping.com/api/push"
	body, _, err = utils.HttpRequest(url, "POST", "application/json", nil, bytes.NewBuffer(body), time.Duration(20*time.Second), true)
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
