package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/types/users"
	"github.com/statping/statping/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/slack"
	"net/http"
	"time"
)

type oAuth struct {
	ID           string
	Email        string
	Username     string
	Token        string
	RefreshToken string
	Valid        bool
	Type         string
}

func oauthHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provider := vars["provider"]
	code := r.URL.Query().Get("code")
	fmt.Println("code: ", code)
	fmt.Println("client: ", core.App.OAuth.SlackClientID)
	fmt.Println("secret: ", core.App.OAuth.SlackClientSecret)

	var err error
	var oauth *oAuth
	switch provider {
	case "google":
		oauth, err = googleOAuth(r)
	case "github":
		oauth, err = githubOAuth(r)
	case "slack":
		oauth, err = slackOAuth(r)
	}

	if err != nil {
		log.Error(err)
		sendErrorJson(err, w, r)
		return
	}

	oauthLogin(oauth, w, r)
}

func oauthLogin(oauth *oAuth, w http.ResponseWriter, r *http.Request) {
	log.Infoln(oauth)
	user := &users.User{
		Id:       0,
		Username: oauth.Username,
		Email:    oauth.Email,
		Admin:    null.NewNullBool(true),
	}
	log.Infoln(fmt.Sprintf("OAuth User %s logged in from IP %s", oauth.Email, r.RemoteAddr))
	setJwtToken(user, w)

	//returnJson(user, w, r)
	http.Redirect(w, r, core.App.Domain, http.StatusPermanentRedirect)
}

func githubOAuth(r *http.Request) (*oAuth, error) {
	c := *core.App
	code := r.URL.Query().Get("code")

	config := &oauth2.Config{
		ClientID:     c.OAuth.GithubClientID,
		ClientSecret: c.OAuth.GithubClientSecret,
		Endpoint:     github.Endpoint,
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

	return oauther.slackIdentity()
}

// slackIdentity will query the Slack API to fetch the users ID, username, and email address.
func (a *oAuth) slackIdentity() (*oAuth, error) {
	url := fmt.Sprintf("https://slack.com/api/users.identity?token=%s", a.Token)
	out, resp, err := utils.HttpRequest(url, "GET", "application/x-www-form-urlencoded", nil, nil, 10*time.Second, true, nil)
	if err != nil {
		return a, err
	}
	defer resp.Body.Close()

	var i *slackIdentity
	if err := json.Unmarshal(out, &i); err != nil {
		return a, err
	}
	a.Email = i.User.Email
	a.ID = i.User.ID
	a.Username = i.User.Name
	return a, nil
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

func secureToken(w http.ResponseWriter, r *http.Request) {

}
