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
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/source"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"net/http"
	"net/url"
	"strconv"
)

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	ExecuteResponse(w, r, "settings.html", core.CoreApp, nil)
}

func saveSettingsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var err error
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
		app.Footer = types.NewNullString(footer)
	}
	domain := r.PostForm.Get("domain")
	if domain != app.Domain {
		app.Domain = domain
	}
	timezone := r.PostForm.Get("timezone")
	timeFloat, _ := strconv.ParseFloat(timezone, 10)
	app.Timezone = float32(timeFloat)

	app.UseCdn = types.NewNullBool(r.PostForm.Get("enable_cdn") == "on")
	core.CoreApp, err = core.UpdateCore(app)
	if err != nil {
		utils.Log(3, fmt.Sprintf("issue updating Core: %v", err.Error()))
	}
	//notifiers.OnSettingsSaved(core.CoreApp.ToCore())
	ExecuteResponse(w, r, "settings.html", core.CoreApp, "/settings")
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
	ExecuteResponse(w, r, "settings.html", core.CoreApp, "/settings")
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
		sendErrorJson(err, w, r)
		return
	}
	err = source.CompileSASS(dir)
	if err != nil {
		source.CopyToPublic(source.CssBox, dir+"/assets/css", "base.css")
		utils.Log(3, "Default 'base.css' was inserted because SASS did not work.")
	}
	resetRouter()
	ExecuteResponse(w, r, "settings.html", core.CoreApp, "/settings")
}

func deleteAssetsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	source.DeleteAllAssets(utils.Directory)
	resetRouter()
	ExecuteResponse(w, r, "settings.html", core.CoreApp, "/settings")
}

func parseId(r *http.Request) int64 {
	vars := mux.Vars(r)
	return utils.ToInt(vars["id"])
}

func parseForm(r *http.Request) url.Values {
	r.ParseForm()
	return r.PostForm
}

func parseGet(r *http.Request) url.Values {
	r.ParseForm()
	return r.Form
}
