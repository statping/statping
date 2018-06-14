package main

import (
	"github.com/gorilla/sessions"
	"github.com/hunterlong/statup/plugin"
)

type Core struct {
	Name         string
	Config       string
	Key          string
	Secret       string
	Version      string
	Plugins      []plugin.Info
	Repos        []PluginJSON
	PluginFields []PluginSelect
}


func SelectCore() (*Core, error) {
	var core Core
	rows, err := db.Query("SELECT * FROM core")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&core.Name, &core.Config, &core.Key, &core.Secret, &core.Version)
		if err != nil {
			return nil, err
		}
	}
	store = sessions.NewCookieStore([]byte(core.Secret))
	return &core, err
}
