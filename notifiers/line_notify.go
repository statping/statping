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
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	LINE_NOTIFY_ID     = 4
	LINE_NOTIFY_METHOD = "line notify"
)

var (
	lineNotify         *LineNotify
	lineNotifyMessages []string
)

type LineNotify struct {
	*Notification
}

// DEFINE YOUR NOTIFICATION HERE.
func init() {
	lineNotify = &LineNotify{&Notification{
		Id:     LINE_NOTIFY_ID,
		Method: LINE_NOTIFY_METHOD,
		Form: []NotificationForm{{
			Type:        "text",
			Title:       "Access Token",
			Placeholder: "Insert your Line Notify Access Token here.",
			DbField:     "api_secret",
		}}},
	}
	err := AddNotifier(lineNotify)
	if err != nil {
		utils.Log(3, err)
	}
}

func (u *LineNotify) postUrl() string {
	return fmt.Sprintf("https://notify-api.line.me/api/notify")
}

// WHEN NOTIFIER LOADS
func (u *LineNotify) Init() error {
	err := u.Install()
	if err == nil {
		notifier, _ := SelectNotification(u.Id)
		forms := u.Form
		u.Notification = notifier
		u.Form = forms
		if u.Enabled {
			go u.Run()
		}
	}

	return err
}

func (u *LineNotify) Test() error {
	msg := fmt.Sprintf("You're Statup Line Notify Notifier is working correctly!")
	SendLineNotify(msg)
	return nil
}

// AFTER NOTIFIER LOADS, IF ENABLED, START A QUEUE PROCESS
func (u *LineNotify) Run() error {
	lineNotifyMessages = uniqueStrings(lineNotifyMessages)
	for _, msg := range lineNotifyMessages {

		if u.CanSend() {
			utils.Log(1, fmt.Sprintf("Sending Line Notify Message"))

			lineNotifyUrl := u.postUrl()
			client := &http.Client{}
			v := url.Values{}
			v.Set("message", msg)
			rb := *strings.NewReader(v.Encode())

			req, err := http.NewRequest("POST", lineNotifyUrl, &rb)
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", u.ApiSecret))
			req.Header.Add("Accept", "application/json")
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			client.Do(req)

			if err != nil {
				utils.Log(3, fmt.Sprintf("Issue sending Line Notify notification: %v", err))
			}
			u.Log(msg)
		}
	}
	lineNotifyMessages = []string{}
	time.Sleep(60 * time.Second)
	if u.Enabled {
		u.Run()
	}
	return nil
}

// CUSTOM FUNCTION FO SENDING LINE NOTIFY MESSAGES
func SendLineNotify(data string) error {
	lineNotifyMessages = append(lineNotifyMessages, data)
	return nil
}

// ON SERVICE FAILURE, DO YOUR OWN FUNCTIONS
func (u *LineNotify) OnFailure(s *types.Service, f *types.Failure) {
	if u.Enabled {
		msg := fmt.Sprintf("Your service '%v' is currently offline!", s.Name)
		SendLineNotify(msg)
	}
}

// ON SERVICE SUCCESS, DO YOUR OWN FUNCTIONS
func (u *LineNotify) OnSuccess(s *types.Service) {

}

// ON SAVE OR UPDATE OF THE NOTIFIER FORM
func (u *LineNotify) OnSave() error {
	utils.Log(1, fmt.Sprintf("Notification %v is receiving updated information.", u.Method))
	// Do updating stuff here
	return nil
}

// ON SERVICE FAILURE, DO YOUR OWN FUNCTIONS
func (u *LineNotify) Install() error {
	var err error
	inDb := lineNotify.Notification.IsInDatabase()
	if !inDb {
		newNotifer, err := InsertDatabase(u.Notification)
		if err != nil {
			utils.Log(3, err)
			return err
		}
		utils.Log(1, fmt.Sprintf("new notifier #%v installed: %v", newNotifer, u.Method))
	}
	return err
}
