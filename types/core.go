package types

import (
	"time"
)

type Core struct {
	Name           string          `gorm:"column:name" json:"name"`
	Description    string          `gorm:"column:description" json:"description,omitempty"`
	Config         string          `gorm:"column:config" json:"-"`
	ApiKey         string          `gorm:"column:api_key" json:"-"`
	ApiSecret      string          `gorm:"column:api_secret" json:"-"`
	Style          string          `gorm:"column:style" json:"style,omitempty"`
	Footer         string          `gorm:"column:footer" json:"footer,omitempty"`
	Domain         string          `gorm:"column:domain" json:"domain,omitempty"`
	Version        string          `gorm:"column:version" json:"version"`
	MigrationId    int64           `gorm:"column:migration_id" json:"migration_id,omitempty"`
	UseCdn         bool            `gorm:"column:use_cdn" json:"using_cdn,omitempty"`
	DbConnection   string          `gorm:"-" json:database"`
	Started        time.Time       `gorm:"-" json:started_on"`
	dbServices     []*Service      `gorm:"-" json:services,omitempty"`
	Plugins        []Info          `gorm:"-" json:-"`
	Repos          []PluginJSON    `gorm:"-" json:-"`
	AllPlugins     []PluginActions `gorm:"-" json:-"`
	Communications []AllNotifiers  `gorm:"-" json:-"`
	CoreInterface  `gorm:"-" json:-"`
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
	c.dbServices = append(slice[:s], slice[s+1:]...)
	return c.dbServices
}

func (c *Core) GetServices() []*Service {
	return c.dbServices
}

type CoreInterface interface {
	SelectAllServices() ([]*Service, error)
	Services() []*Service
}
