package notifiers

import (
	"github.com/statping/statping/types/services"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAllNotifiers(t *testing.T) {

	notifiers := []notifierTest{
		{
			Notifier:    Command,
			RequiredENV: nil,
		},
		{
			Notifier:    Discorder,
			RequiredENV: []string{"DISCORD_URL"},
		},
		{
			Notifier:    email,
			RequiredENV: []string{"EMAIL_HOST", "EMAIL_USER", "EMAIL_PASS", "EMAIL_OUTGOING", "EMAIL_SEND_TO", "EMAIL_PORT"},
		},
		{
			Notifier:    Mobile,
			RequiredENV: []string{"MOBILE_ID", "MOBILE_NUMBER"},
		},
		{
			Notifier:    Pushover,
			RequiredENV: []string{"PUSHOVER_TOKEN", "PUSHOVER_API"},
		},
		{
			Notifier:    slacker,
			RequiredENV: []string{"SLACK_URL"},
		},
		{
			Notifier:    Telegram,
			RequiredENV: []string{"TELEGRAM_TOKEN", "TELEGRAM_CHANNEL"},
		},
		{
			Notifier:    Twilio,
			RequiredENV: []string{"TWILIO_SID", "TWILIO_SECRET", "TWILIO_FROM", "TWILIO_TO"},
		},
		{
			Notifier:    Webhook,
			RequiredENV: nil,
		},
	}

	for _, n := range notifiers {

		if !getEnvs(n.RequiredENV) {
			t.Skip()
			continue
		}

		Add(n.Notifier)

		err := n.Notifier.OnSuccess(exampleService)
		assert.Nil(t, err)

		err = n.Notifier.OnFailure(exampleService, exampleFailure)
		assert.Nil(t, err)

		err = n.Notifier.OnTest()
		assert.Nil(t, err)

	}

}

func getEnvs(env []string) bool {
	for _, v := range env {
		if os.Getenv(v) == "" {
			return false
		}
	}
	return true
}

type notifierTest struct {
	Notifier    services.ServiceNotifier
	RequiredENV []string
}
