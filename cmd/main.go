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

package main

import (
	"flag"
	"fmt"
	"github.com/fatih/structs"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/handlers"
	"github.com/hunterlong/statup/source"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
	plg "plugin"
	"strings"
)

var (
	VERSION   string
	COMMIT    string
	usingEnv  bool
	directory string
	ipAddress string
	port      int
)

func init() {
	core.VERSION = VERSION
}

func parseFlags() {
	ip := flag.String("ip", "0.0.0.0", "IP address to run the Statup HTTP server")
	p := flag.Int("port", 8080, "Port to run the HTTP server")
	flag.Parse()
	ipAddress = *ip
	port = *p
}

func main() {
	var err error
	parseFlags()

	if len(os.Args) >= 2 {
		err := CatchCLI(os.Args)
		if err != nil {
			if err.Error() == "end" {
				os.Exit(0)
			}
			fmt.Println(err)
			os.Exit(1)
		}
	}

	utils.InitLogs()
	source.Assets()
	LoadDotEnvs()

	utils.Log(1, fmt.Sprintf("Starting Statup v%v", VERSION))

	core.Configs, err = core.LoadConfig()
	if err != nil {
		utils.Log(3, err)
		core.SetupMode = true
		fmt.Println(handlers.RunHTTPServer(ipAddress, port))
		os.Exit(1)
	}
	mainProcess()
}

func LoadDotEnvs() error {
	err := godotenv.Load()
	if err == nil {
		utils.Log(1, "Environment file '.env' Loaded")
		usingEnv = true
	}
	return err
}

func mainProcess() {
	var err error
	err = core.DbConnection(core.Configs.Connection, false, ".")
	if err != nil {
		utils.Log(4, fmt.Sprintf("could not connect to database: %v", err))
	}

	core.RunDatabaseUpgrades()
	core.InitApp()

	if !core.SetupMode {
		LoadPlugins(false)
		fmt.Println(handlers.RunHTTPServer(ipAddress, port))
		os.Exit(1)
	}
}

func ForEachPlugin() {
	if len(core.CoreApp.Plugins) > 0 {
		//for _, p := range core.Plugins {
		//	p.OnShutdown()
		//}
	}
}

func LoadPlugins(debug bool) {
	utils.Log(1, fmt.Sprintf("Loading any available Plugins from /plugins directory"))
	if _, err := os.Stat("./plugins"); os.IsNotExist(err) {
		os.Mkdir("./plugins", os.ModePerm)
	}

	//ForEachPlugin()
	files, err := ioutil.ReadDir("./plugins")
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
		plug, err := plg.Open("plugins/" + f.Name())
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
			utils.Log(1, structs.Map(symPlugin))
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

		if debug {
			TestPlugin(plugActions)
		} else {
			plugActions.OnLoad(core.DbSession)
			core.CoreApp.Plugins = append(core.CoreApp.Plugins, plugActions.GetInfo())
			core.CoreApp.AllPlugins = append(core.CoreApp.AllPlugins, plugActions)
		}
	}
	if !debug {
		utils.Log(1, fmt.Sprintf("Loaded %v Plugins\n", len(core.CoreApp.Plugins)))
	}
}
