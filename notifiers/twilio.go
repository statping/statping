package notifiers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/notifier"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
)

var _ notifier.Notifier = (*twilio)(nil)

type twilio struct {
	*notifications.Notification
}

func (t *twilio) Select() *notifications.Notification {
	return t.Notification
}

func (t *twilio) Valid(values notifications.Values) error {
	return nil
}

var Twilio = &twilio{&notifications.Notification{
	Method:      "twilio",
	Title:       "Twilio",
	Description: "Receive SMS text messages directly to your cellphone when a service is offline. You can use a Twilio test account with limits. This notifier uses the <a href=\"https://www.twilio.com/docs/usage/api\">Twilio API</a>.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Icon:        "far fa-comment-alt",
	Delay:       time.Duration(10 * time.Second),
	SuccessData: null.NewNullString("Your service '{{.Service.Name}}' is currently online!"),
	FailureData: null.NewNullString("Your service '{{.Service.Name}}' is currently offline!"),
	DataType:    "text",
	Limits:      15,
	Form: []notifications.NotificationForm{{
		Type:        "text",
		Title:       "Account SID",
		Placeholder: "Insert your Twilio Account SID",
		DbField:     "api_key",
		Required:    true,
	}, {
		Type:        "text",
		Title:       "Account Token",
		Placeholder: "Insert your Twilio Account Token",
		DbField:     "api_secret",
		Required:    true,
	}, {
		Type:        "number",
		Title:       "SMS to Phone Number",
		Placeholder: "18555555555",
		DbField:     "Var1",
		Required:    true,
	}, {
		Type:        "number",
		Title:       "From Phone Number",
		Placeholder: "18555555555",
		DbField:     "Var2",
		Required:    true,
	}}},
}

// Send will send a HTTP Post to the Twilio SMS API. It accepts type: string
func (t *twilio) sendMessage(message string) (string, error) {
	twilioUrl := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", t.ApiKey.String)

	v := url.Values{}
	v.Set("To", "+"+t.Var1.String)
	v.Set("From", "+"+t.Var2.String)
	v.Set("Body", message)
	rb := strings.NewReader(v.Encode())

	authHeader := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", t.ApiKey.String, t.ApiSecret.String)))

	contents, _, err := utils.HttpRequest(twilioUrl, "POST", "application/x-www-form-urlencoded", []string{"Authorization=Basic " + authHeader}, rb, 10*time.Second, true, nil)
	success, _ := twilioSuccess(contents)
	if !success {
		errorOut := twilioError(contents)
		if errorOut.Code == 0 {
			return string(contents), nil
		}
		out := fmt.Sprintf("Error code %v - %s", errorOut.Code, errorOut.Message)
		return string(contents), errors.New(out)
	}
	return string(contents), err
}

// OnFailure will trigger failing service
func (t *twilio) OnFailure(s services.Service, f failures.Failure) (string, error) {
	msg := ReplaceVars(t.FailureData.String, s, f)
	return t.sendMessage(msg)
}

// OnSuccess will trigger successful service
func (t *twilio) OnSuccess(s services.Service) (string, error) {
	msg := ReplaceVars(t.SuccessData.String, s, failures.Failure{})
	return t.sendMessage(msg)
}

// OnTest will test the Twilio SMS messaging
func (t *twilio) OnTest() (string, error) {
	msg := fmt.Sprintf("Testing the Twilio SMS Notifier")
	return t.sendMessage(msg)
}

// OnSave will trigger when this notifier is saved
func (t *twilio) OnSave() (string, error) {
	return "", nil
}

func twilioSuccess(res []byte) (bool, twilioResponse) {
	var obj twilioResponse
	json.Unmarshal(res, &obj)
	if obj.Status == "queued" {
		return true, obj
	}
	return false, obj
}

func twilioError(res []byte) twilioErrorObj {
	var obj twilioErrorObj
	json.Unmarshal(res, &obj)
	return obj
}

type twilioErrorObj struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}

type twilioResponse struct {
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
