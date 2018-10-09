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

package plugin

import (
	"fmt"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"io/ioutil"
	"net/http"
	"os"
	"plugin"
	"strings"
)

//
//     STATUP PLUGIN INTERFACE
//
//            v0.1
//
//       https://statup.io
//
//
// An expandable plugin framework that will still
// work even if there's an update or addition.
//

var (
	AllPlugins []*PluginObject
	dir        string
)

func init() {
	utils.InitLogs()
	dir = utils.Directory
}

func Add(p Pluginer) *PluginObject {
	return &PluginObject{}
}

func (p *PluginObject) AddRoute(s string, i string, f http.HandlerFunc) {

}

func (p *PluginInfo) Form() string {
	return "okkokokkok"
}

func LoadPlugins(debug bool) {
	pluginDir := dir + "/plugins"
	utils.Log(1, fmt.Sprintf("Loading any available Plugins from /plugins directory"))
	if _, err := os.Stat(pluginDir); os.IsNotExist(err) {
		os.Mkdir(pluginDir, os.ModePerm)
	}

	//ForEachPlugin()
	files, err := ioutil.ReadDir(pluginDir)
	if err != nil {
		utils.Log(2, fmt.Sprintf("Plugins directory was not found. Error: %v\n", err))
		return
	}
	for _, f := range files {
		utils.Log(1, fmt.Sprintf("Attempting to load plugin '%v'", f.Name()))
		ext := strings.Split(f.Name(), ".")
		if len(ext) != 2 {
			utils.Log(3, fmt.Sprintf("Plugin '%v' must end in .so extension", f.Name()))
			continue
		}
		if ext[1] != "so" {
			utils.Log(3, fmt.Sprintf("Plugin '%v' must end in .so extension", f.Name()))
			continue
		}
		plug, err := plugin.Open("plugins/" + f.Name())
		if err != nil {
			utils.Log(3, fmt.Sprintf("Plugin '%v' could not load correctly. %v", f.Name(), err))
			continue
		}
		symPlugin, err := plug.Lookup("Plugin")
		if err != nil {
			utils.Log(3, fmt.Sprintf("Plugin '%v' could not load correctly. %v", f.Name(), err))
			continue
		}

		if debug {
			utils.Log(1, fmt.Sprintf("Plugin '%v' struct:", f.Name()))
			//utils.Log(1, structs.Map(symPlugin))
		}

		var plugActions types.PluginActions
		plugActions, ok := symPlugin.(types.PluginActions)
		if !ok {
			utils.Log(3, fmt.Sprintf("Plugin '%v' could not load correctly, error: %v", f.Name(), err))
			if debug {
				//fmt.Println(symPlugin.(plugin.PluginActions))
			}
			continue
		}

		plugActions.OnLoad(*core.DbSession)
		core.CoreApp.Plugins = append(core.CoreApp.Plugins, plugActions.GetInfo())
		core.CoreApp.AllPlugins = append(core.CoreApp.AllPlugins, plugActions)
	}
	if !debug {
		utils.Log(1, fmt.Sprintf("Loaded %v Plugins\n", len(core.CoreApp.Plugins)))
	}
}
