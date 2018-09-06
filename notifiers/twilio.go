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
	TWILIO_ID     = 3
	TWILIO_METHOD = "twilio"
)

var (
	twilio         *Twilio
	twilioMessages []string
)

type Twilio struct {
	*Notification
}

type twilioMessage struct {
	Service *types.Service
	Time    int64
}

// DEFINE YOUR NOTIFICATION HERE.
func init() {
	twilio = &Twilio{&Notification{
		Id:     TWILIO_ID,
		Method: TWILIO_METHOD,
		Form: []NotificationForm{{
			Id:          3,
			Type:        "text",
			Title:       "Account Sid",
			Placeholder: "Insert your Twilio Account Sid",
			DbField:     "api_key",
		}, {
			Id:          3,
			Type:        "text",
			Title:       "Account Token",
			Placeholder: "Insert your Twilio Account Token",
			DbField:     "api_secret",
		}, {
			Id:          3,
			Type:        "text",
			Title:       "SMS to Phone Number",
			Placeholder: "+18555555555",
			DbField:     "Var1",
		}, {
			Id:          3,
			Type:        "text",
			Title:       "From Phone Number",
			Placeholder: "+18555555555",
			DbField:     "Var2",
		}}},
	}
	add(twilio)
}

// Select Obj
func (u *Twilio) Select() *Notification {
	return u.Notification
}

func (u *Twilio) postUrl() string {
	return fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%v/Messages.json", u.ApiKey)
}

// WHEN NOTIFIER LOADS
func (u *Twilio) Init() error {
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

func (u *Twilio) Test() error {
	msg := fmt.Sprintf("You're Statup Twilio Notifier is working correctly!")
	SendTwilio(msg)
	return nil
}

// AFTER NOTIFIER LOADS, IF ENABLED, START A QUEUE PROCESS
func (u *Twilio) Run() error {
	twilioMessages = uniqueStrings(twilioMessages)
	for _, msg := range twilioMessages {

		if u.CanSend() {
			utils.Log(1, fmt.Sprintf("Sending Twilio Message"))

			twilioUrl := u.postUrl()
			client := &http.Client{}
			v := url.Values{}
			v.Set("To", u.Var1)
			v.Set("From", u.Var2)
			v.Set("Body", msg)
			rb := *strings.NewReader(v.Encode())

			req, err := http.NewRequest("POST", twilioUrl, &rb)
			req.SetBasicAuth(u.ApiKey, u.ApiSecret)
			req.Header.Add("Accept", "application/json")
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			client.Do(req)

			if err != nil {
				utils.Log(3, fmt.Sprintf("Issue sending Twilio notification: %v", err))
			}
			u.Log(msg)
		}
	}
	twilioMessages = []string{}
	time.Sleep(60 * time.Second)
	if u.Enabled {
		u.Run()
	}
	return nil
}

// CUSTOM FUNCTION FO SENDING TWILIO MESSAGES
func SendTwilio(data string) error {
	twilioMessages = append(twilioMessages, data)
	return nil
}

// ON SERVICE FAILURE, DO YOUR OWN FUNCTIONS
func (u *Twilio) OnFailure(s *types.Service) error {
	if u.Enabled {
		msg := fmt.Sprintf("Your service '%v' is currently offline!", s.Name)
		SendTwilio(msg)
	}
	return nil
}

// ON SERVICE SUCCESS, DO YOUR OWN FUNCTIONS
func (u *Twilio) OnSuccess(s *types.Service) error {
	if u.Enabled {

	}
	return nil
}

// ON SAVE OR UPDATE OF THE NOTIFIER FORM
func (u *Twilio) OnSave() error {
	utils.Log(1, fmt.Sprintf("Notification %v is receiving updated information.", u.Method))

	// Do updating stuff here

	return nil
}

// ON SERVICE FAILURE, DO YOUR OWN FUNCTIONS
func (u *Twilio) Install() error {
	inDb := twilio.Notification.IsInDatabase()
	if !inDb {
		newNotifer, err := InsertDatabase(u.Notification)
		if err != nil {
			utils.Log(3, err)
			return err
		}
		utils.Log(1, fmt.Sprintf("new notifier #%v installed: %v", newNotifer, u.Method))
	}
	return nil
}
