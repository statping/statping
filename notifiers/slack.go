package notifiers

import (
	"bytes"
	"fmt"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"time"
)

const (
	SLACK_ID     int64 = 2
	SLACK_METHOD       = "slack"
)

var (
	slacker       *Notification
	slackMessages []string
)

// DEFINE YOUR NOTIFICATION HERE.
func init() {
	slacker = &Notification{
		Id:     SLACK_ID,
		Method: SLACK_METHOD,
		Host:   "https://webhooksurl.slack.com/***",
		Form: []NotificationForm{{
			Type:        "text",
			Title:       "Incoming Webhook Url",
			Placeholder: "Insert your Slack webhook URL here.",
		}},
	}
	add(slacker)
}

// WHEN NOTIFIER LOADS
func (u *Notification) Init() error {
	err := SendSlack("its online")
	go u.Run()
	return err
}

// AFTER NOTIFIER LOADS, IF ENABLED, START A QUEUE PROCESS
func (u *Notification) Run() error {
	for _, msg := range slackMessages {
		utils.Log(1, fmt.Sprintf("Sending JSON to Slack Webhook: %v", msg))
		client := http.Client{Timeout: 15 * time.Second}
		_, err := client.Post("https://hooks.slack.com/services/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", "application/json", bytes.NewBuffer([]byte(msg)))
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
func SendSlack(msg string) error {
	//if slackUrl == "" {
	//	return errors.New("Slack Webhook URL has not been set in settings")
	//}
	fullMessage := fmt.Sprintf("{\"text\":\"%v\"}", msg)
	slackMessages = append(slackMessages, fullMessage)
	return nil
}

// ON SERVICE FAILURE, DO YOUR OWN FUNCTIONS
func (u *Notification) OnFailure() error {
	utils.Log(1, fmt.Sprintf("Notification %v is receiving a failure notification.", u.Method))

	// Do failing stuff here!

	return nil
}

// ON SERVICE SUCCESS, DO YOUR OWN FUNCTIONS
func (u *Notification) OnSuccess() error {
	utils.Log(1, fmt.Sprintf("Notification %v is receiving a successful notification.", u.Method))

	// Do checking or any successful things here

	return nil
}

// ON SAVE OR UPDATE OF THE NOTIFIER FORM
func (u *Notification) OnSave() error {
	utils.Log(1, fmt.Sprintf("Notification %v is receiving updated information.", u.Method))

	// Do updating stuff here

	return nil
}
