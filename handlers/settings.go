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
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/core/notifier"
	"github.com/hunterlong/statup/source"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"net/url"
	"strconv"
)

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	executeResponse(w, r, "settings.html", core.CoreApp, nil)
}

func saveSettingsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	r.ParseForm()
	app := core.CoreApp
	name := r.PostForm.Get("project")
	if name != "" {
		app.Name = name
	}
	description := r.PostForm.Get("description")
	if description != app.Description {
		app.Description = description
	}
	style := r.PostForm.Get("style")
	if style != app.Style {
		app.Style = style
	}
	footer := r.PostForm.Get("footer")
	if footer != app.Footer.String {
		app.Footer = utils.NullString(footer)
	}
	domain := r.PostForm.Get("domain")
	if domain != app.Domain {
		app.Domain = domain
	}
	timezone := r.PostForm.Get("timezone")
	timeFloat, _ := strconv.ParseFloat(timezone, 10)
	app.Timezone = float32(timeFloat)

	app.UseCdn = utils.NullBool(r.PostForm.Get("enable_cdn") == "on")
	core.CoreApp, _ = core.UpdateCore(app)
	//notifiers.OnSettingsSaved(core.CoreApp.ToCore())
	executeResponse(w, r, "settings.html", core.CoreApp, "/settings")
}

func saveSASSHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	r.ParseForm()
	theme := r.PostForm.Get("theme")
	variables := r.PostForm.Get("variables")
	mobile := r.PostForm.Get("mobile")
	source.SaveAsset([]byte(theme), utils.Directory, "scss/base.scss")
	source.SaveAsset([]byte(variables), utils.Directory, "scss/variables.scss")
	source.SaveAsset([]byte(mobile), utils.Directory, "scss/mobile.scss")
	source.CompileSASS(utils.Directory)
	resetRouter()
	executeResponse(w, r, "settings.html", core.CoreApp, "/settings")
}

func saveAssetsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	dir := utils.Directory
	err := source.CreateAllAssets(dir)
	if err != nil {
		utils.Log(3, err)
		return
	}
	err = source.CompileSASS(dir)
	if err != nil {
		source.CopyToPublic(source.CssBox, dir+"/assets/css", "base.css")
		utils.Log(3, "Default 'base.css' was inserted because SASS did not work.")
	}
	resetRouter()
	executeResponse(w, r, "settings.html", core.CoreApp, "/settings")
}

func deleteAssetsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	source.DeleteAllAssets(utils.Directory)
	resetRouter()
	executeResponse(w, r, "settings.html", core.CoreApp, "/settings")
}

func parseId(r *http.Request) int64 {
	vars := mux.Vars(r)
	return utils.StringInt(vars["id"])
}

func parseForm(r *http.Request) url.Values {
	r.ParseForm()
	return r.PostForm
}

func parseGet(r *http.Request) url.Values {
	r.ParseForm()
	return r.Form
}

func saveNotificationHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	form := parseForm(r)
	vars := mux.Vars(r)
	method := vars["method"]
	enabled := form.Get("enable")
	host := form.Get("host")
	port := int(utils.StringInt(form.Get("port")))
	username := form.Get("username")
	password := form.Get("password")
	var1 := form.Get("var1")
	var2 := form.Get("var2")
	apiKey := form.Get("api_key")
	apiSecret := form.Get("api_secret")
	limits := int(utils.StringInt(form.Get("limits")))

	notifer, notif, err := notifier.SelectNotifier(method)
	if err != nil {
		utils.Log(3, fmt.Sprintf("issue saving notifier %v: %v", method, err))
		executeResponse(w, r, "settings.html", core.CoreApp, "/settings")
		return
	}

	if host != "" {
		notifer.Host = host
	}
	if port != 0 {
		notifer.Port = port
	}
	if username != "" {
		notifer.Username = username
	}
	if password != "" && password != "##########" {
		notifer.Password = password
	}
	if var1 != "" {
		notifer.Var1 = var1
	}
	if var2 != "" {
		notifer.Var2 = var2
	}
	if apiKey != "" {
		notifer.ApiKey = apiKey
	}
	if apiSecret != "" {
		notifer.ApiSecret = apiSecret
	}
	if limits != 0 {
		notifer.Limits = limits
	}
	notifer.Enabled = sql.NullBool{enabled == "on", true}

	_, err = notifier.Update(notif, notifer)
	if err != nil {
		utils.Log(3, fmt.Sprintf("issue updating notifier: %v", err))
	}
	notifier.OnSave(notifer.Method)
	executeResponse(w, r, "settings.html", core.CoreApp, "/settings")
}

func testNotificationHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	form := parseForm(r)
	vars := mux.Vars(r)
	method := vars["method"]
	enabled := form.Get("enable")
	host := form.Get("host")
	port := int(utils.StringInt(form.Get("port")))
	username := form.Get("username")
	password := form.Get("password")
	var1 := form.Get("var1")
	var2 := form.Get("var2")
	apiKey := form.Get("api_key")
	apiSecret := form.Get("api_secret")
	limits := int(utils.StringInt(form.Get("limits")))

	fakeNotifer, notif, err := notifier.SelectNotifier(method)
	if err != nil {
		utils.Log(3, fmt.Sprintf("issue saving notifier %v: %v", method, err))
		executeResponse(w, r, "settings.html", core.CoreApp, "/settings")
		return
	}

	notifer := *fakeNotifer

	if host != "" {
		notifer.Host = host
	}
	if port != 0 {
		notifer.Port = port
	}
	if username != "" {
		notifer.Username = username
	}
	if password != "" && password != "##########" {
		notifer.Password = password
	}
	if var1 != "" {
		notifer.Var1 = var1
	}
	if var2 != "" {
		notifer.Var2 = var2
	}
	if apiKey != "" {
		notifer.ApiKey = apiKey
	}
	if apiSecret != "" {
		notifer.ApiSecret = apiSecret
	}
	if limits != 0 {
		notifer.Limits = limits
	}
	notifer.Enabled = sql.NullBool{enabled == "on", true}

	err = notif.(notifier.Tester).OnTest()
	if err == nil {
		w.Write([]byte("ok"))
	} else {
		w.Write([]byte(err.Error()))
	}
}
