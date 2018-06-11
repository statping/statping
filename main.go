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
	"github.com/hunterlong/statup/plugin"
	plg "plugin"
	"strings"
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
	allPlugins []*plugin.Plugin
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


type Greeter interface {
	Greet()
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
		LoadPlugins()
		RunHTTPServer()
	}
}


func LoadPlugins() {
	files, err := ioutil.ReadDir("./plugins")
	if err != nil {
		fmt.Printf("Plugins directory was not found. Error: %v\n", err)
		return
	}

	for _, f := range files {

		ext := strings.Split(f.Name(), ".")
		if len(ext) != 2 {
			continue
		}

		if ext[1] == "so" {

			plug, err := plg.Open("plugins/"+f.Name())
			if err != nil {
				fmt.Printf("Plugin '%v' could not load correctly.\n", f.Name())
				continue
			}

			symPlugin, err := plug.Lookup("Plugin")
			var plugActions plugin.PluginActions

			plugActions, ok := symPlugin.(plugin.PluginActions)
			if !ok {
				fmt.Printf("Plugin '%v' could not load correctly, error: %v\n", f.Name(), "unexpected type from module symbol")
				continue
			}
			plugin := plugActions.Plugin()

			fmt.Printf("Plugin Loaded '%v' created by: %v\n", plugin.Name, plugin.Creator)
			plugActions.OnLoad()

			fmt.Println(plugActions.Form())

			allPlugins = append(allPlugins, plugin)

		}

	}

	core.Plugins = allPlugins

	fmt.Printf("Loaded %v Plugins\n", len(allPlugins))
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
	return &config
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
