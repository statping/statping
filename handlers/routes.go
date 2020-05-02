package handlers

import (
	"fmt"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/gorilla/mux"
	"github.com/statping/statping/source"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/utils"
	"net/http"
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
	CacheStorage = NewStorage()
	r := mux.NewRouter().StrictSlash(true)

	authUser := utils.Params.GetString("AUTH_USERNAME")
	authPass := utils.Params.GetString("AUTH_PASSWORD")

	if authUser != "" && authPass != "" {
		r.Use(basicAuthHandler)
	}

	bPath := utils.Params.GetString("BASE_PATH")
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	if bPath != "" {
		basePath = "/" + bPath + "/"
		r = r.PathPrefix("/" + bPath).Subrouter()
		r.Handle("", sentryHandler.Handle(http.HandlerFunc(indexHandler)))
	} else {
		r.Handle("/", sentryHandler.Handle(http.HandlerFunc(indexHandler)))
	}

	r.Use(sendLog)

	if source.UsingAssets(dir) {
		indexHandler := http.FileServer(http.Dir(dir + "/assets/"))

		r.PathPrefix("/css/").Handler(http.StripPrefix(basePath, Gzip(staticAssets("css"))))
		r.PathPrefix("/js/").Handler(http.StripPrefix(basePath, Gzip(staticAssets("js"))))
		r.PathPrefix("/scss/").Handler(http.StripPrefix(basePath, Gzip(staticAssets("scss"))))
		r.PathPrefix("/robots.txt").Handler(http.StripPrefix(basePath, indexHandler))
		r.PathPrefix("/favicon.ico").Handler(http.StripPrefix(basePath, indexHandler))
		r.PathPrefix("/banner.png").Handler(http.StripPrefix(basePath, indexHandler))
	} else {
		tmplFileSrv := http.FileServer(source.TmplBox.HTTPBox())
		tmplBoxHandler := http.StripPrefix(basePath, tmplFileSrv)

		r.PathPrefix("/css/").Handler(http.StripPrefix(basePath, Gzip(tmplFileSrv)))
		r.PathPrefix("/scss/").Handler(http.StripPrefix(basePath, Gzip(tmplFileSrv)))
		r.PathPrefix("/js/").Handler(http.StripPrefix(basePath, Gzip(tmplFileSrv)))
		r.PathPrefix("/robots.txt").Handler(tmplBoxHandler)
		r.PathPrefix("/favicon.ico").Handler(tmplBoxHandler)
		r.PathPrefix("/banner.png").Handler(tmplBoxHandler)
	}

	api := r.NewRoute().Subrouter()
	api.Use(apiMiddleware)

	// API Routes
	r.Handle("/api", scoped(apiIndexHandler))
	r.Handle("/api/setup", http.HandlerFunc(processSetupHandler)).Methods("POST")
	api.Handle("/api/login", http.HandlerFunc(apiLoginHandler)).Methods("POST")
	api.Handle("/api/logout", http.HandlerFunc(logoutHandler))
	api.Handle("/api/renew", authenticated(apiRenewHandler, false))
	api.Handle("/api/cache", authenticated(apiCacheHandler, false)).Methods("GET")
	api.Handle("/api/clear_cache", authenticated(apiClearCacheHandler, false))
	api.Handle("/api/core", authenticated(apiCoreHandler, false)).Methods("POST")
	api.Handle("/api/oauth", scoped(apiOAuthHandler)).Methods("GET")
	api.Handle("/oauth/{provider}", http.HandlerFunc(oauthHandler))
	api.Handle("/api/logs", authenticated(logsHandler, false)).Methods("GET")
	api.Handle("/api/logs/last", authenticated(logsLineHandler, false)).Methods("GET")

	// API SCSS and ASSETS Routes
	api.Handle("/api/theme", authenticated(apiThemeViewHandler, false)).Methods("GET")
	api.Handle("/api/theme", authenticated(apiThemeSaveHandler, false)).Methods("POST")
	api.Handle("/api/theme/create", authenticated(apiThemeCreateHandler, false)).Methods("GET")
	api.Handle("/api/theme", authenticated(apiThemeRemoveHandler, false)).Methods("DELETE")

	// API GROUPS Routes
	api.Handle("/api/groups", scoped(apiAllGroupHandler)).Methods("GET")
	api.Handle("/api/groups", authenticated(apiCreateGroupHandler, false)).Methods("POST")
	api.Handle("/api/groups/{id}", readOnly(apiGroupHandler, false)).Methods("GET")
	api.Handle("/api/groups/{id}", authenticated(apiGroupUpdateHandler, false)).Methods("POST")
	api.Handle("/api/groups/{id}", authenticated(apiGroupDeleteHandler, false)).Methods("DELETE")
	api.Handle("/api/reorder/groups", authenticated(apiGroupReorderHandler, false)).Methods("POST")

	// API SERVICE Routes
	api.Handle("/api/services", scoped(apiAllServicesHandler)).Methods("GET")
	api.Handle("/api/services", authenticated(apiCreateServiceHandler, false)).Methods("POST")
	api.Handle("/api/services/{id}", scoped(apiServiceHandler)).Methods("GET")
	api.Handle("/api/reorder/services", authenticated(reorderServiceHandler, false)).Methods("POST")
	api.Handle("/api/services/{id}", authenticated(apiServiceUpdateHandler, false)).Methods("POST")
	api.Handle("/api/services/{id}", authenticated(apiServiceDeleteHandler, false)).Methods("DELETE")
	api.Handle("/api/services/{id}/failures", scoped(apiServiceFailuresHandler)).Methods("GET")
	api.Handle("/api/services/{id}/failures", authenticated(servicesDeleteFailuresHandler, false)).Methods("DELETE")
	api.Handle("/api/services/{id}/hits", scoped(apiServiceHitsHandler)).Methods("GET")

	// API SERVICE CHART DATA Routes
	api.Handle("/api/services/{id}/hits_data", cached("30s", "application/json", apiServiceDataHandler)).Methods("GET")
	api.Handle("/api/services/{id}/failure_data", cached("30s", "application/json", apiServiceFailureDataHandler)).Methods("GET")
	api.Handle("/api/services/{id}/ping_data", cached("30s", "application/json", apiServicePingDataHandler)).Methods("GET")
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
	api.Handle("/api/users", authenticated(apiAllUsersHandler, false)).Methods("GET")
	api.Handle("/api/users", authenticated(apiCreateUsersHandler, false)).Methods("POST")
	api.Handle("/api/users/{id}", authenticated(apiUserHandler, false)).Methods("GET")
	api.Handle("/api/users/{id}", authenticated(apiUserUpdateHandler, false)).Methods("POST")
	api.Handle("/api/users/{id}", authenticated(apiUserDeleteHandler, false)).Methods("DELETE")

	// API NOTIFIER Routes
	api.Handle("/api/notifiers", authenticated(apiNotifiersHandler, false)).Methods("GET")
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
	api.Handle("/api/checkins", authenticated(apiAllCheckinsHandler, false)).Methods("GET")
	api.Handle("/api/checkins", authenticated(checkinCreateHandler, false)).Methods("POST")
	api.Handle("/api/checkins/{api}", authenticated(apiCheckinHandler, false)).Methods("GET")
	api.Handle("/api/checkins/{api}", authenticated(checkinDeleteHandler, false)).Methods("DELETE")
	r.Handle("/checkin/{api}", http.HandlerFunc(checkinHitHandler))

	// API Generic Routes
	r.Handle("/metrics", readOnly(prometheusHandler, false))
	r.Handle("/health", http.HandlerFunc(healthCheckHandler))
	api.Handle("/api/oauth/{provider}", http.HandlerFunc(oauthHandler))
	r.Handle("/.well-known/", http.StripPrefix("/.well-known/", http.FileServer(http.Dir(dir+"/.well-known"))))
	r.NotFoundHandler = http.HandlerFunc(error404Handler)
	return r
}

func resetRouter() {
	router = Router()
	httpServer.Handler = router
}

func resetCookies() {
	jwtKey = fmt.Sprintf("%s_%d", core.App.ApiSecret, utils.Now().Nanosecond())
}
