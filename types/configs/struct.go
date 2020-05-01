package configs

import "github.com/statping/statping/database"

const SqliteFilename = "statping.db"

// DbConfig struct is used for the Db connection and creates the 'config.yml' file
type DbConfig struct {
	DbConn      string `yaml:"connection" json:"connection"`
	DbHost      string `yaml:"host" json:"-"`
	DbUser      string `yaml:"user" json:"-"`
	DbPass      string `yaml:"password" json:"-"`
	DbData      string `yaml:"database" json:"-"`
	DbPort      int    `yaml:"port" json:"-"`
	ApiSecret   string `yaml:"api_secret" json:"-"`
	Project     string `yaml:"-" json:"-"`
	Description string `yaml:"-" json:"-"`
	Domain      string `yaml:"-" json:"-"`
	Username    string `yaml:"-" json:"-"`
	Password    string `yaml:"-" json:"-"`
	Email       string `yaml:"-" json:"-"`
	Error       error  `yaml:"-" json:"-"`
	Location    string `yaml:"location" json:"-"`
	SqlFile     string `yaml:"sqlfile,omitempty" json:"-"`
	LocalIP     string `yaml:"-" json:"-"`
	filename    string `yaml:"-" json:"-"`

	Db database.Database `yaml:"-" json:"-"`
}
