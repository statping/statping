package handlers

import (
	"github.com/statping/statping/source"
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnAuthenticatedThemeRoutes(t *testing.T) {
	t.SkipNow()
	tests := []HTTPTest{
		{
			Name:           "No Authentication - Create Themes",
			URL:            "/api/theme/create",
			Method:         "GET",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - View Themes",
			URL:            "/api/theme",
			Method:         "GET",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - Update Themes",
			URL:            "/api/theme",
			Method:         "POST",
			ExpectedStatus: 401,
			BeforeTest:     UnsetTestENV,
		},
		{
			Name:           "No Authentication - Delete Themes",
			URL:            "/api/theme",
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

func TestThemeRoutes(t *testing.T) {
	t.SkipNow()
	tests := []HTTPTest{
		{
			Name:             "Create Theme Assets",
			URL:              "/api/theme/create",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{`"status":"success"`},
			BeforeTest:       SetTestENV,
			AfterTest: func(t *testing.T) error {
				assert.True(t, source.UsingAssets(utils.Params.GetString("STATPING_DIR")))
				return nil
			},
		},
		{
			Name:           "Get Theme",
			URL:            "/api/theme",
			Method:         "GET",
			ExpectedStatus: 200,
			ExpectedContains: []string{
				`"base":"@import 'variables';`,
				`"mobile":"@import 'variables'`,
				`"variables":"/*`,
			},
			BeforeTest: SetTestENV,
		},
		{
			Name:   "Save Theme",
			URL:    "/api/theme",
			Method: "POST",
			Body: `{
					"base": ".base { }",
					"variables": "$variable: #bababa",
					"mobile": ".mobile { }"
				}`,
			ExpectedStatus: 200,
			BeforeTest:     SetTestENV,
		},
		{
			Name:           "Delete Theme",
			URL:            "/api/theme",
			Method:         "DELETE",
			ExpectedStatus: 200,
			BeforeTest:     SetTestENV,
			AfterTest: func(t *testing.T) error {
				assert.False(t, source.UsingAssets(utils.Directory))
				return nil
			},
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
