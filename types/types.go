package types

import (
	"time"
)

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
	Id        int       `db:"id,omitempty"`
	Issue     string    `db:"issue"`
	Method    string    `db:"method"`
	Service   int64     `db:"service"`
	CreatedAt time.Time `db:"created_at"`
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

type Communication struct {
	Id        int64     `db:"id,omitempty" json:"id"`
	Method    string    `db:"method" json:"method"`
	Host      string    `db:"host" json:"host"`
	Port      int       `db:"port" json:"port"`
	Username  string    `db:"username" json:"user"`
	Password  string    `db:"password" json:"-"`
	Var1      string    `db:"var1" json:"var1"`
	Var2      string    `db:"var2" json:"var2"`
	ApiKey    string    `db:"api_key" json:"api_key"`
	ApiSecret string    `db:"api_secret" json:"api_secret"`
	Enabled   bool      `db:"enabled" json:"enabled"`
	Limits    int64     `db:"limits" json:"limits"`
	Removable bool      `db:"removable" json:"removable"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Email struct {
	To       string
	Subject  string
	Template string
	Data     interface{}
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
