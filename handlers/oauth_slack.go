package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/errors"
	"github.com/statping/statping/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/slack"
	"net/http"
	"time"
)

func slackOAuth(r *http.Request) (*oAuth, error) {
	c := core.App
	code := r.URL.Query().Get("code")

	config := &oauth2.Config{
		ClientID:     c.OAuth.SlackClientID,
		ClientSecret: c.OAuth.SlackClientSecret,
		Endpoint:     slack.Endpoint,
		RedirectURL:  c.Domain + basePath + "oauth/slack",
		Scopes:       []string{"identity.basic"},
	}

	gg, err := config.Exchange(r.Context(), code)
	if err != nil {
		return nil, err
	}

	oauther := &oAuth{
		Token:        gg.AccessToken,
		RefreshToken: gg.RefreshToken,
		Valid:        gg.Valid(),
		Type:         gg.Type(),
	}

	identity, err := returnSlackIdentity(gg.AccessToken)
	if err != nil {
		return nil, err
	}

	if !identity.Ok {
		return nil, errors.New("slack identity is invalid")
	}

	oauther.Username = identity.User.Name
	oauther.Email = identity.User.Email

	return oauther, nil
}

// slackIdentity will query the Slack API to fetch the users ID, username, and email address.
func returnSlackIdentity(token string) (slackIdentity, error) {
	url := fmt.Sprintf("https://slack.com/api/users.identity?token=%s", token)
	out, _, err := utils.HttpRequest(url, "GET", "application/x-www-form-urlencoded", nil, nil, 10*time.Second, true, nil)
	if err != nil {
		return slackIdentity{}, err
	}

	var i slackIdentity
	if err := json.Unmarshal(out, &i); err != nil {
		return slackIdentity{}, err
	}
	return i, nil
}

type slackIdentity struct {
	Ok   bool `json:"ok"`
	User struct {
		Name  string `json:"name"`
		ID    string `json:"id"`
		Email string `json:"email"`
	} `json:"user"`
	Team struct {
		ID string `json:"id"`
	} `json:"team"`
}
