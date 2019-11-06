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
	"fmt"
	"github.com/hunterlong/statping/core/notifier"
	"github.com/hunterlong/statping/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var (
	EMAIL_HOST     string
	EMAIL_USER     string
	EMAIL_PASS     string
	EMAIL_OUTGOING string
	EMAIL_SEND_TO  string
	EMAIL_PORT     int64
)

var testEmail *emailOutgoing

func init() {
	EMAIL_HOST = os.Getenv("EMAIL_HOST")
	EMAIL_USER = os.Getenv("EMAIL_USER")
	EMAIL_PASS = os.Getenv("EMAIL_PASS")
	EMAIL_OUTGOING = os.Getenv("EMAIL_OUTGOING")
	EMAIL_SEND_TO = os.Getenv("EMAIL_SEND_TO")
	EMAIL_PORT = utils.ToInt(os.Getenv("EMAIL_PORT"))

	Emailer.Host = EMAIL_HOST
	Emailer.Username = EMAIL_USER
	Emailer.Password = EMAIL_PASS
	Emailer.Var1 = EMAIL_OUTGOING
	Emailer.Var2 = EMAIL_SEND_TO
	Emailer.Port = int(EMAIL_PORT)
}

func TestEmailNotifier(t *testing.T) {
	t.Parallel()
	if EMAIL_HOST == "" || EMAIL_USER == "" || EMAIL_PASS == "" {
		t.Log("email notifier testing skipped, missing EMAIL_ environment variables")
		t.SkipNow()
	}
	currentCount = CountNotifiers()

	t.Run("New Emailer", func(t *testing.T) {
		Emailer.Host = EMAIL_HOST
		Emailer.Username = EMAIL_USER
		Emailer.Password = EMAIL_PASS
		Emailer.Var1 = EMAIL_OUTGOING
		Emailer.Var2 = EMAIL_SEND_TO
		Emailer.Port = int(EMAIL_PORT)
		Emailer.Delay = time.Duration(100 * time.Millisecond)

		testEmail = &emailOutgoing{
			To:       Emailer.GetValue("var2"),
			Subject:  fmt.Sprintf("Service %v is Failing", TestService.Name),
			Template: mainEmailTemplate,
			Data:     TestService,
			From:     Emailer.GetValue("var1"),
		}
	})

	t.Run("Add email Notifier", func(t *testing.T) {
		err := notifier.AddNotifiers(Emailer)
		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", Emailer.Author)
		assert.Equal(t, EMAIL_HOST, Emailer.Host)
	})

	t.Run("email Within Limits", func(t *testing.T) {
		ok, err := Emailer.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("email Test Source", func(t *testing.T) {
		emailSource(testEmail)
		assert.NotEmpty(t, testEmail.Source)
	})

	t.Run("email OnFailure", func(t *testing.T) {
		Emailer.OnFailure(TestService, TestFailure)
		assert.Equal(t, 1, len(Emailer.Queue))
	})

	t.Run("email OnSuccess", func(t *testing.T) {
		Emailer.OnSuccess(TestService)
		assert.Equal(t, 1, len(Emailer.Queue))
	})

	t.Run("email Check Back Online", func(t *testing.T) {
		assert.True(t, TestService.Online)
	})

	t.Run("email OnSuccess Again", func(t *testing.T) {
		Emailer.OnSuccess(TestService)
		assert.Equal(t, 1, len(Emailer.Queue))
	})

	t.Run("email Send", func(t *testing.T) {
		err := Emailer.Send(testEmail)
		assert.Nil(t, err)
	})

	t.Run("Emailer Test", func(t *testing.T) {
		t.SkipNow()
		err := Emailer.OnTest()
		assert.Nil(t, err)
	})

	t.Run("email Run Queue", func(t *testing.T) {
		go notifier.Queue(Emailer)
		time.Sleep(6 * time.Second)
		assert.Equal(t, EMAIL_HOST, Emailer.Host)
		assert.Equal(t, 0, len(Emailer.Queue))
	})

}
