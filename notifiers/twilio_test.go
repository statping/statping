package notifiers

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	TWILIO_SID    string
	TWILIO_SECRET string
)

func TestTwilioNotifier(t *testing.T) {
	err := utils.InitLogs()
	require.Nil(t, err)

	TWILIO_SID = utils.Params.GetString("TWILIO_SID")
	TWILIO_SECRET = utils.Params.GetString("TWILIO_SECRET")

	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)
	core.Example()

	if TWILIO_SID == "" || TWILIO_SECRET == "" {
		t.Log("twilio notifier testing skipped, missing TWILIO_SID and TWILIO_SECRET environment variable")
		t.SkipNow()
	}

	t.Run("Load Twilio", func(t *testing.T) {
		Twilio.ApiKey = TWILIO_SID
		Twilio.ApiSecret = TWILIO_SECRET
		Twilio.Var1 = "15005550006"
		Twilio.Var2 = "15005550006"
		Twilio.Delay = 100 * time.Millisecond
		Twilio.Enabled = null.NewNullBool(true)

		Add(Twilio)

		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", Twilio.Author)
		assert.Equal(t, TWILIO_SID, Twilio.ApiKey)
	})

	t.Run("Twilio Within Limits", func(t *testing.T) {
		assert.True(t, Twilio.CanSend())
	})

	t.Run("Twilio OnSave", func(t *testing.T) {
		_, err := Twilio.OnSave()
		assert.Nil(t, err)
	})

	t.Run("Twilio OnFailure", func(t *testing.T) {
		_, err := Twilio.OnFailure(services.Example(false), failures.Example())
		assert.Nil(t, err)
	})

	t.Run("Twilio OnSuccess", func(t *testing.T) {
		_, err := Twilio.OnSuccess(services.Example(true))
		assert.Nil(t, err)
	})

	t.Run("Twilio Test", func(t *testing.T) {
		_, err := Twilio.OnTest()
		assert.Nil(t, err)
	})

}
