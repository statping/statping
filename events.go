package main

import (
	"github.com/fatih/structs"
	"github.com/hunterlong/statup/plugin"
	"upper.io/db.v3/lib/sqlbuilder"
)

func OnLoad(db sqlbuilder.Database) {
	for _, p := range allPlugins {
		p.OnLoad(db)
	}
}

func OnSuccess(s *Service) {
	for _, p := range allPlugins {
		p.OnSuccess(structs.Map(s))
	}
}

func OnFailure(s *Service) {
	for _, p := range allPlugins {
		p.OnFailure(structs.Map(s))
	}
}

func OnSettingsSaved(c *Core) {
	for _, p := range allPlugins {
		p.OnSettingsSaved(structs.Map(c))
	}
}

func OnNewUser(u *User) {
	for _, p := range allPlugins {
		p.OnNewUser(structs.Map(u))
	}
}

func OnNewService(s *Service) {
	for _, p := range allPlugins {
		p.OnNewService(structs.Map(s))
	}
}

func OnDeletedService(s *Service) {
	for _, p := range allPlugins {
		p.OnDeletedService(structs.Map(s))
	}
}

func OnUpdateService(s *Service) {
	for _, p := range allPlugins {
		p.OnUpdatedService(structs.Map(s))
	}
}

func SelectPlugin(name string) plugin.PluginActions {
	for _, p := range allPlugins {
		if p.GetInfo().Name == name {
			return p
		}
	}
	return plugin.PluginInfo{}
}
