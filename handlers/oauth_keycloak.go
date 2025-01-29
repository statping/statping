package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/statping-ng/statping-ng/types/core"
	"golang.org/x/oauth2"
)

type keycloakUserInfo struct {
	Username string   `json:"preferred_username"`
	Email    string   `json:"email"`
	Groups   []string `json:"groups"`
}

func keycloakOAuth(r *http.Request) (*oAuth, error) {
	auth := core.App.OAuth
	code := r.URL.Query().Get("code")

	config := &oauth2.Config{
		ClientID:     auth.KeycloakClientID,
		ClientSecret: auth.KeycloakClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  auth.KeycloakAuthURL,
			TokenURL: auth.KeycloakTokenURL,
		},
		RedirectURL: core.App.Domain + basePath + "oauth/keycloak",
		Scopes:      strings.Split(auth.KeycloakScopes, ","),
	}

	token, err := config.Exchange(r.Context(), code)
	if err != nil {
		log.Errorln("Error exchanging token:", err)
		return nil, err
	}

	client := config.Client(r.Context(), token)
	userInfoResp, err := client.Get(auth.KeycloakUserInfoURL)
	if err != nil {
		log.Errorln("Error getting user info:", err)
		return nil, err
	}
	defer userInfoResp.Body.Close()

	var user keycloakUserInfo
	if err := json.NewDecoder(userInfoResp.Body).Decode(&user); err != nil {
		log.Errorln("Error decoding user info:", err)
		return nil, err
	}

	return &oAuth{
		Token:    token,
		Username: user.Username,
		Email:    user.Email,
		Groups:   user.Groups,
	}, nil
}