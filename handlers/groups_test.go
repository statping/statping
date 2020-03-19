package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
			Name:           "Statping Public and Private Groups",
			URL:            "/api/groups",
			Method:         "GET",
			ExpectedStatus: 200,
			ResponseLen:    2,
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
			Name:           "Statping Delete Group",
			URL:            "/api/groups/1",
			Method:         "DELETE",
			ExpectedStatus: 200,
			AfterTest:      UnsetTestENV,
			SecureRoute:    true,
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			assert.Nil(t, err)
		})
	}
}
