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
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"net/http"
	"strconv"
	"time"
)

func processSetupHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if core.CoreApp.Setup {
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
	dir := utils.Directory

	configs := &types.DbConfig{
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

	log.WithFields(utils.ToFields(core.CoreApp, configs)).Debugln("new configs posted")

	if err := core.SaveConfig(configs); err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	if _, err = core.LoadConfigFile(dir); err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	if err = core.CoreApp.Connect(configs, false, dir); err != nil {
		log.Errorln(err)
		core.DeleteConfig()
		sendErrorJson(err, w, r)
		return
	}

	if err = core.CoreApp.DropDatabase(); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	if err = core.CoreApp.CreateDatabase(); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	core.CoreApp, err = core.InsertCore(configs)
	if err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	admin := &types.User{
		Username: configs.Username,
		Password: configs.Password,
		Email:    configs.Email,
		Admin:    types.NewNullBool(true),
	}
	database.Create(admin)

	if sample {
		if err = core.SampleData(); err != nil {
			sendErrorJson(err, w, r)
			return
		}
	}
	core.InitApp()
	CacheStorage.Delete("/")
	resetCookies()
	time.Sleep(1 * time.Second)
	out := struct {
		Message string          `json:"message"`
		Config  *types.DbConfig `json:"config"`
	}{
		"success",
		configs,
	}
	returnJson(out, w, r)
}

func setupResponseError(w http.ResponseWriter, r *http.Request, a interface{}) {
	ExecuteResponse(w, r, "setup.gohtml", a, nil)
}
