package handlers

import (
	"github.com/hunterlong/statping/utils"
	"github.com/stretchr/testify/require"
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

	form3 := url.Values{}
	form3.Add("variables", "$background-color: #fcfcfc;")

	tests := []HTTPTest{
		{
			Name:           "Statping Index",
			URL:            "/",
			Method:         "GET",
			ExpectedStatus: 200,
			ExpectedContains: []string{
				`<title>Statping Sample Data Status</title>`,
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
			ExpectedContains: []string{"statping_total_services 15"},
		},
		{
			Name:           "Last Log Line",
			URL:            "/logs/line",
			Method:         "GET",
			ExpectedStatus: 200,
		},
		{
			Name:           "Export JSON file of all objcts",
			URL:            "/settings/export",
			Method:         "GET",
			ExpectedStatus: 200,
		},
		{
			Name:           "Export Static Assets",
			URL:            "/settings/build",
			Method:         "GET",
			ExpectedStatus: 303,
			ExpectedFiles:  []string{utils.Directory + "/assets/css/base.css"},
		},
		{
			Name:           "Save SCSS",
			URL:            "/settings/css",
			Method:         "POST",
			Body:           form3.Encode(),
			ExpectedStatus: 303,
			HttpHeaders:    []string{"Content-Type=application/x-www-form-urlencoded"},
		},
		{
			Name:           "Delete Assets",
			URL:            "/settings/delete_assets",
			Method:         "GET",
			ExpectedStatus: 303,
		},
	}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			require.Nil(t, err)
		})
	}
}
