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
	Style        string `db:"style"`
	Footer       string `db:"footer"`
	Domain       string `db:"domain"`
	Version      string `db:"version"`
	Plugins      []plugin.Info
	Repos        []PluginJSON
	PluginFields []PluginSelect
}

func (c *Core) Update() (*Core, error) {
	res := dbSession.Collection("core").Find().Limit(1)
	res.Update(c)
	core = c
	return c, nil
}

func (c Core) AllOnline() bool {
	for _, s := range services {
		if !s.Online {
			return false
		}
	}
	return true
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
