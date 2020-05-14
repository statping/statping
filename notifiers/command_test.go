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
		_, err := Command.OnTest()
		assert.Nil(t, err)
	})

}
