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
	"os"
	"testing"
	"time"
)

var (
	DISCORD_URL    = os.Getenv("DISCORD_URL")
	discordMessage = `{"content": "The discord notifier on Statping has been tested!"}`
)

func init() {
	DISCORD_URL = os.Getenv("DISCORD_URL")
	Discorder.Host = DISCORD_URL
}

func TestDiscordNotifier(t *testing.T) {
	t.Parallel()
	if DISCORD_URL == "" {
		t.Log("discord notifier testing skipped, missing DISCORD_URL environment variable")
		t.SkipNow()
	}
	currentCount = CountNotifiers()

	t.Run("Load discord", func(t *testing.T) {
		Discorder.Host = DISCORD_URL
		Discorder.Delay = time.Duration(100 * time.Millisecond)
		err := notifier.AddNotifiers(Discorder)
		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", Discorder.Author)
		assert.Equal(t, DISCORD_URL, Discorder.Host)
	})

	t.Run("discord Notifier Tester", func(t *testing.T) {
		assert.True(t, Discorder.CanTest())
	})

	t.Run("discord Within Limits", func(t *testing.T) {
		ok, err := Discorder.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("discord OnFailure", func(t *testing.T) {
		Discorder.OnFailure(TestService, TestFailure)
		assert.Equal(t, 1, len(Discorder.Queue))
	})

	t.Run("discord OnSuccess", func(t *testing.T) {
		Discorder.OnSuccess(TestService)
		assert.Equal(t, 1, len(Discorder.Queue))
	})

	t.Run("discord Check Back Online", func(t *testing.T) {
		assert.True(t, TestService.Online)
	})

	t.Run("discord OnSuccess Again", func(t *testing.T) {
		Discorder.OnSuccess(TestService)
		assert.Equal(t, 1, len(Discorder.Queue))
	})

	t.Run("discord Send", func(t *testing.T) {
		err := Discorder.Send(discordMessage)
		assert.Nil(t, err)
	})

	t.Run("discord Test", func(t *testing.T) {
		err := Discorder.OnTest()
		assert.Nil(t, err)
	})

	t.Run("discord Queue", func(t *testing.T) {
		go notifier.Queue(Discorder)
		time.Sleep(1 * time.Second)
		assert.Equal(t, DISCORD_URL, Discorder.Host)
		assert.Equal(t, 0, len(Discorder.Queue))
	})

}
