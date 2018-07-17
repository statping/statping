package notifiers

import (
	"bytes"
	"fmt"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"sync"
	"text/template"
	"time"
)

const (
	SLACK_ID         = 2
	SLACK_METHOD     = "slack"
	FAILING_TEMPLATE = `{ "attachments": [ { "fallback": "Service {{.Service.Name}} - is currently failing", "text": "<{{.Service.Domain}}|{{.Service.Name}}> - Your Statup service '{{.Service.Name}}' has just received a Failure notification with a HTTP Status code of {{.Service.LastStatusCode}}.", "fields": [ { "title": "Expected", "value": "{{.Service.Expected}}", "short": true }, { "title": "Status Code", "value": "{{.Service.LastStatusCode}}", "short": true } ], "color": "#FF0000", "thumb_url": "https://statup.io", "footer": "Statup", "footer_icon": "https://img.cjx.io/statuplogo32.png" } ] }`
	SUCCESS_TEMPLATE = `{ "attachments": [ { "fallback": "Service {{.Service.Name}} - is now back online", "text": "<{{.Service.Domain}}|{{.Service.Name}}> - Your Statup service '{{.Service.Name}}' has just received a Failure notification.", "fields": [ { "title": "Issue", "value": "Awesome Project", "short": true }, { "title": "Status Code", "value": "{{.Service.LastStatusCode}}", "short": true } ], "color": "#00FF00", "thumb_url": "https://statup.io", "footer": "Statup", "footer_icon": "https://img.cjx.io/statuplogo32.png" } ] }`
	TEST_TEMPLATE    = `{"text":"{{.}}"}`
)

var (
	slacker       *Slack
	slackMessages []string
	messageLock   *sync.Mutex
)

type Slack struct {
	*Notification
}

type slackMessage struct {
	Service *types.Service
	Time    int64
}

// DEFINE YOUR NOTIFICATION HERE.
func init() {
	slacker = &Slack{&Notification{
		Id:     SLACK_ID,
		Method: SLACK_METHOD,
		Host:   "https://webhooksurl.slack.com/***",
		Form: []NotificationForm{{
			Id:          2,
			Type:        "text",
			Title:       "Incoming Webhook Url",
			Placeholder: "Insert your Slack webhook URL here.",
			DbField:     "Host",
		}}},
	}
	add(slacker)
	messageLock = new(sync.Mutex)
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
	msg := fmt.Sprintf("You're Statup Slack Notifier is working correctly!")
	SendSlack(TEST_TEMPLATE, msg)
	return nil
}

// AFTER NOTIFIER LOADS, IF ENABLED, START A QUEUE PROCESS
func (u *Slack) Run() error {
	messageLock.Lock()
	slackMessages = uniqueStrings(slackMessages)
	for _, msg := range slackMessages {

		if u.CanSend() {
			utils.Log(1, fmt.Sprintf("Sending JSON to Slack Webhook"))
			client := http.Client{Timeout: 15 * time.Second}
			_, err := client.Post(u.Host, "application/json", bytes.NewBuffer([]byte(msg)))
			if err != nil {
				utils.Log(3, fmt.Sprintf("Issue sending Slack notification: %v", err))
			}
			u.Log(msg)
		}
	}
	slackMessages = []string{}
	messageLock.Unlock()
	time.Sleep(60 * time.Second)
	if u.Enabled {
		u.Run()
	}
	return nil
}

// CUSTOM FUNCTION FO SENDING SLACK MESSAGES
func SendSlack(temp string, data interface{}) error {
	messageLock.Lock()
	buf := new(bytes.Buffer)
	slackTemp, _ := template.New("slack").Parse(temp)
	slackTemp.Execute(buf, data)
	slackMessages = append(slackMessages, buf.String())
	messageLock.Unlock()
	return nil
}

// ON SERVICE FAILURE, DO YOUR OWN FUNCTIONS
func (u *Slack) OnFailure(s *types.Service) error {
	if u.Enabled {
		message := slackMessage{
			Service: s,
			Time:    time.Now().Unix(),
		}
		SendSlack(FAILING_TEMPLATE, message)
	}
	return nil
}

// ON SERVICE SUCCESS, DO YOUR OWN FUNCTIONS
func (u *Slack) OnSuccess(s *types.Service) error {
	if u.Enabled {
		//message := slackMessage{
		//	Service: s,
		//	Time:    time.Now().Unix(),
		//}
		//SendSlack(SUCCESS_TEMPLATE, message)
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
	inDb, err := slacker.Notification.IsInDatabase()
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
