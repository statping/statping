package handlers

import (
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/errors"
	"golang.org/x/oauth2"
	"net/http"
	"strings"
)

func customOAuth(r *http.Request) (*oAuth, error) {
	auth := core.App.OAuth
	code := r.URL.Query().Get("code")

	scopes := strings.Split(auth.CustomScopes, ",")

	config := &oauth2.Config{
		ClientID:     auth.CustomClientID,
		ClientSecret: auth.CustomClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  auth.CustomEndpointAuth,
			TokenURL: auth.CustomEndpointToken,
		},
		RedirectURL: core.App.Domain + basePath + "oauth/custom",
		Scopes:      scopes,
	}

	gg, err := config.Exchange(r.Context(), code)
	if err != nil {
		return nil, err
	}

	if !gg.Valid() {
		return nil, errors.New("oauth token is not valid")
	}

	return &oAuth{
		Token: gg,
	}, nil
}
