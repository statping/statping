package handlers

import (
	"github.com/stretchr/testify/assert"
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
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			body, t, err := RunHTTPTest(v, t)
			assert.Nil(t, err)
			if err != nil {
				t.FailNow()
			}
			t.Logf("Test %v got: %v\n", v.Name, string(body))
		})
	}
}
