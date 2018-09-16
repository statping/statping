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
	twilioMessage = "The twilioNotifier notifier on Statup has been tested!"
)

func init() {
	twilioNotifier.ApiKey = TWILIO_SID
	twilioNotifier.ApiSecret = TWILIO_SECRET
	twilioNotifier.Var1 = TWILIO_TO
	twilioNotifier.Var2 = TWILIO_FROM
}

func TestTwilioNotifier(t *testing.T) {
	t.Parallel()
	if TWILIO_SID == "" || TWILIO_SECRET == "" || TWILIO_FROM == "" {
		t.Log("twilioNotifier notifier testing skipped, missing TWILIO_SID environment variable")
		t.SkipNow()
	}
	currentCount = CountNotifiers()

	t.Run("Load twilioNotifier", func(t *testing.T) {
		twilioNotifier.ApiKey = TWILIO_SID
		twilioNotifier.Delay = time.Duration(100 * time.Millisecond)
		err := notifier.AddNotifier(twilioNotifier)
		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", twilioNotifier.Author)
		assert.Equal(t, TWILIO_SID, twilioNotifier.ApiKey)
	})

	t.Run("Load twilioNotifier Notifier", func(t *testing.T) {
		notifier.Load()
	})

	t.Run("twilioNotifier Within Limits", func(t *testing.T) {
		ok, err := twilioNotifier.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("twilioNotifier Send", func(t *testing.T) {
		err := twilioNotifier.Send(twilioMessage)
		assert.Nil(t, err)
	})

	t.Run("twilioNotifier Queue", func(t *testing.T) {
		go notifier.Queue(twilioNotifier)
		time.Sleep(1 * time.Second)
		assert.Equal(t, TWILIO_SID, twilioNotifier.ApiKey)
		assert.Equal(t, 0, len(twilioNotifier.Queue))
	})

}
