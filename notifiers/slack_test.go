package notifiers

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	SLACK_URL string
)

func TestSlackNotifier(t *testing.T) {
	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)

	SLACK_URL = utils.Params.GetString("SLACK_URL")
	slacker.Host = SLACK_URL
	slacker.Enabled = null.NewNullBool(true)

	if SLACK_URL == "" {
		t.Log("slack notifier testing skipped, missing SLACK_URL environment variable")
		t.SkipNow()
	}

	t.Run("Load slack", func(t *testing.T) {
		slacker.Host = SLACK_URL
		slacker.Delay = time.Duration(100 * time.Millisecond)
		slacker.Limits = 3
		Add(slacker)
		assert.Equal(t, "Hunter Long", slacker.Author)
		assert.Equal(t, SLACK_URL, slacker.Host)
	})

	t.Run("slack Within Limits", func(t *testing.T) {
		ok := slacker.CanSend()
		assert.True(t, ok)
	})

	t.Run("slack OnFailure", func(t *testing.T) {
		err := slacker.OnFailure(exampleService, exampleFailure)
		assert.Nil(t, err)
	})

	t.Run("slack OnSuccess", func(t *testing.T) {
		err := slacker.OnSuccess(exampleService)
		assert.Nil(t, err)
	})

}
