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
	"errors"
	"github.com/hunterlong/statping/types/configs"
	"github.com/hunterlong/statping/types/core"
	"github.com/hunterlong/statping/types/null"
	"github.com/hunterlong/statping/types/users"
	"github.com/hunterlong/statping/utils"
	"net/http"
	"strconv"
	"time"
)

func processSetupHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if core.App.Setup {
		sendErrorJson(errors.New("Statping has already been setup"), w, r)
		return
	}
	if err = r.ParseForm(); err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}
	dbHost := r.PostForm.Get("db_host")
	dbUser := r.PostForm.Get("db_user")
	dbPass := r.PostForm.Get("db_password")
	dbDatabase := r.PostForm.Get("db_database")
	dbConn := r.PostForm.Get("db_connection")
	dbPort := utils.ToInt(r.PostForm.Get("db_port"))
	project := r.PostForm.Get("project")
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	description := r.PostForm.Get("description")
	domain := r.PostForm.Get("domain")
	email := r.PostForm.Get("email")
	sample, _ := strconv.ParseBool(r.PostForm.Get("sample_data"))

	confg := &configs.DbConfig{
		DbConn:      dbConn,
		DbHost:      dbHost,
		DbUser:      dbUser,
		DbPass:      dbPass,
		DbData:      dbDatabase,
		DbPort:      int(dbPort),
		Project:     project,
		Description: description,
		Domain:      domain,
		Username:    username,
		Password:    password,
		Email:       email,
		Error:       nil,
		Location:    utils.Directory,
	}

	log.WithFields(utils.ToFields(core.App, confg)).Debugln("new configs posted")

	if err := confg.Save(utils.Directory); err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	if _, err = configs.LoadConfigs(); err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	if err = configs.ConnectConfigs(confg); err != nil {
		log.Errorln(err)
		if err := confg.Delete(); err != nil {
			log.Errorln(err)
			sendErrorJson(err, w, r)
		}
	}

	if err = configs.MigrateDatabase(); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	c := &core.Core{
		Name:        "Statping Sample Data",
		Description: "This data is only used to testing",
		//ApiKey:      apiKey.(string),
		//ApiSecret:   apiSecret.(string),
		Domain:    "http://localhost:8080",
		Version:   "test",
		CreatedAt: time.Now().UTC(),
		UseCdn:    null.NewNullBool(false),
		Footer:    null.NewNullString(""),
	}

	if err := c.Create(); err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	core.App = c

	admin := &users.User{
		Username: confg.Username,
		Password: confg.Password,
		Email:    confg.Email,
		Admin:    null.NewNullBool(true),
	}

	if err := admin.Create(); err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	if sample {
		if err = configs.TriggerSamples(); err != nil {
			sendErrorJson(err, w, r)
			return
		}
	}

	core.InitApp()
	CacheStorage.Delete("/")
	resetCookies()
	time.Sleep(1 * time.Second)
	out := struct {
		Message string            `json:"message"`
		Config  *configs.DbConfig `json:"config"`
	}{
		"success",
		confg,
	}
	returnJson(out, w, r)
}

func setupResponseError(w http.ResponseWriter, r *http.Request, a interface{}) {
	ExecuteResponse(w, r, "setup.gohtml", a, nil)
}
