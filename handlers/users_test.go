// +build int

package handlers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestApiUsersRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "Statping All Users",
			URL:            "/api/users",
			Method:         "GET",
			ExpectedStatus: 200,
			ResponseLen:    4,
		}, {
			Name:        "Statping Create User",
			URL:         "/api/users",
			HttpHeaders: []string{"Content-Type=application/json"},
			Method:      "POST",
			Body: `{
    "username": "adminuser2",
    "email": "info@adminemail.com",
    "password": "passsword123",
    "admin": true
}`,
			ExpectedStatus: 200,
		}, {
			Name:           "Statping View User",
			URL:            "/api/users/1",
			Method:         "GET",
			ExpectedStatus: 200,
		}, {
			Name:   "Statping Update User",
			URL:    "/api/users/1",
			Method: "POST",
			Body: `{
    "username": "adminupdated",
    "email": "info@email.com",
    "password": "password12345",
    "admin": true
}`,
			ExpectedStatus: 200,
		}, {
			Name:           "Statping Delete User",
			URL:            "/api/users/1",
			Method:         "DELETE",
			ExpectedStatus: 200,
		}, {
			Name:   "Statping Login User",
			URL:    "/api/login",
			Method: "POST",
			Body: `{
    "username": "admin",
    "password": "admin"
}`, ExpectedContains: []string{`"token"`},
			ExpectedStatus: 200,
		}, {
			Name:   "Statping Bad Login User",
			URL:    "/api/login",
			Method: "POST",
			Body: `{
    "username": "admin",
    "password": "wrongpassword"
}`, ExpectedContains: []string{`"token"`},
			ExpectedStatus: 200,
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			require.Nil(t, err)
		})
	}
}
