package handlers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGroupAPIRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "Statping Groups",
			URL:            "/api/groups",
			Method:         "GET",
			ExpectedStatus: 200,
			ResponseLen:    3,
		},
		{
			Name:           "Statping View Group",
			URL:            "/api/groups/1",
			Method:         "GET",
			ExpectedStatus: 200,
		},
		{
			Name:        "Statping Create Group",
			URL:         "/api/groups",
			HttpHeaders: []string{"Content-Type=application/json"},
			Body: `{
    "name": "New Group",
    "public": true
}`,
			Method:         "POST",
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Delete Group",
			URL:            "/api/groups/1",
			Method:         "DELETE",
			ExpectedStatus: 200,
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			require.Nil(t, err)
		})
	}
}
