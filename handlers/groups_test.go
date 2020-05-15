package handlers

import (
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/groups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnAuthenticatedGroupRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "No Authentication - New Group",
			URL:            "/api/groups",
			Method:         "POST",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - Update Group",
			URL:            "/api/groups/1",
			Method:         "POST",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - Delete Group",
			URL:            "/api/groups/1",
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

func TestGroupAPIRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "Statping Public Groups",
			URL:            "/api/groups",
			Method:         "GET",
			ExpectedStatus: 200,
			ResponseLen:    3,
			BeforeTest:     SetTestENV,
			AfterTest:      UnsetTestENV,
		},
		{
			Name:           "Statping View Public Group",
			URL:            "/api/groups/1",
			Method:         "GET",
			ExpectedStatus: 200,
			BeforeTest:     SetTestENV,
			AfterTest:      UnsetTestENV,
		},
		{
			Name:        "Statping Create Public Group",
			URL:         "/api/groups",
			HttpHeaders: []string{"Content-Type=application/json"},
			Body: `{
					"name": "New Group",
					"public": true
				}`,
			Method:         "POST",
			ExpectedStatus: 200,
			BeforeTest:     SetTestENV,
		},
		{
			Name:        "Statping Create Private Group",
			URL:         "/api/groups",
			HttpHeaders: []string{"Content-Type=application/json"},
			Body: `{
					"name": "New Private Group",
					"public": false
				}`,
			Method:         "POST",
			ExpectedStatus: 200,
		},
		{
			Name:             "Incorrect JSON POST",
			URL:              "/api/groups",
			Body:             BadJSON,
			ExpectedContains: []string{BadJSONResponse},
			Method:           "POST",
			ExpectedStatus:   422,
		},
		{
			Name:           "Statping Public and Private Groups",
			URL:            "/api/groups",
			Method:         "GET",
			ExpectedStatus: 200,
			ResponseLen:    3,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "Statping View Private Group",
			URL:            "/api/groups/2",
			Method:         "GET",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "Statping View Private Group Allowed",
			URL:            "/api/groups/2",
			Method:         "GET",
			ExpectedStatus: 200,
			BeforeTest:     SetTestENV,
		},
		{
			Name:           "Statping View Private Group with API Key",
			URL:            "/api/groups/2?api=" + core.App.ApiSecret,
			Method:         "GET",
			ExpectedStatus: 200,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "Statping View Private Group with API Header",
			URL:            "/api/groups/2",
			Method:         "GET",
			HttpHeaders:    []string{"Authorization=" + core.App.ApiSecret},
			ExpectedStatus: 200,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "Statping Reorder Groups",
			URL:            "/api/reorder/groups",
			Method:         "POST",
			Body:           `[{"group":1,"order":2},{"group":2,"order":1}]`,
			ExpectedStatus: 200,
			HttpHeaders:    []string{"Content-Type=application/json"},
			BeforeTest:     SetTestENV,
			SecureRoute:    true,
		},
		{
			Name:           "Statping View Unknown Group",
			URL:            "/api/groups/8383883838",
			Method:         "GET",
			BeforeTest:     SetTestENV,
			ExpectedStatus: 404,
		},
		{
			Name:   "Statping Update Group",
			URL:    "/api/groups/1",
			Method: "POST",
			Body: `{
					"name": "Updated Group",
					"public": false
					}`,
			ExpectedStatus:   200,
			ExpectedContains: []string{Success, MethodUpdate},
			BeforeTest:       SetTestENV,
			SecureRoute:      true,
			AfterTest: func(t *testing.T) error {
				g, err := groups.Find(1)
				require.Nil(t, err)
				assert.Equal(t, "Updated Group", g.Name)
				return nil
			},
		},
		{
			Name:             "Statping Delete Group",
			URL:              "/api/groups/1",
			Method:           "DELETE",
			ExpectedStatus:   200,
			ExpectedContains: []string{Success, MethodDelete},
			AfterTest:        UnsetTestENV,
			SecureRoute:      true,
		},
	}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			assert.Nil(t, err)
		})
	}
}
