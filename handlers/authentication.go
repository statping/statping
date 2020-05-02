package handlers

import (
	"crypto/subtle"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/utils"
	"net/http"
	"strings"
)

// hasSetupEnv checks to see if the GO_ENV is set to 'true'
// or if the Statping instance has not been setup yet
func hasSetupEnv() bool {
	if utils.Params.Get("GO_ENV") == "test" {
		return true
	}
	if core.App == nil {
		return true
	}
	if !core.App.Setup {
		return false
	}
	return false
}

// hasAPIQuery checks the `api` query parameter against the API Secret Key
func hasAPIQuery(r *http.Request) bool {
	query := r.URL.Query()
	key := query.Get("api")
	if key == "" {
		return false
	}
	if subtle.ConstantTimeCompare([]byte(key), []byte(core.App.ApiSecret)) == 1 {
		return true
	}
	return false
}

// hasAuthorizationHeader check to see if the Authorization header is the correct API Secret Key
func hasAuthorizationHeader(r *http.Request) bool {
	var token string
	tokens, ok := r.Header["Authorization"]
	if ok && len(tokens) >= 1 {
		token = tokens[0]
		token = strings.TrimPrefix(token, "Bearer ")
		if subtle.ConstantTimeCompare([]byte(token), []byte(core.App.ApiSecret)) == 1 {
			return true
		}
	}
	return false
}
