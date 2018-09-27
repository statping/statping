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
	"os"
	"testing"
	"time"
)

var (
	DISCORD_URL    = os.Getenv("DISCORD_URL")
	discordMessage = `{"content": "The Discord notifier on Statup has been tested!"}`
)

func init() {
	DISCORD_URL = os.Getenv("DISCORD_URL")
	discorder.Host = DISCORD_URL
}

func TestDiscordNotifier(t *testing.T) {
	t.Parallel()
	if DISCORD_URL == "" {
		t.Log("Discord notifier testing skipped, missing DISCORD_URL environment variable")
		t.SkipNow()
	}
	currentCount = CountNotifiers()

	t.Run("Load Discord", func(t *testing.T) {
		discorder.Host = DISCORD_URL
		discorder.Delay = time.Duration(100 * time.Millisecond)
		err := notifier.AddNotifier(discorder)
		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", discorder.Author)
		assert.Equal(t, DISCORD_URL, discorder.Host)
	})

	t.Run("Load Discord Notifier", func(t *testing.T) {
		notifier.Load()
	})

	t.Run("Discord Notifier Tester", func(t *testing.T) {
		assert.True(t, discorder.CanTest())
	})

	t.Run("Discord Within Limits", func(t *testing.T) {
		ok, err := discorder.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("Discord OnFailure", func(t *testing.T) {
		discorder.OnFailure(TestService, TestFailure)
		assert.Len(t, discorder.Queue, 1)
	})

	t.Run("Discord Check Offline", func(t *testing.T) {
		assert.False(t, discorder.Online)
	})

	t.Run("Discord OnSuccess", func(t *testing.T) {
		discorder.OnSuccess(TestService)
		assert.Len(t, discorder.Queue, 2)
	})

	t.Run("Discord Check Back Online", func(t *testing.T) {
		assert.True(t, discorder.Online)
	})

	t.Run("Discord OnSuccess Again", func(t *testing.T) {
		discorder.OnSuccess(TestService)
		assert.Len(t, discorder.Queue, 2)
	})

	t.Run("Discord Send", func(t *testing.T) {
		err := discorder.Send(discordMessage)
		assert.Nil(t, err)
	})

	t.Run("Discord Queue", func(t *testing.T) {
		go notifier.Queue(discorder)
		time.Sleep(1 * time.Second)
		assert.Equal(t, DISCORD_URL, discorder.Host)
		assert.Equal(t, 0, len(discorder.Queue))
	})

}
