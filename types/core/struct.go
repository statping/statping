package core

import (
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/utils"
	"time"
)

var (
	App *Core
)

func New(version string) {
	App = new(Core)
	App.Version = version
	App.Started = utils.Now()
}

// Core struct contains all the required fields for Statping. All application settings
// will be saved into 1 row in the 'core' table. You can use the core.CoreApp
// global variable to interact with the attributes to the application, such as services.
type Core struct {
	Name          string          `gorm:"not null;column:name" json:"name,omitempty"`
	Description   string          `gorm:"not null;column:description" json:"description,omitempty"`
	ConfigFile    string          `gorm:"column:config" json:"-"`
	ApiSecret     string          `gorm:"column:api_secret" json:"api_secret" scope:"admin"`
	Style         string          `gorm:"not null;column:style" json:"style,omitempty"`
	Footer        null.NullString `gorm:"column:footer" json:"footer"`
	Domain        string          `gorm:"not null;column:domain" json:"domain"`
	Version       string          `gorm:"column:version" json:"version"`
	Setup         bool            `gorm:"-" json:"setup"`
	MigrationId   int64           `gorm:"column:migration_id" json:"migration_id,omitempty"`
	UseCdn        null.NullBool   `gorm:"column:use_cdn;default:false" json:"using_cdn,omitempty"`
	LoggedIn      bool            `gorm:"-" json:"logged_in"`
	IsAdmin       bool            `gorm:"-" json:"admin"`
	AllowReports  null.NullBool   `gorm:"column:allow_reports;default:false" json:"allow_reports,omitempty"`
	CreatedAt     time.Time       `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time       `gorm:"column:updated_at" json:"updated_at"`
	Started       time.Time       `gorm:"-" json:"started_on"`
	Notifications []AllNotifiers  `gorm:"-" json:"-"`
	Integrations  []Integrator    `gorm:"-" json:"-"`

	OAuth `json:"-"`
}

type OAuth struct {
	Domains            string `gorm:"column:oauth_domains" json:"oauth_domains" scope:"admin"`
	Providers          string `gorm:"column:oauth_providers;" json:"oauth_providers"`
	GithubClientID     string `gorm:"column:gh_client_id" json:"gh_client_id"`
	GithubClientSecret string `gorm:"column:gh_client_secret" json:"gh_client_secret" scope:"admin"`
	GoogleClientID     string `gorm:"column:google_client_id" json:"google_client_id"`
	GoogleClientSecret string `gorm:"column:google_client_secret" json:"google_client_secret" scope:"admin"`
	SlackClientID      string `gorm:"column:slack_client_id" json:"slack_client_id"`
	SlackClientSecret  string `gorm:"column:slack_client_secret" json:"slack_client_secret" scope:"admin"`
	SlackTeam          string `gorm:"column:slack_team" json:"slack_team"`
}

// AllNotifiers contains all the Notifiers loaded
type AllNotifiers interface{}

type Integrator interface{}

func (Core) TableName() string {
	return "core"
}
