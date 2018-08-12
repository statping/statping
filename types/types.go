package types

import (
	"net/http"
	"time"
	"upper.io/db.v3/lib/sqlbuilder"
)

type PluginInfo struct {
	Info Info
	PluginActions
}

type Routing struct {
	URL     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

type Info struct {
	Name        string
	Description string
	Form        string
}

type PluginActions interface {
	GetInfo() Info
	GetForm() string
	OnLoad(sqlbuilder.Database)
	SetInfo(map[string]interface{}) Info
	Routes() []Routing
	OnSave(map[string]interface{})
	OnFailure(map[string]interface{})
	OnSuccess(map[string]interface{})
	OnSettingsSaved(map[string]interface{})
	OnNewUser(map[string]interface{})
	OnNewService(map[string]interface{})
	OnUpdatedService(map[string]interface{})
	OnDeletedService(map[string]interface{})
	OnInstall(map[string]interface{})
	OnUninstall(map[string]interface{})
	OnBeforeRequest(map[string]interface{})
	OnAfterRequest(map[string]interface{})
	OnShutdown()
}

type AllNotifiers interface{}

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
	Services       []*Service `json:"services,omitempty"`
	Plugins        []Info
	Repos          []PluginJSON
	AllPlugins     []PluginActions
	Communications []AllNotifiers
	DbConnection   string
	Started        time.Time
}

type User struct {
	Id        int64     `db:"id,omitempty" json:"id"`
	Username  string    `db:"username" json:"username"`
	Password  string    `db:"password" json:"-"`
	Email     string    `db:"email" json:"-"`
	ApiKey    string    `db:"api_key" json:"api_key"`
	ApiSecret string    `db:"api_secret" json:"-"`
	Admin     bool      `db:"administrator" json:"admin"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Hit struct {
	Id        int       `db:"id,omitempty"`
	Service   int64     `db:"service"`
	Latency   float64   `db:"latency"`
	CreatedAt time.Time `db:"created_at"`
}

type Failure struct {
	Id        int       `db:"id,omitempty" json:"id"`
	Issue     string    `db:"issue" json:"issue"`
	Method    string    `db:"method" json:"method,omitempty"`
	Service   int64     `db:"service" json:"service_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Checkin struct {
	Id        int       `db:"id,omitempty"`
	Service   int64     `db:"service"`
	Interval  int64     `db:"check_interval"`
	Api       string    `db:"api"`
	CreatedAt time.Time `db:"created_at"`
	Hits      int64     `json:"hits"`
	Last      time.Time `json:"last"`
}

type Config struct {
	Connection string `yaml:"connection"`
	Host       string `yaml:"host"`
	Database   string `yaml:"database"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Port       string `yaml:"port"`
	Secret     string `yaml:"secret"`
}

type DbConfig struct {
	DbConn      string `yaml:"connection"`
	DbHost      string `yaml:"host"`
	DbUser      string `yaml:"user"`
	DbPass      string `yaml:"password"`
	DbData      string `yaml:"database"`
	DbPort      int    `yaml:"port"`
	Project     string `yaml:"-"`
	Description string `yaml:"-"`
	Domain      string `yaml:"-"`
	Username    string `yaml:"-"`
	Password    string `yaml:"-"`
	Email       string `yaml:"-"`
	Error       error  `yaml:"-"`
	Location    string `yaml:"location"`
}

type PluginRepos struct {
	Plugins []PluginJSON
}

type PluginJSON struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Repo        string `json:"repo"`
	Author      string `json:"author"`
	Namespace   string `json:"namespace"`
}

type FailureData struct {
	Issue string
}
