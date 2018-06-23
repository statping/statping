package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/hunterlong/statup/plugin"
	"github.com/hunterlong/statup/types"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	Domain      string `yaml:"-"`
	Username    string `yaml:"-"`
	Password    string `yaml:"-"`
	Error       error  `yaml:"-"`
}

func ProcessSetupHandler(w http.ResponseWriter, r *http.Request) {
	if core != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
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
	domain := r.PostForm.Get("domain")

	config := &DbConfig{
		dbConn,
		dbHost,
		dbUser,
		dbPass,
		dbDatabase,
		dbPort,
		project,
		description,
		domain,
		username,
		password,
		nil,
	}
	err := config.Save()
	if err != nil {
		config.Error = err
		SetupResponseError(w, r, config)
		return
	}

	configs, err = LoadConfig()
	if err != nil {
		config.Error = err
		SetupResponseError(w, r, config)
		return
	}

	err = DbConnection(configs.Connection)
	if err != nil {
		DeleteConfig()
		config.Error = err
		SetupResponseError(w, r, config)
		return
	}

	admin := &User{
		Username: config.Username,
		Password: config.Password,
		Admin:    true,
	}
	admin.Create()

	InsertDefaultComms()

	if sample == "on" {
		go LoadSampleData()
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	time.Sleep(2 * time.Second)
	mainProcess()
}

func InsertDefaultComms() {
	emailer := &types.Communication{
		Method:    "email",
		Removable: false,
		Enabled:   false,
	}
	Create(emailer)
}

func DeleteConfig() {
	err := os.Remove("./config.yml")
	if err != nil {
		throw(err)
	}
}

type ErrorResponse struct {
	Error string
}

func SetupResponseError(w http.ResponseWriter, r *http.Request, a interface{}) {
	ExecuteResponse(w, r, "setup.html", a)
}

func (c *DbConfig) Clean() *DbConfig {
	if os.Getenv("DB_PORT") != "" {
		if c.DbConn == "postgres" {
			c.DbHost = c.DbHost + ":" + os.Getenv("DB_PORT")
		}
	}
	return c
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
		"",
		"",
		"",
		VERSION,
		[]plugin.Info{},
		[]PluginJSON{},
		[]PluginSelect{},
		nil,
	}

	col := dbSession.Collection("core")
	_, err = col.Insert(newCore)

	return err
}

func DropDatabase() {
	fmt.Println("Dropping Tables...")
	down, _ := sqlBox.String("down.sql")
	requests := strings.Split(down, ";")
	for _, request := range requests {
		_, err := dbSession.Exec(request)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func CreateDatabase() {
	fmt.Println("Creating Tables...")
	sql := "postgres_up.sql"
	if dbServer == "mysql" {
		sql = "mysql_up.sql"
	} else if dbServer == "sqlite3" {
		sql = "sqlite_up.sql"
	}
	up, _ := sqlBox.String(sql)
	requests := strings.Split(up, ";")
	for _, request := range requests {
		_, err := dbSession.Exec(request)
		if err != nil {
			fmt.Println(err)
		}
	}
	//secret := NewSHA1Hash()
	//db.QueryRow("INSERT INTO core (secret, version) VALUES ($1, $2);", secret, VERSION).Scan()
	fmt.Println("Database Created")
	//SampleData()
}

func LoadSampleData() error {
	fmt.Println("Inserting Sample Data...")
	s1 := &Service{
		Name:           "Google",
		Domain:         "https://google.com",
		ExpectedStatus: 200,
		Interval:       10,
		Port:           0,
		Type:           "https",
		Method:         "GET",
	}
	s2 := &Service{
		Name:           "Statup.io",
		Domain:         "https://statup.io",
		ExpectedStatus: 200,
		Interval:       15,
		Port:           0,
		Type:           "https",
		Method:         "GET",
	}
	s3 := &Service{
		Name:           "Statup.io SSL Check",
		Domain:         "https://statup.io",
		ExpectedStatus: 200,
		Interval:       15,
		Port:           443,
		Type:           "tcp",
	}
	s4 := &Service{
		Name:           "Github Failing Check",
		Domain:         "https://github.com/thisisnotausernamemaybeitis",
		ExpectedStatus: 200,
		Interval:       15,
		Port:           0,
		Type:           "https",
		Method:         "GET",
	}
	s1.Create()
	s2.Create()
	s3.Create()
	s4.Create()

	checkin := &Checkin{
		Service:  s2.Id,
		Interval: 30,
		Api:      NewSHA1Hash(18),
	}
	checkin.Create()

	for i := 0; i < 20; i++ {
		s1.Check()
		s2.Check()
		s3.Check()
		s4.Check()
	}

	return nil
}
