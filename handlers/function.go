package handlers

import (
	"github.com/hunterlong/statping/core"
	"html/template"
	"net/http"
)

var (
	basePath = "/"
)

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
		"BasePath": func() string {
			return basePath
		},
	}
}
