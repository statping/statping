package types

import "time"

type Core struct {
	Name           string     `db:"name" json:"name"`
	Description    string     `db:"description" json:"description"`
	Config         string     `db:"config" json:"-"`
	ApiKey         string     `db:"api_key" json:"api_key"`
	ApiSecret      string     `db:"api_secret" json:"api_secret"`
	Style          string     `db:"style" json:"-"`
	Footer         string     `db:"footer" json:"-"`
	Domain         string     `db:"domain" json:"domain,omitempty"`
	Version        string     `db:"version" json:"version,omitempty"`
	MigrationId    int64      `db:"migration_id" json:"-"`
	UseCdn         bool       `db:"use_cdn" json:"-"`
	DbServices     []*Service `json:"services,omitempty"`
	Plugins        []Info
	Repos          []PluginJSON
	AllPlugins     []PluginActions
	Communications []AllNotifiers
	DbConnection   string
	Started        time.Time
	CoreInterface
}

type CoreInterface interface {
	SelectAllServices() ([]*Service, error)
	Services() []*Service
}
