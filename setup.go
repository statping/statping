package main

import (
	"github.com/go-yaml/yaml"
	"github.com/hunterlong/statup/plugin"
	"net/http"
	"os"
	"strconv"
	"time"
)

type DbConfig struct {
	DbConn      string `yaml:"connection"`
	DbHost      string `yaml:"host"`
	DbUser      string `yaml:"user"`
	DbPass      string `yaml:"password"`
	DbData      string `yaml:"database"`
	DbPort      int    `yaml:"port"`
	Project     string `yaml:"-"`
	Description string `yaml:"-"`
	Username    string `yaml:"-"`
	Password    string `yaml:"-"`
}

func ProcessSetupHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dbHost := r.PostForm.Get("db_host")
	dbUser := r.PostForm.Get("db_user")
	dbPass := r.PostForm.Get("db_password")
	dbDatabase := r.PostForm.Get("db_database")
	dbConn := r.PostForm.Get("db_connection")
	dbPort, _ := strconv.Atoi(r.PostForm.Get("db_port"))
	project := r.PostForm.Get("project")
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	sample := r.PostForm.Get("sample_data")
	description := r.PostForm.Get("description")

	config := &DbConfig{
		dbConn,
		dbHost,
		dbUser,
		dbPass,
		dbDatabase,
		dbPort,
		project,
		description,
		username,
		password,
	}
	err := config.Save()
	if err != nil {
		throw(err)
	}

	configs, err = LoadConfig()
	if err != nil {
		throw(err)
	}

	err = DbConnection(configs.Connection)
	if err != nil {
		throw(err)
	}

	if sample == "on" {
		LoadSampleData()
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	time.Sleep(2 * time.Second)
	mainProcess()
}

func (c *DbConfig) Save() error {
	var err error
	config, err := os.Create("config.yml")
	if err != nil {
		return err
	}
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	config.WriteString(string(data))
	config.Close()

	configs, err = LoadConfig()
	if err != nil {
		return err
	}
	err = DbConnection(configs.Connection)
	if err != nil {
		return err
	}
	DropDatabase()
	CreateDatabase()

	newCore := Core{
		c.Project,
		c.Description,
		"config.yml",
		NewSHA1Hash(5),
		NewSHA1Hash(10),
		VERSION,
		[]plugin.Info{},
		[]PluginJSON{},
		[]PluginSelect{},
	}

	col := dbSession.Collection("core")
	_, err = col.Insert(newCore)

	return err
}
