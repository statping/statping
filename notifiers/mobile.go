// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
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
	"fmt"
	"github.com/hunterlong/statup/core/notifier"
	"github.com/hunterlong/statup/types"
	"github.com/oliveroneill/exponent-server-sdk-golang/sdk"
	"strings"
	"time"
)

type mobilePush struct {
	*notifier.Notification
}

var mobile = &mobilePush{&notifier.Notification{
	Method:      "mobile",
	Title:       "Mobile Notifications",
	Description: "Receive push notifications on your Android or iPhone devices using the Statup App. You can scan the Authentication QR Code found in Settings to get the mobile app setup in seconds.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(5 * time.Second),
	Icon:        "fas fa-mobile-alt",
	Form: []notifier.NotificationForm{{
		Type:        "text",
		Title:       "Device Identifiers",
		Placeholder: "A list of your mobile device push notification ID's.",
		DbField:     "var1",
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

// OnFailure will trigger failing service
func (u *mobilePush) OnFailure(s *types.Service, f *types.Failure) {
	msg := &expo.PushMessage{
		Body:     fmt.Sprintf("Your service '%v' is currently failing! Reason: %v", s.Name, f.Issue),
		Sound:    "default",
		Title:    "Service Offline",
		Priority: expo.DefaultPriority,
	}
	u.AddQueue(s.Id, msg)
	u.Online = false
}

// OnSuccess will trigger successful service
func (u *mobilePush) OnSuccess(s *types.Service) {
	if !u.Online {
		u.ResetUniqueQueue(s.Id)
		msg := &expo.PushMessage{
			Body:     fmt.Sprintf("Your service '%v' is back online!", s.Name),
			Sound:    "default",
			Title:    "Service Online",
			Priority: expo.DefaultPriority,
		}
		u.AddQueue(s.Id, msg)
	}
	u.Online = true
}

// OnSave triggers when this notifier has been saved
func (u *mobilePush) OnSave() error {
	msg := &expo.PushMessage{
		Body:     "This is a test notification",
		Sound:    "default",
		Title:    "Notification Test",
		Priority: expo.DefaultPriority,
	}
	u.AddQueue(0, msg)
	return nil
}

// OnTest triggers when this notifier has been saved
func (u *mobilePush) OnTest() error {
	return nil
}

// Send will send message to expo mobile push notifications endpoint
func (u *mobilePush) Send(msg interface{}) error {
	pushMessage := msg.(*expo.PushMessage)
	client := expo.NewPushClient(nil)
	splitIds := strings.Split(u.Var1, ",")

	for _, id := range splitIds {
		pushToken, err := expo.NewExponentPushToken(expo.ExponentPushToken(id))
		if err != nil {
			return err
		}
		pushMessage.To = pushToken
		response, err := client.Publish(pushMessage)
		if err != nil {
			return err
		}
		if response.ValidateResponse() != nil {
			fmt.Println(response.PushMessage.To, "failed")
		}
	}

	return nil
}
