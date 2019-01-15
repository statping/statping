package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestGenericRoutes(t *testing.T) {

	form := url.Values{}
	form.Add("username", "admin")
	form.Add("password", "password123")

	form2 := url.Values{}
	form2.Add("username", "admin")
	form2.Add("password", "wrongpassword")

	tests := []HTTPTest{
		{
			Name:           "Statping Index",
			URL:            "/",
			Method:         "GET",
			ExpectedStatus: 200,
			ExpectedContains: []string{
				`<title>Tester Status</title>`,
				`<footer>`,
			},
		},
		{
			Name:             "Statping Incorrect Login",
			URL:              "/dashboard",
			Method:           "POST",
			Body:             form.Encode(),
			HttpHeaders:      []string{"Content-Type=application/x-www-form-urlencoded"},
			ExpectedContains: []string{"Incorrect login information submitted, try again."},
			ExpectedStatus:   200,
		},
		{
			Name:             "Dashboard Routes",
			URL:              "/dashboard",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{"<title>Statping | Dashboard</title>"},
		},
		{
			Name:             "Services Routes",
			URL:              "/services",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{"<title>Statping | Services</title>"},
		},
		{
			Name:             "Users Routes",
			URL:              "/users",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{"<title>Statping | Users</title>"},
		},
		{
			Name:             "Messages Routes",
			URL:              "/messages",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{"<title>Statping Messages</title>"},
		},
		{
			Name:             "Settings Routes",
			URL:              "/settings",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{"<title>Statping | Settings</title>"},
		},
		{
			Name:             "Logs Routes",
			URL:              "/logs",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{"<title>Statping | Logs</title>"},
		},
		{
			Name:             "Help Routes",
			URL:              "/help",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{"<title>Statping | Help</title>"},
		},
		{
			Name:           "Logout",
			URL:            "/logout",
			Method:         "GET",
			ExpectedStatus: 303,
		},
		{
			Name:             "Prometheus Metrics Routes",
			URL:              "/metrics",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{"statping_total_services 5"},
		},
	}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			body, t, err := RunHTTPTest(v, t)
			assert.Nil(t, err)
			if err != nil {
				t.FailNow()
			}
			t.Logf("Test %v got: %v\n", v.Name, string(body))
		})
	}
}
