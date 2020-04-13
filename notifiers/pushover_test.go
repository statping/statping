package notifiers

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/null"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var (
	PUSHOVER_TOKEN = os.Getenv("PUSHOVER_TOKEN")
	PUSHOVER_API   = os.Getenv("PUSHOVER_API")
)

func TestPushoverNotifier(t *testing.T) {
	t.SkipNow()
	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)

	if PUSHOVER_TOKEN == "" || PUSHOVER_API == "" {
		t.Log("Pushover notifier testing skipped, missing PUSHOVER_TOKEN and PUSHOVER_API environment variable")
		t.SkipNow()
	}

	t.Run("Load Pushover", func(t *testing.T) {
		Pushover.ApiKey = PUSHOVER_TOKEN
		Pushover.ApiSecret = PUSHOVER_API
		Pushover.Enabled = null.NewNullBool(true)

		Add(Pushover)

		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", Pushover.Author)
		assert.Equal(t, PUSHOVER_TOKEN, Pushover.ApiKey)
	})

	t.Run("Pushover Within Limits", func(t *testing.T) {
		assert.True(t, Pushover.CanSend())
	})

	t.Run("Pushover OnFailure", func(t *testing.T) {
		err := Pushover.OnFailure(exampleService, exampleFailure)
		assert.Nil(t, err)
	})

	t.Run("Pushover OnSuccess", func(t *testing.T) {
		err := Pushover.OnSuccess(exampleService)
		assert.Nil(t, err)
	})

	t.Run("Pushover Test", func(t *testing.T) {
		_, err := Pushover.OnTest()
		assert.Nil(t, err)
	})

}
