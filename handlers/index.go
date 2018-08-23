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
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"net/http"
)

type index struct {
	Core *core.Core
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if core.CoreApp.DbConnection == "" {
		http.Redirect(w, r, "/setup", http.StatusSeeOther)
		return
	}
	ExecuteResponse(w, r, "index.html", core.CoreApp)
}

func TrayHandler(w http.ResponseWriter, r *http.Request) {
	ExecuteResponse(w, r, "tray.html", core.CoreApp)
}

func DesktopInit(ip string, port int) {
	config := &core.DbConfig{DbConfig: &types.DbConfig{
		DbConn:      "sqlite",
		Project:     "Statup",
		Description: "Statup running as an App!",
		Domain:      "http://localhost",
		Username:    "admin",
		Password:    "admin",
		Email:       "user@email.com",
		Error:       nil,
		Location:    utils.Directory,
	}}

	fmt.Println(config)

	err := config.Save()
	if err != nil {
		utils.Log(4, err)
	}

	if err != nil {
		utils.Log(3, err)
		return
	}

	core.Configs, err = core.LoadConfig()
	if err != nil {
		utils.Log(3, err)
		config.Error = err
		return
	}

	err = core.DbConnection(core.Configs.Connection, false, utils.Directory)
	if err != nil {
		utils.Log(3, err)
		core.DeleteConfig()
		config.Error = err
		return
	}

	admin := core.ReturnUser(&types.User{
		Username: config.Username,
		Password: config.Password,
		Email:    config.Email,
		Admin:    true,
	})
	admin.Create()

	core.LoadSampleData()

	core.InitApp()
	RunHTTPServer(ip, port)
}
