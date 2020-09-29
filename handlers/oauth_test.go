package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOAuthRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name: "OAuth Save",
			URL:  "/api/oauth",
			Body: `{
						"gh_client_id": "githubid",
						"gh_client_secret": "githubsecret",
						"google_client_id": "googleid",
						"google_client_secret": "googlesecret",
						"oauth_domains": "gmail.com,yahoo.com,socialeck.com",
						"oauth_providers": "local,slack,google,github",
						"slack_client_id": "example.iddd",
						"slack_client_secret": "exampleeesecret",
						"slack_team": "dev"
					}`,
			Method:         "POST",
			ExpectedStatus: 200,
			BeforeTest:     SetTestENV,
		},
		{
			Name:             "OAuth Values",
			URL:              "/api/oauth",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{`"slack_client_id":"example.iddd"`},
			AfterTest:        UnsetTestENV,
		},
	}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			res, t, err := RunHTTPTest(v, t)
			assert.Nil(t, err)
			t.Log(res)
		})
	}
}
