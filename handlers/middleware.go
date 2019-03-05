package handlers

import (
	"github.com/hunterlong/statping/core"
	"net/http"
	"net/http/httptest"
	"time"
)

// authenticated is a middleware function to check if user is an Admin before running original request
func authenticated(handler func(w http.ResponseWriter, r *http.Request), redirect bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !IsFullAuthenticated(r) {
			if redirect {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			} else {
				sendUnauthorizedJson(w, r)
			}
			return
		}
		handler(w, r)
	})
}

// readOnly is a middleware function to check if user is a User before running original request
func readOnly(handler func(w http.ResponseWriter, r *http.Request), redirect bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !IsReadAuthenticated(r) {
			if redirect {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			} else {
				sendUnauthorizedJson(w, r)
			}
			return
		}
		handler(w, r)
	})
}

// cached is a middleware function that accepts a duration and content type and will cache the response of the original request
func cached(duration, contentType string, handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content := CacheStorage.Get(r.RequestURI)
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if core.Configs == nil {
			handler(w, r)
			return
		}
		if content != nil {
			w.Write(content)
		} else {
			c := httptest.NewRecorder()
			handler(c, r)
			content := c.Body.Bytes()
			result := c.Result()
			if result.StatusCode != 200 {
				w.WriteHeader(result.StatusCode)
				w.Write(content)
				return
			}
			w.Write(content)
			if d, err := time.ParseDuration(duration); err == nil {
				go CacheStorage.Set(r.RequestURI, content, d)
			}
		}
	})
}
