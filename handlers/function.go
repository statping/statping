package handlers

import (
	"bytes"
	"github.com/hunterlong/statping/core"
	"html/template"
	"net/http"
	"net/url"
)

var (
	basePath = "/"
)

type HandlerFunc func(Responder, *Request)

func (f HandlerFunc) ServeHTTP(w Responder, r *Request) {
	f(w, r)
}

type Handler interface {
	ServeHTTP(Responder, *Request)
}

type Responder struct {
	Code      int           // the HTTP response code from WriteHeader
	HeaderMap http.Header   // the HTTP response headers
	Body      *bytes.Buffer // if non-nil, the bytes.Buffer to append written data to
	Flushed   bool
}

type Request struct {
	*http.Request
}

func (r Responder) Header() http.Header {
	return r.HeaderMap
}

func (r Responder) Write(p []byte) (int, error) {
	return r.Body.Write(p)
}

func (r Responder) WriteHeader(statusCode int) {
	r.Code = statusCode
}

func parseForm(r *http.Request) url.Values {
	r.ParseForm()
	return r.PostForm
}

func parseGet(r *http.Request) url.Values {
	r.ParseForm()
	return r.Form
}

var handlerFuncs = func(w http.ResponseWriter, r *http.Request) template.FuncMap {
	return template.FuncMap{
		"VERSION": func() string {
			return core.VERSION
		},
		"CoreApp": func() core.Core {
			c := *core.CoreApp
			if c.Name == "" {
				c.Name = "Statping"
			}
			return c
		},
		"USE_CDN": func() bool {
			return core.CoreApp.UseCdn.Bool
		},
		"USING_ASSETS": func() bool {
			return core.CoreApp.UsingAssets()
		},
		"BasePath": func() string {
			return basePath
		},
	}
}
