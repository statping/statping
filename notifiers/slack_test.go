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
	SLACK_URL        string
	slackTestMessage = `{"text":"this is a test from the slack notifier!"}`
)

func init() {
	SLACK_URL = os.Getenv("SLACK_URL")
	slacker.Host = SLACK_URL
}

func TestSlackNotifier(t *testing.T) {
	t.Parallel()
	SLACK_URL = os.Getenv("SLACK_URL")
	slacker.Host = SLACK_URL
	if SLACK_URL == "" {
		t.Log("slack notifier testing skipped, missing SLACK_URL environment variable")
		t.SkipNow()
	}
	currentCount = CountNotifiers()

	t.Run("Load slack", func(t *testing.T) {
		slacker.Host = SLACK_URL
		slacker.Delay = time.Duration(100 * time.Millisecond)
		slacker.Limits = 3
		err := notifier.AddNotifier(slacker)
		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", slacker.Author)
		assert.Equal(t, SLACK_URL, slacker.Host)
	})

	t.Run("Load slack Notifier", func(t *testing.T) {
		notifier.Load()
	})

	t.Run("slack Notifier Tester", func(t *testing.T) {
		assert.True(t, slacker.CanTest())
	})

	//t.Run("slack parse message", func(t *testing.T) {
	//	err := parseSlackMessage(slackText, "this is a test!")
	//	assert.Nil(t, err)
	//	assert.Equal(t, 1, len(slacker.Queue))
	//})

	t.Run("slack Within Limits", func(t *testing.T) {
		ok, err := slacker.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("slack OnFailure", func(t *testing.T) {
		slacker.OnFailure(TestService, TestFailure)
		assert.Equal(t, 1, len(slacker.Queue))
	})

	t.Run("slack OnFailure multiple times", func(t *testing.T) {
		for i := 0; i <= 50; i++ {
			slacker.OnFailure(TestService, TestFailure)
		}
		assert.Equal(t, 52, len(slacker.Queue))
	})

	t.Run("slack Check Offline", func(t *testing.T) {
		assert.False(t, TestService.Online)
	})

	t.Run("slack OnSuccess", func(t *testing.T) {
		slacker.OnSuccess(TestService)
		assert.Equal(t, 1, len(slacker.Queue))
	})

	t.Run("slack Queue after being online", func(t *testing.T) {
		assert.True(t, TestService.Online)
		assert.Equal(t, 1, len(slacker.Queue))
	})

	t.Run("slack OnSuccess Again", func(t *testing.T) {
		assert.True(t, TestService.Online)
		slacker.OnSuccess(TestService)
		assert.Equal(t, 1, len(slacker.Queue))
		go notifier.Queue(slacker)
		time.Sleep(6 * time.Second)
		assert.Equal(t, 0, len(slacker.Queue))
	})

	t.Run("slack Within Limits again", func(t *testing.T) {
		ok, err := slacker.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("slack Send", func(t *testing.T) {
		err := slacker.Send(slackTestMessage)
		assert.Nil(t, err)
		assert.Equal(t, 0, len(slacker.Queue))
	})

	t.Run("slack Test", func(t *testing.T) {
		err := slacker.OnTest()
		assert.Nil(t, err)
	})

	t.Run("slack Queue", func(t *testing.T) {
		go notifier.Queue(slacker)
		time.Sleep(5 * time.Second)
		assert.Equal(t, SLACK_URL, slacker.Host)
		assert.Equal(t, 0, len(slacker.Queue))
	})

}
