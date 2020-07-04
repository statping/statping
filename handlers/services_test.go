package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/statping/statping/types"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUnAuthenticatedServicesRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "No Authentication - New Service",
			URL:            "/api/services",
			Method:         "POST",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - Update Service",
			URL:            "/api/services/1",
			Method:         "POST",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - Delete Service",
			URL:            "/api/services/1",
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

func TestApiServiceRoutes(t *testing.T) {
	since := utils.Now().Add(-30 * types.Day)
	end := utils.Now().Add(-30 * time.Minute)
	startEndQuery := fmt.Sprintf("?start=%d&end=%d", since.Unix(), end.Unix()+15)

	tests := []HTTPTest{
		{
			Name:             "Statping All Public and Private Services",
			URL:              "/api/services",
			Method:           "GET",
			ExpectedContains: []string{`"name":"Google"`},
			ExpectedStatus:   200,
			ResponseLen:      6,
			BeforeTest:       SetTestENV,
			FuncTest: func(t *testing.T) error {
				count := len(services.Services())
				if count != 6 {
					return errors.Errorf("incorrect services count: %d", count)
				}
				return nil
			},
		},
		{
			Name:             "Statping All Public Services",
			URL:              "/api/services",
			Method:           "GET",
			ExpectedContains: []string{`"name":"Google"`},
			ExpectedStatus:   200,
			ResponseLen:      5,
			BeforeTest:       UnsetTestENV,
			FuncTest: func(t *testing.T) error {
				count := len(services.Services())
				if count != 6 {
					return errors.Errorf("incorrect services count: %d", count)
				}
				return nil
			},
		},
		{
			Name:             "Statping Public Service 1",
			URL:              "/api/services/1",
			Method:           "GET",
			ExpectedContains: []string{`"name":"Google"`},
			ExpectedStatus:   200,
			BeforeTest:       UnsetTestENV,
		},
		{
			Name:             "Statping Private Service 6",
			URL:              "/api/services/6",
			Method:           "GET",
			ExpectedContains: []string{`"error":"user not authenticated"`},
			ExpectedStatus:   401,
			BeforeTest:       UnsetTestENV,
		},
		{
			Name:           "Statping Authenticated Private Service 6",
			URL:            "/api/services/6",
			Method:         "GET",
			ExpectedStatus: 200,
			BeforeTest:     SetTestENV,
		},
		{
			Name:           "Statping Private Service with API Key",
			URL:            "/api/services/6?api=" + core.App.ApiSecret,
			Method:         "GET",
			ExpectedStatus: 200,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "Statping Private Service with API Header",
			URL:            "/api/services/6?api=" + core.App.ApiSecret,
			Method:         "GET",
			HttpHeaders:    []string{"Authorization=" + core.App.ApiSecret},
			ExpectedStatus: 200,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:             "Statping Service 1 with Private responses",
			URL:              "/api/services/1",
			Method:           "GET",
			ExpectedContains: []string{`"name":"Google"`},
			ExpectedStatus:   200,
			BeforeTest:       SetTestENV,
		},
		{
			Name:           "Statping Service Failures",
			URL:            "/api/services/1/failures" + startEndQuery,
			Method:         "GET",
			GreaterThan:    120,
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Service 1 Hits",
			URL:            "/api/services/1/hits" + startEndQuery,
			Method:         "GET",
			GreaterThan:    8580,
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Service 2 Hits",
			URL:            "/api/services/2/hits" + startEndQuery,
			Method:         "GET",
			GreaterThan:    8580,
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Service Failures Limited",
			URL:            "/api/services/1/failures?limit=1",
			Method:         "GET",
			ResponseLen:    1,
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Service 1 Hits Data",
			URL:            "/api/services/1/hits_data" + startEndQuery,
			Method:         "GET",
			GreaterThan:    70,
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Service 1 Ping Data",
			URL:            "/api/services/1/ping_data" + startEndQuery,
			Method:         "GET",
			ExpectedStatus: 200,
			GreaterThan:    70,
		},
		{
			Name:           "Statping Service 1 Failure Data - 24 Hour",
			URL:            "/api/services/1/failure_data" + startEndQuery + "&group=24h",
			Method:         "GET",
			ExpectedStatus: 200,
			GreaterThan:    3,
		},
		{
			Name:           "Statping Service 1 Failure Data - 12 Hour",
			URL:            "/api/services/1/failure_data" + startEndQuery + "&group=12h",
			Method:         "GET",
			ExpectedStatus: 200,
			GreaterThan:    6,
		},
		{
			Name:           "Statping Service 1 Failure Data - 1 Hour",
			URL:            "/api/services/1/failure_data" + startEndQuery + "&group=1h",
			Method:         "GET",
			ExpectedStatus: 200,
			GreaterThan:    70,
		},
		{
			Name:           "Statping Service 1 Failure Data - 15 Minute",
			URL:            "/api/services/1/failure_data" + startEndQuery + "&group=15m",
			Method:         "GET",
			GreaterThan:    120,
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Service 1 Hits",
			URL:            "/api/services/1/hits_data" + startEndQuery,
			Method:         "GET",
			GreaterThan:    70,
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Service 1 Uptime",
			URL:            "/api/services/1/uptime_data" + startEndQuery,
			Method:         "GET",
			ExpectedStatus: 200,
			ResponseFunc: func(req *httptest.ResponseRecorder, t *testing.T, resp []byte) error {
				var uptime *services.UptimeSeries
				if err := json.Unmarshal(resp, &uptime); err != nil {
					return err
				}
				assert.GreaterOrEqual(t, uptime.Uptime, int64(200000000))
				return nil
			},
		},
		{
			Name:           "Statping Service 1 Failure Data",
			URL:            "/api/services/1/failure_data" + startEndQuery,
			Method:         "GET",
			GreaterThan:    70,
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Reorder Services",
			URL:            "/api/reorder/services",
			Method:         "POST",
			Body:           `[{"service":1,"order":1},{"service":4,"order":2},{"service":2,"order":3},{"service":3,"order":4}]`,
			ExpectedStatus: 200,
			HttpHeaders:    []string{"Content-Type=application/json"},
			SecureRoute:    true,
		},
		{
			Name:        "Statping Create Service",
			URL:         "/api/services",
			HttpHeaders: []string{"Content-Type=application/json"},
			Method:      "POST",
			Body: `{
					"name": "New Private Service",
					"domain": "https://statping.com",
					"expected": "",
					"expected_status": 200,
					"check_interval": 30,
					"type": "http",
					"public": false,
					"group_id": 1,
					"method": "GET",
					"post_data": "",
					"port": 0,
					"timeout": 30,
					"order_id": 0
				}`,
			ExpectedStatus:   200,
			ExpectedContains: []string{Success, `"type":"service","method":"create"`, `"public":false`, `"group_id":1`},
			FuncTest: func(t *testing.T) error {
				count := len(services.Services())
				if count != 7 {
					return errors.Errorf("incorrect services count: %d", count)
				}
				return nil
			},
			SecureRoute: true,
		},
		{
			Name:        "Statping Update Service",
			URL:         "/api/services/1",
			HttpHeaders: []string{"Content-Type=application/json"},
			Method:      "POST",
			Body: `{
					"name": "Updated New Service",
					"domain": "https://google.com",
					"expected": "",
					"expected_status": 200,
					"check_interval": 60,
					"type": "http",
					"method": "GET",
					"post_data": "",
					"port": 0,
					"timeout": 10,
					"order_id": 0
				}`,
			ExpectedStatus:   200,
			ExpectedContains: []string{Success, `"name":"Updated New Service"`, MethodUpdate},
			FuncTest: func(t *testing.T) error {
				item, err := services.Find(1)
				require.Nil(t, err)
				if item.Interval != 60 {
					return errors.Errorf("incorrect service check interval: %d", item.Interval)
				}
				return nil
			},
			SecureRoute: true,
		},
		{
			Name:             "Statping Delete Failures",
			URL:              "/api/services/1/failures",
			Method:           "DELETE",
			ExpectedStatus:   200,
			ExpectedContains: []string{Success, `"method":"delete_failures"`},
			FuncTest: func(t *testing.T) error {
				item, err := services.Find(1)
				require.Nil(t, err)
				fails := item.AllFailures().Count()
				if fails != 0 {
					return errors.Errorf("incorrect service failures count: %d", fails)
				}
				return nil
			},
			SecureRoute: true,
		},
		{
			Name:             "Statping Delete Service",
			URL:              "/api/services/1",
			Method:           "DELETE",
			ExpectedStatus:   200,
			ExpectedContains: []string{Success, MethodDelete},
			FuncTest: func(t *testing.T) error {
				count := len(services.Services())
				if count != 6 {
					return errors.Errorf("incorrect services count: %d", count)
				}
				return nil
			},
			SecureRoute: true,
		},
		{
			Name:             "Incorrect JSON POST",
			URL:              "/api/services",
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
			assert.Nil(t, err)
		})
	}
}
