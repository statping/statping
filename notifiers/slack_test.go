package notifiers

import (
	"github.com/hunterlong/statup/core/notifier"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var (
	SLACK_URL        string
	slackMessage     = `{"text":"this is a test from the Slack notifier!"}`
	slackTestMessage = SlackMessage{
		Service:  TestService,
		Template: FAILURE,
		Time:     time.Now().Unix(),
	}
)

func init() {
	SLACK_URL = os.Getenv("SLACK_URL")
	slacker.Host = SLACK_URL
}

func TestSlackNotifier(t *testing.T) {
	t.Parallel()
	if SLACK_URL == "" {
		t.Log("Slack notifier testing skipped, missing SLACK_URL environment variable")
		t.SkipNow()
	}
	currentCount = CountNotifiers()

	t.Run("Load Slack", func(t *testing.T) {
		slacker.Host = SLACK_URL
		slacker.Delay = time.Duration(100 * time.Millisecond)
		err := notifier.AddNotifier(slacker)
		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", slacker.Author)
		assert.Equal(t, SLACK_URL, slacker.Host)
	})

	t.Run("Load Slack Notifier", func(t *testing.T) {
		notifier.Load()
	})

	t.Run("Slack parse message", func(t *testing.T) {
		err := parseSlackMessage("this is a test message!", slackTestMessage)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(slacker.Queue))
	})

	t.Run("Slack Within Limits", func(t *testing.T) {
		ok, err := slacker.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("Slack Send", func(t *testing.T) {
		err := slacker.Send(slackMessage)
		assert.Nil(t, err)
	})

	t.Run("Slack Queue", func(t *testing.T) {
		go notifier.Queue(slacker)
		time.Sleep(1 * time.Second)
		assert.Equal(t, SLACK_URL, slacker.Host)
		assert.Equal(t, 0, len(slacker.Queue))
	})

}
