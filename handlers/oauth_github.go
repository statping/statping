package handlers

import (
	"encoding/json"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/errors"
	"github.com/statping/statping/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"net/http"
	"strings"
	"time"
)

func githubOAuth(r *http.Request) (*oAuth, error) {
	auth := core.App.OAuth
	code := r.URL.Query().Get("code")

	config := &oauth2.Config{
		ClientID:     auth.GithubClientID,
		ClientSecret: auth.GithubClientSecret,
		Endpoint:     github.Endpoint,
		RedirectURL:  core.App.Domain + basePath + "oauth/github",
	}

	gg, err := config.Exchange(r.Context(), code)
	if err != nil {
		return nil, err
	}

	if !gg.Valid() {
		return nil, errors.New("oauth token is not valid")
	}

	user, err := returnGithubUser(gg.AccessToken)
	if err != nil {
		return nil, err
	}

	orgs, err := returnGithubOrganizations(gg.AccessToken, user.Login)
	if err != nil {
		return nil, err
	}

	if !validateGithub(user, orgs) {
		return nil, errors.New("github user is not allowed to login")
	}

	return &oAuth{
		Token:    gg,
		Username: strings.ToLower(user.Name),
		Email:    strings.ToLower(user.Email),
	}, nil
}

func returnGithubUser(token string) (githubUser, error) {
	headers := []string{
		"Accept=application/vnd.github.machine-man-preview+json",
		"Authorization=token " + token,
	}
	resp, _, err := utils.HttpRequest("https://api.github.com/user", "GET", nil, headers, nil, 10*time.Second, true, nil)
	if err != nil {
		return githubUser{}, err
	}
	var user githubUser
	if err := json.Unmarshal(resp, &user); err != nil {
		return githubUser{}, err
	}
	return user, nil
}

func returnGithubOrganizations(token, username string) ([]githubOrgs, error) {
	headers := []string{
		"Accept=application/vnd.github.machine-man-preview+json",
		"Authorization=token " + token,
	}
	resp, _, err := utils.HttpRequest("https://api.github.com/users/"+username+"/orgs", "GET", nil, headers, nil, 10*time.Second, true, nil)
	if err != nil {
		return nil, err
	}
	var orgs []githubOrgs
	if err := json.Unmarshal(resp, &orgs); err != nil {
		return nil, err
	}
	return orgs, nil
}

func validateGithub(ghUser githubUser, orgs []githubOrgs) bool {
	auth := core.App.OAuth
	if auth.GithubUsers == "" && auth.GithubOrgs == "" {
		return true
	}

	if auth.GithubUsers != "" {
		users := strings.Split(auth.GithubUsers, ",")
		for _, u := range users {
			if strings.ToLower(ghUser.Login) == strings.ToLower(u) {
				return true
			}
		}
	}
	if auth.GithubOrgs != "" {
		orgsAllowed := strings.Split(auth.GithubOrgs, ",")
		for _, o := range orgsAllowed {
			for _, org := range orgs {
				if strings.ToLower(o) == strings.ToLower(org.Login) {
					return true
				}
			}
		}
	}
	return false
}

type githubOrgs struct {
	Login            string `json:"login"`
	ID               int    `json:"id"`
	NodeID           string `json:"node_id"`
	URL              string `json:"url"`
	ReposURL         string `json:"repos_url"`
	EventsURL        string `json:"events_url"`
	HooksURL         string `json:"hooks_url"`
	IssuesURL        string `json:"issues_url"`
	MembersURL       string `json:"members_url"`
	PublicMembersURL string `json:"public_members_url"`
	AvatarURL        string `json:"avatar_url"`
	Description      string `json:"description"`
}

type githubUser struct {
	Login                   string    `json:"login"`
	ID                      int       `json:"id"`
	NodeID                  string    `json:"node_id"`
	AvatarURL               string    `json:"avatar_url"`
	GravatarID              string    `json:"gravatar_id"`
	URL                     string    `json:"url"`
	HTMLURL                 string    `json:"html_url"`
	FollowersURL            string    `json:"followers_url"`
	FollowingURL            string    `json:"following_url"`
	GistsURL                string    `json:"gists_url"`
	StarredURL              string    `json:"starred_url"`
	SubscriptionsURL        string    `json:"subscriptions_url"`
	OrganizationsURL        string    `json:"organizations_url"`
	ReposURL                string    `json:"repos_url"`
	EventsURL               string    `json:"events_url"`
	ReceivedEventsURL       string    `json:"received_events_url"`
	Type                    string    `json:"type"`
	SiteAdmin               bool      `json:"site_admin"`
	Name                    string    `json:"name"`
	Company                 string    `json:"company"`
	Blog                    string    `json:"blog"`
	Location                string    `json:"location"`
	Email                   string    `json:"email"`
	Hireable                bool      `json:"hireable"`
	Bio                     string    `json:"bio"`
	TwitterUsername         string    `json:"twitter_username"`
	PublicRepos             int       `json:"public_repos"`
	PublicGists             int       `json:"public_gists"`
	Followers               int       `json:"followers"`
	Following               int       `json:"following"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	PrivateGists            int       `json:"private_gists"`
	TotalPrivateRepos       int       `json:"total_private_repos"`
	OwnedPrivateRepos       int       `json:"owned_private_repos"`
	DiskUsage               int       `json:"disk_usage"`
	Collaborators           int       `json:"collaborators"`
	TwoFactorAuthentication bool      `json:"two_factor_authentication"`
	Plan                    struct {
		Name          string `json:"name"`
		Space         int    `json:"space"`
		PrivateRepos  int    `json:"private_repos"`
		Collaborators int    `json:"collaborators"`
	} `json:"plan"`
}
