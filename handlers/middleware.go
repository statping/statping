package handlers

import (
	"compress/gzip"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/errors"
	"github.com/statping/statping/utils"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

var (
	authUser string
	authPass string
)

// Gzip Compression
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func Gzip(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			handler.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzw := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		handler.ServeHTTP(gzw, r)
	})
}

// basicAuthHandler is a middleware to implement HTTP basic authentication using
// AUTH_USERNAME and AUTH_PASSWORD environment variables
func basicAuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user),
			[]byte(authUser)) != 1 || subtle.ConstantTimeCompare([]byte(pass),
			[]byte(authPass)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="statping"`)
			w.WriteHeader(401)
			w.Write([]byte("You are unauthorized to access the application.\n"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

// apiMiddleware will confirm if Core has been setup
func apiMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !core.App.Setup {
			sendErrorJson(errors.New("statping has not been setup"), w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// sendLog is a http middleware that will log the duration of request and other useful fields
func sendLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := utils.Now()
		if r.RequestURI == "/api/logs" || r.RequestURI == "/api/logs/last" {
			next.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
		t2 := utils.Now().Sub(t1)
		log.WithFields(utils.ToFields(w, r)).
			WithField("url", r.RequestURI).
			WithField("method", r.Method).
			WithField("load_micro_seconds", t2.Microseconds()).
			Infoln(fmt.Sprintf("%v (%v) | IP: %v", r.RequestURI, r.Method, r.Host))
	})
}

// scoped is a middleware handler that will remove private fields based on struct tags
// this will look for the `scope:"user,admin"` tag and remove the JSON field from response
// if user is not authenticated based on the scope.
func scoped(handler func(r *http.Request) interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := handler(r)
		err, ok := data.(error)
		if ok {
			sendErrorJson(err, w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(scope{data: data, scope: ScopeName(r)})
	})
}

// authenticated is a middleware function to check if user is an Admin before running original request
func authenticated(handler func(w http.ResponseWriter, r *http.Request), redirect bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !IsFullAuthenticated(r) {
			if redirect {
				http.Redirect(w, r, basePath, http.StatusSeeOther)
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
				http.Redirect(w, r, basePath, http.StatusSeeOther)
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
		if !core.App.Setup {
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

func DecodeJSON(r *http.Request, obj interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&obj)
	if err != nil {
		return errors.DecodeJSON
	}
	defer r.Body.Close()
	return nil
}
