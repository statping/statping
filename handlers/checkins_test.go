package handlers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestApiCheckinRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "Statping Checkins",
			URL:            "/api/checkins",
			Method:         "GET",
			ExpectedStatus: 200,
			BeforeTest:     SetTestENV,
		}, {
			Name:           "Statping Checkins",
			URL:            "/api/checkins",
			Method:         "GET",
			ExpectedStatus: 200,
			BeforeTest:     UnsetTestENV,
		}, {
			Name:   "Statping Create Checkin",
			URL:    "/api/checkin",
			Method: "POST",
			Body: `{
					"service_id": 2,
					"name": "Server Checkin",
					"interval": 900,
					"grace": 60
				}`,
			ExpectedStatus:   200,
			ExpectedContains: []string{`"status":"success","type":"checkin","method":"create"`},
			BeforeTest:       SetTestENV,
		},
		{
			Name:           "Statping Checkins",
			URL:            "/api/checkins",
			Method:         "GET",
			ExpectedStatus: 200,
			AfterTest:      UnsetTestENV,
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			str, t, err := RunHTTPTest(v, t)
			t.Logf("Test %s: \n %v\n", v.Name, str)
			require.Nil(t, err)
		})
	}
}
