package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/handlers"
	"github.com/hunterlong/statup/plugin"
	"github.com/hunterlong/statup/utils"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
	plg "plugin"
	"strings"
)

var (
	VERSION string
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
	utils.Log(1, fmt.Sprintf("Starting Statup v%v\n", VERSION))
	RenderBoxes()
	core.HasAssets()

	core.Configs, err = core.LoadConfig()
	if err != nil {
		utils.Log(2, "config.yml file not found - starting in setup mode")
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
	core.EmailBox = rice.MustFindBox("source/emails")
}

func LoadDotEnvs() {
	err := godotenv.Load()
	if err == nil {
		utils.Log(1, "Environment file '.env' Loaded")
	}
}

func mainProcess() {
	var err error
	err = core.DbConnection(core.Configs.Connection)
	if err != nil {
		utils.Log(3, err)
	}
	core.RunDatabaseUpgrades()
	core.CoreApp, err = core.SelectCore()
	if err != nil {
		utils.Log(2, "Core database was not found, Statup is not setup yet.")
		handlers.RunHTTPServer()
	}

	core.CheckServices()
	core.CoreApp.Communications, err = core.SelectAllCommunications()
	if err != nil {
		utils.Log(2, err)
	}
	core.LoadDefaultCommunications()

	go core.DatabaseMaintence()

	if !core.SetupMode {
		LoadPlugins()
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

func LoadPlugins() {
	if _, err := os.Stat("./plugins"); os.IsNotExist(err) {
		os.Mkdir("./plugins", os.ModePerm)
	}

	ForEachPlugin()

	files, err := ioutil.ReadDir("./plugins")
	if err != nil {
		utils.Log(2, fmt.Sprintf("Plugins directory was not found. Error: %v\n", err))
		return
	}
	for _, f := range files {
		ext := strings.Split(f.Name(), ".")
		if len(ext) != 2 {
			continue
		}
		if ext[1] != "so" {
			continue
		}
		plug, err := plg.Open("plugins/" + f.Name())
		if err != nil {
			utils.Log(2, fmt.Sprintf("Plugin '%v' could not load correctly.\n", f.Name()))
			continue
		}
		symPlugin, err := plug.Lookup("Plugin")

		var plugActions plugin.PluginActions
		plugActions, ok := symPlugin.(plugin.PluginActions)
		if !ok {
			utils.Log(2, fmt.Sprintf("Plugin '%v' could not load correctly, error: %v\n", f.Name(), "unexpected type from module symbol"))
			continue
		}

		//allPlugins = append(allPlugins, plugActions)
		core.CoreApp.Plugins = append(core.CoreApp.Plugins, plugActions.GetInfo())
	}

	core.OnLoad(core.DbSession)

	//utils.Log(1, fmt.Sprintf("Loaded %v Plugins\n", len(allPlugins)))
	ForEachPlugin()
}
