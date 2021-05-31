package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/errors"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/types/users"
	"golang.org/x/oauth2"
	"net/http"
)

type oAuth struct {
	Email    string
	Username string
	*oauth2.Token
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
	log.Infoln(fmt.Sprintf("OAuth %s User %s logged in from IP %s", oauth.Type(), oauth.Email, r.RemoteAddr))
	setJwtToken(user, w)

	http.Redirect(w, r, core.App.Domain+"/dashboard", http.StatusPermanentRedirect)
}
