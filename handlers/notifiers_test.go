// +build int

package handlers

import (
	"github.com/statping/statping/notifiers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAttachment(t *testing.T) {
	err := notifiers.AttachNotifiers()
	require.Nil(t, err)
}

func TestApiNotifiersRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "Statping Notifiers",
			URL:            "/api/notifiers",
			Method:         "GET",
			ExpectedStatus: 200,
			BeforeTest:     SetTestENV,
		}, {
			Name:           "Statping Mobile Notifier",
			URL:            "/api/notifier/mobile",
			Method:         "GET",
			ExpectedStatus: 200,
			BeforeTest:     SetTestENV,
		}, {
			Name:   "Statping Update Notifier",
			URL:    "/api/notifier/mobile",
			Method: "POST",
			Body: `{
					"method": "mobile",
					"var1": "ExponentPushToken[ToBadIWillError123456]",
					"enabled": true,
					"limits": 55
				}`,
			ExpectedStatus: 200,
			BeforeTest:     SetTestENV,
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			assert.Nil(t, err)
		})
	}
}
