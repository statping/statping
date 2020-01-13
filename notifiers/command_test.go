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
	"testing"
	"time"
)

const (
	commandTest = "curl -I https://statping.com/"
)

func TestCommandNotifier(t *testing.T) {
	t.Parallel()
	Command.Host = "sh"
	Command.Var1 = commandTest
	Command.Var2 = commandTest
	currentCount = CountNotifiers()

	t.Run("Load Command", func(t *testing.T) {
		Command.Host = "sh"
		Command.Var1 = commandTest
		Command.Var2 = commandTest
		Command.Delay = time.Duration(100 * time.Millisecond)
		Command.Limits = 99
		err := notifier.AddNotifiers(Command)
		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", Command.Author)
		assert.Equal(t, "sh", Command.Host)
		assert.Equal(t, commandTest, Command.Var1)
		assert.Equal(t, commandTest, Command.Var2)
	})

	t.Run("Command Notifier Tester", func(t *testing.T) {
		assert.True(t, Command.CanTest())
	})

	t.Run("Command Within Limits", func(t *testing.T) {
		ok, err := Command.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("Command OnFailure", func(t *testing.T) {
		Command.OnFailure(TestService, TestFailure)
		assert.Equal(t, 1, len(Command.Queue))
	})

	t.Run("Command OnSuccess", func(t *testing.T) {
		Command.OnSuccess(TestService)
		assert.Equal(t, 1, len(Command.Queue))
	})

	t.Run("Command OnSuccess Again", func(t *testing.T) {
		Command.OnSuccess(TestService)
		assert.Equal(t, 1, len(Command.Queue))
		go notifier.Queue(Command)
		time.Sleep(20 * time.Second)
		assert.Equal(t, 0, len(Command.Queue))
	})

	t.Run("Command Within Limits again", func(t *testing.T) {
		ok, err := Command.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("Command Send", func(t *testing.T) {
		Command.Send(commandTest)
		assert.Equal(t, 0, len(Command.Queue))
	})

	t.Run("Command Test", func(t *testing.T) {
		Command.OnTest()
	})

	t.Run("Command Queue", func(t *testing.T) {
		go notifier.Queue(Command)
		time.Sleep(5 * time.Second)
		assert.Equal(t, "sh", Command.Host)
		assert.Equal(t, commandTest, Command.Var1)
		assert.Equal(t, commandTest, Command.Var2)
		assert.Equal(t, 0, len(Command.Queue))
	})

}
