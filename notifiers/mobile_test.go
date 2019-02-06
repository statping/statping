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
	MOBILE_ID     string
	MOBILE_NUMBER string
)

func init() {
	MOBILE_ID = os.Getenv("MOBILE_ID")
	MOBILE_NUMBER = os.Getenv("MOBILE_NUMBER")
	mobile.Var1 = MOBILE_ID
}

func TestMobileNotifier(t *testing.T) {
	t.Parallel()
	mobile.Var1 = MOBILE_ID
	mobile.Var2 = os.Getenv("MOBILE_NUMBER")
	if MOBILE_ID == "" {
		t.Log("mobile notifier testing skipped, missing MOBILE_ID environment variable")
		t.SkipNow()
	}
	currentCount = CountNotifiers()

	t.Run("Load mobile", func(t *testing.T) {
		mobile.Var1 = MOBILE_ID
		mobile.Var2 = MOBILE_NUMBER
		mobile.Delay = time.Duration(100 * time.Millisecond)
		mobile.Limits = 10
		err := notifier.AddNotifier(mobile)
		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", mobile.Author)
		assert.Equal(t, MOBILE_ID, mobile.Var1)
		assert.Equal(t, MOBILE_NUMBER, mobile.Var2)
	})

	t.Run("Load mobile Notifier", func(t *testing.T) {
		notifier.Load()
	})

	t.Run("mobile Notifier Tester", func(t *testing.T) {
		assert.True(t, mobile.CanTest())
	})

	t.Run("mobile Within Limits", func(t *testing.T) {
		ok, err := mobile.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("mobile OnFailure", func(t *testing.T) {
		mobile.OnFailure(TestService, TestFailure)
		assert.Equal(t, 1, len(mobile.Queue))
	})

	t.Run("mobile OnFailure multiple times", func(t *testing.T) {
		for i := 0; i <= 50; i++ {
			mobile.OnFailure(TestService, TestFailure)
		}
		assert.Equal(t, 52, len(mobile.Queue))
	})

	t.Run("mobile Check Offline", func(t *testing.T) {
		assert.False(t, mobile.Online)
	})

	t.Run("mobile OnSuccess", func(t *testing.T) {
		mobile.OnSuccess(TestService)
		assert.Equal(t, 1, len(mobile.Queue))
	})

	t.Run("mobile Queue after being online", func(t *testing.T) {
		assert.True(t, mobile.Online)
		assert.Equal(t, 1, len(mobile.Queue))
	})

	t.Run("mobile OnSuccess Again", func(t *testing.T) {
		t.SkipNow()
		assert.True(t, mobile.Online)
		mobile.OnSuccess(TestService)
		assert.Equal(t, 1, len(mobile.Queue))
		go notifier.Queue(mobile)
		time.Sleep(20 * time.Second)
		assert.Equal(t, 1, len(mobile.Queue))
	})

	t.Run("mobile Within Limits again", func(t *testing.T) {
		ok, err := mobile.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("mobile Test", func(t *testing.T) {
		t.SkipNow()
		err := mobile.OnTest()
		assert.Nil(t, err)
	})

	t.Run("mobile Queue", func(t *testing.T) {
		t.SkipNow()
		go notifier.Queue(mobile)
		time.Sleep(15 * time.Second)
		assert.Equal(t, MOBILE_ID, mobile.Var1)
		assert.Equal(t, 0, len(mobile.Queue))
	})

}
