package handlers

import (
	"github.com/stretchr/testify/assert"
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
			ExpectedContains: []string{`<title>Statping | adminuser2</title>`},
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			assert.Nil(t, err)
			if err != nil {
				t.FailNow()
			}
		})
	}
}
