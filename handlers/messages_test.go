package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessageRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:             "Messages",
			URL:              "/messages",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{`<title>Statping Messages</title>`},
		},
		{
			Name:             "Message 2",
			URL:              "/message/2",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{`<title>Statping | Server Reboot</title>`},
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			assert.Nil(t, err)
			if err != nil {
				t.FailNow()
			}
		})
	}
}
