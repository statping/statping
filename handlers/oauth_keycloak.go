package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/statping-ng/statping-ng/types/core"
	"github.com/statping-ng/statping-ng/types/errors"
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
	log.Infof("Keycloak code: %s", code)
	if code == "" {
		return nil, errors.New("code not found")
	}
	scopes := strings.Split(auth.KeycloakScopes, ",")

	conf := &oauth2.Config{
		ClientID:     core.App.OAuth.KeycloakClientID,
		ClientSecret: core.App.OAuth.KeycloakClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  auth.KeycloakEndpointAuth,
			TokenURL: auth.KeycloakEndpointToken,
		},
		RedirectURL: core.App.Domain + basePath + "oauth/keycloak",
		Scopes:      scopes,
	}
	log.Infof("Keycloak config: %+v", conf)

	token, err := conf.Exchange(context.Background(), code)
	log.Infof("Keycloak token: %+v", token)
	if err != nil {
		log.Error("Failed to exchange token: ", err)
		return nil, err
	}

	client := conf.Client(context.Background(), token)
	log.Infof("Keycloak client: %+v", client)
	resp, err := client.Get(core.App.OAuth.KeycloakEndpointUserinfo)
	log.Infof("Keycloak user info URL: %s", core.App.OAuth.KeycloakEndpointUserinfo)
	if err != nil {
		log.Error("Failed to get user info: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo struct {
		Email    string   `json:"email"`
		Username string   `json:"preferred_username"`
		Groups   []string `json:"groups"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		log.Error("Failed to decode user info: ", err)
		return nil, err
	}

	log.Infof("Keycloak user info: %+v", userInfo)
	log.Infof("Keycloak user groups: %+v", userInfo.Groups)
	return &oAuth{
		Email:    userInfo.Email,
		Username: userInfo.Username,
		Token:    token,
		Groups:   userInfo.Groups,
	}, nil
}