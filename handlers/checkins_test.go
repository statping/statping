package handlers

import (
	"encoding/json"
	"github.com/statping/statping/types/checkins"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestApiCheckinRoutes(t *testing.T) {
	var apiToken string
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
						"checkin_interval": 300,
						"grace_period": 60
					}`,
		},
		{
			Name:           "Statping Checkins",
			URL:            "/api/checkins",
			Method:         "GET",
			ExpectedStatus: 200,
			ResponseLen:    3,
			BeforeTest:     SetTestENV,
			ResponseFunc: func(t *testing.T, resp []byte) error {
				var checkin []*checkins.Checkin
				if err := json.Unmarshal(resp, &checkin); err != nil {
					return err
				}
				require.Len(t, checkin, 3)
				last := checkin[len(checkin)-1]
				apiToken = last.ApiKey
				require.NotEmpty(t, apiToken)
				return nil
			},
		},
		{
			Name:             "Statping View Checkin",
			URL:              "/api/checkins/" + apiToken,
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{Success, `"type":"checkin","method":"create"`},
			BeforeTest:       SetTestENV,
			SecureRoute:      true,
			Skip:             true,
		},
		{
			Name:             "Statping Trigger Checkin",
			URL:              "/checkin/" + apiToken,
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{Success, `"type":"checkin","method":"create"`},
			SecureRoute:      true,
			BeforeTest:       SetTestENV,
			Skip:             true,
		},
		{
			Name:             "Statping Delete Checkin",
			URL:              "/api/checkins/" + apiToken,
			Method:           "DELETE",
			ExpectedContains: []string{Success, `"type":"checkin","method":"create"`},
			BeforeTest:       SetTestENV,
			Skip:             true,
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
