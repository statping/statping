package notifiers

import (
	"testing"

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
	webhookTestUrl = "https://statping.com"
	webhookMessage = `{"id": {{.Service.Id}},"name": "{{.Service.Name}}","online": {{.Service.Online}},"issue": "{{.Failure.Issue}}"}`
	apiKey         = "application/json"
	fullMsg        string
)

func TestWebhookNotifier(t *testing.T) {
	err := utils.InitLogs()
	require.Nil(t, err)

	t.Parallel()
	t.SkipNow()

	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)
	core.Example()

	t.Run("Load webhooker", func(t *testing.T) {
		Webhook.Host = null.NewNullString(webhookTestUrl)
		Webhook.Var1 = null.NewNullString("POST")
		Webhook.Var2 = null.NewNullString(webhookMessage)
		Webhook.ApiKey = null.NewNullString("application/json")
		Webhook.Enabled = null.NewNullBool(true)

		Add(Webhook)

		assert.Equal(t, "Hunter Long", Webhook.Author)
		assert.Equal(t, webhookTestUrl, Webhook.Host)
		assert.Equal(t, apiKey, Webhook.ApiKey)
	})

	t.Run("webhooker Notifier Tester", func(t *testing.T) {
		assert.True(t, Webhook.CanSend())
	})

	t.Run("webhooker OnSave", func(t *testing.T) {
		_, err := Webhook.OnSave()
		assert.Nil(t, err)
	})

	t.Run("webhooker OnFailure", func(t *testing.T) {
		_, err := Webhook.OnFailure(services.Example(false), failures.Example())
		assert.Nil(t, err)
	})

	t.Run("webhooker OnSuccess", func(t *testing.T) {
		_, err := Webhook.OnSuccess(services.Example(true))
		assert.Nil(t, err)
	})

	t.Run("webhooker Send", func(t *testing.T) {
		err := Webhook.Send(fullMsg)
		assert.Nil(t, err)
	})

}
