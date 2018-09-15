package notifiers

import (
	"github.com/hunterlong/statup/core/notifier"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var (
	TWILIO_SID    = os.Getenv("TWILIO_SID")
	TWILIO_SECRET = os.Getenv("TWILIO_SECRET")
	TWILIO_FROM   = os.Getenv("TWILIO_FROM")
	TWILIO_TO     = os.Getenv("TWILIO_TO")
	twilioMessage = "The Twilio notifier on Statup has been tested!"
)

func init() {
	twilio.ApiKey = TWILIO_SID
	twilio.ApiSecret = TWILIO_SECRET
	twilio.Var1 = TWILIO_TO
	twilio.Var2 = TWILIO_FROM
}

func TestTwilioNotifier(t *testing.T) {
	if TWILIO_SID == "" || TWILIO_SECRET == "" || TWILIO_FROM == "" {
		t.Log("Twilio notifier testing skipped, missing TWILIO_SID environment variable")
		t.SkipNow()
	}
	currentCount = CountNotifiers()

	t.Run("Load Twilio", func(t *testing.T) {
		twilio.ApiKey = TWILIO_SID
		twilio.Delay = time.Duration(100 * time.Millisecond)
		err := notifier.AddNotifier(twilio)
		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", twilio.Author)
		assert.Equal(t, TWILIO_SID, twilio.ApiKey)
		assert.Equal(t, currentCount+1, CountNotifiers())
	})

	t.Run("Load Twilio Notifier", func(t *testing.T) {
		count := notifier.Load()
		assert.Equal(t, currentCount+1, len(count))
	})

	t.Run("Twilio Within Limits", func(t *testing.T) {
		ok, err := twilio.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("Twilio Send", func(t *testing.T) {
		err := twilio.Send(twilioMessage)
		assert.Nil(t, err)
	})

	t.Run("Twilio Queue", func(t *testing.T) {
		go notifier.Queue(twilio)
		time.Sleep(1 * time.Second)
		assert.Equal(t, TWILIO_SID, twilio.ApiKey)
		assert.Equal(t, 0, len(twilio.Queue))
	})

}
