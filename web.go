package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func RunHTTPServer() {

	r := mux.NewRouter()

	fmt.Println("Statup HTTP Server running on http://localhost:8080")
	r.Handle("/", http.HandlerFunc(IndexHandler))

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(cssBox.HTTPBox())))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(jsBox.HTTPBox())))

	r.Handle("/setup", http.HandlerFunc(SetupHandler))
	r.Handle("/setup/save", http.HandlerFunc(ProcessSetupHandler))
	r.Handle("/dashboard", http.HandlerFunc(DashboardHandler))
	r.Handle("/login", http.HandlerFunc(LoginHandler))
	r.Handle("/logout", http.HandlerFunc(LogoutHandler))

	r.Handle("/services", http.HandlerFunc(ServicesHandler))
	r.Handle("/services", http.HandlerFunc(CreateServiceHandler)).Methods("POST")
	r.Handle("/service/{id}", http.HandlerFunc(ServicesViewHandler))
	r.Handle("/service/{id}", http.HandlerFunc(ServicesUpdateHandler)).Methods("POST")
	r.Handle("/service/{id}/edit", http.HandlerFunc(ServicesViewHandler))
	r.Handle("/service/{id}/delete", http.HandlerFunc(ServicesDeleteHandler))

	r.Handle("/users", http.HandlerFunc(UsersHandler))
	r.Handle("/users", http.HandlerFunc(CreateUserHandler)).Methods("POST")

	r.Handle("/settings", http.HandlerFunc(SettingsHandler))
	r.Handle("/plugins", http.HandlerFunc(PluginsHandler))
	r.Handle("/help", http.HandlerFunc(HelpHandler))

	for _, route := range Routes() {
		fmt.Printf("Adding plugin route: /plugins/%v\n", route.URL)
		r.Handle("/plugins/"+route.URL, http.HandlerFunc(route.Handler)).Methods(route.Method)
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
	session, _ := store.Get(r, "apizer_auth")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "apizer_auth")
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	user, auth := AuthUser(username, password)
	fmt.Println(user)
	fmt.Println(auth)
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

//func AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	key := r.PostForm.Get("key")
//	secret := r.PostForm.Get("secret")
//	token := SelectToken(key, secret)
//	if token.Id != 0 {
//		go token.Hit(r)
//		w.WriteHeader(200)
//		w.Header().Set("Content-Type", "plain/text")
//		fmt.Fprintln(w, token.Id)
//	} else {
//		w.WriteHeader(502)
//		w.Header().Set("Content-Type", "plain/text")
//		fmt.Fprintln(w, "bad")
//	}
//}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	user := &User{
		Username: username,
		Password: password,
	}
	user.Create()
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func CreateServiceHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.PostForm.Get("name")
	domain := r.PostForm.Get("domain")
	method := r.PostForm.Get("method")
	expected := r.PostForm.Get("expected")
	status, _ := strconv.Atoi(r.PostForm.Get("expected_status"))
	interval, _ := strconv.Atoi(r.PostForm.Get("interval"))

	service := &Service{
		Name:           name,
		Domain:         domain,
		Method:         method,
		Expected:       expected,
		ExpectedStatus: status,
		Interval:       interval,
	}

	fmt.Println(service)

	service.Create()
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
	Count24Failures int
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "apizer_auth")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		tmpl := Parse("login.html")
		tmpl.Execute(w, nil)
	} else {
		tmpl := Parse("dashboard.html")
		out := dashboard{services, core, CountOnline(), len(services), CountFailures()}
		tmpl.Execute(w, out)
	}

}

type serviceHandler struct {
	Service *Service
	Auth    bool
}

func ServicesHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "apizer_auth")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl := Parse("services.html")
	tmpl.Execute(w, services)
}

func ServicesDeleteHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "apizer_auth")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	service := SelectService(vars["id"])

	service.Delete()
	services = SelectAllServices()
	http.Redirect(w, r, "/services", http.StatusSeeOther)
}

func IsAuthenticated(r *http.Request) bool {
	session, _ := store.Get(r, "apizer_auth")
	return session.Values["authenticated"].(bool)
}

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl := Parse("settings.html")
	tmpl.Execute(w, core)
}

func PluginsHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl := ParsePlugins("plugins.html")
	tmpl.Execute(w, core)
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

func ServicesViewHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	vars := mux.Vars(r)
	service := SelectService(vars["id"])

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
	})
	t, _ = t.Parse(nav)
	t.Parse(slack)
	t.Parse(render)
	return t
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "apizer_auth")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl := Parse("users.html")
	tmpl.Execute(w, SelectAllUsers())
}

func PermissionsHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "apizer_auth")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	permsFile, err := tmplBox.String("permissions.html")
	if err != nil {
		panic(err)
	}
	permsTmpl, err := template.New("message").Parse(permsFile)
	if err != nil {
		panic(err)
	}
	permsTmpl.Execute(w, SelectAllUsers())
}
