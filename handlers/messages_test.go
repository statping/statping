package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnAuthenticatedMessageRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "No Authentication - New Message",
			URL:            "/api/messages",
			Method:         "POST",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - Update Message",
			URL:            "/api/messages/1",
			Method:         "POST",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - Delete Message",
			URL:            "/api/messages/1",
			Method:         "DELETE",
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

func TestMessagesApiRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:             "Statping Messages",
			URL:              "/api/messages",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{`"title":"Routine Downtime"`},
		}, {
			Name:   "Statping Create Message",
			URL:    "/api/messages",
			Method: "POST",
			Body: `{
					"title": "API Message",
					"description": "This is an example a upcoming message for a service!",
					"start_on": "2022-11-17T03:28:16.323797-08:00",
					"end_on": "2022-11-17T05:13:16.323798-08:00",
					"service": 1,
					"notify_users": true,
					"notify_method": "email",
					"notify_before": 6,
					"notify_before_scale": "hour"
				}`,
			ExpectedStatus:   200,
			ExpectedContains: []string{Success, `"type":"message"`, `"method":"create"`, `"title":"API Message"`},
			BeforeTest:       SetTestENV,
			AfterTest:        UnsetTestENV,
			SecureRoute:      true,
		},
		{
			Name:             "Statping View Message",
			URL:              "/api/messages/1",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{`"title":"Routine Downtime"`},
		},
		{
			Name:   "Statping Update Message",
			URL:    "/api/messages/1",
			Method: "POST",
			Body: `{
					"title": "Updated Message",
					"description": "This message was updated",
					"start_on": "2022-11-17T03:28:16.323797-08:00",
					"end_on": "2022-11-17T05:13:16.323798-08:00",
					"service": 1,
					"notify_users": true,
					"notify_method": "email",
					"notify_before": 3,
					"notify_before_scale": "hour"
				}`,
			ExpectedStatus:   200,
			ExpectedContains: []string{Success, `"type":"message"`, MethodUpdate},
			BeforeTest:       SetTestENV,
			SecureRoute:      true,
		},
		{
			Name:             "Statping Delete Message",
			URL:              "/api/messages/1",
			Method:           "DELETE",
			ExpectedStatus:   200,
			ExpectedContains: []string{Success, MethodDelete},
			BeforeTest:       SetTestENV,
			SecureRoute:      true,
		},
		{
			Name:           "Statping Missing Message",
			URL:            "/api/messages/999999",
			Method:         "GET",
			ExpectedStatus: 404,
		},
		{
			Name:             "Incorrect JSON POST",
			URL:              "/api/messages",
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
