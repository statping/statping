package main

import (
	"github.com/hunterlong/statup/plugins"
)

var (
	pluginRoutes []*plugins.Routing
)

func InitPluginsDatabase() {
	plugins.InitDB(db)
}

func Routes() []*plugins.Routing {
	return plugins.PluginRoutes
}
