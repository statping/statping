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
		Admin:    null.NewNullBool(true),
	}
	log.Infof("OAuth User: %+v", user)
	log.Infof("OAuth User Groups: %+v", oauth.Groups)
	log.Infof("core.App.OAuth.KeycloakAdminGroups: %+v", core.App.OAuth.KeycloakAdminGroups)

	isAdmin := false

	// Check if the user is in the Keycloak admin groups
	if oauth.Groups != nil && core.App.OAuth.KeycloakAdminGroups != "" {
		adminGroups := strings.Split(core.App.OAuth.KeycloakAdminGroups, ",")
		log.Infof("Admin Groups: %+v", adminGroups)
	
		for _, keycloakAdminGroup := range adminGroups {
			for _, userGroup := range oauth.Groups {
				log.Infof("Checking if user group '%s' is in admin group '%s'", userGroup, keycloakAdminGroup)
				if userGroup == keycloakAdminGroup {
					isAdmin = true // Set the flag to true if a group match is found
					break
				}
			}
			if user.Admin.Valid && user.Admin.Bool {
				break
			}
		}
	}
	user.Admin = null.NewNullBool(isAdmin)
	log.Infof("OAuth User Admin: %+v", user.Admin)
	log.Infoln(fmt.Sprintf("OAuth %s User %s logged in from IP %s", oauth.Type(), oauth.Email, r.RemoteAddr))
	setJwtToken(user, w)

	http.Redirect(w, r, core.App.Domain+"/dashboard", http.StatusPermanentRedirect)
}


