package main

import (
	"github.com/gorilla/sessions"
	"github.com/hunterlong/statup/plugin"
	"github.com/hunterlong/statup/types"
)

type Core struct {
	Name           string `db:"name"`
	Description    string `db:"description"`
	Config         string `db:"config"`
	ApiKey         string `db:"api_key"`
	ApiSecret      string `db:"api_secret"`
	Style          string `db:"style"`
	Footer         string `db:"footer"`
	Domain         string `db:"domain"`
	Version        string `db:"version"`
	Plugins        []plugin.Info
	Repos          []PluginJSON
	PluginFields   []PluginSelect
	Communications []*types.Communication
	OfflineAssets  bool
}

func (c *Core) Update() (*Core, error) {
	res := dbSession.Collection("core").Find().Limit(1)
	res.Update(c)
	core = c
	return c, nil
}

func (c Core) UsingAssets() bool {
	return useAssets
}

func (c Core) SassVars() string {
	if !useAssets {
		return ""
	}
	return OpenAsset("scss/_variables.scss")
}

func (c Core) BaseSASS() string {
	if !useAssets {
		return ""
	}
	return OpenAsset("scss/base.scss")
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
	var c *Core
	err := dbSession.Collection("core").Find().One(&c)
	if err != nil {
		return nil, err
	}
	core = c
	store = sessions.NewCookieStore([]byte(core.ApiSecret))
	return core, err
}
