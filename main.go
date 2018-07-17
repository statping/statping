package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/fatih/structs"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/handlers"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
	plg "plugin"
	"strings"
)

var (
	VERSION  string
	usingEnv bool
)

func init() {
	LoadDotEnvs()
	core.VERSION = VERSION
}

func main() {
	var err error
	if len(os.Args) >= 2 {
		CatchCLI(os.Args)
		os.Exit(0)
	}
	utils.Log(1, fmt.Sprintf("Starting Statup v%v", VERSION))
	RenderBoxes()
	core.HasAssets()

	core.Configs, err = core.LoadConfig()
	if err != nil {
		utils.Log(3, err)
		core.SetupMode = true
		handlers.RunHTTPServer()
	}
	mainProcess()
}

func RenderBoxes() {
	core.SqlBox = rice.MustFindBox("source/sql")
	core.CssBox = rice.MustFindBox("source/css")
	core.ScssBox = rice.MustFindBox("source/scss")
	core.JsBox = rice.MustFindBox("source/js")
	core.TmplBox = rice.MustFindBox("source/tmpl")
}

func LoadDotEnvs() {
	err := godotenv.Load()
	if err == nil {
		utils.Log(1, "Environment file '.env' Loaded")
		usingEnv = true
	}
}

func mainProcess() {
	var err error
	err = core.DbConnection(core.Configs.Connection)
	if err != nil {
		utils.Log(4, fmt.Sprintf("could not connect to database: %v", err))
	}

	core.RunDatabaseUpgrades()
	core.InitApp()

	if !core.SetupMode {
		LoadPlugins(false)
		handlers.RunHTTPServer()
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
