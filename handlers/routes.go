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
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/source"
	"github.com/hunterlong/statping/utils"
	"net/http"
	"time"
)

var (
	router *mux.Router
)

// Router returns all of the routes used in Statping
func Router() *mux.Router {
	dir := utils.Directory
	CacheStorage = NewStorage()
	r := mux.NewRouter()
	r.Handle("/", cached("120s", "text/html", http.HandlerFunc(indexHandler)))
	if source.UsingAssets(dir) {
		indexHandler := http.FileServer(http.Dir(dir + "/assets/"))
		r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(dir+"/assets/css"))))
		r.PathPrefix("/font/").Handler(http.StripPrefix("/font/", http.FileServer(http.Dir(dir+"/assets/font"))))
		r.PathPrefix("/robots.txt").Handler(indexHandler)
		r.PathPrefix("/favicon.ico").Handler(indexHandler)
		r.PathPrefix("/banner.png").Handler(indexHandler)
	} else {
		r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(source.CssBox.HTTPBox())))
		r.PathPrefix("/font/").Handler(http.StripPrefix("/font/", http.FileServer(source.FontBox.HTTPBox())))
		r.PathPrefix("/robots.txt").Handler(http.FileServer(source.TmplBox.HTTPBox()))
		r.PathPrefix("/favicon.ico").Handler(http.FileServer(source.TmplBox.HTTPBox()))
		r.PathPrefix("/banner.png").Handler(http.FileServer(source.TmplBox.HTTPBox()))
	}
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(source.JsBox.HTTPBox())))
	r.Handle("/charts.js", http.HandlerFunc(renderServiceChartsHandler))
	r.Handle("/setup", http.HandlerFunc(setupHandler)).Methods("GET")
	r.Handle("/setup", http.HandlerFunc(processSetupHandler)).Methods("POST")
	r.Handle("/dashboard", http.HandlerFunc(dashboardHandler)).Methods("GET")
	r.Handle("/dashboard", http.HandlerFunc(loginHandler)).Methods("POST")
	r.Handle("/logout", http.HandlerFunc(logoutHandler))
	r.Handle("/plugins/download/{name}", http.HandlerFunc(pluginsDownloadHandler))
	r.Handle("/plugins/{name}/save", http.HandlerFunc(pluginSavedHandler)).Methods("POST")
	r.Handle("/help", http.HandlerFunc(helpHandler))
	r.Handle("/logs", http.HandlerFunc(logsHandler))
	r.Handle("/logs/line", http.HandlerFunc(logsLineHandler))

	// USER Routes
	r.Handle("/users", http.HandlerFunc(usersHandler)).Methods("GET")
	r.Handle("/user/{id}", http.HandlerFunc(usersEditHandler)).Methods("GET")

	// MESSAGES Routes
	r.Handle("/messages", http.HandlerFunc(messagesHandler)).Methods("GET")
	r.Handle("/message/{id}", http.HandlerFunc(viewMessageHandler)).Methods("GET")

	// SETTINGS Routes
	r.Handle("/settings", http.HandlerFunc(settingsHandler)).Methods("GET")
	r.Handle("/settings", http.HandlerFunc(saveSettingsHandler)).Methods("POST")
	r.Handle("/settings/css", http.HandlerFunc(saveSASSHandler)).Methods("POST")
	r.Handle("/settings/build", http.HandlerFunc(saveAssetsHandler)).Methods("GET")
	r.Handle("/settings/delete_assets", http.HandlerFunc(deleteAssetsHandler)).Methods("GET")
	r.Handle("/settings/export", http.HandlerFunc(exportHandler)).Methods("GET")

	// SERVICE Routes
	r.Handle("/services", http.HandlerFunc(servicesHandler)).Methods("GET")
	r.Handle("/service/{id}", http.HandlerFunc(servicesViewHandler)).Methods("GET")
	r.Handle("/service/{id}/edit", http.HandlerFunc(servicesViewHandler))
	r.Handle("/service/{id}/delete_failures", http.HandlerFunc(servicesDeleteFailuresHandler)).Methods("GET")

	// API Routes
	r.Handle("/api", http.HandlerFunc(apiIndexHandler))
	r.Handle("/api/renew", http.HandlerFunc(apiRenewHandler))

	// API SERVICE Routes
	r.Handle("/api/services", http.HandlerFunc(apiAllServicesHandler)).Methods("GET")
	r.Handle("/api/services", http.HandlerFunc(apiCreateServiceHandler)).Methods("POST")
	r.Handle("/api/services/{id}", http.HandlerFunc(apiServiceHandler)).Methods("GET")
	r.Handle("/api/services/reorder", http.HandlerFunc(reorderServiceHandler)).Methods("POST")
	r.Handle("/api/services/{id}/data", cached("120s", "application/json", http.HandlerFunc(apiServiceDataHandler))).Methods("GET")
	r.Handle("/api/services/{id}/ping", http.HandlerFunc(apiServicePingDataHandler)).Methods("GET")
	r.Handle("/api/services/{id}", http.HandlerFunc(apiServiceUpdateHandler)).Methods("POST")
	r.Handle("/api/services/{id}", http.HandlerFunc(apiServiceDeleteHandler)).Methods("DELETE")

	// API USER Routes
	r.Handle("/api/users", http.HandlerFunc(apiAllUsersHandler)).Methods("GET")
	r.Handle("/api/users", http.HandlerFunc(apiCreateUsersHandler)).Methods("POST")
	r.Handle("/api/users/{id}", http.HandlerFunc(apiUserHandler)).Methods("GET")
	r.Handle("/api/users/{id}", http.HandlerFunc(apiUserUpdateHandler)).Methods("POST")
	r.Handle("/api/users/{id}", http.HandlerFunc(apiUserDeleteHandler)).Methods("DELETE")

	// API NOTIFIER Routes
	r.Handle("/api/notifiers", http.HandlerFunc(apiNotifiersHandler)).Methods("GET")
	r.Handle("/api/notifier/{notifier}", http.HandlerFunc(apiNotifierGetHandler)).Methods("GET")
	r.Handle("/api/notifier/{notifier}", http.HandlerFunc(apiNotifierUpdateHandler)).Methods("POST")
	r.Handle("/api/notifier/{method}/test", http.HandlerFunc(testNotificationHandler)).Methods("POST")

	// API MESSAGES Routes
	r.Handle("/api/messages", http.HandlerFunc(apiAllMessagesHandler)).Methods("GET")
	r.Handle("/api/messages", http.HandlerFunc(apiMessageCreateHandler)).Methods("POST")
	r.Handle("/api/messages/{id}", http.HandlerFunc(apiMessageGetHandler)).Methods("GET")
	r.Handle("/api/messages/{id}", http.HandlerFunc(apiMessageUpdateHandler)).Methods("POST")
	r.Handle("/api/messages/{id}", http.HandlerFunc(apiMessageDeleteHandler)).Methods("DELETE")

	// API CHECKIN Routes
	r.Handle("/api/checkins", http.HandlerFunc(apiAllCheckinsHandler)).Methods("GET")
	r.Handle("/api/checkin/{api}", http.HandlerFunc(apiCheckinHandler)).Methods("GET")
	r.Handle("/api/checkin", http.HandlerFunc(checkinCreateHandler)).Methods("POST")
	r.Handle("/api/checkin/{api}", http.HandlerFunc(checkinDeleteHandler)).Methods("DELETE")
	r.Handle("/checkin/{api}", http.HandlerFunc(checkinHitHandler))

	// Static Files Routes
	r.PathPrefix("/files/postman.json").Handler(http.StripPrefix("/files/", http.FileServer(source.TmplBox.HTTPBox())))
	r.PathPrefix("/files/swagger.json").Handler(http.StripPrefix("/files/", http.FileServer(source.TmplBox.HTTPBox())))
	r.PathPrefix("/files/grafana.json").Handler(http.StripPrefix("/files/", http.FileServer(source.TmplBox.HTTPBox())))

	// API Generic Routes
	r.Handle("/metrics", http.HandlerFunc(prometheusHandler))
	r.Handle("/health", http.HandlerFunc(healthCheckHandler))
	r.Handle("/tray", http.HandlerFunc(trayHandler))
	r.Handle("/.well-known/", http.StripPrefix("/.well-known/", http.FileServer(http.Dir(dir+"/.well-known"))))
	r.NotFoundHandler = http.HandlerFunc(error404Handler)
	return r
}

func resetRouter() {
	router = Router()
	httpServer.Handler = router
}

func resetCookies() {
	if core.CoreApp != nil {
		cookie := fmt.Sprintf("%v_%v", core.CoreApp.ApiSecret, time.Now().Nanosecond())
		sessionStore = sessions.NewCookieStore([]byte(cookie))
	} else {
		sessionStore = sessions.NewCookieStore([]byte("secretinfo"))
	}
}
