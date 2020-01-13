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

package plugin

import (
	"fmt"
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"io/ioutil"
	"os"
	"plugin"
	"strings"
)

//
//     STATPING PLUGIN INTERFACE
//
//            v0.1
//
//       https://statping.com
//
//
// An expandable plugin framework that will still
// work even if there's an update or addition.
//

var (
	AllPlugins []*types.PluginObject
	dir        string
	log        = utils.Log.WithField("type", "plugin")
)

func init() {
	utils.InitLogs()
	dir = utils.Directory
}

func LoadPlugin(file string) error {
	log.Infoln(fmt.Sprintf("opening file %v", file))
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	fSplit := strings.Split(f.Name(), "/")
	fileBin := fSplit[len(fSplit)-1]

	log.Infoln(fmt.Sprintf("Attempting to load plugin '%v'", fileBin))
	ext := strings.Split(fileBin, ".")
	if len(ext) != 2 {
		log.Errorln(fmt.Sprintf("Plugin '%v' must end in .so extension", fileBin))
		return fmt.Errorf("Plugin '%v' must end in .so extension %v", fileBin, len(ext))
	}
	if ext[1] != "so" {
		log.Errorln(fmt.Sprintf("Plugin '%v' must end in .so extension", fileBin))
		return fmt.Errorf("Plugin '%v' must end in .so extension", fileBin)
	}
	plug, err := plugin.Open(file)
	if err != nil {
		log.Errorln(fmt.Sprintf("Plugin '%v' could not load correctly. %v", fileBin, err))
		return err
	}
	symPlugin, err := plug.Lookup("Plugin")
	if err != nil {
		log.Errorln(fmt.Sprintf("Plugin '%v' could not locate Plugin variable. %v", fileBin, err))
		return err
	}
	plugActions, ok := symPlugin.(types.PluginActions)
	if !ok {
		log.Errorln(fmt.Sprintf("Plugin %v was not type PluginObject", f.Name()))
		return fmt.Errorf("Plugin %v was not type PluginActions %v", f.Name(), plugActions.GetInfo())
	}
	info := plugActions.GetInfo()
	err = plugActions.OnLoad()
	if err != nil {
		return err
	}
	log.Infoln(fmt.Sprintf("Plugin %v loaded from %v", info.Name, f.Name()))
	core.CoreApp.AllPlugins = append(core.CoreApp.AllPlugins, plugActions)
	return nil
}

func LoadPlugins() {
	pluginDir := dir + "/plugins"
	log.Infoln(fmt.Sprintf("Loading any available Plugins from /plugins directory"))
	if _, err := os.Stat(pluginDir); os.IsNotExist(err) {
		os.Mkdir(pluginDir, os.ModePerm)
	}
	files, err := ioutil.ReadDir(pluginDir)
	if err != nil {
		log.Warnln(fmt.Sprintf("Plugins directory was not found. Error: %v", err))
		return
	}
	for _, f := range files {
		err := LoadPlugin(f.Name())
		if err != nil {
			log.Errorln(err)
			continue
		}
	}
	log.Infoln(fmt.Sprintf("Loaded %v Plugins", len(core.CoreApp.Plugins)))
}
