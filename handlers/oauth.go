package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/types/users"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"net/http"
)

type oAuth struct {
	Email        string
	Token        string
	RefreshToken string
	Valid        bool
}

func oauthHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provider := vars["provider"]

	var err error
	var oauth *oAuth
	switch provider {
	case "google":
		err, oauth = googleOAuth(r)
	case "github":
		err, oauth = githubOAuth(r)
	}

	if err != nil {
		log.Error(err)
		return
	}

	oauthLogin(oauth, w, r)
}

func oauthLogin(oauth *oAuth, w http.ResponseWriter, r *http.Request) {
	user := &users.User{
		Id:       0,
		Username: oauth.Email,
		Email:    oauth.Email,
		Admin:    null.NewNullBool(true),
	}
	log.Infoln(fmt.Sprintf("OAuth User %v logged in from IP %v", oauth.Email, r.RemoteAddr))
	setJwtToken(user, w)

	http.Redirect(w, r, basePath+"dashboard", http.StatusSeeOther)
}

func githubOAuth(r *http.Request) (error, *oAuth) {
	c := *core.App
	code := r.URL.Query().Get("code")

	config := &oauth2.Config{
		ClientID:     c.OAuth.GithubClientID,
		ClientSecret: c.OAuth.GithubClientSecret,
		Endpoint:     github.Endpoint,
	}

	gg, err := config.Exchange(r.Context(), code)
	if err != nil {
		return err, nil
	}

	return nil, &oAuth{
		Token:        gg.AccessToken,
		RefreshToken: gg.RefreshToken,
		Valid:        gg.Valid(),
	}
}

func googleOAuth(r *http.Request) (error, *oAuth) {
	c := *core.App
	code := r.URL.Query().Get("code")

	config := &oauth2.Config{
		ClientID:     c.OAuth.GithubClientID,
		ClientSecret: c.OAuth.GithubClientSecret,
		Endpoint:     google.Endpoint,
	}

	gg, err := config.Exchange(r.Context(), code)
	if err != nil {
		return err, nil
	}

	return nil, &oAuth{
		Token:        gg.AccessToken,
		RefreshToken: gg.RefreshToken,
		Valid:        gg.Valid(),
	}
}
