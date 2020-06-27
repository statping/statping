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
	mobileToken string
)

func TestMobileNotifier(t *testing.T) {
	t.SkipNow()
	err := utils.InitLogs()
	require.Nil(t, err)

	mobileToken = utils.Params.GetString("MOBILE_TOKEN")
	Mobile.Var1 = mobileToken

	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)
	core.Example()

	Mobile.Var1 = mobileToken
	if mobileToken == "" {
		t.Log("Mobile notifier testing skipped, missing MOBILE_ID environment variable")
		t.SkipNow()
	}

	t.Run("Load Mobile", func(t *testing.T) {
		Mobile.Var1 = mobileToken
		Mobile.Delay = time.Duration(100 * time.Millisecond)
		Mobile.Limits = 10
		Mobile.Enabled = null.NewNullBool(true)

		Add(Mobile)

		assert.Equal(t, "Hunter Long", Mobile.Author)
		assert.Equal(t, mobileToken, Mobile.Var1)
	})

	t.Run("Mobile Notifier Tester", func(t *testing.T) {
		assert.True(t, Mobile.CanSend())
	})

	t.Run("Mobile OnSave", func(t *testing.T) {
		_, err := Mobile.OnSave()
		assert.Nil(t, err)
	})

	t.Run("Mobile OnFailure", func(t *testing.T) {
		_, err := Mobile.OnFailure(services.Example(false), failures.Example())
		assert.Nil(t, err)
	})

	t.Run("Mobile OnSuccess", func(t *testing.T) {
		_, err := Mobile.OnSuccess(services.Example(true))
		assert.Nil(t, err)
	})

	t.Run("Mobile Test", func(t *testing.T) {
		_, err := Mobile.OnTest()
		assert.Nil(t, err)
	})

}
