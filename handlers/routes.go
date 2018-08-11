package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/source"
	"net/http"
	"time"
)

var (
	r *mux.Router
)

func Router() *mux.Router {
	r = mux.NewRouter()
	r.Handle("/", http.HandlerFunc(IndexHandler))
	LocalizedAssets(r)
	r.Handle("/charts.js", http.HandlerFunc(RenderServiceChartsHandler))
	r.Handle("/setup", http.HandlerFunc(SetupHandler)).Methods("GET")
	r.Handle("/setup", http.HandlerFunc(ProcessSetupHandler)).Methods("POST")
	r.Handle("/dashboard", http.HandlerFunc(DashboardHandler)).Methods("GET")
	//r.Handle("/backups/create", http.HandlerFunc(BackupCreateHandler)).Methods("GET")
	r.Handle("/dashboard", http.HandlerFunc(LoginHandler)).Methods("POST")
	r.Handle("/logout", http.HandlerFunc(LogoutHandler))
	r.Handle("/services", http.HandlerFunc(ServicesHandler)).Methods("GET")
	r.Handle("/services", http.HandlerFunc(CreateServiceHandler)).Methods("POST")
	r.Handle("/service/{id}", http.HandlerFunc(ServicesViewHandler)).Methods("GET")
	r.Handle("/service/{id}", http.HandlerFunc(ServicesUpdateHandler)).Methods("POST")
	r.Handle("/service/{id}/edit", http.HandlerFunc(ServicesViewHandler))
	r.Handle("/service/{id}/delete", http.HandlerFunc(ServicesDeleteHandler))
	r.Handle("/service/{id}/delete_failures", http.HandlerFunc(ServicesDeleteFailuresHandler)).Methods("GET")
	r.Handle("/service/{id}/checkin", http.HandlerFunc(CheckinCreateUpdateHandler)).Methods("POST")
	r.Handle("/users", http.HandlerFunc(UsersHandler)).Methods("GET")
	r.Handle("/users", http.HandlerFunc(CreateUserHandler)).Methods("POST")
	r.Handle("/user/{id}", http.HandlerFunc(UsersEditHandler)).Methods("GET")
	r.Handle("/user/{id}", http.HandlerFunc(UpdateUserHandler)).Methods("POST")
	r.Handle("/user/{id}/delete", http.HandlerFunc(UsersDeleteHandler)).Methods("GET")
	r.Handle("/settings", http.HandlerFunc(SettingsHandler)).Methods("GET")
	r.Handle("/settings", http.HandlerFunc(SaveSettingsHandler)).Methods("POST")
	r.Handle("/settings/css", http.HandlerFunc(SaveSASSHandler)).Methods("POST")
	r.Handle("/settings/build", http.HandlerFunc(SaveAssetsHandler)).Methods("GET")
	r.Handle("/settings/delete_assets", http.HandlerFunc(DeleteAssetsHandler)).Methods("GET")
	r.Handle("/settings/notifier/{id}", http.HandlerFunc(SaveNotificationHandler)).Methods("POST")
	r.Handle("/plugins/download/{name}", http.HandlerFunc(PluginsDownloadHandler))
	r.Handle("/plugins/{name}/save", http.HandlerFunc(PluginSavedHandler)).Methods("POST")
	r.Handle("/help", http.HandlerFunc(HelpHandler))
	r.Handle("/logs", http.HandlerFunc(LogsHandler))
	r.Handle("/logs/line", http.HandlerFunc(LogsLineHandler))
	r.Handle("/api", http.HandlerFunc(ApiIndexHandler))
	r.Handle("/api/renew", http.HandlerFunc(ApiRenewHandler))
	r.Handle("/api/checkin/{api}", http.HandlerFunc(ApiCheckinHandler))
	r.Handle("/api/services", http.HandlerFunc(ApiAllServicesHandler))
	r.Handle("/api/services/{id}", http.HandlerFunc(ApiServiceHandler)).Methods("GET")
	r.Handle("/api/services/{id}", http.HandlerFunc(ApiServiceUpdateHandler)).Methods("POST")
	r.Handle("/api/users", http.HandlerFunc(ApiAllUsersHandler))
	r.Handle("/api/users/{id}", http.HandlerFunc(ApiUserHandler))
	r.Handle("/metrics", http.HandlerFunc(PrometheusHandler))
	r.NotFoundHandler = http.HandlerFunc(Error404Handler)
	if core.CoreApp != nil {
		cookie := fmt.Sprintf("%v_%v", core.CoreApp.ApiSecret, time.Now().Nanosecond())
		Store = sessions.NewCookieStore([]byte(cookie))
	} else {
		Store = sessions.NewCookieStore([]byte("secretinfo"))
	}
	return r
}

func LocalizedAssets(r *mux.Router) *mux.Router {
	if source.UsingAssets {
		cssHandler := http.FileServer(http.Dir("./assets/css"))
		jsHandler := http.FileServer(http.Dir("./assets/js"))
		indexHandler := http.FileServer(http.Dir("./assets/"))
		r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", cssHandler))
		r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", jsHandler))
		r.PathPrefix("/robots.txt").Handler(indexHandler)
		r.PathPrefix("/favicon.ico").Handler(indexHandler)
		r.PathPrefix("/statup.png").Handler(indexHandler)
	} else {
		r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(source.CssBox.HTTPBox())))
		r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(source.JsBox.HTTPBox())))
		r.PathPrefix("/robots.txt").Handler(http.FileServer(source.TmplBox.HTTPBox()))
		r.PathPrefix("/favicon.ico").Handler(http.FileServer(source.TmplBox.HTTPBox()))
		r.PathPrefix("/statup.png").Handler(http.FileServer(source.TmplBox.HTTPBox()))
	}
	return r
}
