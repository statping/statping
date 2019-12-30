package handlers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestServiceRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:             "Services",
			URL:              "/services",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{`<title>Statping | Services</title>`},
		},
		{
			Name:             "Services 2",
			URL:              "/service/2",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{`<title>Statping Github Status</title>`},
		}, {
			Name:           "chart.js index file",
			URL:            "/charts.js",
			Method:         "GET",
			ExpectedStatus: 200,
		},
	}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			require.Nil(t, err)
		})
	}
}
