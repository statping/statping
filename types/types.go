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

type Email struct {
	To       string
	Subject  string
	Template string
	From     string
	Data     interface{}
	Source   string
	Sent     bool
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
