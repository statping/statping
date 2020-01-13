// Statup
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
	t.SkipNow()
	t.Parallel()
	if telegramToken == "" || telegramChannel == "" {
		t.Log("Telegram notifier testing skipped, missing TELEGRAM_TOKEN and TELEGRAM_CHANNEL environment variable")
		t.SkipNow()
	}
	currentCount = CountNotifiers()

	t.Run("Load Telegram", func(t *testing.T) {
		Telegram.ApiSecret = telegramToken
		Telegram.Var1 = telegramChannel
		Telegram.Delay = time.Duration(1 * time.Second)
		err := notifier.AddNotifiers(Telegram)
		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", Telegram.Author)
		assert.Equal(t, telegramToken, Telegram.ApiSecret)
		assert.Equal(t, telegramChannel, Telegram.Var1)
	})

	t.Run("Telegram Within Limits", func(t *testing.T) {
		ok, err := Telegram.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("Telegram OnFailure", func(t *testing.T) {
		Telegram.OnFailure(TestService, TestFailure)
		assert.Equal(t, 1, len(Telegram.Queue))
	})

	t.Run("Telegram Check Offline", func(t *testing.T) {
		assert.False(t, TestService.Online)
	})

	t.Run("Telegram OnSuccess", func(t *testing.T) {
		Telegram.OnSuccess(TestService)
		assert.Equal(t, 1, len(Telegram.Queue))
	})

	t.Run("Telegram Check Back Online", func(t *testing.T) {
		assert.True(t, TestService.Online)
	})

	t.Run("Telegram OnSuccess Again", func(t *testing.T) {
		Telegram.OnSuccess(TestService)
		assert.Equal(t, 1, len(Telegram.Queue))
	})

	t.Run("Telegram Send", func(t *testing.T) {
		err := Telegram.Send(telegramMessage)
		assert.Nil(t, err)
	})

	t.Run("Telegram Test", func(t *testing.T) {
		err := Telegram.OnTest()
		assert.Nil(t, err)
	})

	t.Run("Telegram Queue", func(t *testing.T) {
		go notifier.Queue(Telegram)
		time.Sleep(3 * time.Second)
		assert.Equal(t, telegramToken, Telegram.ApiSecret)
		assert.Equal(t, 0, len(Telegram.Queue))
	})

}
