// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
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
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/source"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"time"
)

var (
	router *mux.Router
)

func Router() *mux.Router {
	dir := utils.Directory
	r := mux.NewRouter()
	r.Handle("/", http.HandlerFunc(IndexHandler))
	if source.UsingAssets(dir) {
		indexHandler := http.FileServer(http.Dir(dir + "/assets/"))
		r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(dir+"/assets/css"))))
		r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(dir+"/assets/js"))))
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
	r.Handle("/charts.js", http.HandlerFunc(RenderServiceChartsHandler))
	r.Handle("/setup", http.HandlerFunc(SetupHandler)).Methods("GET")
	r.Handle("/setup", http.HandlerFunc(ProcessSetupHandler)).Methods("POST")
	r.Handle("/dashboard", http.HandlerFunc(DashboardHandler)).Methods("GET")
	//r.Handle("/backups/create", http.HandlerFunc(BackupCreateHandler)).Methods("GET")
	r.Handle("/dashboard", http.HandlerFunc(LoginHandler)).Methods("POST")
	r.Handle("/logout", http.HandlerFunc(LogoutHandler))
	r.Handle("/services", http.HandlerFunc(ServicesHandler)).Methods("GET")
	r.Handle("/services", http.HandlerFunc(CreateServiceHandler)).Methods("POST")
	r.Handle("/services/reorder", http.HandlerFunc(ReorderServiceHandler)).Methods("POST")
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

	// SERVICE API Routes
	r.Handle("/api/services", http.HandlerFunc(ApiAllServicesHandler)).Methods("GET")
	r.Handle("/api/services", http.HandlerFunc(ApiCreateServiceHandler)).Methods("POST")
	r.Handle("/api/services/{id}", http.HandlerFunc(ApiServiceHandler)).Methods("GET")
	r.Handle("/api/services/{id}", http.HandlerFunc(ApiServiceUpdateHandler)).Methods("POST")
	r.Handle("/api/services/{id}", http.HandlerFunc(ApiServiceDeleteHandler)).Methods("DELETE")

	// USER API Routes
	r.Handle("/api/users", http.HandlerFunc(ApiAllUsersHandler)).Methods("GET")
	r.Handle("/api/users", http.HandlerFunc(ApiCreateUsersHandler)).Methods("POST")
	r.Handle("/api/users/{id}", http.HandlerFunc(ApiUserHandler)).Methods("GET")
	r.Handle("/api/users/{id}", http.HandlerFunc(ApiUserUpdateHandler)).Methods("POST")
	r.Handle("/api/users/{id}", http.HandlerFunc(ApiUserDeleteHandler)).Methods("DELETE")

	// Generic API Routes
	r.Handle("/api", http.HandlerFunc(ApiIndexHandler))
	r.Handle("/api/renew", http.HandlerFunc(ApiRenewHandler))
	r.Handle("/api/checkin/{api}", http.HandlerFunc(ApiCheckinHandler))
	r.Handle("/metrics", http.HandlerFunc(PrometheusHandler))
	r.NotFoundHandler = http.HandlerFunc(Error404Handler)
	r.Handle("/tray", http.HandlerFunc(TrayHandler))
	return r
}

func ResetRouter() {
	router = Router()
	httpServer.Handler = router
}

func resetCookies() {
	if core.CoreApp != nil {
		cookie := fmt.Sprintf("%v_%v", core.CoreApp.ApiSecret, time.Now().Nanosecond())
		Store = sessions.NewCookieStore([]byte(cookie))
	} else {
		Store = sessions.NewCookieStore([]byte("secretinfo"))
	}
}
