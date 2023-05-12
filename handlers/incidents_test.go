package handlers

import (
	"github.com/razorpay/statping/types/incidents"
	"github.com/razorpay/statping/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUnAuthenticatedIncidentRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "No Authentication - New Incident",
			URL:            "/api/services/1/incidents",
			Method:         "POST",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - New Incident Update",
			URL:            "/api/incidents/updates",
			Method:         "POST",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - Update Incident",
			URL:            "/api/incidents/1",
			Method:         "POST",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - Delete Incident",
			URL:            "/api/incidents/1",
			Method:         "DELETE",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - Delete Incident Update",
			URL:            "/api/incidents/1/updates/1",
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

func TestIncidentsAPIRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "Statping Create Incident",
			URL:            "/api/services/1/incidents",
			Method:         "POST",
			ExpectedStatus: 200,
			BeforeTest:     SetTestENV,
			AfterTest:      UnsetTestENV,
			Body: `{
					"title": "New Incident",
					"description": "This is a test for incidents"
				    }`,
			ExpectedContains: []string{Success},
		},
		{
			Name:             "Statping Service 1 Incidents",
			URL:              "/api/services/1/incidents",
			Method:           "GET",
			ExpectedStatus:   200,
			ResponseLen:      1,
			BeforeTest:       SetTestENV,
			AfterTest:        UnsetTestENV,
			ExpectedContains: []string{`"title":"New Incident"`},
		},
		{
			Name: "Statping Update Incident",
			URL:  "/api/incidents/1",
			Body: `{
					"title": "Updated Incident",
					"description": "This is an updated incidents"
				    }`,
			Method:           "POST",
			ExpectedStatus:   200,
			BeforeTest:       SetTestENV,
			ExpectedContains: []string{Success},
		},
		{
			Name:             "Statping View Incident Updates",
			URL:              "/api/incidents/1/updates",
			Method:           "GET",
			ExpectedStatus:   200,
			ResponseLen:      3,
			BeforeTest:       SetTestENV,
			ExpectedContains: []string{`"type":"investigating"`},
		},
		{
			Name:   "Statping Create Incident Update",
			URL:    "/api/incidents/1/updates",
			Method: "POST",
			Body: `{
								"message": "Test message here",
								"type": "Update"
								}`,
			ExpectedStatus:   200,
			BeforeTest:       SetTestENV,
			ExpectedContains: []string{Success},
		},
		{
			Name:             "Incorrect Checkin JSON POST",
			URL:              "/api/incidents/1/updates",
			Body:             BadJSON,
			ExpectedContains: []string{BadJSONResponse},
			BeforeTest:       SetTestENV,
			Method:           "POST",
			ExpectedStatus:   422,
		},
		{
			Name:             "Statping Delete Incident Update",
			URL:              "/api/incidents/1/updates/1",
			Method:           "DELETE",
			ExpectedStatus:   200,
			BeforeTest:       SetTestENV,
			ExpectedContains: []string{Success},
		},
		{
			Name:             "Statping Delete Incident",
			URL:              "/api/incidents/1",
			Method:           "DELETE",
			ExpectedStatus:   200,
			BeforeTest:       SetTestENV,
			ExpectedContains: []string{Success},
		},
		{
			Name:             "Incorrect JSON POST",
			URL:              "/api/services/1/incidents",
			Body:             BadJSON,
			ExpectedContains: []string{BadJSONResponse},
			Method:           "POST",
			ExpectedStatus:   422,
		},
	}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			assert.Nil(t, err)
		})
	}
}

func TestIncidentUpdateOrder(t *testing.T) {
	now := utils.Now()
	t0 := now.Add(-5 * time.Minute)
	t1 := now.Add(-5*time.Minute + 1*time.Second)
	t2 := now.Add(-5*time.Minute + 2*time.Second)

	updates := []*incidents.IncidentUpdate{
		{
			Id:         78,
			IncidentId: 23,
			Message:    "Update 2",
			Type:       "update",
			CreatedAt:  t1,
			UpdatedAt:  t2,
		},
		{
			Id:         79,
			IncidentId: 23,
			Message:    "Resolved 1",
			Type:       "resolved",
			CreatedAt:  t2,
			UpdatedAt:  t2,
		},
		{
			Id:         76,
			IncidentId: 23,
			Message:    "Update 1",
			Type:       "update",
			CreatedAt:  t0,
			UpdatedAt:  t2,
		},
	}
	sortUpdates(updates, "DESC")
	assert.Equal(t, int64(79), updates[0].Id)
	assert.Equal(t, int64(78), updates[1].Id)
	assert.Equal(t, int64(76), updates[2].Id)
}
