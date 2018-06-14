package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/go-yaml/yaml"
	"github.com/gorilla/sessions"
	"github.com/hunterlong/statup/plugin"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	plg "plugin"
	"strings"
)

var (
	db         *sql.DB
	configs    *Config
	core       *Core
	store      *sessions.CookieStore
	VERSION    string
	sqlBox     *rice.Box
	cssBox     *rice.Box
	jsBox      *rice.Box
	tmplBox    *rice.Box
	setupMode  bool
	allPlugins []plugin.PluginActions
)

const (
	pluginsRepo = "https://raw.githubusercontent.com/hunterlong/statup/master/plugins.json"
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

func (c *Core) FetchPluginRepo() []PluginJSON {
	resp, err := http.Get(pluginsRepo)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var pk []PluginJSON
	json.Unmarshal(body, &pk)
	c.Repos = pk
	return pk
}

func DownloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func SelectSettings(p plugin.Info) map[string]string {

	data := make(map[string]string)
	var tableInput []string

	for _, v := range p.Form {
		val := fmt.Sprintf("%v", v.InputName)
		tableInput = append(tableInput, val)
	}

	ins := strings.Join(tableInput, ", ")

	sql := fmt.Sprintf("SELECT %v FROM settings_%v LIMIT 1", ins, p.Name)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("SQL ERROR: ", err)
		return map[string]string{}
	}

	count := len(p.Form)
	valuePtrs := make([]interface{}, count)
	values := make([]interface{}, count)

	for rows.Next() {

		for i, _ := range p.Form {
			valuePtrs[i] = &values[i]
		}

		err = rows.Scan(valuePtrs...)
		if err != nil {
			panic(err)
		}

		for i, col := range p.Form {

			val := values[i]

			b, _ := val.([]byte)

			realVal := string(b)

			//col.ChangeVal(realVal)

			fmt.Println(col.Value, realVal)

			data[col.InputName] = realVal
		}

	}

	return data
}

func CreateSettingsTable(p plugin.Info) string {
	var tableValues []string
	tableValues = append(tableValues, "plugin text")
	tableValues = append(tableValues, "enabled bool")
	for _, v := range p.Form {
		tb := fmt.Sprintf("%v %v", v.InputName, v.InputType)
		tableValues = append(tableValues, tb)
	}
	vals := strings.Join(tableValues, ", ")
	out := fmt.Sprintf("CREATE TABLE settings_%v (%v);", p.Name, vals)
	smtp, _ := db.Prepare(out)
	_, _ = smtp.Exec()
	InitalSettings(p)
	return out
}

func InitalSettings(p plugin.Info) {
	var tableValues []string
	var tableInput []string

	tableValues = append(tableValues, "plugin")
	tableInput = append(tableInput, fmt.Sprintf("'%v'", p.Name))

	tableValues = append(tableValues, "enabled")
	tableInput = append(tableInput, "false")

	for _, v := range p.Form {
		val := fmt.Sprintf("'%v'", v.Value)
		tableValues = append(tableValues, v.InputName)
		tableInput = append(tableInput, val)
	}

	vals := strings.Join(tableValues, ",")
	ins := strings.Join(tableInput, ",")
	sql := fmt.Sprintf("INSERT INTO settings_%v(%v) VALUES(%v);", p.Name, vals, ins)
	smtp, _ := db.Prepare(sql)
	_, _ = smtp.Exec()
}

func UpdateSettings(p plugin.Info, data map[string]string) {
	var tableInput []string

	for _, v := range p.Form {
		newValue := data[v.InputName]
		val := fmt.Sprintf("%v='%v'", v.InputName, newValue)
		tableInput = append(tableInput, val)
	}

	ins := strings.Join(tableInput, ", ")
	sql := fmt.Sprintf("UPDATE settings_%v SET %v WHERE plugin='%v';", p.Name, ins, p.Name)
	smtp, _ := db.Prepare(sql)
	_, _ = smtp.Exec()
}

//func DownloadPlugin(name string) {
//	plugin := SelectPlugin(name)
//	var _, err = os.Stat("plugins/" + plugin.Namespace)
//	if err != nil {
//	}
//	if os.IsNotExist(err) {
//		var file, _ = os.Create("plugins/" + plugin.Namespace)
//		defer file.Close()
//	}
//	resp, err := http.Get("https://raw.githubusercontent.com/hunterlong/statup/master/plugins.json")
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		panic(err)
//	}
//	file, err := os.OpenFile("plugins/"+plugin.Namespace, os.O_RDWR, 0644)
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//
//	_, err = file.Write(body)
//	if err != nil {
//		panic(err)
//	}
//	err = file.Sync()
//}

func main() {
	var err error
	VERSION = "1.1.1"
	fmt.Printf("Starting Statup v%v\n", VERSION)
	RenderBoxes()
	configs, err = LoadConfig()
	if err != nil {
		fmt.Println("config.yml file not found - starting in setup mode")
		setupMode = true
		RunHTTPServer()
	}
	mainProcess()
}

func mainProcess() {
	var err error
	err = DbConnection(configs.Connection)
	if err != nil {
		throw(err)
	}
	core, err = SelectCore()
	if err != nil {
		throw(err)
	}
	go CheckServices()
	if !setupMode {
		LoadPlugins()
		RunHTTPServer()
	}
}

func throw(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func ForEachPlugin() {
	if len(core.Plugins) > 0 {
		//for _, p := range core.Plugins {
		//	p.OnShutdown()
		//}
	}
}

func LoadPlugins() {
	if _, err := os.Stat("./plugins"); os.IsNotExist(err) {
		os.Mkdir("./plugins", os.ModePerm)
	}

	ForEachPlugin()

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
		if ext[1] != "so" {
			continue
		}
		plug, err := plg.Open("plugins/" + f.Name())
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

		plugActions.OnLoad()

		allPlugins = append(allPlugins, plugActions)
		core.Plugins = append(core.Plugins, plugActions.GetInfo())
	}

	fmt.Printf("Loaded %v Plugins\n", len(allPlugins))

	ForEachPlugin()
}

func RenderBoxes() {
	sqlBox = rice.MustFindBox("sql")
	cssBox = rice.MustFindBox("html/css")
	jsBox = rice.MustFindBox("html/js")
	tmplBox = rice.MustFindBox("html/tmpl")
}

func LoadConfig() (*Config, error) {
	var config Config
	file, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &config)
	return &config, err
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
