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
	"os"
	"time"
)

type mobilePush struct {
	*notifier.Notification
}

var mobile = &mobilePush{&notifier.Notification{
	Method: "mobile",
	Title:  "Mobile Notifications",
	Description: `Receive push notifications on your Android or iPhone devices using the Statping App. You can scan the Authentication QR Code found in Settings to get the mobile app setup in seconds.
				 <p align="center"><a href="https://play.google.com/store/apps/details?id=com.statping"><img src="https://img.cjx.io/google-play.svg"></a><a href="https://itunes.apple.com/us/app/apple-store/id1445513219"><img src="https://img.cjx.io/app-store-badge.svg"></a></p>`,
	Author:    "Hunter Long",
	AuthorUrl: "https://github.com/hunterlong",
	Delay:     time.Duration(5 * time.Second),
	Icon:      "fas fa-mobile-alt",
	Form: []notifier.NotificationForm{{
		Type:        "text",
		Title:       "Device Identifiers",
		Placeholder: "A list of your mobile device push notification ID's.",
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

// init the discord notifier
func init() {
	err := notifier.AddNotifier(mobile)
	if err != nil {
		panic(err)
	}
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
	msg := &PushArray{
		Message: fmt.Sprintf("Your service '%v' is currently failing! Reason: %v", s.Name, f.Issue),
		Title:   "Service Offline",
		Data:    data,
	}
	u.AddQueue(s.Id, msg)
	u.Online = false
}

// OnSuccess will trigger successful service
func (u *mobilePush) OnSuccess(s *types.Service) {
	data := dataJson(s, nil)
	if !u.Online {
		u.ResetUniqueQueue(s.Id)
		msg := &PushArray{
			Message: fmt.Sprintf("Your service '%v' is back online!", s.Name),
			Title:   "Service Online",
			Data:    data,
		}
		u.AddQueue(s.Id, msg)
	}
	u.Online = true
}

// OnSave triggers when this notifier has been saved
func (u *mobilePush) OnSave() error {
	msg := &PushArray{
		Message: "The Mobile Notifier has been saved",
		Title:   "Notification Saved",
	}
	u.AddQueue(0, msg)
	return nil
}

// OnTest triggers when this notifier has been saved
func (u *mobilePush) OnTest() error {
	return nil
}

// Send will send message to Statping push notifications endpoint
func (u *mobilePush) Send(msg interface{}) error {
	pushMessage := msg.(*PushArray)
	pushMessage.Tokens = []string{u.Var1}
	pushMessage.Platform = utils.ToInt(u.Var2)
	err := pushRequest(pushMessage)
	if err != nil {
		return err
	}
	return nil
}

func pushRequest(msg *PushArray) error {
	if msg.Platform == 1 {
		msg.Title = ""
	}
	body, _ := json.Marshal(&PushNotification{[]*PushArray{msg}})
	url := "https://push.statping.com/api/push"
	if os.Getenv("GO_ENV") == "test" {
		url = "https://pushdev.statping.com/api/push"
	}
	_, _, err := utils.HttpRequest(url, "POST", "application/json", nil, bytes.NewBuffer(body), time.Duration(20*time.Second))
	return err
}

type PushNotification struct {
	Array []*PushArray `json:"notifications"`
}

type PushArray struct {
	Tokens   []string               `json:"tokens"`
	Platform int64                  `json:"platform"`
	Message  string                 `json:"message"`
	Title    string                 `json:"title,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty"`
}

type MobileResponse struct {
	Counts  int           `json:"counts"`
	Logs    []interface{} `json:"logs"`
	Success string        `json:"success"`
}
