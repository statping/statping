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
	"github.com/hunterlong/statup/utils"
	"net/http"
	"net/url"
	"strings"
)

const (
	lineNotifyMethod = "line notify"
)

type lineNotifier struct {
	*notifier.Notification
}

var lineNotify = &lineNotifier{&notifier.Notification{
	Method:      lineNotifyMethod,
	Title:       "LINE Notify",
	Description: "LINE Notify will send notifications to your LINE Notify account when services are offline or online. Baed on the <a href=\"https://notify-bot.line.me/doc/en/\">LINE Notify API</a>.",
	Author:      "Kanin Peanviriyakulkit",
	AuthorUrl:   "https://github.com/dogrocker",
	Icon:        "far fa-bell",
	Form: []notifier.NotificationForm{{
		Type:        "text",
		Title:       "Access Token",
		Placeholder: "Insert your Line Notify Access Token here.",
		DbField:     "api_secret",
	}}},
}

// DEFINE YOUR NOTIFICATION HERE.
func init() {
	err := notifier.AddNotifier(lineNotify)
	if err != nil {
		panic(err)
	}
}

// Send will send a HTTP Post with the Authorization to the notify-api.line.me server. It accepts type: string
func (u *lineNotifier) Send(msg interface{}) error {
	message := msg.(string)
	client := new(http.Client)
	v := url.Values{}
	v.Set("message", message)
	req, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", u.GetValue("api_secret")))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func (u *lineNotifier) Select() *notifier.Notification {
	return u.Notification
}

// OnFailure will trigger failing service
func (u *lineNotifier) OnFailure(s *types.Service, f *types.Failure) {
	msg := fmt.Sprintf("Your service '%v' is currently offline!", s.Name)
	u.AddQueue(s.Id, msg)
	u.Online = false
}

// OnSuccess will trigger successful service
func (u *lineNotifier) OnSuccess(s *types.Service) {
	if !u.Online {
		u.ResetUniqueQueue(s.Id)
		msg := fmt.Sprintf("Your service '%v' is back online!", s.Name)
		u.AddQueue(s.Id, msg)
	}
	u.Online = true
}

// OnSave triggers when this notifier has been saved
func (u *lineNotifier) OnSave() error {
	utils.Log(1, fmt.Sprintf("Notification %v is receiving updated information.", u.Method))
	// Do updating stuff here
	return nil
}
