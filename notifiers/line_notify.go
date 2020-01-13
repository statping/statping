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
	"fmt"
	"github.com/hunterlong/statping/core/notifier"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"net/url"
	"strings"
	"time"
)

const (
	lineNotifyMethod = "line notify"
)

type lineNotifier struct {
	*notifier.Notification
}

var LineNotify = &lineNotifier{&notifier.Notification{
	Method:      lineNotifyMethod,
	Title:       "LINE Notify",
	Description: "LINE Notify will send notifications to your LINE Notify account when services are offline or online. Based on the <a href=\"https://notify-bot.line.me/doc/en/\">LINE Notify API</a>.",
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

// Send will send a HTTP Post with the Authorization to the notify-api.line.me server. It accepts type: string
func (u *lineNotifier) Send(msg interface{}) error {
	message := msg.(string)
	v := url.Values{}
	v.Set("message", message)
	headers := []string{fmt.Sprintf("Authorization=Bearer %v", u.ApiSecret)}
	_, _, err := utils.HttpRequest("https://notify-api.line.me/api/notify", "POST", "application/x-www-form-urlencoded", headers, strings.NewReader(v.Encode()), time.Duration(10*time.Second), true)
	return err
}

func (u *lineNotifier) Select() *notifier.Notification {
	return u.Notification
}

// OnFailure will trigger failing service
func (u *lineNotifier) OnFailure(s *types.Service, f *types.Failure) {
	msg := fmt.Sprintf("Your service '%v' is currently offline!", s.Name)
	u.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnSuccess will trigger successful service
func (u *lineNotifier) OnSuccess(s *types.Service) {
	if !s.Online || !s.SuccessNotified {
		var msg string
		if s.UpdateNotify {
			s.UpdateNotify = false
		}
		msg = s.DownText

		u.ResetUniqueQueue(fmt.Sprintf("service_%v", s.Id))
		u.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
	}
}

// OnSave triggers when this notifier has been saved
func (u *lineNotifier) OnSave() error {
	msg := fmt.Sprintf("Notification %v is receiving updated information.", u.Method)
	utils.Log.Infoln(msg)
	u.AddQueue("saved", msg)
	return nil
}
