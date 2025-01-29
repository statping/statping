package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/statping-ng/statping-ng/types/core"
	"github.com/statping-ng/statping-ng/types/errors"
	"github.com/statping-ng/statping-ng/types/null"
	"github.com/statping-ng/statping-ng/types/users"
	"golang.org/x/oauth2"
	"net/http"
	"strings"
)

type oAuth struct {
	Email    string
	Username string
	*oauth2.Token
	Groups []string
}

func oauthHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provider := vars["provider"]

	var err error
	var oauth *oAuth
	switch provider {
	case "google":
		oauth, err = googleOAuth(r)
	case "github":
		oauth, err = githubOAuth(r)
	case "slack":
		oauth, err = slackOAuth(r)
	case "custom":
		oauth, err = customOAuth(r)
	case "keycloak":
		oauth, err = keycloakOAuth(r)
	default:
		err = errors.New("unknown oauth provider")
	}

	if err != nil {
		log.Error(err)
		sendErrorJson(err, w, r)
		return
	}

	oauthLogin(oauth, w, r)
}

func oauthLogin(oauth *oAuth, w http.ResponseWriter, r *http.Request) {
	user := &users.User{
		Id:       0,
		Username: oauth.Username,
		Email:    oauth.Email,
		Admin:    null.NewNullBool(false),
	}

	// Check if the user is in the Keycloak admin groups
	if oauth.Groups != nil && core.App.OAuth.KeycloakAdminGroups != "" {
		adminGroups := strings.Split(core.App.OAuth.KeycloakAdminGroups, ",")
		for _, group := range adminGroups {
			if contains(oauth.Groups, group) {
				user.Admin = null.NewNullBool(true)
				break
			}
		}
	}
	log.Infoln(fmt.Sprintf("OAuth %s User %s logged in from IP %s", oauth.Type(), oauth.Email, r.RemoteAddr))
	setJwtToken(user, w)

	http.Redirect(w, r, core.App.Domain+"/dashboard", http.StatusPermanentRedirect)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
