package handlers

import (
	"github.com/statping/statping/types/core"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
)

func googleOAuth(r *http.Request) (*oAuth, error) {
	c := core.App
	code := r.URL.Query().Get("code")

	config := &oauth2.Config{
		ClientID:     c.OAuth.GoogleClientID,
		ClientSecret: c.OAuth.GoogleClientSecret,
		Endpoint:     google.Endpoint,
	}

	gg, err := config.Exchange(r.Context(), code)
	if err != nil {
		return nil, err
	}

	return &oAuth{
		Token:        gg.AccessToken,
		RefreshToken: gg.RefreshToken,
		Valid:        gg.Valid(),
	}, nil
}
