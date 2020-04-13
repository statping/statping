package notifiers

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/null"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

var (
	TWILIO_SID    = os.Getenv("TWILIO_SID")
	TWILIO_SECRET = os.Getenv("TWILIO_SECRET")
	TWILIO_FROM   = os.Getenv("TWILIO_FROM")
	TWILIO_TO     = os.Getenv("TWILIO_TO")
)

func init() {
	TWILIO_SID = os.Getenv("TWILIO_SID")
	TWILIO_SECRET = os.Getenv("TWILIO_SECRET")
	TWILIO_FROM = os.Getenv("TWILIO_FROM")
	TWILIO_TO = os.Getenv("TWILIO_TO")
}

func TestTwilioNotifier(t *testing.T) {
	t.SkipNow()

	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)

	if TWILIO_SID == "" || TWILIO_SECRET == "" || TWILIO_FROM == "" {
		t.Log("twilio notifier testing skipped, missing TWILIO_SID environment variable")
		t.SkipNow()
	}

	t.Run("Load Twilio", func(t *testing.T) {
		Twilio.ApiKey = TWILIO_SID
		Twilio.ApiSecret = TWILIO_SECRET
		Twilio.Var1 = TWILIO_TO
		Twilio.Var2 = TWILIO_FROM
		Twilio.Delay = time.Duration(100 * time.Millisecond)
		Twilio.Enabled = null.NewNullBool(true)

		Add(Twilio)

		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", Twilio.Author)
		assert.Equal(t, TWILIO_SID, Twilio.ApiKey)
	})

	t.Run("Twilio Within Limits", func(t *testing.T) {
		assert.True(t, Twilio.CanSend())
	})

	t.Run("Twilio OnFailure", func(t *testing.T) {
		err := Twilio.OnFailure(exampleService, exampleFailure)
		assert.Nil(t, err)
	})

	t.Run("Twilio OnSuccess", func(t *testing.T) {
		err := Twilio.OnSuccess(exampleService)
		assert.Nil(t, err)
	})

	t.Run("Twilio Test", func(t *testing.T) {
		_, err := Twilio.OnTest()
		assert.Nil(t, err)
	})

}
