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
	"bytes"
	"fmt"
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/source"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	ExecuteResponse(w, r, "settings.gohtml", core.CoreApp, nil)
}

func saveSettingsHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	form := parseForm(r)
	app := core.CoreApp
	name := form.Get("project")
	if name != "" {
		app.Name = name
	}
	description := form.Get("description")
	if description != app.Description {
		app.Description = description
	}
	style := form.Get("style")
	if style != app.Style {
		app.Style = style
	}
	footer := form.Get("footer")
	if footer != app.Footer.String {
		app.Footer = types.NewNullString(footer)
	}
	domain := form.Get("domain")
	if domain != app.Domain {
		app.Domain = domain
	}
	timezone := form.Get("timezone")
	timeFloat, _ := strconv.ParseFloat(timezone, 10)
	app.Timezone = float32(timeFloat)

	app.UseCdn = types.NewNullBool(form.Get("enable_cdn") == "on")
	core.CoreApp, err = core.UpdateCore(app)
	if err != nil {
		utils.Log(3, fmt.Sprintf("issue updating Core: %v", err.Error()))
	}
	//notifiers.OnSettingsSaved(core.CoreApp.ToCore())
	ExecuteResponse(w, r, "settings.gohtml", core.CoreApp, "/settings")
}

func saveSASSHandler(w http.ResponseWriter, r *http.Request) {
	form := parseForm(r)
	theme := form.Get("theme")
	variables := form.Get("variables")
	mobile := form.Get("mobile")
	source.SaveAsset([]byte(theme), utils.Directory, "scss/base.scss")
	source.SaveAsset([]byte(variables), utils.Directory, "scss/variables.scss")
	source.SaveAsset([]byte(mobile), utils.Directory, "scss/mobile.scss")
	source.CompileSASS(utils.Directory)
	resetRouter()
	ExecuteResponse(w, r, "settings.gohtml", core.CoreApp, "/settings")
}

func saveAssetsHandler(w http.ResponseWriter, r *http.Request) {
	dir := utils.Directory
	if err := source.CreateAllAssets(dir); err != nil {
		utils.Log(3, err)
		sendErrorJson(err, w, r)
		return
	}
	if err := source.CompileSASS(dir); err != nil {
		source.CopyToPublic(source.CssBox, dir+"/assets/css", "base.css")
		utils.Log(3, "Default 'base.css' was inserted because SASS did not work.")
	}
	resetRouter()
	ExecuteResponse(w, r, "settings.gohtml", core.CoreApp, "/settings")
}

func deleteAssetsHandler(w http.ResponseWriter, r *http.Request) {
	if err := source.DeleteAllAssets(utils.Directory); err != nil {
		utils.Log(3, fmt.Errorf("error deleting all assets %v", err))
	}
	resetRouter()
	ExecuteResponse(w, r, "settings.gohtml", core.CoreApp, "/settings")
}

func bulkImportHandler(w http.ResponseWriter, r *http.Request) {
	var fileData bytes.Buffer
	file, _, err := r.FormFile("file")
	if err != nil {
		utils.Log(3, fmt.Errorf("error bulk import services: %v", err))
		w.Write([]byte(err.Error()))
		return
	}
	defer file.Close()

	io.Copy(&fileData, file)
	data := fileData.String()

	for i, line := range strings.Split(strings.TrimSuffix(data, "\n"), "\n")[1:] {
		col := strings.Split(line, ",")

		newService, err := commaToService(col)
		if err != nil {
			utils.Log(3, fmt.Errorf("issue with row %v: %v", i, err))
			return
		}

		service := core.ReturnService(newService)
		_, err = service.Create(true)
		if err != nil {
			utils.Log(3, fmt.Errorf("cannot create service %v: %v", col[0], err))
			return
		}
		utils.Log(1, fmt.Sprintf("Created new service %v", service.Name))
	}

	ExecuteResponse(w, r, "settings.gohtml", core.CoreApp, "/settings")
}

// commaToService will convert a CSV comma delimited string slice to a Service type
// this function is used for the bulk import services feature
func commaToService(s []string) (*types.Service, error) {
	if len(s) != 16 {
		err := fmt.Errorf("does not have the expected amount of %v columns for a service", 16)
		return nil, err
	}

	interval, err := time.ParseDuration(s[4])
	if err != nil {
		return nil, err
	}

	timeout, err := time.ParseDuration(s[9])
	if err != nil {
		return nil, err
	}

	allowNotifications, err := strconv.ParseBool(s[10])
	if err != nil {
		return nil, err
	}

	public, err := strconv.ParseBool(s[11])
	if err != nil {
		return nil, err
	}

	newService := &types.Service{
		Name:               s[0],
		Domain:             s[1],
		Expected:           types.NewNullString(s[2]),
		ExpectedStatus:     int(utils.ToInt(s[3])),
		Interval:           int(utils.ToInt(interval.Seconds())),
		Type:               s[5],
		Method:             s[6],
		PostData:           types.NewNullString(s[7]),
		Port:               int(utils.ToInt(s[8])),
		Timeout:            int(utils.ToInt(timeout.Seconds())),
		AllowNotifications: types.NewNullBool(allowNotifications),
		Public:             types.NewNullBool(public),
		GroupId:            int(utils.ToInt(s[12])),
		Headers:            types.NewNullString(s[13]),
		Permalink:          types.NewNullString(s[14]),
	}

	return newService, nil

}

func parseForm(r *http.Request) url.Values {
	r.ParseForm()
	return r.PostForm
}

func parseGet(r *http.Request) url.Values {
	r.ParseForm()
	return r.Form
}
