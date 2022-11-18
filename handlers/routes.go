package handlers

import (
	"net/http"
	"net/http/pprof"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/statping-ng/statping-ng/source"
	"github.com/statping-ng/statping-ng/types/core"
	"github.com/statping-ng/statping-ng/utils"

	_ "github.com/statping-ng/statping-ng/types/metrics"
)

var (
	router *mux.Router
	log    = utils.Log.WithField("type", "handlers")
)

func staticAssets(src string) http.Handler {
	return http.StripPrefix(src+"/", http.FileServer(http.Dir(utils.Directory+"/assets/"+src)))
}

// Router returns all of the routes used in Statping.
// Server will use static assets if the 'assets' directory is found in the root directory.
func Router() *mux.Router {
	dir := utils.Directory

	r := mux.NewRouter().StrictSlash(true)
	r.Use(prometheusMiddleware)

	authUser := utils.Params.GetString("AUTH_USERNAME")
	authPass := utils.Params.GetString("AUTH_PASSWORD")
	if authUser != "" && authPass != "" {
		r.Use(basicAuthHandler)
	}

	bPath := utils.Params.GetString("BASE_PATH")

	if bPath != "" {
		basePath = "/" + bPath + "/"
		r = r.PathPrefix("/" + bPath).Subrouter()
		r.Handle("", http.HandlerFunc(indexHandler))
	} else {
		r.Handle("/", http.HandlerFunc(indexHandler))
	}

	if !utils.Params.GetBool("DISABLE_LOGS") {
		r.Use(sendLog)
	}

	if utils.Params.GetBool("DEBUG") {
		go func() {
			log.Infoln("Starting pprof web server on http://0.0.0.0:9090")
			r := http.NewServeMux()
			r.HandleFunc("/debug/pprof/", pprof.Index)
			r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
			r.HandleFunc("/debug/pprof/profile", pprof.Profile)
			r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
			r.HandleFunc("/debug/pprof/trace", pprof.Trace)
			http.ListenAndServe(":9090", r)
			// pprof -http=:9000 http://localhost:9090/debug/pprof/heap?debug=1
		}()
	}

	if source.UsingAssets(dir) {
		indexHandler := http.FileServer(http.Dir(dir + "/assets/"))

		r.PathPrefix("/css/").Handler(http.StripPrefix(basePath, staticAssets("css")))
		r.PathPrefix("/robots.txt").Handler(http.StripPrefix(basePath, indexHandler))
	} else {
		tmplFileSrv := http.FileServer(source.TmplBox.HTTPBox())
		tmplBoxHandler := http.StripPrefix(basePath, tmplFileSrv)

		r.PathPrefix("/css/").Handler(http.StripPrefix(basePath, tmplFileSrv))
		r.PathPrefix("/robots.txt").Handler(tmplBoxHandler)
	}

	r.PathPrefix("/js/").Handler(http.StripPrefix(basePath, http.FileServer(source.TmplBox.HTTPBox())))

	api := r.NewRoute().Subrouter()
	api.Use(apiMiddleware)
	api.Use(prometheusMiddleware)

	// API Routes
	r.Handle("/api", scoped(apiIndexHandler))
	r.Handle("/api/setup", http.HandlerFunc(processSetupHandler)).Methods("POST")
	api.Handle("/api/login", http.HandlerFunc(apiLoginHandler)).Methods("POST")
	api.Handle("/api/logout", http.HandlerFunc(logoutHandler))
	api.Handle("/api/renew", authenticated(apiRenewHandler, false))
	api.Handle("/api/core", authenticated(apiCoreHandler, false)).Methods("POST")
	api.Handle("/api/logs", authenticated(logsHandler, false)).Methods("GET")
	api.Handle("/api/logs/last", authenticated(logsLineHandler, false)).Methods("GET")
	api.Handle("/api/settings/import", authenticated(settingsImportHandler, false)).Methods("POST")
	api.Handle("/api/settings/export", authenticated(settingsExportHandler, false)).Methods("GET")
	api.Handle("/api/settings/configs", authenticated(configsViewHandler, false)).Methods("GET")
	api.Handle("/api/settings/configs", authenticated(configsSaveHandler, false)).Methods("POST")

	// API OAUTH Routes
	api.Handle("/api/oauth", scoped(apiOAuthHandler)).Methods("GET")
	api.Handle("/api/oauth", authenticated(apiUpdateOAuthHandler, false)).Methods("POST")
	api.Handle("/oauth/{provider}", http.HandlerFunc(oauthHandler))

	// API SCSS and ASSETS Routes
	api.Handle("/api/theme", authenticated(apiThemeViewHandler, false)).Methods("GET")
	api.Handle("/api/theme", authenticated(apiThemeSaveHandler, false)).Methods("POST")
	api.Handle("/api/theme/create", authenticated(apiThemeCreateHandler, false)).Methods("GET")
	api.Handle("/api/theme", authenticated(apiThemeRemoveHandler, false)).Methods("DELETE")

	// API GROUPS Routes
	api.Handle("/api/groups", scoped(apiAllGroupHandler)).Methods("GET")
	api.Handle("/api/groups", authenticated(apiCreateGroupHandler, false)).Methods("POST")
	api.Handle("/api/groups/{id}", readOnly(http.HandlerFunc(apiGroupHandler), false)).Methods("GET")
	api.Handle("/api/groups/{id}", authenticated(apiGroupUpdateHandler, false)).Methods("POST")
	api.Handle("/api/groups/{id}", authenticated(apiGroupDeleteHandler, false)).Methods("DELETE")
	api.Handle("/api/reorder/groups", authenticated(apiGroupReorderHandler, false)).Methods("POST")

	// API SERVICE Routes
	api.Handle("/api/services", scoped(apiAllServicesHandler)).Methods("GET")
	api.Handle("/api/services", authenticated(apiCreateServiceHandler, false)).Methods("POST")
	api.Handle("/api/services/{id}", scoped(apiServiceHandler)).Methods("GET")
	api.Handle("/api/reorder/services", authenticated(reorderServiceHandler, false)).Methods("POST")
	api.Handle("/api/services/{id}", authenticated(apiServiceUpdateHandler, false)).Methods("POST")
	api.Handle("/api/services/{id}", authenticated(apiServicePatchHandler, false)).Methods("PATCH")
	api.Handle("/api/services/{id}", authenticated(apiServiceDeleteHandler, false)).Methods("DELETE")
	api.Handle("/api/services/{id}/failures", scoped(apiServiceFailuresHandler)).Methods("GET")
	api.Handle("/api/services/{id}/failures", authenticated(servicesDeleteFailuresHandler, false)).Methods("DELETE")
	api.Handle("/api/services/{id}/hits", scoped(apiServiceHitsHandler)).Methods("GET")
	api.Handle("/api/services/{id}/hits", authenticated(apiServiceHitsDeleteHandler, false)).Methods("DELETE")

	// API SERVICE CHART DATA Routes
	api.Handle("/api/services/{id}/hits_data", http.HandlerFunc(apiServiceDataHandler)).Methods("GET")
	api.Handle("/api/services/{id}/failure_data", http.HandlerFunc(apiServiceFailureDataHandler)).Methods("GET")
	api.Handle("/api/services/{id}/ping_data", http.HandlerFunc(apiServicePingDataHandler)).Methods("GET")
	api.Handle("/api/services/{id}/uptime_data", http.HandlerFunc(apiServiceTimeDataHandler)).Methods("GET")

	// API INCIDENTS Routes
	api.Handle("/api/services/{id}/incidents", http.HandlerFunc(apiServiceIncidentsHandler)).Methods("GET")
	api.Handle("/api/services/{id}/incidents", authenticated(apiCreateIncidentHandler, false)).Methods("POST")
	api.Handle("/api/incidents/{id}", authenticated(apiIncidentUpdateHandler, false)).Methods("POST")
	api.Handle("/api/incidents/{id}", authenticated(apiDeleteIncidentHandler, false)).Methods("DELETE")

	// API INCIDENTS UPDATES Routes
	api.Handle("/api/incidents/{id}/updates", http.HandlerFunc(apiIncidentUpdatesHandler)).Methods("GET")
	api.Handle("/api/incidents/{id}/updates", authenticated(apiCreateIncidentUpdateHandler, false)).Methods("POST")
	api.Handle("/api/incidents/{id}/updates/{uid}", authenticated(apiDeleteIncidentUpdateHandler, false)).Methods("DELETE")

	// API USER Routes
	api.Handle("/api/users", scoped(apiAllUsersHandler)).Methods("GET")
	api.Handle("/api/users", authenticated(apiCreateUsersHandler, false)).Methods("POST")
	api.Handle("/api/users/token", http.HandlerFunc(apiCheckUserTokenHandler)).Methods("POST")
	api.Handle("/api/users/{id}", authenticated(apiUserHandler, false)).Methods("GET")
	api.Handle("/api/users/{id}", authenticated(apiUserUpdateHandler, false)).Methods("POST")
	api.Handle("/api/users/{id}", authenticated(apiUserDeleteHandler, false)).Methods("DELETE")

	// API NOTIFIER Routes
	api.Handle("/api/notifiers", scoped(apiAllNotifiersHandler)).Methods("GET")
	api.Handle("/api/notifier/{notifier}", authenticated(apiNotifierGetHandler, false)).Methods("GET")
	api.Handle("/api/notifier/{notifier}", authenticated(apiNotifierUpdateHandler, false)).Methods("POST")
	api.Handle("/api/notifier/{notifier}/test", authenticated(testNotificationHandler, false)).Methods("POST")

	// API MESSAGES Routes
	api.Handle("/api/messages", scoped(apiAllMessagesHandler)).Methods("GET")
	api.Handle("/api/messages", authenticated(apiMessageCreateHandler, false)).Methods("POST")
	api.Handle("/api/messages/{id}", scoped(apiMessageGetHandler)).Methods("GET")
	api.Handle("/api/messages/{id}", authenticated(apiMessageUpdateHandler, false)).Methods("POST")
	api.Handle("/api/messages/{id}", authenticated(apiMessageDeleteHandler, false)).Methods("DELETE")

	// API CHECKIN Routes
	api.Handle("/api/checkins", scoped(apiAllCheckinsHandler)).Methods("GET")
	api.Handle("/api/checkins", authenticated(checkinCreateHandler, false)).Methods("POST")
	api.Handle("/api/checkins/{api}", authenticated(apiCheckinHandler, false)).Methods("GET")
	api.Handle("/api/checkins/{api}", authenticated(checkinDeleteHandler, false)).Methods("DELETE")
	r.Handle("/checkin/{api}", http.HandlerFunc(checkinHitHandler))

	// API Generic Routes
	r.Handle("/metrics", readOnly(promhttp.Handler(), false))
	r.Handle("/health", http.HandlerFunc(healthCheckHandler))
	r.NotFoundHandler = http.HandlerFunc(baseHandler)
	return r
}

func resetRouter() {
	log.Infoln("Restarting HTTP Router")
	router = Router()
	httpServer.Handler = router
}

func resetCookies() {
	if core.App == nil {
		jwtKey = []byte(utils.NewSHA256Hash())
		return
	}
	jwtKey = []byte(utils.Sha256Hash(core.App.ApiSecret))
}
