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
	Mobile.Var1 = MOBILE_ID
}

func TestMobileNotifier(t *testing.T) {
	t.Parallel()
	Mobile.Var1 = MOBILE_ID
	Mobile.Var2 = os.Getenv("MOBILE_NUMBER")
	if MOBILE_ID == "" {
		t.Log("Mobile notifier testing skipped, missing MOBILE_ID environment variable")
		t.SkipNow()
	}
	currentCount = CountNotifiers()

	t.Run("Load Mobile", func(t *testing.T) {
		Mobile.Var1 = MOBILE_ID
		Mobile.Var2 = MOBILE_NUMBER
		Mobile.Delay = time.Duration(100 * time.Millisecond)
		Mobile.Limits = 10
		err := notifier.AddNotifiers(Mobile)
		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", Mobile.Author)
		assert.Equal(t, MOBILE_ID, Mobile.Var1)
		assert.Equal(t, MOBILE_NUMBER, Mobile.Var2)
	})

	t.Run("Mobile Notifier Tester", func(t *testing.T) {
		assert.True(t, Mobile.CanTest())
	})

	t.Run("Mobile Within Limits", func(t *testing.T) {
		ok, err := Mobile.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("Mobile OnFailure", func(t *testing.T) {
		Mobile.OnFailure(TestService, TestFailure)
		assert.Equal(t, 1, len(Mobile.Queue))
	})

	t.Run("Mobile OnSuccess", func(t *testing.T) {
		Mobile.OnSuccess(TestService)
		assert.Equal(t, 1, len(Mobile.Queue))
	})

	t.Run("Mobile OnSuccess Again", func(t *testing.T) {
		t.SkipNow()
		assert.True(t, TestService.Online)
		Mobile.OnSuccess(TestService)
		assert.Equal(t, 1, len(Mobile.Queue))
		go notifier.Queue(Mobile)
		time.Sleep(20 * time.Second)
		assert.Equal(t, 1, len(Mobile.Queue))
	})

	t.Run("Mobile Within Limits again", func(t *testing.T) {
		ok, err := Mobile.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("Mobile Test", func(t *testing.T) {
		t.SkipNow()
		err := Mobile.OnTest()
		assert.Nil(t, err)
	})

	t.Run("Mobile Queue", func(t *testing.T) {
		t.SkipNow()
		go notifier.Queue(Mobile)
		time.Sleep(15 * time.Second)
		assert.Equal(t, MOBILE_ID, Mobile.Var1)
		assert.Equal(t, 0, len(Mobile.Queue))
	})

}
