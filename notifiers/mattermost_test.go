package notifiers

import (
	"github.com/statping-ng/statping-ng/database"
	"github.com/statping-ng/statping-ng/types/core"
	"github.com/statping-ng/statping-ng/types/failures"
	"github.com/statping-ng/statping-ng/types/notifications"
	"github.com/statping-ng/statping-ng/types/null"
	"github.com/statping-ng/statping-ng/types/services"
	"github.com/statping-ng/statping-ng/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	MATTERMOST_URL string
)

func TestMattermostNotifier(t *testing.T) {
	err := utils.InitLogs()
	require.Nil(t, err)

	t.Parallel()
	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)
	core.Example()

	MATTERMOST_URL = utils.Params.GetString("MATTERMOST_URL")
	if MATTERMOST_URL == "" {
		t.Log("mattermost notifier testing skipped, missing MATTERMOST_URL environment variable")
		t.SkipNow()
	}

	mattermoster.Host = null.NewNullString(MATTERMOST_URL)
	mattermoster.Enabled = null.NewNullBool(true)

	t.Run("Load mattermost", func(t *testing.T) {
		mattermoster.Host = null.NewNullString(MATTERMOST_URL)
		mattermoster.Delay = 100 * time.Millisecond
		mattermoster.Limits = 3
		Add(mattermoster)
		assert.Equal(t, "Adam Boutcher", mattermoster.Author)
		assert.Equal(t, MATTERMOST_URL, mattermoster.Host.String)
	})

	t.Run("mattermost Within Limits", func(t *testing.T) {
		ok := mattermoster.CanSend()
		assert.True(t, ok)
	})

	t.Run("mattermost OnSave", func(t *testing.T) {
		_, err := mattermoster.OnSave()
		assert.Nil(t, err)
	})

	t.Run("mattermost OnFailure", func(t *testing.T) {
		_, err := mattermoster.OnFailure(services.Example(false), failures.Example())
		assert.Nil(t, err)
	})

	t.Run("mattermost OnSuccess", func(t *testing.T) {
		_, err := mattermoster.OnSuccess(services.Example(true))
		assert.Nil(t, err)
	})

}
