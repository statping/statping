package handlers

import (
	"github.com/hunterlong/statping/types/services"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestApiServiceRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:             "Statping All Services",
			URL:              "/api/services",
			Method:           "GET",
			ExpectedContains: []string{`"name":"Google"`},
			ExpectedStatus:   200,
			ResponseLen:      5,
			FuncTest: func() error {
				count := len(services.Services())
				if count != 5 {
					return errors.Errorf("incorrect services count: %d", count)
				}
				return nil
			},
		},
		{
			Name:             "Statping Service 1",
			URL:              "/api/services/1",
			Method:           "GET",
			ExpectedContains: []string{`"name":"Google"`},
			ExpectedStatus:   200,
		},
		{
			Name:           "Statping Service 1 Data",
			URL:            "/api/services/1/hits_data",
			Method:         "GET",
			Body:           "",
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Service 1 Ping Data",
			URL:            "/api/services/1/ping_data",
			Method:         "GET",
			Body:           "",
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Service 1 Failure Data",
			URL:            "/api/services/1/failure_data",
			Method:         "GET",
			Body:           "",
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Service 1 Hits",
			URL:            "/api/services/1/hits_data",
			Method:         "GET",
			Body:           "",
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Service 1 Failures",
			URL:            "/api/services/1/failure_data",
			Method:         "GET",
			Body:           "",
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Reorder Services",
			URL:            "/api/services/reorder",
			Method:         "POST",
			Body:           `[{"service":1,"order":1},{"service":5,"order":2},{"service":2,"order":3},{"service":3,"order":4},{"service":4,"order":5}]`,
			ExpectedStatus: 200,
			HttpHeaders:    []string{"Content-Type=application/json"},
		},
		{
			Name:        "Statping Create Service",
			URL:         "/api/services",
			HttpHeaders: []string{"Content-Type=application/json"},
			Method:      "POST",
			Body: `{
    "name": "New Service",
    "domain": "https://statping.com",
    "expected": "",
    "expected_status": 200,
    "check_interval": 30,
    "type": "http",
    "method": "GET",
    "post_data": "",
    "port": 0,
    "timeout": 30,
    "order_id": 0
}`,
			ExpectedStatus:   200,
			ExpectedContains: []string{`"status":"success","type":"service","method":"create"`},
			FuncTest: func() error {
				count := len(services.Services())
				if count != 6 {
					return errors.Errorf("incorrect services count: %d", count)
				}
				return nil
			},
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
			ExpectedContains: []string{`"status":"success"`, `"name":"Updated New Service"`, `"method":"update"`},
		},
		{
			Name:             "Statping Delete Service",
			URL:              "/api/services/1",
			Method:           "DELETE",
			ExpectedStatus:   200,
			ExpectedContains: []string{`"status":"success"`, `"method":"delete"`},
			FuncTest: func() error {
				count := len(services.Services())
				if count != 5 {
					return errors.Errorf("incorrect services count: %d", count)
				}
				return nil
			},
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			require.Nil(t, err)
		})
	}
}
