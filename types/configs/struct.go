package configs

import "github.com/statping/statping/database"

const SqliteFilename = "statping.db"

// DbConfig struct is used for the Db connection and creates the 'config.yml' file
type DbConfig struct {
	DbConn            string `yaml:"connection,omitempty" json:"connection"`
	DbHost            string `yaml:"host,omitempty" json:"-"`
	DbUser            string `yaml:"user,omitempty" json:"-"`
	DbPass            string `yaml:"password,omitempty" json:"-"`
	DbData            string `yaml:"database,omitempty" json:"-"`
	DbPort            int    `yaml:"port,omitempty" json:"-"`
	ApiSecret         string `yaml:"api_secret,omitempty" json:"-"`
	Language          string `yaml:"language,omitempty" json:"language"`
	AllowReports      bool   `yaml:"allow_reports,omitempty" json:"allow_reports"`
	Project           string `yaml:"-" json:"-"`
	Description       string `yaml:"-" json:"-"`
	Domain            string `yaml:"-" json:"-"`
	Username          string `yaml:"-" json:"-"`
	Password          string `yaml:"-" json:"-"`
	Email             string `yaml:"-" json:"-"`
	Error             error  `yaml:"-" json:"-"`
	Location          string `yaml:"location,omitempty" json:"-"`
	SqlFile           string `yaml:"sqlfile,omitempty" json:"-"`
	LetsEncryptHost   string `yaml:"letsencrypt_host,omitempty" json:"letsencrypt_host"`
	LetsEncryptEmail  string `yaml:"letsencrypt_email,omitempty" json:"letsencrypt_email"`
	LetsEncryptEnable bool   `yaml:"letsencrypt_enable,omitempty" json:"letsencrypt_enable"`
	LocalIP           string `yaml:"-" json:"-"`

	DisableHTTP bool   `yaml:"disable_http" json:"disable_http"`
	DemoMode    bool   `yaml:"demo_mode" json:"demo_mode"`
	DisableLogs bool   `yaml:"disable_logs" json:"disable_logs"`
	UseAssets   bool   `yaml:"use_assets" json:"use_assets"`
	BasePath    string `yaml:"base_path,omitempty" json:"base_path"`

	AdminUser     string `yaml:"admin_user,omitempty" json:"admin_user"`
	AdminPassword string `yaml:"admin_password,omitempty" json:"admin_password"`
	AdminEmail    string `yaml:"admin_email,omitempty" json:"admin_email"`

	MaxOpenConnections int `yaml:"db_open_connections,omitempty" json:"db_open_connections"`
	MaxIdleConnections int `yaml:"db_idle_connections,omitempty" json:"db_idle_connections"`
	MaxLifeConnections int `yaml:"db_max_life_connections,omitempty" json:"db_max_life_connections"`

	SampleData    bool `yaml:"sample_data" json:"sample_data"`
	UseCDN        bool `yaml:"use_cdn" json:"use_cdn"`
	DisableColors bool `yaml:"disable_colors" json:"disable_colors"`

	PostgresSSLMode string `yaml:"postgres_ssl,omitempty" json:"postgres_ssl"`

	Db database.Database `yaml:"-" json:"-"`
}
