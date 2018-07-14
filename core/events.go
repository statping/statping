package core

import (
	"github.com/fatih/structs"
	"github.com/hunterlong/statup/notifiers"
	"github.com/hunterlong/statup/plugin"
	"github.com/hunterlong/statup/types"
	"upper.io/db.v3/lib/sqlbuilder"
)

func OnLoad(db sqlbuilder.Database) {
	for _, p := range CoreApp.AllPlugins {
		p.OnLoad(db)
	}
}

func OnSuccess(s *types.Service) {
	for _, p := range CoreApp.AllPlugins {
		p.OnSuccess(structs.Map(s))
	}
	notifiers.OnSuccess(s)
}

func OnFailure(s *types.Service, f FailureData) {
	for _, p := range CoreApp.AllPlugins {
		p.OnFailure(structs.Map(s))
	}
	notifiers.OnFailure(s)
}

func OnSettingsSaved(c *Core) {
	for _, p := range CoreApp.AllPlugins {
		p.OnSettingsSaved(structs.Map(c))
	}
}

func OnNewUser(u *types.User) {
	for _, p := range CoreApp.AllPlugins {
		p.OnNewUser(structs.Map(u))
	}
}

func OnNewService(s *types.Service) {
	for _, p := range CoreApp.AllPlugins {
		p.OnNewService(structs.Map(s))
	}
}

func OnDeletedService(s *types.Service) {
	for _, p := range CoreApp.AllPlugins {
		p.OnDeletedService(structs.Map(s))
	}
}

func OnUpdateService(s *types.Service) {
	for _, p := range CoreApp.AllPlugins {
		p.OnUpdatedService(structs.Map(s))
	}
}

func SelectPlugin(name string) plugin.PluginActions {
	for _, p := range CoreApp.AllPlugins {
		if p.GetInfo().Name == name {
			return p
		}
	}
	return plugin.PluginInfo{}
}
