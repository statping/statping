package notifiers

import (
	"bytes"
	"fmt"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"time"
	"text/template"
)

const (
	SLACK_ID     = 2
	SLACK_METHOD = "slack"
	SERVICE_TEMPLATE = `{ "attachments": [ { "fallback": "ReferenceError - UI is not defined: https://honeybadger.io/path/to/event/", "text": "<https://honeybadger.io/path/to/event/|Google> - Your Statup service 'Google' has just received a Failure notification.", "fields": [ { "title": "Issue", "value": "Awesome Project", "short": true }, { "title": "Response", "value": "production", "short": true } ], "color": "#FF0000", "thumb_url": "http://example.com/path/to/thumb.png", "footer": "Statup", "footer_icon": "https://img.cjx.io/statuplogo32.png", "ts": 123456789 } ] }`
	TEST_TEMPLATE = `{"text":"%{{.Message}}"}`
)

var (
	slacker       *Slack
	slackMessages []string
)

type Slack struct {
	*Notification
}

// DEFINE YOUR NOTIFICATION HERE.
func init() {
	slacker = &Slack{&Notification{
		Id:     SLACK_ID,
		Method: SLACK_METHOD,
		Host:   "https://webhooksurl.slack.com/***",
		Form: []NotificationForm{{
			id:          2,
			Type:        "text",
			Title:       "Incoming Webhook Url",
			Placeholder: "Insert your Slack webhook URL here.",
			DbField:     "Host",
		}}},
	}
	add(slacker)
}

// Select Obj
func (u *Slack) Select() *Notification {
	return u.Notification
}

// WHEN NOTIFIER LOADS
func (u *Slack) Init() error {

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

func (u *Slack) Test() error {
	SendSlack(TEST_TEMPLATE, nil)
	return nil
}

// AFTER NOTIFIER LOADS, IF ENABLED, START A QUEUE PROCESS
func (u *Slack) Run() error {
	for _, msg := range slackMessages {
		utils.Log(1, fmt.Sprintf("Sending JSON to Slack Webhook: %v", msg))
		client := http.Client{Timeout: 15 * time.Second}
		_, err := client.Post(u.Host, "application/json", bytes.NewBuffer([]byte(msg)))
		if err != nil {
			utils.Log(3, fmt.Sprintf("Issue sending Slack notification: %v", err))
		}
		slackMessages = uniqueMessages(slackMessages, msg)
	}
	time.Sleep(60 * time.Second)
	if u.Enabled {
		u.Run()
	}
	return nil
}

// CUSTOM FUNCTION FO SENDING SLACK MESSAGES
func SendSlack(temp string, data ...interface{}) error {
	buf := new(bytes.Buffer)
	slackTemp, _ := template.New("slack").Parse(temp)
	slackTemp.Execute(buf, data)
	slackMessages = append(slackMessages, buf.String())
	return nil
}

// ON SERVICE FAILURE, DO YOUR OWN FUNCTIONS
func (u *Slack) OnFailure(data map[string]interface{}) error {
	if u.Enabled {
		utils.Log(1, fmt.Sprintf("Notification %v is receiving a failure notification.", u.Method))
		// Do failing stuff here!
	}
	return nil
}

// ON SERVICE SUCCESS, DO YOUR OWN FUNCTIONS
func (u *Slack) OnSuccess(data map[string]interface{}) error {
	if u.Enabled {
		utils.Log(1, fmt.Sprintf("Notification %v is receiving a successful notification. %v", u.Method, data))

		//domain := data["Domain"]
		//expected := data["Expected"]
		//expectedStatus := data["ExpectedStatus"]
		failures := data["Failures"]
		response := data["LastResponse"]

		fullMessage := fmt.Sprintf(`{ "attachments": [ { "fallback": "Service is currently offline", "text": "Service is currently offline", "fields": [ { "title": "Issue", "value": "%v", "short": true }, { "title": "Response", "value": "%v", "short": true } ], "color": "#FF0000", "thumb_url": "http://example.com/path/to/thumb.png", "footer": "Statup", "footer_icon": "https://img.cjx.io/statuplogo32.png", "ts": %v } ] }`, failures, response, time.Now().Unix())
		slackMessages = append(slackMessages, fullMessage)

		// Do checking or any successful things here
	}
	return nil
}

// ON SAVE OR UPDATE OF THE NOTIFIER FORM
func (u *Slack) OnSave() error {
	utils.Log(1, fmt.Sprintf("Notification %v is receiving updated information.", u.Method))

	// Do updating stuff here

	return nil
}

// ON SERVICE FAILURE, DO YOUR OWN FUNCTIONS
func (u *Slack) Install() error {
	inDb, err := slacker.Notification.isInDatabase()
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
