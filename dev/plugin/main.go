package main

import (
	"github.com/hunterlong/statping/types"
	"net/http"
)

type PluginObj types.PluginInfo

var Plugin = PluginObj{
	Info: &types.Info{
		Name:        "Example Plugin",
		Description: "This is an example plugin for Statping Status Page application. It can be implemented pretty quick!",
	},
	Routes: []*types.PluginRoute{{
		Url:    "/setuper",
		Method: "GET",
		Func:   SampleHandler,
	}},
}

func main() {

}

func SampleHandler(w http.ResponseWriter, r *http.Request) {

}

func (e *PluginObj) OnLoad() error {
	return nil
}

func (e *PluginObj) GetInfo() *types.Info {
	return e.Info
}

func (e *PluginObj) Router() []*types.PluginRoute {
	return e.Routes
}
