package main

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/hunterlong/statup/types"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	session *sessions.CookieStore
)

const (
	cookieKey = "statup_auth"
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
	r.Handle("/service/{id}", http.HandlerFunc(ServicesViewHandler)).Methods("GET")
	r.Handle("/service/{id}", http.HandlerFunc(ServicesUpdateHandler)).Methods("POST")
	r.Handle("/service/{id}/edit", http.HandlerFunc(ServicesViewHandler))
	r.Handle("/service/{id}/delete", http.HandlerFunc(ServicesDeleteHandler))
	r.Handle("/service/{id}/badge.svg", http.HandlerFunc(ServicesBadgeHandler))
	r.Handle("/service/{id}/delete_failures", http.HandlerFunc(ServicesDeleteFailuresHandler)).Methods("GET")
	r.Handle("/service/{id}/checkin", http.HandlerFunc(CheckinCreateUpdateHandler)).Methods("POST")
	r.Handle("/users", http.HandlerFunc(UsersHandler)).Methods("GET")
	r.Handle("/users", http.HandlerFunc(CreateUserHandler)).Methods("POST")
	r.Handle("/users/{id}/delete", http.HandlerFunc(UsersDeleteHandler)).Methods("GET")
	r.Handle("/settings", http.HandlerFunc(PluginsHandler)).Methods("GET")
	r.Handle("/settings", http.HandlerFunc(SaveSettingsHandler)).Methods("POST")
	r.Handle("/settings/email", http.HandlerFunc(SaveEmailSettingsHandler)).Methods("POST")
	r.Handle("/plugins/download/{name}", http.HandlerFunc(PluginsDownloadHandler))
	r.Handle("/plugins/{name}/save", http.HandlerFunc(PluginSavedHandler)).Methods("POST")
	r.Handle("/help", http.HandlerFunc(HelpHandler))
	r.Handle("/api", http.HandlerFunc(ApiIndexHandler))
	r.Handle("/api/checkin/{api}", http.HandlerFunc(ApiCheckinHandler))
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
	user, auth := AuthUser(username, password)
	if auth {
		session.Values["authenticated"] = true
		session.Values["user_id"] = user.Id
		session.Save(r, w)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		err := ErrorResponse{Error: "Incorrect login information submitted, try again."}
		ExecuteResponse(w, r, "login.html", err)
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
	port, _ := strconv.Atoi(r.PostForm.Get("port"))
	checkType := r.PostForm.Get("check_type")

	service := &Service{
		Name:           name,
		Domain:         domain,
		Method:         method,
		Expected:       expected,
		ExpectedStatus: status,
		Interval:       interval,
		Type:           checkType,
		Port:           port,
	}

	_, err := service.Create()
	if err != nil {
		go service.CheckQueue()
	}
	http.Redirect(w, r, "/services", http.StatusSeeOther)
}

func SetupHandler(w http.ResponseWriter, r *http.Request) {
	if core != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	port := 5432
	if os.Getenv("DB_CONN") == "mysql" {
		port = 3306
	}
	var data interface{}
	if os.Getenv("DB_CONN") != "" {
		data = &DbConfig{
			DbConn:      os.Getenv("DB_CONN"),
			DbHost:      os.Getenv("DB_HOST"),
			DbUser:      os.Getenv("DB_USER"),
			DbPass:      os.Getenv("DB_PASS"),
			DbData:      os.Getenv("DB_DATABASE"),
			DbPort:      port,
			Project:     os.Getenv("NAME"),
			Description: os.Getenv("DESCRIPTION"),
			Email:       "",
			Username:    "admin",
			Password:    "",
		}
	}
	ExecuteResponse(w, r, "setup.html", data)
}

type index struct {
	Core     Core
	Services []*Service
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if core == nil {
		http.Redirect(w, r, "/setup", http.StatusSeeOther)
		return
	}
	out := index{*core, services}
	ExecuteResponse(w, r, "index.html", out)
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
		err := ErrorResponse{}
		ExecuteResponse(w, r, "login.html", err)
	} else {
		fails, _ := CountFailures()
		out := dashboard{services, core, CountOnline(), len(services), fails}
		ExecuteResponse(w, r, "dashboard.html", out)
	}
}

type serviceHandler struct {
	Service Service
	Auth    bool
}

func ServicesHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	ExecuteResponse(w, r, "services.html", services)
}

func ServicesDeleteHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	service := SelectService(StringInt(vars["id"]))
	service.Delete()
	http.Redirect(w, r, "/services", http.StatusSeeOther)
}

func ServicesDeleteFailuresHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	service := SelectService(StringInt(vars["id"]))

	service.DeleteFailures()
	services, _ = SelectAllServices()
	http.Redirect(w, r, "/services", http.StatusSeeOther)
}

func IsAuthenticated(r *http.Request) bool {
	if store == nil {
		return false
	}
	session, err := store.Get(r, cookieKey)
	if err != nil {
		return false
	}
	if session.Values["authenticated"] == nil {
		return false
	}
	return session.Values["authenticated"].(bool)
}

func SaveEmailSettingsHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	emailer := SelectCommunication(1)

	r.ParseForm()
	emailer.Host = r.PostForm.Get("host")
	emailer.Username = r.PostForm.Get("username")
	emailer.Password = r.PostForm.Get("password")
	emailer.Port = int(StringInt(r.PostForm.Get("port")))
	emailer.Var1 = r.PostForm.Get("address")
	Update(emailer)

	sample := &types.Email{
		To:       SessionUser(r).Email,
		Subject:  "Sample Email",
		Template: "error.html",
	}
	AddEmail(sample)

	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}

func SaveSettingsHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	r.ParseForm()
	name := r.PostForm.Get("project")
	if name != "" {
		core.Name = name
	}
	description := r.PostForm.Get("description")
	if description != core.Description {
		core.Description = description
	}
	style := r.PostForm.Get("style")
	if style != core.Style {
		core.Style = style
	}
	footer := r.PostForm.Get("footer")
	if footer != core.Footer {
		core.Footer = footer
	}
	core.Update()
	OnSettingsSaved(core)
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}

func PluginsHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	core.FetchPluginRepo()

	var pluginFields []PluginSelect

	for _, p := range allPlugins {
		fields := structs.Map(p.GetInfo())

		pluginFields = append(pluginFields, PluginSelect{p.GetInfo().Name, p.GetForm(), fields})
	}

	core.PluginFields = pluginFields
	fmt.Println(core.Communications)

	ExecuteResponse(w, r, "plugins.html", core)
}

type PluginSelect struct {
	Plugin string
	Form   string
	Params map[string]interface{}
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
	plug.OnSave(structs.Map(data))
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
	ExecuteResponse(w, r, "help.html", nil)
}

func CheckinCreateUpdateHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	interval := StringInt(r.PostForm.Get("interval"))
	service := SelectService(StringInt(vars["id"]))
	checkin := &Checkin{
		Service:  service.Id,
		Interval: interval,
		Api:      NewSHA1Hash(18),
	}
	checkin.Create()
	fmt.Println(checkin.Create())
	ExecuteResponse(w, r, "service.html", service)
}

func ServicesUpdateHandler(w http.ResponseWriter, r *http.Request) {
	//auth := IsAuthenticated(r)
	//
	//vars := mux.Vars(r)
	//service := SelectService(vars["id"])
}

func ServicesBadgeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := SelectService(StringInt(vars["id"]))

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
	vars := mux.Vars(r)
	service := SelectService(StringInt(vars["id"]))
	ExecuteResponse(w, r, "service.html", service)
}

func ExecuteResponse(w http.ResponseWriter, r *http.Request, file string, data interface{}) {
	nav, _ := tmplBox.String("nav.html")
	footer, _ := tmplBox.String("footer.html")
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
		"Auth": func() bool {
			return IsAuthenticated(r)
		},
		"VERSION": func() string {
			return VERSION
		},
		"underscore": func(html string) string {
			return UnderScoreString(html)
		},
		"User": func() *User {
			return SessionUser(r)
		},
	})
	t, _ = t.Parse(nav)
	t, _ = t.Parse(footer)
	t.Parse(render)
	t.Execute(w, data)
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	users, _ := SelectAllUsers()
	ExecuteResponse(w, r, "users.html", users)
}

func UsersDeleteHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user, _ := SelectUser(int64(id))

	users, _ := SelectAllUsers()
	if len(users) == 1 {
		http.Redirect(w, r, "/users", http.StatusSeeOther)
		return
	}
	user.Delete()
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func UnderScoreString(str string) string {

	// convert every letter to lower case
	newStr := strings.ToLower(str)

	// convert all spaces/tab to underscore
	regExp := regexp.MustCompile("[[:space:][:blank:]]")
	newStrByte := regExp.ReplaceAll([]byte(newStr), []byte("_"))

	regExp = regexp.MustCompile("`[^a-z0-9]`i")
	newStrByte = regExp.ReplaceAll(newStrByte, []byte("_"))

	regExp = regexp.MustCompile("[!/']")
	newStrByte = regExp.ReplaceAll(newStrByte, []byte("_"))

	// and remove underscore from beginning and ending

	newStr = strings.TrimPrefix(string(newStrByte), "_")
	newStr = strings.TrimSuffix(newStr, "_")

	return newStr
}
