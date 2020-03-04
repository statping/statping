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
		},
		{
			Name:           "Statping Checkins",
			URL:            "/api/checkins",
			Method:         "GET",
			ExpectedStatus: 200,
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			require.Nil(t, err)
		})
	}
}
