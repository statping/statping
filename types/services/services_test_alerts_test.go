package services

import (
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestServiceNotifications(t *testing.T) {
	t.Run("Strategy #1 - Startup - [online, always notify changes, notify after first", func(t *testing.T) {
		allNotifiers[notification.Method] = notification

		service := Example(true)
		service.prevOnline = true // set online during startup
		failure := failures.Example()

		service.UpdateNotify = null.NewNullBool(true)
		service.NotifyAfter = 0

		tests := []notifyTest{
			{
				Name:             "service already online",
				OnSuccess:        true,
				Service:          &service,
				Failure:          &failure,
				Notifier:         notification,
				ExpectedSuccess:  0,
				ExpectedFailures: 0,
				CountLogs:        0,
			},
			{
				Name:             "service triggers online",
				OnSuccess:        true,
				Service:          &service,
				Failure:          &failure,
				Notifier:         notification,
				ExpectedSuccess:  0,
				ExpectedFailures: 0,
				CountLogs:        0,
			},
			{
				Name:             "service triggers offline, was online",
				OnSuccess:        false,
				Service:          &service,
				Failure:          &failure,
				Notifier:         notification,
				ExpectedSuccess:  0,
				ExpectedFailures: 1,
				CountLogs:        1,
			},
			{
				Name:             "service triggers offline again, already was offline, notify",
				OnSuccess:        false,
				Service:          &service,
				Failure:          &failure,
				Notifier:         notification,
				ExpectedSuccess:  0,
				ExpectedFailures: 2,
				CountLogs:        2,
			},
		}

		runNotifyTests(t, notification, tests...)
	})

	t.Run("Strategy #2 - Delayed Failure - [online, notify only 1 time on change, notify after 2 changes", func(t *testing.T) {
		allNotifiers[notification.Method] = notification

		service := Example(true)
		service.prevOnline = true // set online during startup
		failure := failures.Example()
		notif := notification

		service.UpdateNotify = null.NewNullBool(false)
		service.NotifyAfter = 2

		assert.True(t, notif.CanSend())
		assert.True(t, notif.Enabled.Bool)

		tests := []notifyTest{
			{
				Name:             "service already online",
				OnSuccess:        true,
				Service:          &service,
				Failure:          &failure,
				Notifier:         notif,
				ExpectedSuccess:  0,
				ExpectedFailures: 2,
				CountLogs:        2,
			},
			{
				Name:             "service triggers offline (1st)",
				OnSuccess:        false,
				Service:          &service,
				Failure:          &failure,
				Notifier:         notif,
				ExpectedSuccess:  0,
				ExpectedFailures: 2,
				CountLogs:        2,
			},
			{
				Name:             "service triggers offline (2nd) ignore",
				OnSuccess:        false,
				Service:          &service,
				Failure:          &failure,
				Notifier:         notif,
				ExpectedSuccess:  0,
				ExpectedFailures: 2,
				CountLogs:        2,
			},
			{
				Name:             "service triggers offline (3rd) NOTIFY",
				OnSuccess:        false,
				Service:          &service,
				Failure:          &failure,
				Notifier:         notif,
				ExpectedSuccess:  0,
				ExpectedFailures: 3,
				CountLogs:        3,
			},
			{
				Name:             "service triggers back online, notify",
				OnSuccess:        true,
				Service:          &service,
				Failure:          &failure,
				Notifier:         notif,
				ExpectedSuccess:  1,
				ExpectedFailures: 3,
				CountLogs:        4,
			},
		}

		runNotifyTests(t, notif, tests...)
	})

	t.Run("Strategy #3 - Back Online - [offline, notify once for changes, notify after 2 changes", func(t *testing.T) {
		allNotifiers[notification.Method] = notification

		service := Example(false)
		failure := failures.Example()
		notif := notification
		service.prevOnline = false // set offline

		service.UpdateNotify = null.NewNullBool(false)
		service.NotifyAfter = 2

		assert.True(t, notif.CanSend())
		assert.True(t, notif.Enabled.Bool)

		tests := []notifyTest{
			{
				Name:             "service already offline",
				OnSuccess:        false,
				Service:          &service,
				Failure:          &failure,
				Notifier:         notif,
				ExpectedSuccess:  1,
				ExpectedFailures: 3,
				CountLogs:        4,
			},
			{
				Name:             "service triggers offline again, ignore",
				OnSuccess:        false,
				Service:          &service,
				Failure:          &failure,
				Notifier:         notif,
				ExpectedSuccess:  1,
				ExpectedFailures: 3,
				CountLogs:        4,
			},
			{
				Name:             "service triggers offline again, ignore",
				OnSuccess:        false,
				Service:          &service,
				Failure:          &failure,
				Notifier:         notif,
				ExpectedSuccess:  1,
				ExpectedFailures: 3,
				CountLogs:        4,
			},
			{
				Name:             "service triggers back online, NOTIFY",
				OnSuccess:        true,
				Service:          &service,
				Failure:          &failure,
				Notifier:         notif,
				ExpectedSuccess:  2,
				ExpectedFailures: 3,
				CountLogs:        5,
			},
			{
				Name:             "service triggers online, ignore",
				OnSuccess:        true,
				Service:          &service,
				Failure:          &failure,
				Notifier:         notif,
				ExpectedSuccess:  2,
				ExpectedFailures: 3,
				CountLogs:        5,
			},
		}

		runNotifyTests(t, notif, tests...)
	})

	t.Run("Strategy #4 - Disabled - [online, notifications are disabled", func(t *testing.T) {
		allNotifiers[notification.Method] = notification
		service := Example(false)
		service.prevOnline = true // set online during startup
		service.AllowNotifications = null.NewNullBool(false)
		failure := failures.Example()
		notif := notification

		tests := []notifyTest{
			{
				Name:             "service offline",
				OnSuccess:        true,
				Service:          &service,
				Failure:          &failure,
				Notifier:         notif,
				ExpectedSuccess:  2,
				ExpectedFailures: 3,
				CountLogs:        5,
			},
			{
				Name:             "service online",
				OnSuccess:        false,
				Service:          &service,
				Failure:          &failure,
				Notifier:         notif,
				ExpectedSuccess:  2,
				ExpectedFailures: 3,
				CountLogs:        5,
			},
			{
				Name:             "service offline",
				OnSuccess:        true,
				Service:          &service,
				Failure:          &failure,
				Notifier:         notif,
				ExpectedSuccess:  2,
				ExpectedFailures: 3,
				CountLogs:        5,
			},
		}

		runNotifyTests(t, notif, tests...)
	})

	t.Run("Test Samples", func(t *testing.T) {
		require.Nil(t, Samples())
		assert.Len(t, All(), 11)
	})

	t.Run("Test Close", func(t *testing.T) {
		assert.Nil(t, db.Close())
	})

}

func runNotifyTests(t *testing.T, notif *exampleNotifier, tests ...notifyTest) {
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if test.OnSuccess {
				RecordSuccess(test.Service)
			} else {
				RecordFailure(test.Service, "test issue", "lookup")
			}

			assert.Equal(t, test.ExpectedSuccess, notif.success)
			assert.Equal(t, test.ExpectedFailures, notif.failures)
			assert.Equal(t, test.CountLogs, notif.LastSentCount)
		})
	}
}

var notification = &exampleNotifier{Notification: &notifications.Notification{
	Method:    "test",
	CreatedAt: utils.Now().Add(-5 * time.Second),
	Limits:    60,
	Enabled:   null.NewNullBool(true),
}, failures: 0, success: 0}

type notifyTest struct {
	Name             string
	OnSuccess        bool
	Service          *Service
	Failure          *failures.Failure
	Notifier         ServiceNotifier
	ExpectedSuccess  int
	ExpectedFailures int
	CountLogs        int
}

type exampleNotifier struct {
	*notifications.Notification
	failures int
	success  int
	saves    int
	tests    int
}

func (e *exampleNotifier) OnSuccess(s Service) (string, error) {
	e.success++
	return "", nil
}

func (e *exampleNotifier) OnFailure(s Service, f failures.Failure) (string, error) {
	e.failures++
	return "", nil
}

func (e *exampleNotifier) OnSave() (string, error) {
	e.saves++
	return "", nil
}

func (e *exampleNotifier) Select() *notifications.Notification {
	return e.Notification
}

func (e *exampleNotifier) OnTest() (string, error) {
	e.tests++
	return "", nil
}

func (e *exampleNotifier) Valid(form notifications.Values) error {
	return nil
}
