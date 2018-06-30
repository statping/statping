package core

import (
	"github.com/GeertJohan/go.rice"
	"github.com/hunterlong/statup/plugin"
	"github.com/hunterlong/statup/types"
)

type PluginJSON types.PluginJSON
type PluginRepos types.PluginRepos

type Core struct {
	Name        string `db:"name"`
	Description string `db:"description"`
	Config      string `db:"config"`
	ApiKey      string `db:"api_key"`
	ApiSecret   string `db:"api_secret"`
	Style       string `db:"style"`
	Footer      string `db:"footer"`
	Domain      string `db:"domain"`
	Version     string `db:"version"`
	Services    []*Service
	Plugins     []plugin.Info
	Repos       []PluginJSON
	//PluginFields   []PluginSelect
	Communications []*Communication
	OfflineAssets  bool
}

var (
	Configs     *Config
	CoreApp     *Core
	SqlBox      *rice.Box
	CssBox      *rice.Box
	ScssBox     *rice.Box
	JsBox       *rice.Box
	TmplBox     *rice.Box
	EmailBox    *rice.Box
	SetupMode   bool
	AllPlugins  []plugin.PluginActions
	UsingAssets bool
)

func init() {
	CoreApp = new(Core)
}

func (c *Core) Update() (*Core, error) {
	res := DbSession.Collection("core").Find().Limit(1)
	res.Update(c)
	CoreApp = c
	return c, nil
}

func (c Core) UsingAssets() bool {
	return UsingAssets
}

func (c Core) SassVars() string {
	if !UsingAssets {
		return ""
	}
	return OpenAsset("scss/variables.scss")
}

func (c Core) BaseSASS() string {
	if !UsingAssets {
		return ""
	}
	return OpenAsset("scss/base.scss")
}

func (c Core) AllOnline() bool {
	for _, s := range CoreApp.Services {
		if !s.Online {
			return false
		}
	}
	return true
}

func SelectCore() (*Core, error) {
	var c *Core
	err := DbSession.Collection("core").Find().One(&c)
	if err != nil {
		return nil, err
	}
	CoreApp = c
	//store = sessions.NewCookieStore([]byte(core.ApiSecret))
	return CoreApp, err
}
