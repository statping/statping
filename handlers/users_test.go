package handlers

import (
	"github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

func TestApiUsersRoutes(t *testing.T) {
	form := url.Values{}
	form.Add("username", "adminupdated")
	form.Add("password", "password12345")

	badForm := url.Values{}
	badForm.Add("username", "adminupdated")
	badForm.Add("password", "wrongpassword")

	tests := []HTTPTest{
		{
			Name:           "Statping All Users",
			URL:            "/api/users",
			Method:         "GET",
			ExpectedStatus: 200,
			ResponseLen:    1,
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
			URL:            "/api/users/2",
			Method:         "DELETE",
			ExpectedStatus: 200,
		}, {
			Name:             "Statping Login User",
			URL:              "/api/login",
			Method:           "POST",
			Body:             form.Encode(),
			ExpectedContains: []string{`"token"`},
			ExpectedStatus:   200,
			HttpHeaders:      []string{"Content-Type=application/x-www-form-urlencoded"},
		}, {
			Name:             "Statping Bad Login User",
			URL:              "/api/login",
			Method:           "POST",
			Body:             badForm.Encode(),
			ExpectedContains: []string{`incorrect authentication`},
			ExpectedStatus:   200,
			HttpHeaders:      []string{"Content-Type=application/x-www-form-urlencoded"},
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			require.Nil(t, err)
		})
	}
}
