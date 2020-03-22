// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/statping/statping
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
	return http.StripPrefix(basePath+src+"/", http.FileServer(http.Dir(utils.Directory+"/assets/"+src)))
}

// Router returns all of the routes used in Statping.
// Server will use static assets if the 'assets' directory is found in the root directory.
func Router() *mux.Router {
	dir := utils.Directory
	CacheStorage = NewStorage()
	r := mux.NewRouter().StrictSlash(true)

	authUser := utils.Getenv("AUTH_USERNAME", "").(string)
	authPass := utils.Getenv("AUTH_PASSWORD", "").(string)

	if authUser != "" && authPass != "" {
		r.Use(basicAuthHandler)
	}

	bPath := utils.Getenv("BASE_PATH", "").(string)
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

		r.PathPrefix("/css/").Handler(Gzip(staticAssets("css")))
		r.PathPrefix("/font/").Handler(staticAssets("font"))
		r.PathPrefix("/js/").Handler(Gzip(staticAssets("js")))
		r.PathPrefix("/robots.txt").Handler(http.StripPrefix(basePath, indexHandler))
		r.PathPrefix("/favicon.ico").Handler(http.StripPrefix(basePath, indexHandler))
		r.PathPrefix("/banner.png").Handler(http.StripPrefix(basePath, indexHandler))
	} else {
		tmplFileSrv := http.FileServer(source.TmplBox.HTTPBox())
		tmplBoxHandler := http.StripPrefix(basePath, tmplFileSrv)

		r.PathPrefix("/css/").Handler(Gzip(tmplFileSrv))
		r.PathPrefix("/scss/").Handler(Gzip(tmplFileSrv))
		r.PathPrefix("/font/").Handler(tmplFileSrv)
		r.PathPrefix("/js/").Handler(Gzip(tmplFileSrv))
		r.PathPrefix("/robots.txt").Handler(tmplBoxHandler)
		r.PathPrefix("/favicon.ico").Handler(tmplBoxHandler)
		r.PathPrefix("/banner.png").Handler(tmplBoxHandler)
	}

	api := r.NewRoute().Subrouter()
	api.Use(apiMiddleware)

	// API Routes
	r.Handle("/api", scoped(apiIndexHandler))
	r.Handle("/api/setup", http.HandlerFunc(processSetupHandler)).Methods("POST")
	//r.Handle("/oauth/callback", http.HandlerFunc(OAuthRedirect))
	api.Handle("/api/login", http.HandlerFunc(apiLoginHandler)).Methods("POST")
	api.Handle("/api/logout", http.HandlerFunc(logoutHandler))
	api.Handle("/api/renew", authenticated(apiRenewHandler, false))
	api.Handle("/api/cache", authenticated(apiCacheHandler, false)).Methods("GET")
	api.Handle("/api/clear_cache", authenticated(apiClearCacheHandler, false))
	api.Handle("/api/core", authenticated(apiCoreHandler, false)).Methods("POST")
	api.Handle("/api/logs", authenticated(logsHandler, false)).Methods("GET")
	api.Handle("/api/logs/last", authenticated(logsLineHandler, false)).Methods("GET")

	// API SCSS and ASSETS Routes
	api.Handle("/api/theme", authenticated(apiThemeHandler, false)).Methods("GET")
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
	api.Handle("/api/services/{id}/running", authenticated(apiServiceRunningHandler, false)).Methods("POST")
	api.Handle("/api/services/{id}", authenticated(apiServiceUpdateHandler, false)).Methods("POST")
	api.Handle("/api/services/{id}", authenticated(apiServiceDeleteHandler, false)).Methods("DELETE")
	api.Handle("/api/services/{id}/failures", scoped(apiServiceFailuresHandler)).Methods("GET")
	api.Handle("/api/services/{id}/failures", authenticated(servicesDeleteFailuresHandler, false)).Methods("DELETE")
	api.Handle("/api/services/{id}/hits", scoped(apiServiceHitsHandler)).Methods("GET")

	// API SERVICE CHART DATA Routes
	api.Handle("/api/services/{id}/hits_data", cached("30s", "application/json", apiServiceDataHandler)).Methods("GET")
	api.Handle("/api/services/{id}/failure_data", cached("30s", "application/json", apiServiceFailureDataHandler)).Methods("GET")
	api.Handle("/api/services/{id}/ping_data", cached("30s", "application/json", apiServicePingDataHandler)).Methods("GET")
	//api.Handle("/api/services/{id}/heatmap", cached("30s", "application/json", apiServiceHeatmapHandler)).Methods("GET")

	// API INCIDENTS Routes
	api.Handle("/api/services/{id}/incidents", http.HandlerFunc(apiServiceIncidentsHandler)).Methods("GET")
	api.Handle("/api/services/{id}/incidents", authenticated(apiCreateIncidentHandler, false)).Methods("POST")
	api.Handle("/api/incidents/{id}", authenticated(apiIncidentUpdateHandler, false)).Methods("POST")
	api.Handle("/api/incidents/{id}", authenticated(apiDeleteIncidentHandler, false)).Methods("DELETE")

	// API INCIDENTS UPDATES Routes
	api.Handle("/api/incidents/{id}/updates", authenticated(apiIncidentUpdatesHandler, false)).Methods("GET")
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
	api.Handle("/api/notifier/{method}/test", authenticated(testNotificationHandler, false)).Methods("POST")

	// API MESSAGES Routes
	api.Handle("/api/messages", scoped(apiAllMessagesHandler)).Methods("GET")
	api.Handle("/api/messages", authenticated(apiMessageCreateHandler, false)).Methods("POST")
	api.Handle("/api/messages/{id}", scoped(apiMessageGetHandler)).Methods("GET")
	api.Handle("/api/messages/{id}", authenticated(apiMessageUpdateHandler, false)).Methods("POST")
	api.Handle("/api/messages/{id}", authenticated(apiMessageDeleteHandler, false)).Methods("DELETE")

	// API CHECKIN Routes
	api.Handle("/api/checkins", authenticated(apiAllCheckinsHandler, false)).Methods("GET")
	api.Handle("/api/checkin/{api}", authenticated(apiCheckinHandler, false)).Methods("GET")
	api.Handle("/api/checkin", authenticated(checkinCreateHandler, false)).Methods("POST")
	api.Handle("/api/checkin/{api}", authenticated(checkinDeleteHandler, false)).Methods("DELETE")
	r.Handle("/checkin/{api}", http.HandlerFunc(checkinHitHandler))

	//r.PathPrefix("/").Handler(http.HandlerFunc(indexHandler))
	//r.Handle("/badge", http.HandlerFunc(badgeHandler)).Methods("GET")

	// Static Files Routes
	r.PathPrefix("/files/postman.json").Handler(http.StripPrefix("/files/", http.FileServer(source.TmplBox.HTTPBox())))
	r.PathPrefix("/files/swagger.json").Handler(http.StripPrefix("/files/", http.FileServer(source.TmplBox.HTTPBox())))
	r.PathPrefix("/files/grafana.json").Handler(http.StripPrefix("/files/", http.FileServer(source.TmplBox.HTTPBox())))

	// API Generic Routes
	r.Handle("/metrics", readOnly(prometheusHandler, false))
	r.Handle("/health", http.HandlerFunc(healthCheckHandler))
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
