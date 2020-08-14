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
	DINGTALK_URL string
)

func TestDingtalkNotifier(t *testing.T) {
	t.Parallel()
	err := utils.InitLogs()
	require.Nil(t, err)
	DINGTALK_URL = utils.Params.GetString("DINGTALK_URL")

	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)
	core.Example()

	if DINGTALK_URL == "" {
		t.Log("dingtalk notifier testing skipped, missing DINGTALK_URL environment variable")
		t.SkipNow()
	}

	t.Run("Load dingtalk", func(t *testing.T) {
		Dingtalker.Host = null.NewNullString(DINGTALK_URL)
		Dingtalker.Delay = time.Duration(100 * time.Millisecond)
		Dingtalker.Enabled = null.NewNullBool(true)

		Add(Dingtalker)

		assert.Equal(t, "Huangchun Zhang", Dingtalker.Author)
		assert.Equal(t, DINGTALK_URL, Dingtalker.Host.String)
	})

	t.Run("dingtalk Notifier Tester", func(t *testing.T) {
		assert.True(t, Dingtalker.CanSend())
	})

	t.Run("dingtalk Notifier Tester OnSave", func(t *testing.T) {
		_, err := Dingtalker.OnSave()
		assert.Nil(t, err)
	})

	t.Run("dingtalk OnFailure", func(t *testing.T) {
		_, err := Dingtalker.OnFailure(services.Example(false), failures.Example())
		assert.Nil(t, err)
	})

	t.Run("dingtalk OnSuccess", func(t *testing.T) {
		_, err := Dingtalker.OnSuccess(services.Example(true))
		assert.Nil(t, err)
	})

	t.Run("dingtalk Test", func(t *testing.T) {
		_, err := Dingtalker.OnTest()
		assert.Nil(t, err)
	})

}
