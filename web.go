package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	session *sessions.CookieStore
)

const (
	cookieKey = "apizer_auth"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/", http.HandlerFunc(IndexHandler))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(cssBox.HTTPBox())))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(jsBox.HTTPBox())))
	r.Handle("/setup", http.HandlerFunc(SetupHandler)).Methods("GET")
	r.Handle("/setup", http.HandlerFunc(ProcessSetupHandler)).Methods("POST")
	r.Handle("/dashboard", http.HandlerFunc(DashboardHandler)).Methods("GET")
	r.Handle("/dashboard", http.HandlerFunc(LoginHandler)).Methods("POST")
	r.Handle("/logout", http.HandlerFunc(LogoutHandler))
	r.Handle("/services", http.HandlerFunc(ServicesHandler)).Methods("GET")
	r.Handle("/services", http.HandlerFunc(CreateServiceHandler)).Methods("POST")
	r.Handle("/service/{id}", http.HandlerFunc(ServicesViewHandler))
	r.Handle("/service/{id}", http.HandlerFunc(ServicesUpdateHandler)).Methods("POST")
	r.Handle("/service/{id}/edit", http.HandlerFunc(ServicesViewHandler))
	r.Handle("/service/{id}/delete", http.HandlerFunc(ServicesDeleteHandler))
	r.Handle("/service/{id}/badge.svg", http.HandlerFunc(ServicesBadgeHandler))
	r.Handle("/users", http.HandlerFunc(UsersHandler)).Methods("GET")
	r.Handle("/users", http.HandlerFunc(CreateUserHandler)).Methods("POST")
	r.Handle("/settings", http.HandlerFunc(PluginsHandler))
	r.Handle("/plugins/download/{name}", http.HandlerFunc(PluginsDownloadHandler))
	r.Handle("/plugins/{name}/save", http.HandlerFunc(PluginSavedHandler)).Methods("POST")
	r.Handle("/help", http.HandlerFunc(HelpHandler))

	r.Handle("/api", http.HandlerFunc(ApiIndexHandler))
	r.Handle("/api/services", http.HandlerFunc(ApiAllServicesHandler))
	r.Handle("/api/services/{id}", http.HandlerFunc(ApiServiceHandler)).Methods("GET")
	r.Handle("/api/services/{id}", http.HandlerFunc(ApiServiceUpdateHandler)).Methods("POST")
	r.Handle("/api/users", http.HandlerFunc(ApiAllUsersHandler))
	r.Handle("/api/users/{id}", http.HandlerFunc(ApiUserHandler))
	return r
}

func RunHTTPServer() {
	fmt.Println("Statup HTTP Server running on http://localhost:8080")
	r := Router()
	for _, p := range allPlugins {
		info := p.GetInfo()
		for _, route := range p.Routes() {
			path := fmt.Sprintf("/plugins/%v/%v", info.Name, route.URL)
			r.Handle(path, http.HandlerFunc(route.Handler)).Methods(route.Method)
			fmt.Printf("Added Route %v for plugin %v\n", path, info.Name)
		}
	}

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	srv.ListenAndServe()
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieKey)
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieKey)
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	_, auth := AuthUser(username, password)
	if auth {
		session.Values["authenticated"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		w.WriteHeader(502)
		w.Header().Set("Content-Type", "plain/text")
		fmt.Fprintln(w, "bad")
	}
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating user")
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	user := &User{
		Username: username,
		Password: password,
	}
	_, err := user.Create()
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func CreateServiceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("service adding")
	r.ParseForm()
	name := r.PostForm.Get("name")
	domain := r.PostForm.Get("domain")
	method := r.PostForm.Get("method")
	expected := r.PostForm.Get("expected")
	status, _ := strconv.Atoi(r.PostForm.Get("expected_status"))
	interval, _ := strconv.Atoi(r.PostForm.Get("interval"))

	fmt.Println(r.PostForm)

	service := Service{
		Name:           name,
		Domain:         domain,
		Method:         method,
		Expected:       expected,
		ExpectedStatus: status,
		Interval:       interval,
	}

	fmt.Println(service)

	_, err := service.Create()
	if err != nil {
		go service.CheckQueue()
	}
	http.Redirect(w, r, "/services", http.StatusSeeOther)
}

func SetupHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := Parse("setup.html")
	tmpl.Execute(w, nil)
}

type index struct {
	Project  string
	Services []*Service
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if core == nil {
		http.Redirect(w, r, "/setup", http.StatusSeeOther)
		return
	}

	tmpl := Parse("index.html")
	out := index{core.Name, services}
	tmpl.Execute(w, out)
}

type dashboard struct {
	Services        []*Service
	Core            *Core
	CountOnline     int
	CountServices   int
	Count24Failures uint64
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieKey)
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		tmpl := Parse("login.html")
		tmpl.Execute(w, nil)
	} else {
		tmpl := Parse("dashboard.html")
		fails, _ := CountFailures()
		out := dashboard{services, core, CountOnline(), len(services), fails}
		tmpl.Execute(w, out)
	}

}

type serviceHandler struct {
	Service Service
	Auth    bool
}

func ServicesHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieKey)
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl := Parse("services.html")
	tmpl.Execute(w, services)
}

func ServicesDeleteHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieKey)
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	service, _ := SelectService(StringInt(vars["id"]))

	service.Delete()
	services, _ = SelectAllServices()
	http.Redirect(w, r, "/services", http.StatusSeeOther)
}

func IsAuthenticated(r *http.Request) bool {
	session, _ := store.Get(r, cookieKey)
	if session.Values["authenticated"] == nil {
		return false
	}
	return session.Values["authenticated"].(bool)
}

func PluginsHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl := ParsePlugins("plugins.html")
	core.FetchPluginRepo()

	var pluginFields []PluginSelect

	for _, p := range allPlugins {
		fields := SelectSettings(p.GetInfo())

		pluginFields = append(pluginFields, PluginSelect{p.GetInfo().Name, fields})
	}

	core.PluginFields = pluginFields

	fmt.Println(&core.PluginFields)

	tmpl.Execute(w, core)
}

type PluginSelect struct {
	Plugin string
	Params map[string]string
}

func PluginSavedHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	r.ParseForm()
	vars := mux.Vars(r)
	plug := SelectPlugin(vars["name"])
	data := make(map[string]string)
	for k, v := range r.PostForm {
		data[k] = strings.Join(v, "")
	}
	UpdateSettings(plug.GetInfo(), data)
	plug.OnSave(data)
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}

func PluginsDownloadHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	//vars := mux.Vars(r)
	//name := vars["name"]
	//DownloadPlugin(name)
	LoadConfig()
	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}

func HelpHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl := Parse("help.html")
	tmpl.Execute(w, nil)
}

func ServicesUpdateHandler(w http.ResponseWriter, r *http.Request) {
	//auth := IsAuthenticated(r)
	//
	//vars := mux.Vars(r)
	//service := SelectService(vars["id"])
}

func ServicesBadgeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service, _ := SelectService(StringInt(vars["id"]))

	var badge []byte
	if service.Online {
		badge = []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="104" height="20"><linearGradient id="b" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><mask id="a"><rect width="104" height="20" rx="3" fill="#fff"/></mask><g mask="url(#a)"><path fill="#555" d="M0 0h54v20H0z"/><path fill="#4c1" d="M54 0h50v20H54z"/><path fill="url(#b)" d="M0 0h104v20H0z"/></g><g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11"><text x="28" y="15" fill="#010101" fill-opacity=".3">` + service.Name + `</text><text x="28" y="14">` + service.Name + `</text><text x="78" y="15" fill="#010101" fill-opacity=".3">online</text><text x="78" y="14">online</text></g></svg>`)
	} else {
		badge = []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="99" height="20"><linearGradient id="b" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><mask id="a"><rect width="99" height="20" rx="3" fill="#fff"/></mask><g mask="url(#a)"><path fill="#555" d="M0 0h54v20H0z"/><path fill="#e05d44" d="M54 0h45v20H54z"/><path fill="url(#b)" d="M0 0h99v20H0z"/></g><g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11"><text x="28" y="15" fill="#010101" fill-opacity=".3">` + service.Name + `</text><text x="28" y="14">` + service.Name + `</text><text x="75.5" y="15" fill="#010101" fill-opacity=".3">offline</text><text x="75.5" y="14">offline</text></g></svg>`)
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	w.Write(badge)

}

func ServicesViewHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	vars := mux.Vars(r)
	service, _ := SelectService(StringInt(vars["id"]))
	tmpl := Parse("service.html")
	serve := &serviceHandler{service, auth}
	tmpl.Execute(w, serve)
}

func Parse(file string) *template.Template {
	nav, _ := tmplBox.String("nav.html")
	render, err := tmplBox.String(file)
	if err != nil {
		panic(err)
	}
	t := template.New("message")
	t.Funcs(template.FuncMap{
		"js": func(html string) template.JS {
			return template.JS(html)
		},
		"safe": func(html string) template.HTML {
			return template.HTML(html)
		},
	})
	t, _ = t.Parse(nav)
	t.Parse(render)
	return t
}

func ParsePlugins(file string) *template.Template {
	nav, _ := tmplBox.String("nav.html")
	slack, _ := tmplBox.String("plugins/slack.html")
	render, err := tmplBox.String(file)
	if err != nil {
		panic(err)
	}
	t := template.New("message")
	t.Funcs(template.FuncMap{
		"js": func(html string) template.JS {
			return template.JS(html)
		},
		"safe": func(html string) template.HTML {
			return template.HTML(html)
		},
	})
	t, _ = t.Parse(nav)
	t.Parse(slack)
	t.Parse(render)
	return t
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("viewing user")
	session, _ := store.Get(r, cookieKey)
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl := Parse("users.html")
	users, _ := SelectAllUsers()
	tmpl.Execute(w, users)
}
