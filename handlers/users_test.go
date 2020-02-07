package handlers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:             "Users",
			URL:              "/users",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{`<title>Statping | Users</title>`},
		},
		{
			Name:             "User 2",
			URL:              "/user/2",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{`<title>Statping | testadmin2</title>`},
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			require.Nil(t, err)
		})
	}
}
