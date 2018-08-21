package types

import "time"

type Core struct {
	Name           string          `db:"name" json:"name"`
	Description    string          `db:"description" json:"description,omitempty"`
	Config         string          `db:"config" json:"-"`
	ApiKey         string          `db:"api_key" json:"-"`
	ApiSecret      string          `db:"api_secret" json:"-"`
	Style          string          `db:"style" json:"style,omitempty"`
	Footer         string          `db:"footer" json:"footer,omitempty"`
	Domain         string          `db:"domain" json:"domain,omitempty"`
	Version        string          `db:"version" json:"version"`
	MigrationId    int64           `db:"migration_id" json:"migration_id,omitempty"`
	UseCdn         bool            `db:"use_cdn" json:"using_cdn,omitempty"`
	DbConnection   string          `json:"database"`
	Started        time.Time       `json:"started_on"`
	dbServices     []*Service      `json:"services,omitempty"`
	Plugins        []Info          `json:"-"`
	Repos          []PluginJSON    `json:"-"`
	AllPlugins     []PluginActions `json:"-"`
	Communications []AllNotifiers  `json:"-"`
	CoreInterface  `json:"-"`
}

func (c *Core) SetServices(s []*Service) {
	c.dbServices = s
}

func (c *Core) UpdateService(index int, s *Service) {
	c.dbServices[index] = s
}

func (c *Core) AddService(s *Service) {
	c.dbServices = append(c.dbServices, s)
}

func (c *Core) RemoveService(s int) []*Service {
	slice := c.dbServices
	return append(slice[:s], slice[s+1:]...)
}

func (c *Core) GetServices() []*Service {
	return c.dbServices
}

type CoreInterface interface {
	SelectAllServices() ([]*Service, error)
	Services() []*Service
}
