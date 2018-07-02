package notifications

import (
	"bytes"
	"fmt"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

var (
	slackUrl      string
	sentLastMin   int
	slackMessages []string
	SlackComm     *types.Communication
)

func LoadSlack(url string) {
	if url == "" {
		utils.Log(1, "Slack Webhook URL is empty")
		return
	}
	slackUrl = url
}

func SlackRoutine() {
	for _, msg := range slackMessages {
		utils.Log(1, fmt.Sprintf("Sending JSON to Slack Webhook: %v", msg))
		client := http.Client{Timeout: 15 * time.Second}
		_, err := client.Post(slackUrl, "application/json", bytes.NewBuffer([]byte(msg)))
		if err != nil {
			utils.Log(3, fmt.Sprintf("Issue sending Slack notification: %v", err))
		}
		slackMessages = removeStrArray(slackMessages, msg)
	}
	time.Sleep(60 * time.Second)
	if SlackComm.Enabled {
		SlackRoutine()
	}
}

func removeStrArray(arr []string, v string) []string {
	var newArray []string
	for _, i := range arr {
		if i != v {
			newArray = append(newArray, v)
		}
	}
	return newArray
}

func SendSlack(msg string) error {
	if slackUrl == "" {
		return errors.New("Slack Webhook URL has not been set in settings")
	}
	fullMessage := fmt.Sprintf("{\"text\":\"%v\"}", msg)
	slackMessages = append(slackMessages, fullMessage)
	return nil
}
