package handlers

import (
	"github.com/statping/statping/notifiers"
	"github.com/statping/statping/types/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAttachment(t *testing.T) {
	notifiers.InitNotifiers()
}

func TestUnAuthenticatedNotifierRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "No Authentication - View All Notifiers",
			URL:            "/api/notifiers",
			Method:         "GET",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - View Notifier",
			URL:            "/api/notifier/slack",
			Method:         "GET",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - Update Notifier",
			URL:            "/api/notifier/slack",
			Method:         "POST",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
	}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			str, t, err := RunHTTPTest(v, t)
			t.Logf("Test %s: \n %v\n", v.Name, str)
			assert.Nil(t, err)
		})
	}
}

func TestApiNotifiersRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "Statping Notifiers",
			URL:            "/api/notifiers",
			Method:         "GET",
			ExpectedStatus: 200,
			ResponseLen:    len(services.AllNotifiers()),
			BeforeTest:     SetTestENV,
			SecureRoute:    true,
		}, {
			Name:           "Statping Slack Notifier",
			URL:            "/api/notifier/slack",
			Method:         "GET",
			ExpectedStatus: 200,
			BeforeTest:     SetTestENV,
			SecureRoute:    true,
		}, {
			Name:   "Statping Update Notifier",
			URL:    "/api/notifier/slack",
			Method: "POST",
			Body: `{
					"method": "slack",
					"host": "https://slack.api/example/12345",
					"enabled": true,
					"limits": 55
				}`,
			ExpectedStatus: 200,
			BeforeTest:     SetTestENV,
			SecureRoute:    true,
		}, {
			Name:             "Statping Slack Notifier",
			URL:              "/api/notifier/slack",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{`"method":"slack"`},
			BeforeTest:       SetTestENV,
			SecureRoute:      true,
		},
		{
			Name:             "Incorrect JSON POST",
			URL:              "/api/notifier/slack",
			Body:             BadJSON,
			ExpectedContains: []string{BadJSONResponse},
			BeforeTest:       SetTestENV,
			Method:           "POST",
			ExpectedStatus:   422,
		},
	}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			assert.Nil(t, err)
		})
	}
}
