// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
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
	"github.com/hunterlong/statup/core/notifier"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	WEBHOOK_URL    = "https://jsonplaceholder.typicode.com/posts"
	webhookMessage = `{ "title": "%service.Id", "body": "%service.Name", "online": %service.Online, "userId": 19999 }`
	fullMsg        string
)

func init() {
	webhook.Host = WEBHOOK_URL
	webhook.Var1 = "POST"
}

func TestWebhookNotifier(t *testing.T) {
	t.SkipNow()
	t.Parallel()
	currentCount = CountNotifiers()

	t.Run("Load Webhook", func(t *testing.T) {
		webhook.Host = WEBHOOK_URL
		webhook.Delay = time.Duration(100 * time.Millisecond)
		err := notifier.AddNotifier(webhook)
		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", webhook.Author)
		assert.Equal(t, WEBHOOK_URL, webhook.Host)
	})

	t.Run("Load Webhook Notifier", func(t *testing.T) {
		notifier.Load()
	})

	t.Run("Webhook Notifier Tester", func(t *testing.T) {
		assert.True(t, webhook.CanTest())
	})

	t.Run("Webhook Replace Body Text", func(t *testing.T) {
		fullMsg = replaceBodyText(webhookMessage, TestService, TestFailure)
		assert.Equal(t, "{ \"title\": \"1\", \"body\": \"Interpol - All The Rage Back Home\", \"online\": false, \"userId\": 19999 }", fullMsg)
	})

	t.Run("Webhook Within Limits", func(t *testing.T) {
		ok, err := webhook.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("Webhook OnFailure", func(t *testing.T) {
		webhook.OnFailure(TestService, TestFailure)
		assert.Len(t, webhook.Queue, 1)
	})

	t.Run("Webhook Check Offline", func(t *testing.T) {
		assert.False(t, webhook.Online)
	})

	t.Run("Webhook OnSuccess", func(t *testing.T) {
		webhook.OnSuccess(TestService)
		assert.Len(t, webhook.Queue, 2)
	})

	t.Run("Webhook Check Back Online", func(t *testing.T) {
		assert.True(t, webhook.Online)
	})

	t.Run("Webhook OnSuccess Again", func(t *testing.T) {
		webhook.OnSuccess(TestService)
		assert.Len(t, webhook.Queue, 2)
	})

	t.Run("Webhook Send", func(t *testing.T) {
		err := webhook.Send(fullMsg)
		assert.Nil(t, err)
		assert.Len(t, webhook.Queue, 2)
	})

	t.Run("Webhook Queue", func(t *testing.T) {
		go notifier.Queue(webhook)
		time.Sleep(8 * time.Second)
		assert.Equal(t, WEBHOOK_URL, webhook.Host)
		assert.Equal(t, 1, len(webhook.Queue))
	})

}
