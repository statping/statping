package core

import (
	"github.com/statping/statping/types/null"
	"time"
)

var (
	App *Core
)

func New(version string) {
	App = new(Core)
	App.Version = version
}

// Core struct contains all the required fields for Statping. All application settings
// will be saved into 1 row in the 'core' table. You can use the core.CoreApp
// global variable to interact with the attributes to the application, such as services.
type Core struct {
	Name          string          `gorm:"not null;column:name" json:"name"`
	Description   string          `gorm:"not null;column:description" json:"description,omitempty"`
	ConfigFile    string          `gorm:"column:config" json:"-"`
	ApiKey        string          `gorm:"column:api_key" json:"api_key" scope:"admin"`
	ApiSecret     string          `gorm:"column:api_secret" json:"api_secret" scope:"admin"`
	Style         string          `gorm:"not null;column:style" json:"style,omitempty"`
	Footer        null.NullString `gorm:"column:footer" json:"footer"`
	Domain        string          `gorm:"not null;column:domain" json:"domain"`
	Version       string          `gorm:"column:version" json:"version"`
	Setup         bool            `gorm:"-" json:"setup"`
	MigrationId   int64           `gorm:"column:migration_id" json:"migration_id,omitempty"`
	UseCdn        null.NullBool   `gorm:"column:use_cdn;default:false" json:"using_cdn,omitempty"`
	Timezone      float32         `gorm:"column:timezone;default:-8.0" json:"timezone,omitempty"`
	LoggedIn      bool            `gorm:"-" json:"logged_in"`
	IsAdmin       bool            `gorm:"-" json:"admin"`
	CreatedAt     time.Time       `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time       `gorm:"column:updated_at" json:"updated_at"`
	Started       time.Time       `gorm:"-" json:"started_on"`
	Image         string          `gorm:"image" json:"started_on"`
	Notifications []AllNotifiers  `gorm:"-" json:"-"`
	Integrations  []Integrator    `gorm:"-" json:"-"`

	GHAuth
}

type GHAuth struct {
	GithubClientID     string `gorm:"gh_client_id" json:"gh_client_id"`
	GithubClientSecret string `gorm:"gh_client_secret" json:"-"`
}

// AllNotifiers contains all the Notifiers loaded
type AllNotifiers interface{}

type Integrator interface{}

func (Core) TableName() string {
	return "core"
}
