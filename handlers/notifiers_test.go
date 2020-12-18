package handlers

import (
	"testing"

	"github.com/statping/statping/notifiers"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAttachment(t *testing.T) {
	notifiers.InitNotifiers()
}

func TestAuthenticatedNotifierRoutes(t *testing.T) {
	slackWebhookUrl := utils.Params.GetString("SLACK_URL")

	tests := []HTTPTest{
		{
			Name:           "No Authentication - View All Notifiers",
			URL:            "/api/notifiers",
			Method:         "GET",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "View All Notifiers",
			URL:            "/api/notifiers",
			Method:         "GET",
			ExpectedStatus: 200,
			BeforeTest:     SetTestENV,
		},
		{
			Name:           "No Authentication - View Notifier",
			URL:            "/api/notifier/slack",
			Method:         "GET",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "View Notifier",
			URL:            "/api/notifier/slack",
			Method:         "GET",
			ExpectedStatus: 200,
			BeforeTest: func(t *testing.T) error {
				notif := services.FindNotifier("slack")
				require.NotNil(t, notif)
				assert.Equal(t, "slack", notif.Method)
				return SetTestENV(t)
			},
		},
		{
			Name:   "No Authentication - Update Notifier",
			URL:    "/api/notifier/slack",
			Method: "POST",
			Body: `{
					"method": "slack",
					"host": "` + slackWebhookUrl + `",
					"enabled": true,
					"limits": 55
				}`,
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:   "Update Notifier",
			URL:    "/api/notifier/slack",
			Method: "POST",
			Body: `{
					"method": "slack",
					"host": "` + slackWebhookUrl + `",
					"enabled": true,
					"limits": 55
				}`,
			ExpectedStatus: 200,
			BeforeTest: func(t *testing.T) error {
				notif := services.FindNotifier("slack")
				require.NotNil(t, notif)
				assert.Equal(t, "slack", notif.Method)
				assert.False(t, notif.Enabled.Bool)
				return SetTestENV(t)
			},
			AfterTest: func(t *testing.T) error {
				notif := services.FindNotifier("slack")
				require.NotNil(t, notif)
				assert.Equal(t, "slack", notif.Method)
				assert.True(t, notif.Enabled.Bool)
				return UnsetTestENV(t)
			},
		},
		{
			Name:   "Test Notifier (OnSuccess)",
			URL:    "/api/notifier/slack/test",
			Method: "POST",
			Body: `{
				"method": "success",
				"notifier": {
					"enabled": false,
					"limits": 60,
					"method": "slack",
					"host": "` + slackWebhookUrl + `",
					"success_data": "{\n  \"blocks\": [{\n    \"type\": \"section\",\n    \"text\": {\n      \"type\": \"mrkdwn\",\n      \"text\": \"The service {{.Service.Name}} is back online.\"\n    }\n  }, {\n    \"type\": \"actions\",\n    \"elements\": [{\n      \"type\": \"button\",\n      \"text\": {\n        \"type\": \"plain_text\",\n        \"text\": \"View Service\",\n        \"emoji\": true\n      },\n      \"style\": \"primary\",\n      \"url\": \"{{.Core.Domain}}/service/{{.Service.Id}}\"\n    }, {\n      \"type\": \"button\",\n      \"text\": {\n        \"type\": \"plain_text\",\n        \"text\": \"Go to Statping\",\n        \"emoji\": true\n      },\n      \"url\": \"{{.Core.Domain}}\"\n    }]\n  }]\n}",
					"failure_data": "{\n  \"blocks\": [{\n    \"type\": \"section\",\n    \"text\": {\n      \"type\": \"mrkdwn\",\n      \"text\": \":warning: The service {{.Service.Name}} is currently offline! :warning:\"\n    }\n  }, {\n    \"type\": \"divider\"\n  }, {\n    \"type\": \"section\",\n    \"fields\": [{\n      \"type\": \"mrkdwn\",\n      \"text\": \"*Service:*\\n{{.Service.Name}}\"\n    }, {\n      \"type\": \"mrkdwn\",\n      \"text\": \"*URL:*\\n{{.Service.Domain}}\"\n    }, {\n      \"type\": \"mrkdwn\",\n      \"text\": \"*Status Code:*\\n{{.Service.LastStatusCode}}\"\n    }, {\n      \"type\": \"mrkdwn\",\n      \"text\": \"*When:*\\n{{.Failure.CreatedAt}}\"\n    }, {\n      \"type\": \"mrkdwn\",\n      \"text\": \"*Downtime:*\\n{{.Service.Downtime.Human}}\"\n    }, {\n      \"type\": \"plain_text\",\n      \"text\": \"*Error:*\\n{{.Failure.Issue}}\"\n    }]\n  }, {\n    \"type\": \"divider\"\n  }, {\n    \"type\": \"actions\",\n    \"elements\": [{\n      \"type\": \"button\",\n      \"text\": {\n        \"type\": \"plain_text\",\n        \"text\": \"View Offline Service\",\n        \"emoji\": true\n      },\n      \"style\": \"danger\",\n      \"url\": \"{{.Core.Domain}}/service/{{.Service.Id}}\"\n    }, {\n      \"type\": \"button\",\n      \"text\": {\n        \"type\": \"plain_text\",\n        \"text\": \"Go to Statping\",\n        \"emoji\": true\n      },\n      \"url\": \"{{.Core.Domain}}\"\n    }]\n  }]\n}"
				}
			}`,
			ExpectedStatus:   200,
			ExpectedContains: []string{`"success":true`},
			BeforeTest:       SetTestENV,
		},
		{
			Name:   "Test Notifier (OnFailure)",
			URL:    "/api/notifier/slack/test",
			Method: "POST",
			Body: `{
				"method": "failure",
				"notifier": {
					"enabled": false,
					"limits": 60,
					"method": "slack",
					"host": "` + slackWebhookUrl + `",
					"success_data": "{\n  \"blocks\": [{\n    \"type\": \"section\",\n    \"text\": {\n      \"type\": \"mrkdwn\",\n      \"text\": \"The service {{.Service.Name}} is back online.\"\n    }\n  }, {\n    \"type\": \"actions\",\n    \"elements\": [{\n      \"type\": \"button\",\n      \"text\": {\n        \"type\": \"plain_text\",\n        \"text\": \"View Service\",\n        \"emoji\": true\n      },\n      \"style\": \"primary\",\n      \"url\": \"{{.Core.Domain}}/service/{{.Service.Id}}\"\n    }, {\n      \"type\": \"button\",\n      \"text\": {\n        \"type\": \"plain_text\",\n        \"text\": \"Go to Statping\",\n        \"emoji\": true\n      },\n      \"url\": \"{{.Core.Domain}}\"\n    }]\n  }]\n}",
					"failure_data": "{\n  \"blocks\": [{\n    \"type\": \"section\",\n    \"text\": {\n      \"type\": \"mrkdwn\",\n      \"text\": \":warning: The service {{.Service.Name}} is currently offline! :warning:\"\n    }\n  }, {\n    \"type\": \"divider\"\n  }, {\n    \"type\": \"section\",\n    \"fields\": [{\n      \"type\": \"mrkdwn\",\n      \"text\": \"*Service:*\\n{{.Service.Name}}\"\n    }, {\n      \"type\": \"mrkdwn\",\n      \"text\": \"*URL:*\\n{{.Service.Domain}}\"\n    }, {\n      \"type\": \"mrkdwn\",\n      \"text\": \"*Status Code:*\\n{{.Service.LastStatusCode}}\"\n    }, {\n      \"type\": \"mrkdwn\",\n      \"text\": \"*When:*\\n{{.Failure.CreatedAt}}\"\n    }, {\n      \"type\": \"mrkdwn\",\n      \"text\": \"*Downtime:*\\n{{.Service.Downtime.Human}}\"\n    }, {\n      \"type\": \"plain_text\",\n      \"text\": \"*Error:*\\n{{.Failure.Issue}}\"\n    }]\n  }, {\n    \"type\": \"divider\"\n  }, {\n    \"type\": \"actions\",\n    \"elements\": [{\n      \"type\": \"button\",\n      \"text\": {\n        \"type\": \"plain_text\",\n        \"text\": \"View Offline Service\",\n        \"emoji\": true\n      },\n      \"style\": \"danger\",\n      \"url\": \"{{.Core.Domain}}/service/{{.Service.Id}}\"\n    }, {\n      \"type\": \"button\",\n      \"text\": {\n        \"type\": \"plain_text\",\n        \"text\": \"Go to Statping\",\n        \"emoji\": true\n      },\n      \"url\": \"{{.Core.Domain}}\"\n    }]\n  }]\n}"
				}
			}`,
			ExpectedStatus:   200,
			ExpectedContains: []string{`"success":true`},
			BeforeTest:       SetTestENV,
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
					"host": "https://hooks.slack.com/services/TTJ1B49DP/XBNU09O9M/9uI2123SUnYBuGcxLopZomz9H",
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
			ExpectedContains: []string{`"method":"slack"`, `"host":"https://hooks.slack.com/services/TTJ1B49DP/XBNU09O9M/9uI2123SUnYBuGcxLopZomz9H"`},
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
