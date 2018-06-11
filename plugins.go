package main

import (
	"github.com/hunterlong/statup/plugins"
)

var (
	pluginRoutes []*plugins.Routing
	allPlugins   []*plugins.Plugin
)

func InitPluginsDatabase() {
	plugins.InitDB(db)
}

func SetAuthorized() {

}

func AfterInstall() {

}

func AfterUninstall() {

}

func AfterSave() {

}

func Routes() []*plugins.Routing {
	return plugins.PluginRoutes
}

func AllPlugins() []*plugins.Plugin {
	return plugins.Plugins
}
