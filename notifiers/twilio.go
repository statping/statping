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
	"time"
)

const (
	TWILIO_ID     = 3
	TWILIO_METHOD = "twilio"
)

var (
	twilioMessages []string
)

type Twilio struct {
	*notifier.Notification
}

var twilio = &Twilio{&notifier.Notification{
	Method: TWILIO_METHOD,
	Form: []notifier.NotificationForm{{
		Type:        "text",
		Title:       "Account Sid",
		Placeholder: "Insert your Twilio Account Sid",
		DbField:     "api_key",
	}, {
		Type:        "text",
		Title:       "Account Token",
		Placeholder: "Insert your Twilio Account Token",
		DbField:     "api_secret",
	}, {
		Type:        "text",
		Title:       "SMS to Phone Number",
		Placeholder: "+18555555555",
		DbField:     "Var1",
	}, {
		Type:        "text",
		Title:       "From Phone Number",
		Placeholder: "+18555555555",
		DbField:     "Var2",
	}}},
}

// DEFINE YOUR NOTIFICATION HERE.
func init() {
	err := notifier.AddNotifier(twilio)
	if err != nil {
		panic(err)
	}
}

func (u *Twilio) postUrl() string {
	return fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%v/Messages.json", u.ApiKey)
}

func (u *Twilio) Test() error {
	utils.Log(1, "Twilio notifier loaded")
	msg := fmt.Sprintf("You're Statup Twilio Notifier is working correctly!")
	SendTwilio(msg)
	return nil
}

// AFTER NOTIFIER LOADS, IF ENABLED, START A QUEUE PROCESS
func (u *Twilio) Run() error {
	twilioMessages = notifier.UniqueStrings(twilioMessages)
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
func (u *Twilio) OnFailure(s *types.Service, f *types.Failure) {
	if u.Enabled {
		msg := fmt.Sprintf("Your service '%v' is currently offline!", s.Name)
		SendTwilio(msg)
	}
}

// ON SERVICE SUCCESS, DO YOUR OWN FUNCTIONS
func (u *Twilio) OnSuccess(s *types.Service) {

}

// ON SAVE OR UPDATE OF THE NOTIFIER FORM
func (u *Twilio) OnSave() error {
	utils.Log(1, fmt.Sprintf("Notification %v is receiving updated information.", u.Method))

	// Do updating stuff here

	return nil
}
