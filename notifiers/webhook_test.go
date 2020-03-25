// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package notifiers

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/null"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	webhookTestUrl = "https://statping.com"
	webhookMessage = `{"id": {{.Service.Id}},"name": "{{.Service.Name}}","online": {{.Service.Online}},"issue": "{{.Failure.Issue}}"}`
	apiKey         = "application/json"
	fullMsg        string
)

func TestWebhookNotifier(t *testing.T) {
	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)

	t.Run("Load webhooker", func(t *testing.T) {
		Webhook.Host = webhookTestUrl
		Webhook.Var1 = "POST"
		Webhook.Var2 = webhookMessage
		Webhook.ApiKey = "application/json"
		Webhook.Enabled = null.NewNullBool(true)

		Add(Webhook)

		assert.Equal(t, "Hunter Long", Webhook.Author)
		assert.Equal(t, webhookTestUrl, Webhook.Host)
		assert.Equal(t, apiKey, Webhook.ApiKey)
	})

	t.Run("webhooker Notifier Tester", func(t *testing.T) {
		assert.True(t, Webhook.CanSend())
	})

	t.Run("webhooker OnFailure", func(t *testing.T) {
		err := Webhook.OnFailure(exampleService, exampleFailure)
		assert.Nil(t, err)
	})

	t.Run("webhooker OnSuccess", func(t *testing.T) {
		err := Webhook.OnSuccess(exampleService)
		assert.Nil(t, err)
	})

	t.Run("webhooker Send", func(t *testing.T) {
		err := Webhook.Send(fullMsg)
		assert.Nil(t, err)
	})

}
