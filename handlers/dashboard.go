package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/statping/statping/source"
	"github.com/statping/statping/types/checkins"
	"github.com/statping/statping/types/configs"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/errors"
	"github.com/statping/statping/types/groups"
	"github.com/statping/statping/types/incidents"
	"github.com/statping/statping/types/messages"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/types/users"
	"github.com/statping/statping/utils"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	removeJwtToken(w)
	out := make(map[string]string)
	out["status"] = "success"
	returnJson(out, w, r)
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	utils.LockLines.Lock()
	logs := make([]string, 0)
	length := len(utils.LastLines)
	// We need string log lines from end to start.
	for i := length - 1; i >= 0; i-- {
		logs = append(logs, utils.LastLines[i].FormatForHtml()+"\r\n")
	}
	utils.LockLines.Unlock()
	returnJson(logs, w, r)
}

type themeApi struct {
	Directory string `json:"directory,omitempty"`
	Base      string `json:"base"`
	Forms     string `json:"forms"`
	Layout    string `json:"layout"`
	Mixins    string `json:"mixins"`
	Mobile    string `json:"mobile"`
	Variables string `json:"variables"`
}

func apiThemeViewHandler(w http.ResponseWriter, r *http.Request) {
	var base, forms, layout, mixins, variables, mobile, dir string
	assets := utils.Directory + "/assets"

	if _, err := os.Stat(assets); err == nil {
		dir = assets
	}

	if dir != "" {
		base, _ = utils.OpenFile(dir + "/scss/base.scss")
		variables, _ = utils.OpenFile(dir + "/scss/variables.scss")
		mobile, _ = utils.OpenFile(dir + "/scss/mobile.scss")
		layout, _ = utils.OpenFile(dir + "/scss/layout.scss")
		forms, _ = utils.OpenFile(dir + "/scss/forms.scss")
		mixins, _ = utils.OpenFile(dir + "/scss/mixin.scss")
	} else {
		base, _ = source.TmplBox.String("scss/base.scss")
		variables, _ = source.TmplBox.String("scss/variables.scss")
		mobile, _ = source.TmplBox.String("scss/mobile.scss")
		layout, _ = source.TmplBox.String("scss/layout.scss")
		forms, _ = source.TmplBox.String("scss/forms.scss")
		mixins, _ = source.TmplBox.String("scss/mixin.scss")
	}

	resp := &themeApi{
		Directory: dir,
		Base:      base,
		Variables: variables,
		Mobile:    mobile,
		Layout:    layout,
		Forms:     forms,
		Mixins:    mixins,
	}
	returnJson(resp, w, r)
}

func apiThemeSaveHandler(w http.ResponseWriter, r *http.Request) {
	var themes themeApi
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&themes)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	defer r.Body.Close()

	fmt.Println(themes.Variables)

	if err := source.SaveAsset([]byte(themes.Base), "scss/base.scss"); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	if err := source.SaveAsset([]byte(themes.Layout), "scss/layout.scss"); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	if err := source.SaveAsset([]byte(themes.Variables), "scss/variables.scss"); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	if err := source.SaveAsset([]byte(themes.Forms), "scss/forms.scss"); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	if err := source.SaveAsset([]byte(themes.Mixins), "scss/mixin.scss"); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	if err := source.SaveAsset([]byte(themes.Mobile), "scss/mobile.scss"); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	if err := source.CompileSASS(); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	resetRouter()
	sendJsonAction(themes, "saved", w, r)
}

func apiThemeCreateHandler(w http.ResponseWriter, r *http.Request) {
	dir := utils.Params.GetString("STATPING_DIR")
	if source.UsingAssets(dir) {
		err := errors.New("assets have already been created")
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}
	utils.Log.Infof("creating assets in folder: %s/%s", dir, "assets")
	if err := source.CreateAllAssets(dir); err != nil {
		if err := source.CopyToPublic(source.TmplBox, "css", "style.css"); err != nil {
			log.Errorln(err)
			sendErrorJson(err, w, r)
			return
		} else {
			log.Errorln(err)
			sendErrorJson(err, w, r)
		}
	}
	resetRouter()
	sendJsonAction(dir+"/assets", "created", w, r)
}

func apiThemeRemoveHandler(w http.ResponseWriter, r *http.Request) {
	if err := source.DeleteAllAssets(utils.Directory); err != nil {
		log.Errorln(fmt.Errorf("error deleting all assets %v", err))
	}
	resetRouter()
	sendJsonAction(utils.Directory+"/assets", "deleted", w, r)
}

type ExportData struct {
	Config          *configs.DbConfig            `json:"config,omitempty"`
	Core            *core.Core                   `json:"core"`
	Services        []services.Service           `json:"services"`
	Messages        []*messages.Message          `json:"messages"`
	Incidents       []*incidents.Incident        `json:"incidents"`
	IncidentUpdates []*incidents.IncidentUpdate  `json:"incident_updates"`
	Checkins        []*checkins.Checkin          `json:"checkins"`
	Users           []*users.User                `json:"users"`
	Groups          []*groups.Group              `json:"groups"`
	Notifiers       []notifications.Notification `json:"notifiers"`
}

func (e *ExportData) JSON() []byte {
	d, _ := json.Marshal(e)
	return d
}

func ExportSettings() (*ExportData, error) {
	var notifiers []notifications.Notification
	for _, n := range services.AllNotifiers() {
		notifiers = append(notifiers, *n.Select())
	}

	data := &ExportData{
		Core:      core.App,
		Notifiers: notifiers,
		Checkins:  checkins.All(),
		Users:     users.All(),
		Services:  services.AllInOrder(),
		Groups:    groups.All(),
		Messages:  messages.All(),
	}
	return data, nil
}

func settingsImportHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	defer r.Body.Close()

	var exportData *ExportData
	if err := json.Unmarshal(data, &exportData); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	if exportData.Core != nil {
		core.App = exportData.Core
		if err := core.App.Update(); err != nil {
			sendErrorJson(err, w, r)
			return
		}
	}

	if exportData.Groups != nil {
		for _, s := range exportData.Groups {
			s.Id = 0
			if err := s.Create(); err != nil {
				sendErrorJson(err, w, r)
				return
			}
		}
	}

	if exportData.Services != nil {
		for _, s := range exportData.Services {
			s.Id = 0
			if err := s.Create(); err != nil {
				sendErrorJson(err, w, r)
				return
			}
		}
	}

	if exportData.Users != nil {
		for _, s := range exportData.Users {
			s.Id = 0
			if err := s.Create(); err != nil {
				sendErrorJson(err, w, r)
				return
			}
		}
	}

	if exportData.Notifiers != nil {
		for _, s := range exportData.Notifiers {
			notif := services.ReturnNotifier(s.Method)
			n := notif.Select().UpdateFields(&s)
			if err := n.Update(); err != nil {
				sendErrorJson(err, w, r)
				return
			}
		}
	}

	sendJsonAction(exportData, "import", w, r)
}

func configsSaveHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	defer r.Body.Close()

	var cfg *configs.DbConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	oldCfg, err := configs.LoadConfigs(utils.Directory + "/configs.yml")
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	newCfg := cfg.Merge(oldCfg)
	if err := newCfg.Save(utils.Directory); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	sendJsonAction(newCfg.Clean(), "updated", w, r)
}

func configsViewHandler(w http.ResponseWriter, r *http.Request) {
	db, err := configs.LoadConfigs(utils.Directory + "/configs.yml")
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	w.Write(db.Clean().ToYAML())
}

func settingsExportHandler(w http.ResponseWriter, r *http.Request) {
	exported, err := ExportSettings()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	file := bytes.NewBuffer(exported.JSON())

	w.Header().Set("Content-Disposition", "attachment; filename=statping.json")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", utils.ToString(len(exported.JSON())))

	io.Copy(w, file)
}

func logsLineHandler(w http.ResponseWriter, r *http.Request) {
	if lastLine := utils.GetLastLine(); lastLine != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(lastLine.FormatForHtml()))
	}
}

func apiLoginHandler(w http.ResponseWriter, r *http.Request) {
	form := parseForm(r)
	username := form.Get("username")
	password := form.Get("password")

	user, auth := users.AuthUser(username, password)
	if auth {
		log.Infoln(fmt.Sprintf("User %v logged in from IP %v", user.Username, r.RemoteAddr))
		claim, token := setJwtToken(user, w)
		resp := struct {
			Token   string `json:"token"`
			IsAdmin bool   `json:"admin"`
		}{
			token,
			claim.Admin,
		}
		returnJson(resp, w, r)
	} else {
		resp := struct {
			Error string `json:"error"`
		}{
			"incorrect authentication",
		}
		returnJson(resp, w, r)
	}
}
