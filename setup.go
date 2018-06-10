package main

import (
	"net/http"
	"strconv"
	"os"
	"github.com/go-yaml/yaml"
	"time"
)

type DbConfig struct {
	DbConn  string `yaml:"connection"`
	DbHost  string `yaml:"host"`
	DbUser  string `yaml:"user"`
	DbPass  string `yaml:"password"`
	DbData  string `yaml:"database"`
	DbPort  int `yaml:"port"`
	Project  string `yaml:"-"`
	Username  string `yaml:"-"`
	Password  string `yaml:"-"`
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

	config := &DbConfig{
		dbConn,
		dbHost,
		dbUser,
		dbPass,
		dbDatabase,
		dbPort,
		project,
		username,
		password,
	}

	err := config.Save()
	if err != nil {
		panic(err)
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

	configs = LoadConfig()

	DbConnection()
	DropDatabase()
	CreateDatabase()
	db.QueryRow("INSERT INTO core (name, config, api_key, api_secret, version) VALUES($1,$2,$3,$4,$5);", c.Project, "config.yml", NewSHA1Hash(5), NewSHA1Hash(10), VERSION).Scan()
	return err
}
