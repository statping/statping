package handlers

import (
	"github.com/hunterlong/statping/core"
	"html/template"
	"net/http"
	"net/url"
)

var (
	basePath = "/"
)

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
