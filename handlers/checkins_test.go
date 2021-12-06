package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnAuthenticatedCheckinRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "No Authentication - New Checkin",
			URL:            "/api/checkins",
			Method:         "POST",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - Delete Checkin",
			URL:            "/api/checkins/1",
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

func TestApiCheckinRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:             "Statping Create Checkins",
			URL:              "/api/checkins",
			Method:           "POST",
			ExpectedStatus:   200,
			BeforeTest:       SetTestENV,
			SecureRoute:      true,
			ExpectedContains: []string{Success},
			Body: `{
						"name": "Example Checkin",
						"service_id": 1,
						"interval": 300,
						"api_key": "example"
					}`,
		},
		{
			Name:           "Statping Checkins",
			URL:            "/api/checkins",
			Method:         "GET",
			ExpectedStatus: 200,
			ResponseLen:    3,
			BeforeTest:     SetTestENV,
		},
		{
			Name:           "Statping View Checkin",
			URL:            "/api/checkins/example",
			Method:         "GET",
			ExpectedStatus: 200,
			BeforeTest:     SetTestENV,
			SecureRoute:    true,
		},
		{
			Name:           "Statping Trigger Checkin",
			URL:            "/checkin/example",
			Method:         "GET",
			ExpectedStatus: 200,
			SecureRoute:    true,
			BeforeTest:     SetTestENV,
		},
		{
			Name:           "Statping Missing Trigger Checkin",
			URL:            "/checkin/missing123",
			Method:         "GET",
			BeforeTest:     SetTestENV,
			ExpectedStatus: 404,
		},
		{
			Name:           "Statping Missing Checkin",
			URL:            "/api/checkins/missing123",
			Method:         "GET",
			BeforeTest:     SetTestENV,
			ExpectedStatus: 404,
		},
		{
			Name:             "Statping Delete Checkin",
			URL:              "/api/checkins/example",
			Method:           "DELETE",
			BeforeTest:       SetTestENV,
			ExpectedContains: []string{Success},
			ExpectedStatus:   200,
		},
		{
			Name:             "Incorrect JSON POST",
			URL:              "/api/checkins",
			Body:             BadJSON,
			ExpectedContains: []string{BadJSONResponse},
			Method:           "POST",
			ExpectedStatus:   422,
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
