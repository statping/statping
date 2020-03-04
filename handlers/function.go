package handlers

import (
	"encoding/json"
	"github.com/hunterlong/statping/source"
	"github.com/hunterlong/statping/types/core"
	"github.com/hunterlong/statping/utils"
	"html/template"
	"net/http"
	"net/url"
)

var (
	basePath = "/"
)

type CustomResponseWriter struct {
	body       []byte
	statusCode int
	header     http.Header
}

func NewCustomResponseWriter() *CustomResponseWriter {
	return &CustomResponseWriter{
		header: http.Header{},
	}
}

func (w *CustomResponseWriter) Header() http.Header {
	return w.header
}

func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.body = b
	// implement it as per your requirement
	return 0, nil
}

func (w *CustomResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func parseForm(r *http.Request) url.Values {
	r.ParseForm()
	return r.PostForm
}

func parseGet(r *http.Request) url.Values {
	r.ParseForm()
	return r.Form
}

func decodeRequest(r *http.Request, object interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	return decoder.Decode(&object)
}

type parsedObject struct {
	Error Error
}

func serviceFromID(r *http.Request, object interface{}) error {
	return nil
}

var handlerFuncs = func(w http.ResponseWriter, r *http.Request) template.FuncMap {
	return template.FuncMap{
		"VERSION": func() string {
			return core.App.Version
		},
		"CoreApp": func() core.Core {
			c := *core.App
			if c.Name == "" {
				c.Name = "Statping"
			}
			return c
		},
		"USE_CDN": func() bool {
			return core.App.UseCdn.Bool
		},
		"USING_ASSETS": func() bool {
			return source.UsingAssets(utils.Directory)
		},
		"BasePath": func() string {
			return basePath
		},
	}
}
