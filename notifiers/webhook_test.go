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
	"github.com/hunterlong/statping/core/notifier"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	webhookTestUrl = "https://statping.com"
	webhookMessage = `{"id": "%service.Id","name": "%service.Name","online": "%service.Online","issue": "%failure.Issue"}`
	apiKey         = "application/json"
	fullMsg        string
)

func init() {
	Webhook.Host = webhookTestUrl
	Webhook.Var1 = "POST"
	Webhook.Var2 = webhookMessage
	Webhook.ApiKey = "application/json"
}

func TestWebhookNotifier(t *testing.T) {
	t.Parallel()
	currentCount = CountNotifiers()

	t.Run("Load webhooker", func(t *testing.T) {
		Webhook.Host = webhookTestUrl
		Webhook.Delay = time.Duration(100 * time.Millisecond)
		Webhook.ApiKey = apiKey
		err := notifier.AddNotifiers(Webhook)
		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", Webhook.Author)
		assert.Equal(t, webhookTestUrl, Webhook.Host)
		assert.Equal(t, apiKey, Webhook.ApiKey)
	})

	t.Run("webhooker Notifier Tester", func(t *testing.T) {
		assert.True(t, Webhook.CanTest())
	})

	t.Run("webhooker Replace Body Text", func(t *testing.T) {
		fullMsg = replaceBodyText(webhookMessage, TestService, TestFailure)
		assert.Equal(t, `{"id": "1","name": "Interpol - All The Rage Back Home","online": "true","issue": "testing"}`, fullMsg)
	})

	t.Run("webhooker Within Limits", func(t *testing.T) {
		ok, err := Webhook.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("webhooker OnFailure", func(t *testing.T) {
		Webhook.OnFailure(TestService, TestFailure)
		assert.Len(t, Webhook.Queue, 1)
	})

	t.Run("webhooker OnSuccess", func(t *testing.T) {
		Webhook.OnSuccess(TestService)
		assert.Equal(t, len(Webhook.Queue), 1)
	})

	t.Run("webhooker Check Back Online", func(t *testing.T) {
		assert.True(t, TestService.Online)
	})

	t.Run("webhooker OnSuccess Again", func(t *testing.T) {
		Webhook.OnSuccess(TestService)
		assert.Equal(t, len(Webhook.Queue), 1)
	})

	t.Run("webhooker Send", func(t *testing.T) {
		err := Webhook.Send(fullMsg)
		assert.Nil(t, err)
		assert.Equal(t, len(Webhook.Queue), 1)
	})

	t.Run("webhooker Queue", func(t *testing.T) {
		go notifier.Queue(Webhook)
		time.Sleep(8 * time.Second)
		assert.Equal(t, webhookTestUrl, Webhook.Host)
		assert.Equal(t, len(Webhook.Queue), 0)
	})

}
