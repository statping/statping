package main

import (
	"database/sql"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/go-yaml/yaml"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
)

var (
	db        *sql.DB
	configs   *Config
	core      *Core
	store     *sessions.CookieStore
	VERSION   string
	sqlBox    *rice.Box
	cssBox    *rice.Box
	jsBox     *rice.Box
	tmplBox   *rice.Box
	setupMode bool
)

type Config struct {
	Connection string `yaml:"connection"`
	Host       string `yaml:"host"`
	Database   string `yaml:"database"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Port       string `yaml:"port"`
	Secret     string `yaml:"secret"`
}

func main() {
	VERSION = "1.1.1"
	RenderBoxes()
	configs = LoadConfig()
	if configs == nil {
		fmt.Println("config.yml file not found - starting in setup mode")
		setupMode = true
		RunHTTPServer()
	}
	mainProcess()
}

func mainProcess() {
	var err error
	DbConnection()
	core, err = SelectCore()
	if err != nil {
		panic(err)
	}
	go CheckServices()
	if !setupMode {
		RunHTTPServer()
	}
}

func RenderBoxes() {
	sqlBox = rice.MustFindBox("sql")
	cssBox = rice.MustFindBox("html/css")
	jsBox = rice.MustFindBox("html/js")
	tmplBox = rice.MustFindBox("html/tmpl")
}

func LoadConfig() *Config {
	var config Config
	file, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return nil
	}
	yaml.Unmarshal(file, &config)
	store = sessions.NewCookieStore([]byte(config.Secret))
	return &config
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
