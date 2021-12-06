package notifiers

import (
	"testing"
	"time"

	"github.com/statping/statping/database"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	GOTIFY_URL   string
	GOTIFY_TOKEN string
)

func TestGotifyNotifier(t *testing.T) {
	err := utils.InitLogs()
	require.Nil(t, err)

	t.Parallel()
	GOTIFY_URL = utils.Params.GetString("GOTIFY_URL")
	GOTIFY_TOKEN = utils.Params.GetString("GOTIFY_TOKEN")

	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)
	core.Example()

	if GOTIFY_URL == "" {
		t.Log("gotify notifier testing skipped, missing GOTIFY_URL environment variable")
		t.SkipNow()
	}
	if GOTIFY_TOKEN == "" {
		t.Log("gotify notifier testing skipped, missing GOTIFY_TOKEN environment variable")
		t.SkipNow()
	}

	t.Run("Load gotify", func(t *testing.T) {
		Gotify.Host = null.NewNullString(GOTIFY_URL)
		Gotify.Delay = time.Duration(100 * time.Millisecond)
		Gotify.Enabled = null.NewNullBool(true)

		Add(Gotify)

		assert.Equal(t, "Hugo van Rijswijk", Gotify.Author)
		assert.Equal(t, GOTIFY_URL, Gotify.Host.String)
	})

	t.Run("gotify Notifier Tester", func(t *testing.T) {
		assert.True(t, Gotify.CanSend())
	})

	t.Run("gotify Notifier Tester OnSave", func(t *testing.T) {
		_, err := Gotify.OnSave()
		assert.Nil(t, err)
	})

	t.Run("gotify OnFailure", func(t *testing.T) {
		_, err := Gotify.OnFailure(services.Example(false), failures.Example())
		assert.Nil(t, err)
	})

	t.Run("gotify OnSuccess", func(t *testing.T) {
		_, err := Gotify.OnSuccess(services.Example(true))
		assert.Nil(t, err)
	})

	t.Run("gotify Test", func(t *testing.T) {
		_, err := Gotify.OnTest()
		assert.Nil(t, err)
	})

}
