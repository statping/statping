package handlers

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/source"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"html/template"
	"net/http"
	"os"
	"time"
)

const (
	COOKIE_KEY = "statup_auth"
)

var (
	Store *sessions.CookieStore
)

func RunHTTPServer(ip string, port int) error {
	host := fmt.Sprintf("%v:%v", ip, port)
	utils.Log(1, "Statup HTTP Server running on http://"+host)
	for _, p := range core.CoreApp.AllPlugins {
		info := p.GetInfo()
		for _, route := range p.Routes() {
			path := fmt.Sprintf("%v", route.URL)
			r.Handle(path, http.HandlerFunc(route.Handler)).Methods(route.Method)
			utils.Log(1, fmt.Sprintf("Added Route %v for plugin %v\n", path, info.Name))
		}
	}
	srv := &http.Server{
		Addr:         host,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      Router(),
	}
	return srv.ListenAndServe()
}

func IsAuthenticated(r *http.Request) bool {
	if os.Getenv("GO_ENV") == "test" {
		return true
	}
	if core.CoreApp == nil {
		return false
	}
	if Store == nil {
		return false
	}
	session, err := Store.Get(r, COOKIE_KEY)
	if err != nil {
		return false
	}
	if session.Values["authenticated"] == nil {
		return false
	}
	return session.Values["authenticated"].(bool)
}

func ExecuteResponse(w http.ResponseWriter, r *http.Request, file string, data interface{}) {
	utils.Http(r)
	nav, _ := source.TmplBox.String("nav.html")
	footer, _ := source.TmplBox.String("footer.html")
	render, err := source.TmplBox.String(file)
	if err != nil {
		utils.Log(4, err)
	}
	t := template.New("message")
	t.Funcs(template.FuncMap{
		"js": func(html string) template.JS {
			return template.JS(html)
		},
		"safe": func(html string) template.HTML {
			return template.HTML(html)
		},
		"Auth": func() bool {
			return IsAuthenticated(r)
		},
		"VERSION": func() string {
			return core.VERSION
		},
		"CoreApp": func() *core.Core {
			return core.CoreApp
		},
		"USE_CDN": func() bool {
			return core.CoreApp.UseCdn
		},
		"underscore": func(html string) string {
			return utils.UnderScoreString(html)
		},
	})
	t, _ = t.Parse(nav)
	t, _ = t.Parse(footer)
	t.Parse(render)
	t.Execute(w, data)
}

func ExecuteJSResponse(w http.ResponseWriter, r *http.Request, file string, data interface{}) {
	render, err := source.JsBox.String(file)
	if err != nil {
		utils.Log(4, err)
	}
	t := template.New("charts")
	t.Funcs(template.FuncMap{
		"safe": func(html string) template.HTML {
			return template.HTML(html)
		},
	})
	t.Parse(render)
	t.Execute(w, data)
}

func Error404Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	ExecuteResponse(w, r, "error_404.html", nil)
}

type DbConfig types.DbConfig
