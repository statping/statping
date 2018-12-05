// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/source"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"html/template"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

const (
	cookieKey = "statping_auth"
)

var (
	sessionStore *sessions.CookieStore
	httpServer   *http.Server
)

// RunHTTPServer will start a HTTP server on a specific IP and port
func RunHTTPServer(ip string, port int) error {
	host := fmt.Sprintf("%v:%v", ip, port)
	utils.Log(1, "Statping HTTP Server running on http://"+host)
	//for _, p := range core.CoreApp.AllPlugins {
	//	info := p.GetInfo()
	//	for _, route := range p.Routes() {
	//		path := fmt.Sprintf("%v", route.URL)
	//		router.Handle(path, http.HandlerFunc(route.Handler)).Methods(route.Method)
	//		utils.Log(1, fmt.Sprintf("Added Route %v for plugin %v\n", path, info.Name))
	//	}
	//}
	router = Router()
	httpServer = &http.Server{
		Addr:         host,
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 60,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
	resetCookies()
	return httpServer.ListenAndServe()
}

// IsAuthenticated returns true if the HTTP request is authenticated. You can set the environment variable GO_ENV=test
// to bypass the admin authenticate to the dashboard features.
func IsAuthenticated(r *http.Request) bool {
	if os.Getenv("GO_ENV") == "test" {
		return true
	}
	if core.CoreApp == nil {
		return true
	}
	if sessionStore == nil {
		return true
	}
	var token string
	tokens, ok := r.Header["Authorization"]
	if ok && len(tokens) >= 1 {
		token = tokens[0]
		token = strings.TrimPrefix(token, "Bearer ")
		if token == core.CoreApp.ApiSecret {
			return true
		}
	}
	session, err := sessionStore.Get(r, cookieKey)
	if err != nil {
		return false
	}
	if session.Values["authenticated"] == nil {
		return false
	}
	return session.Values["authenticated"].(bool)
}

var handlerFuncs = func(w http.ResponseWriter, r *http.Request) template.FuncMap {
	return template.FuncMap{
		"js": func(html interface{}) template.JS {
			return template.JS(utils.ToString(html))
		},
		"safe": func(html string) template.HTML {
			return template.HTML(html)
		},
		"safeURL": func(u string) template.URL {
			return template.URL(u)
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
		"Services": func() []types.ServiceInterface {
			return core.CoreApp.Services
		},
		"len": func(g []types.ServiceInterface) int {
			return len(g)
		},
		"IsNil": func(g interface{}) bool {
			return g == nil
		},
		"USE_CDN": func() bool {
			return core.CoreApp.UseCdn.Bool
		},
		"QrAuth": func() string {
			return fmt.Sprintf("statping://setup?domain=%v&api=%v", core.CoreApp.Domain, core.CoreApp.ApiSecret)
		},
		"Type": func(g interface{}) []string {
			fooType := reflect.TypeOf(g)
			var methods []string
			methods = append(methods, fooType.String())
			for i := 0; i < fooType.NumMethod(); i++ {
				method := fooType.Method(i)
				fmt.Println(method.Name)
				methods = append(methods, method.Name)
			}
			return methods
		},
		"ToJSON": func(g interface{}) template.HTML {
			data, _ := json.Marshal(g)
			return template.HTML(string(data))
		},
		"underscore": func(html string) string {
			return utils.UnderScoreString(html)
		},
		"URL": func() string {
			return r.URL.String()
		},
		"CHART_DATA": func() string {
			return ""
		},
		"Error": func() string {
			return ""
		},
		"ToString": func(v interface{}) string {
			return utils.ToString(v)
		},
		"Ago": func(t time.Time) string {
			return utils.Timestamp(t).Ago()
		},
		"Duration": func(t time.Duration) string {
			duration, _ := time.ParseDuration(fmt.Sprintf("%vs", t.Seconds()))
			return utils.FormatDuration(duration)
		},
		"ToUnix": func(t time.Time) int64 {
			return t.UTC().Unix()
		},
		"FromUnix": func(t int64) string {
			return utils.Timezoner(time.Unix(t, 0), core.CoreApp.Timezone).Format("Monday, January 02")
		},
		"NewService": func() *types.Service {
			return new(types.Service)
		},
		"NewUser": func() *types.User {
			return new(types.User)
		},
		"NewCheckin": func() *types.Checkin {
			return new(types.Checkin)
		},
		"NewMessage": func() *types.Message {
			return new(types.Message)
		},
	}
}

var mainTmpl = `{{define "main" }} {{ template "base" . }} {{ end }}`

// ExecuteResponse will render a HTTP response for the front end user
func ExecuteResponse(w http.ResponseWriter, r *http.Request, file string, data interface{}, redirect interface{}) {
	utils.Http(r)
	if url, ok := redirect.(string); ok {
		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	}

	templates := []string{"base.html", "head.html", "nav.html", "footer.html", "scripts.html", "form_service.html", "form_notifier.html", "form_user.html", "form_checkin.html", "form_message.html"}

	javascripts := []string{"charts.js", "chart_index.js"}

	render, err := source.TmplBox.String(file)
	if err != nil {
		utils.Log(4, err)
	}

	// setup the main template and handler funcs
	t := template.New("main")
	t.Funcs(handlerFuncs(w, r))
	t, err = t.Parse(mainTmpl)
	if err != nil {
		utils.Log(4, err)
	}

	// render all templates
	for _, temp := range templates {
		tmp, _ := source.TmplBox.String(temp)
		t, err = t.Parse(tmp)
		if err != nil {
			utils.Log(4, err)
		}
	}

	// render all javascript files
	for _, temp := range javascripts {
		tmp, _ := source.JsBox.String(temp)
		t, err = t.Parse(tmp)
		if err != nil {
			utils.Log(4, err)
		}
	}

	// render the page requested
	_, err = t.Parse(render)
	if err != nil {
		utils.Log(4, err)
	}

	// execute the template
	err = t.Execute(w, data)
	if err != nil {
		utils.Log(4, err)
	}
}

// executeJSResponse will render a Javascript response
func executeJSResponse(w http.ResponseWriter, r *http.Request, file string, data interface{}) {
	render, err := source.JsBox.String(file)
	if err != nil {
		utils.Log(4, err)
	}
	t := template.New("charts")
	t.Funcs(template.FuncMap{
		"safe": func(html string) template.HTML {
			return template.HTML(html)
		},
		"Services": func() []types.ServiceInterface {
			return core.CoreApp.Services
		},
	})
	_, err = t.Parse(render)
	if err != nil {
		utils.Log(4, err)
	}

	err = t.Execute(w, data)
	if err != nil {
		utils.Log(4, err)
	}
}

// error404Handler is a HTTP handler for 404 error pages
func error404Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	ExecuteResponse(w, r, "error_404.html", nil, nil)
}
