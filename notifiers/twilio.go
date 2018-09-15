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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hunterlong/statup/core/notifier"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	twilioMethod = "twilioNotifier"
)

type twilio struct {
	*notifier.Notification
}

var twilioNotifier = &twilio{&notifier.Notification{
	Method:      twilioMethod,
	Title:       "twilioNotifier",
	Description: "Receive SMS text messages directly to your cellphone when a service is offline. You can use a twilioNotifier test account with limits. This notifier uses the <a href=\"https://www.twilioNotifier.com/docs/usage/api\">twilioNotifier API</a>.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(10 * time.Second),
	Form: []notifier.NotificationForm{{
		Type:        "text",
		Title:       "Account Sid",
		Placeholder: "Insert your twilioNotifier Account Sid",
		DbField:     "api_key",
	}, {
		Type:        "text",
		Title:       "Account Token",
		Placeholder: "Insert your twilioNotifier Account Token",
		DbField:     "api_secret",
	}, {
		Type:        "text",
		Title:       "SMS to Phone Number",
		Placeholder: "18555555555",
		DbField:     "Var1",
	}, {
		Type:        "text",
		Title:       "From Phone Number",
		Placeholder: "18555555555",
		DbField:     "Var2",
	}}},
}

// DEFINE YOUR NOTIFICATION HERE.
func init() {
	err := notifier.AddNotifier(twilioNotifier)
	if err != nil {
		panic(err)
	}
}

func (u *twilio) Select() *notifier.Notification {
	return u.Notification
}

// Send will send a HTTP Post to the Twilio SMS API. It accepts type: string
func (u *twilio) Send(msg interface{}) error {
	message := msg.(string)
	twilioUrl := fmt.Sprintf("https://api.twilioNotifier.com/2010-04-01/Accounts/%v/Messages.json", u.GetValue("api_key"))
	client := &http.Client{}
	v := url.Values{}
	v.Set("To", "+"+u.Var1)
	v.Set("From", "+"+u.Var2)
	v.Set("Body", message)
	rb := *strings.NewReader(v.Encode())
	req, err := http.NewRequest("POST", twilioUrl, &rb)
	req.SetBasicAuth(u.ApiKey, u.ApiSecret)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	contents, _ := ioutil.ReadAll(res.Body)
	success, twilioRes := twilioSuccess(contents)
	if !success {
		return errors.New(fmt.Sprintf("twilioNotifier didn't receive the expected status of 'enque' from API got: %v", twilioRes))
	}
	return nil
}

// OnFailure will trigger failing service
func (u *twilio) OnFailure(s *types.Service, f *types.Failure) {
	msg := fmt.Sprintf("Your service '%v' is currently offline!", s.Name)
	u.AddQueue(msg)
}

// OnSuccess will trigger successful service
func (u *twilio) OnSuccess(s *types.Service) {

}

// OnSave triggers when this notifier has been saved
func (u *twilio) OnSave() error {
	utils.Log(1, fmt.Sprintf("Notification %v is receiving updated information.", u.Method))

	// Do updating stuff here

	return nil
}

func twilioSuccess(res []byte) (bool, TwilioResponse) {
	var obj TwilioResponse
	json.Unmarshal(res, &obj)
	if obj.Status == "queued" {
		return true, obj
	}
	return false, obj
}

type TwilioResponse struct {
	Sid                 string      `json:"sid"`
	DateCreated         string      `json:"date_created"`
	DateUpdated         string      `json:"date_updated"`
	DateSent            interface{} `json:"date_sent"`
	AccountSid          string      `json:"account_sid"`
	To                  string      `json:"to"`
	From                string      `json:"from"`
	MessagingServiceSid interface{} `json:"messaging_service_sid"`
	Body                string      `json:"body"`
	Status              string      `json:"status"`
	NumSegments         string      `json:"num_segments"`
	NumMedia            string      `json:"num_media"`
	Direction           string      `json:"direction"`
	APIVersion          string      `json:"api_version"`
	Price               interface{} `json:"price"`
	PriceUnit           string      `json:"price_unit"`
	ErrorCode           interface{} `json:"error_code"`
	ErrorMessage        interface{} `json:"error_message"`
	URI                 string      `json:"uri"`
	SubresourceUris     struct {
		Media string `json:"media"`
	} `json:"subresource_uris"`
}
