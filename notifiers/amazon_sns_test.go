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

func TestAmazonSNSNotifier(t *testing.T) {
	err := utils.InitLogs()
	require.Nil(t, err)
	snsToken := utils.Params.GetString("SNS_TOKEN")
	snsSecret := utils.Params.GetString("SNS_SECRET")
	snsRegion := utils.Params.GetString("SNS_REGION")
	snsTopic := utils.Params.GetString("SNS_TOPIC")

	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)
	core.Example()

	if snsToken == "" || snsSecret == "" || snsRegion == "" || snsTopic == "" {
		t.Log("SNS notifier testing skipped, missing SNS_TOKEN, SNS_SECRET, SNS_REGION, SNS_TOPIC environment variables")
		t.SkipNow()
	}

	t.Run("Load SNS", func(t *testing.T) {
		AmazonSNS.ApiKey = null.NewNullString(snsToken)
		AmazonSNS.ApiSecret = null.NewNullString(snsSecret)
		AmazonSNS.Var1 = null.NewNullString(snsRegion)
		AmazonSNS.Host = null.NewNullString(snsTopic)
		AmazonSNS.Delay = 15 * time.Second
		AmazonSNS.Enabled = null.NewNullBool(true)

		Add(AmazonSNS)

		assert.Equal(t, "Hunter Long", AmazonSNS.Author)
		assert.Equal(t, snsToken, AmazonSNS.ApiKey.String)
		assert.Equal(t, snsSecret, AmazonSNS.ApiSecret.String)
	})

	t.Run("SNS Notifier Tester", func(t *testing.T) {
		assert.True(t, AmazonSNS.CanSend())
	})

	t.Run("SNS Notifier Tester OnSave", func(t *testing.T) {
		_, err := AmazonSNS.OnSave()
		assert.Nil(t, err)
	})

	t.Run("SNS OnFailure", func(t *testing.T) {
		_, err := AmazonSNS.OnFailure(services.Example(false), failures.Example())
		assert.Nil(t, err)
	})

	t.Run("SNS OnSuccess", func(t *testing.T) {
		_, err := AmazonSNS.OnSuccess(services.Example(true))
		assert.Nil(t, err)
	})

	t.Run("SNS Test", func(t *testing.T) {
		_, err := AmazonSNS.OnTest()
		assert.Nil(t, err)
	})

}
