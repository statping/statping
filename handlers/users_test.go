package handlers

import (
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

func TestUnAuthenticatedUserRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "No Authentication - New User",
			URL:            "/api/users",
			Method:         "POST",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - Update User",
			URL:            "/api/users/1",
			Method:         "POST",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - View User",
			URL:            "/api/users/1",
			Method:         "GET",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - Delete User",
			URL:            "/api/users/1",
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

func TestApiUsersRoutes(t *testing.T) {
	form := url.Values{}
	form.Add("username", "adminupdated")
	form.Add("password", "password12345")

	badForm := url.Values{}
	badForm.Add("username", "adminupdated")
	badForm.Add("password", "wrongpassword")

	tests := []HTTPTest{
		{
			Name:           "Check Basic Authentication",
			URL:            "/api",
			Method:         "GET",
			ExpectedStatus: 401,
			BeforeTest: func(t *testing.T) error {
				utils.Params.Set("AUTH_USERNAME", "admin")
				utils.Params.Set("AUTH_PASSWORD", "admin")
				return nil
			},
			AfterTest: func(t *testing.T) error {
				utils.Params.Set("AUTH_USERNAME", "")
				utils.Params.Set("AUTH_PASSWORD", "")
				return nil
			},
		},
		{
			Name:           "Statping All Users",
			URL:            "/api/users",
			Method:         "GET",
			ExpectedStatus: 200,
			ResponseLen:    1,
			BeforeTest:     SetTestENV,
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
			ExpectedStatus:   200,
			ExpectedContains: []string{Success, MethodCreate},
		}, {
			Name:           "Statping View User",
			URL:            "/api/users/1",
			Method:         "GET",
			ExpectedStatus: 200,
		}, {
			Name:           "Statping Incorrect User ID",
			URL:            "/api/users/NOinteger",
			Method:         "GET",
			ExpectedStatus: 422,
		}, {
			Name:           "Statping Missing User",
			URL:            "/api/users/9393939393",
			Method:         "GET",
			ExpectedStatus: 404,
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
			ExpectedStatus:   200,
			ExpectedContains: []string{Success, MethodUpdate},
		}, {
			Name:             "Statping Delete User",
			URL:              "/api/users/2",
			Method:           "DELETE",
			ExpectedStatus:   200,
			ExpectedContains: []string{Success, MethodDelete},
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
		}, {
			Name:           "Statping Logout",
			URL:            "/api/logout",
			Method:         "GET",
			ExpectedStatus: 200,
		}, {
			Name:             "Incorrect JSON POST",
			URL:              "/api/users",
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
			require.Nil(t, err)
		})
	}
}
