package main

import (
	"github.com/hunterlong/statup/plugins"
)

var (
	pluginRoutes []*plugins.Routing
	allPlugins []*plugins.Plugin
)

func InitPluginsDatabase() {
	plugins.InitDB(db)
}

func SetAuthorized() {

}

func Routes() []*plugins.Routing {
	return plugins.PluginRoutes
}
