package main

import (
	"github.com/gorilla/sessions"
	"github.com/hunterlong/statup/plugin"
)

type Core struct {
	Name         string `db:"name"`
	Description  string `db:"description"`
	Config       string `db:"config"`
	ApiKey       string `db:"api_key"`
	ApiSecret    string `db:"api_secret"`
	Version      string `db:"version"`
	Plugins      []plugin.Info
	Repos        []PluginJSON
	PluginFields []PluginSelect
}

func SelectCore() (*Core, error) {
	var core Core
	err := dbSession.Collection("core").Find().One(&core)
	if err != nil {
		return nil, err
	}
	store = sessions.NewCookieStore([]byte(core.ApiSecret))
	return &core, err
}
