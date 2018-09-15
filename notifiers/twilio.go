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
	Method:      TWILIO_METHOD,
	Title:       "Twilio",
	Description: "Receive SMS text messages directly to your cellphone when a service is offline. You can use a Twilio test account with limits. This notifier uses the <a href=\"https://www.twilio.com/docs/usage/api\">Twilio API</a>.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(10 * time.Second),
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

func (u *Twilio) Test() error {
	utils.Log(1, "Twilio notifier loaded")
	msg := fmt.Sprintf("You're Statup Twilio Notifier is working correctly!")
	u.AddQueue(msg)
	return nil
}

func (u *Twilio) Select() *notifier.Notification {
	return u.Notification
}

func (u *Twilio) Send(msg interface{}) error {
	message := msg.(string)
	twilioUrl := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%v/Messages.json", u.GetValue("api_key"))
	client := &http.Client{}
	v := url.Values{}
	v.Set("To", u.Var1)
	v.Set("From", u.Var2)
	v.Set("Body", message)
	rb := *strings.NewReader(v.Encode())
	req, err := http.NewRequest("POST", twilioUrl, &rb)
	req.SetBasicAuth(u.ApiKey, u.ApiSecret)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client.Do(req)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Issue sending Twilio notification: %v", err))
		return err
	}
	return nil
}

// ON SERVICE FAILURE, DO YOUR OWN FUNCTIONS
func (u *Twilio) OnFailure(s *types.Service, f *types.Failure) {
	msg := fmt.Sprintf("Your service '%v' is currently offline!", s.Name)
	u.AddQueue(msg)
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
