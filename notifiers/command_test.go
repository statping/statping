package notifiers

import (
	"testing"
	"time"

	"github.com/statping/statping/database"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommandNotifier(t *testing.T) {
	t.Parallel()
	t.SkipNow()
	err := utils.InitLogs()
	require.Nil(t, err)
	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)
	core.Example()

	t.Run("Load Command", func(t *testing.T) {
		Command.Host = null.NewNullString("/bin/echo")
		Command.Var1 = null.NewNullString("service {{.Service.Domain}} is online")
		Command.Var2 = null.NewNullString("service {{.Service.Domain}} is offline")
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

	t.Run("Command OnSave", func(t *testing.T) {
		_, err := Command.OnSave()
		assert.Nil(t, err)
	})

	t.Run("Command OnFailure", func(t *testing.T) {
		_, err := Command.OnFailure(services.Example(false), failures.Example())
		assert.Nil(t, err)
	})

	t.Run("Command OnSuccess", func(t *testing.T) {
		_, err := Command.OnSuccess(services.Example(true))
		assert.Nil(t, err)
	})

	t.Run("Command Test", func(t *testing.T) {
		_, err := Command.OnTest()
		assert.Nil(t, err)
	})

}
