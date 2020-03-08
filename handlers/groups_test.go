package handlers

import (
	"github.com/stretchr/testify/require"
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
			Name:           "Statping Public and Private Groups",
			URL:            "/api/groups",
			Method:         "GET",
			ExpectedStatus: 200,
			ResponseLen:    3,
			BeforeTest:     UnsetTestENV,
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
			Name:           "Statping View Private Group",
			URL:            "/api/groups/2",
			Method:         "GET",
			ExpectedStatus: 404,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "Statping View Unknown Group",
			URL:            "/api/groups/8383883838",
			Method:         "GET",
			ExpectedStatus: 404,
		},
		{
			Name:           "Statping Delete Group",
			URL:            "/api/groups/1",
			Method:         "DELETE",
			ExpectedStatus: 200,
			AfterTest:      UnsetTestENV,
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			require.Nil(t, err)
		})
	}
}
