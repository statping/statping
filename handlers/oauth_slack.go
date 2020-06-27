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
	"strings"
	"time"
)

func slackOAuth(r *http.Request) (*oAuth, error) {
	auth := core.App.OAuth
	code := r.URL.Query().Get("code")

	config := &oauth2.Config{
		ClientID:     auth.SlackClientID,
		ClientSecret: auth.SlackClientSecret,
		Endpoint:     slack.Endpoint,
		RedirectURL:  core.App.Domain + basePath + "oauth/slack",
		Scopes:       []string{"identity.basic"},
	}

	gg, err := config.Exchange(r.Context(), code)
	if err != nil {
		return nil, err
	}

	if !gg.Valid() {
		return nil, errors.New("oauth token is not valid")
	}

	identity, err := returnSlackIdentity(gg.AccessToken)
	if err != nil {
		return nil, err
	}

	if !identity.Ok {
		return nil, errors.New("slack identity is invalid")
	}

	if !validateSlack(identity) {
		return nil, errors.New("slack user is not whitelisted")
	}

	return &oAuth{
		Token:    gg,
		Username: strings.ToLower(identity.User.Name),
		Email:    strings.ToLower(identity.User.Email),
	}, nil
}

func validateSlack(id slackIdentity) bool {
	auth := core.App.OAuth
	if auth.SlackUsers == "" {
		return true
	}

	if auth.SlackUsers != "" {
		users := strings.Split(auth.SlackUsers, ",")
		for _, u := range users {
			if strings.ToLower(u) == strings.ToLower(id.User.Email) {
				return true
			}
			if strings.ToLower(u) == strings.ToLower(id.User.Name) {
				return true
			}
		}
	}

	return false
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
