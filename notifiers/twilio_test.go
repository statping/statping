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
	TWILIO_SID    = os.Getenv("TWILIO_SID")
	TWILIO_SECRET = os.Getenv("TWILIO_SECRET")
	TWILIO_FROM   = os.Getenv("TWILIO_FROM")
	TWILIO_TO     = os.Getenv("TWILIO_TO")
	twilioMessage = "The Twilio notifier on Statping has been tested!"
)

func init() {
	TWILIO_SID = os.Getenv("TWILIO_SID")
	TWILIO_SECRET = os.Getenv("TWILIO_SECRET")
	TWILIO_FROM = os.Getenv("TWILIO_FROM")
	TWILIO_TO = os.Getenv("TWILIO_TO")

	twilioNotifier.ApiKey = TWILIO_SID
	twilioNotifier.ApiSecret = TWILIO_SECRET
	twilioNotifier.Var1 = TWILIO_TO
	twilioNotifier.Var2 = TWILIO_FROM
}

func TestTwilioNotifier(t *testing.T) {
	t.SkipNow()
	t.Parallel()
	if TWILIO_SID == "" || TWILIO_SECRET == "" || TWILIO_FROM == "" {
		t.Log("twilio notifier testing skipped, missing TWILIO_SID environment variable")
		t.SkipNow()
	}
	currentCount = CountNotifiers()

	t.Run("Load Twilio", func(t *testing.T) {
		twilioNotifier.ApiKey = TWILIO_SID
		twilioNotifier.Delay = time.Duration(100 * time.Millisecond)
		err := notifier.AddNotifier(twilioNotifier)
		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", twilioNotifier.Author)
		assert.Equal(t, TWILIO_SID, twilioNotifier.ApiKey)
	})

	t.Run("Load Twilio Notifier", func(t *testing.T) {
		notifier.Load()
	})

	t.Run("Twilio Within Limits", func(t *testing.T) {
		ok, err := twilioNotifier.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("Twilio OnFailure", func(t *testing.T) {
		twilioNotifier.OnFailure(TestService, TestFailure)
		assert.Len(t, twilioNotifier.Queue, 1)
	})

	t.Run("Twilio Check Offline", func(t *testing.T) {
		assert.False(t, twilioNotifier.Online)
	})

	t.Run("Twilio OnSuccess", func(t *testing.T) {
		twilioNotifier.OnSuccess(TestService)
		assert.Len(t, twilioNotifier.Queue, 2)
	})

	t.Run("Twilio Check Back Online", func(t *testing.T) {
		assert.True(t, twilioNotifier.Online)
	})

	t.Run("Twilio OnSuccess Again", func(t *testing.T) {
		twilioNotifier.OnSuccess(TestService)
		assert.Len(t, twilioNotifier.Queue, 2)
	})

	t.Run("Twilio Send", func(t *testing.T) {
		err := twilioNotifier.Send(twilioMessage)
		assert.Nil(t, err)
	})

	t.Run("Twilio Test", func(t *testing.T) {
		err := twilioNotifier.OnTest()
		assert.Nil(t, err)
	})

	t.Run("Twilio Queue", func(t *testing.T) {
		go notifier.Queue(twilioNotifier)
		time.Sleep(1 * time.Second)
		assert.Equal(t, TWILIO_SID, twilioNotifier.ApiKey)
		assert.Equal(t, 0, len(twilioNotifier.Queue))
	})

}
