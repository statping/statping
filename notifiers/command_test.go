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
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/null"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCommandNotifier(t *testing.T) {
	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)

	t.Run("Load Command", func(t *testing.T) {
		Command.Host = "/bin/echo"
		Command.Var1 = "service {{.Service.Domain}} is online"
		Command.Var2 = "service {{.Service.Domain}} is offline"
		Command.Delay = time.Duration(100 * time.Millisecond)
		Command.Limits = 99
		Command.Enabled = null.NewNullBool(true)

		Add(Command)

		assert.Equal(t, "Hunter Long", Command.Author)
		assert.Equal(t, "/bin/echo", Command.Host)
	})

	t.Run("Command Notifier Tester", func(t *testing.T) {
		assert.True(t, Command.CanSend())
	})

	t.Run("Command OnFailure", func(t *testing.T) {
		err := Command.OnFailure(exampleService, exampleFailure)
		assert.Nil(t, err)
	})

	t.Run("Command OnSuccess", func(t *testing.T) {
		err := Command.OnSuccess(exampleService)
		assert.Nil(t, err)
	})

	t.Run("Command Test", func(t *testing.T) {
		err := Command.OnTest()
		assert.Nil(t, err)
	})

}
