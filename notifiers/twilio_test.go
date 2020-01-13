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

	Twilio.ApiKey = TWILIO_SID
	Twilio.ApiSecret = TWILIO_SECRET
	Twilio.Var1 = TWILIO_TO
	Twilio.Var2 = TWILIO_FROM
}

func TestTwilioNotifier(t *testing.T) {
	t.SkipNow()
	if TWILIO_SID == "" || TWILIO_SECRET == "" || TWILIO_FROM == "" {
		t.Log("twilio notifier testing skipped, missing TWILIO_SID environment variable")
		t.SkipNow()
	}
	currentCount = CountNotifiers()

	t.Run("Load Twilio", func(t *testing.T) {
		Twilio.ApiKey = TWILIO_SID
		Twilio.Delay = time.Duration(100 * time.Millisecond)
		err := notifier.AddNotifiers(Twilio)
		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", Twilio.Author)
		assert.Equal(t, TWILIO_SID, Twilio.ApiKey)
	})

	t.Run("Twilio Within Limits", func(t *testing.T) {
		ok, err := Twilio.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("Twilio OnFailure", func(t *testing.T) {
		Twilio.OnFailure(TestService, TestFailure)
		assert.Len(t, Twilio.Queue, 1)
	})

	t.Run("Twilio Check Offline", func(t *testing.T) {
		assert.False(t, TestService.Online)
	})

	t.Run("Twilio OnSuccess", func(t *testing.T) {
		Twilio.OnSuccess(TestService)
		assert.Len(t, Twilio.Queue, 2)
	})

	t.Run("Twilio Check Back Online", func(t *testing.T) {
		assert.True(t, TestService.Online)
	})

	t.Run("Twilio OnSuccess Again", func(t *testing.T) {
		Twilio.OnSuccess(TestService)
		assert.Len(t, Twilio.Queue, 2)
	})

	t.Run("Twilio Send", func(t *testing.T) {
		err := Twilio.Send(twilioMessage)
		assert.Nil(t, err)
	})

	t.Run("Twilio Test", func(t *testing.T) {
		err := Twilio.OnTest()
		assert.Nil(t, err)
	})

	t.Run("Twilio Queue", func(t *testing.T) {
		go notifier.Queue(Twilio)
		time.Sleep(1 * time.Second)
		assert.Equal(t, TWILIO_SID, Twilio.ApiKey)
		assert.Equal(t, 0, len(Twilio.Queue))
	})

}
