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
	"testing"
	"time"
)

const (
	commandTest = "curl -I https://statup.io/"
)

func TestCommandNotifier(t *testing.T) {
	t.Parallel()
	command.Host = "sh"
	command.Var1 = commandTest
	command.Var2 = commandTest
	currentCount = CountNotifiers()

	t.Run("Load command", func(t *testing.T) {
		command.Host = "sh"
		command.Var1 = commandTest
		command.Var2 = commandTest
		command.Delay = time.Duration(100 * time.Millisecond)
		command.Limits = 99
		err := notifier.AddNotifier(command)
		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", command.Author)
		assert.Equal(t, "sh", command.Host)
		assert.Equal(t, commandTest, command.Var1)
		assert.Equal(t, commandTest, command.Var2)
	})

	t.Run("Load command Notifier", func(t *testing.T) {
		notifier.Load()
	})

	t.Run("command Notifier Tester", func(t *testing.T) {
		assert.True(t, command.CanTest())
	})

	t.Run("command Within Limits", func(t *testing.T) {
		ok, err := command.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("command OnFailure", func(t *testing.T) {
		command.OnFailure(TestService, TestFailure)
		assert.Equal(t, 1, len(command.Queue))
	})

	t.Run("command OnFailure multiple times", func(t *testing.T) {
		for i := 0; i <= 50; i++ {
			command.OnFailure(TestService, TestFailure)
		}
		assert.Equal(t, 52, len(command.Queue))
	})

	t.Run("command Check Offline", func(t *testing.T) {
		assert.False(t, command.Online)
	})

	t.Run("command OnSuccess", func(t *testing.T) {
		command.OnSuccess(TestService)
		assert.Equal(t, 1, len(command.Queue))
	})

	t.Run("command Queue after being online", func(t *testing.T) {
		assert.True(t, command.Online)
		assert.Equal(t, 1, len(command.Queue))
	})

	t.Run("command OnSuccess Again", func(t *testing.T) {
		assert.True(t, command.Online)
		command.OnSuccess(TestService)
		assert.Equal(t, 1, len(command.Queue))
		go notifier.Queue(command)
		time.Sleep(5 * time.Second)
		assert.Equal(t, 0, len(command.Queue))
	})

	t.Run("command Within Limits again", func(t *testing.T) {
		ok, err := command.WithinLimits()
		assert.Nil(t, err)
		assert.True(t, ok)
	})

	t.Run("command Send", func(t *testing.T) {
		err := command.Send(commandTest)
		assert.Nil(t, err)
		assert.Equal(t, 0, len(command.Queue))
	})

	t.Run("command Test", func(t *testing.T) {
		err := command.OnTest()
		assert.Nil(t, err)
	})

	t.Run("command Queue", func(t *testing.T) {
		go notifier.Queue(command)
		time.Sleep(5 * time.Second)
		assert.Equal(t, "sh", command.Host)
		assert.Equal(t, commandTest, command.Var1)
		assert.Equal(t, commandTest, command.Var2)
		assert.Equal(t, 0, len(command.Queue))
	})

}
