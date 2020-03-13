package handlers

import (
	"github.com/pkg/errors"
	"github.com/statping/statping/types/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApiServiceRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:             "Statping All Public and Private Services",
			URL:              "/api/services",
			Method:           "GET",
			ExpectedContains: []string{`"name":"Google"`},
			ExpectedStatus:   200,
			ResponseLen:      5,
			BeforeTest:       SetTestENV,
			FuncTest: func() error {
				count := len(services.Services())
				if count != 5 {
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
			FuncTest: func() error {
				count := len(services.Services())
				if count != 5 {
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
		},
		{
			Name:             "Statping Private Service 1",
			URL:              "/api/services/1",
			Method:           "GET",
			ExpectedContains: []string{`"name":"Google"`},
			ExpectedStatus:   200,
			BeforeTest:       SetTestENV,
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
			ResponseLen:    30,
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Reorder Services",
			URL:            "/api/services/reorder",
			Method:         "POST",
			Body:           `[{"service":1,"order":1},{"service":4,"order":2},{"service":2,"order":3},{"service":3,"order":4}]`,
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
			assert.Nil(t, err)
		})
	}
}
