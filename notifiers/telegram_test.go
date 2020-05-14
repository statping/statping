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
	telegramToken   string
	telegramChannel string
	telegramMessage = "The Telegram notifier on Statping has been tested!"
)

func init() {
	telegramToken = os.Getenv("TELEGRAM_TOKEN")
	telegramChannel = os.Getenv("TELEGRAM_CHANNEL")
	Telegram.ApiSecret = telegramToken
	Telegram.Var1 = telegramChannel
}

func TestTelegramNotifier(t *testing.T) {
	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)

	if telegramToken == "" || telegramChannel == "" {
		t.Log("Telegram notifier testing skipped, missing TELEGRAM_TOKEN and TELEGRAM_CHANNEL environment variable")
		t.SkipNow()
	}

	t.Run("Load Telegram", func(t *testing.T) {
		Telegram.ApiSecret = telegramToken
		Telegram.Var1 = telegramChannel
		Telegram.Delay = time.Duration(1 * time.Second)
		Telegram.Enabled = null.NewNullBool(true)

		Add(Telegram)

		assert.Equal(t, "Hunter Long", Telegram.Author)
		assert.Equal(t, telegramToken, Telegram.ApiSecret)
		assert.Equal(t, telegramChannel, Telegram.Var1)
	})

	t.Run("Telegram Within Limits", func(t *testing.T) {
		assert.True(t, Telegram.CanSend())
	})

	t.Run("Telegram OnFailure", func(t *testing.T) {
		err := Telegram.OnFailure(exampleService, exampleFailure)
		assert.Nil(t, err)
	})

	t.Run("Telegram OnSuccess", func(t *testing.T) {
		err := Telegram.OnSuccess(exampleService)
		assert.Nil(t, err)
	})

	t.Run("Telegram Test", func(t *testing.T) {
		_, err := Telegram.OnTest()
		assert.Nil(t, err)
	})

}
