package core

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"os"
	"strings"
	"time"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
	"upper.io/db.v3/postgresql"
	"upper.io/db.v3/sqlite"
)

var (
	dbServer         string
	sqliteSettings   sqlite.ConnectionURL
	postgresSettings postgresql.ConnectionURL
	mysqlSettings    mysql.ConnectionURL
	DbSession        sqlbuilder.Database
)

type DbConfig types.DbConfig

func DbConnection(dbType string) error {
	var err error
	if dbType == "sqlite" {
		sqliteSettings = sqlite.ConnectionURL{
			Database: "statup.db",
		}
		DbSession, err = sqlite.Open(sqliteSettings)
		if err != nil {
			return err
		}
	} else if dbType == "mysql" {
		if Configs.Port == "" {
			Configs.Port = "3306"
		}
		mysqlSettings = mysql.ConnectionURL{
			Database: Configs.Database,
			Host:     Configs.Host,
			User:     Configs.User,
			Password: Configs.Password,
		}
		DbSession, err = mysql.Open(mysqlSettings)
		if err != nil {
			return err
		}
	} else {
		if Configs.Port == "" {
			Configs.Port = "5432"
		}
		host := fmt.Sprintf("%v:%v", Configs.Host, Configs.Port)
		postgresSettings = postgresql.ConnectionURL{
			Database: Configs.Database,
			Host:     host,
			User:     Configs.User,
			Password: Configs.Password,
		}
		DbSession, err = postgresql.Open(postgresSettings)
		if err != nil {
			return err
		}
	}
	//dbSession.SetLogging(true)
	dbServer = dbType
	OnLoad(DbSession)
	return err
}

func DatabaseMaintence() {
	defer DatabaseMaintence()
	utils.Log(1, "Checking for database records older than 7 days...")
	since := time.Now().AddDate(0, 0, -7)
	DeleteAllSince("failures", since)
	DeleteAllSince("hits", since)
	time.Sleep(60 * time.Minute)
}

func DeleteAllSince(table string, date time.Time) {
	sql := fmt.Sprintf("DELETE FROM %v WHERE created_at < '%v';", table, date.Format("2006-01-02"))
	_, err := DbSession.Exec(db.Raw(sql))
	if err != nil {
		utils.Log(2, err)
	}
}

func (c *DbConfig) Save() error {
	var err error
	config, err := os.Create("config.yml")
	if err != nil {
		utils.Log(4, err)
		return err
	}
	data, err := yaml.Marshal(c)
	if err != nil {
		utils.Log(3, err)
		return err
	}
	config.WriteString(string(data))
	config.Close()

	Configs, err = LoadConfig()
	if err != nil {
		utils.Log(3, err)
		return err
	}
	err = DbConnection(Configs.Connection)
	if err != nil {
		utils.Log(4, err)
		return err
	}
	DropDatabase()
	CreateDatabase()

	newCore := &Core{
		Name:        c.Project,
		Description: c.Description,
		Config:      "config.yml",
		ApiKey:      utils.NewSHA1Hash(9),
		ApiSecret:   utils.NewSHA1Hash(16),
		Domain:      c.Domain,
	}
	col := DbSession.Collection("core")
	_, err = col.Insert(newCore)
	if err == nil {
		CoreApp = newCore
	}
	return err
}

func RunDatabaseUpgrades() {
	utils.Log(1, "Running Database Upgrade from 'upgrade.sql'...")
	upgrade, _ := SqlBox.String("upgrade.sql")
	requests := strings.Split(upgrade, ";")
	for _, request := range requests {
		_, err := DbSession.Exec(db.Raw(request + ";"))
		if err != nil {
			utils.Log(2, err)
		}
	}
	utils.Log(1, "Database Upgraded")
}

func DropDatabase() {
	fmt.Println("Dropping Tables...")
	down, _ := SqlBox.String("down.sql")
	requests := strings.Split(down, ";")
	for _, request := range requests {
		_, err := DbSession.Exec(request)
		if err != nil {
			utils.Log(2, err)
		}
	}
}

func CreateDatabase() {
	fmt.Println("Creating Tables...")
	sql := "postgres_up.sql"
	if dbServer == "mysql" {
		sql = "mysql_up.sql"
	} else if dbServer == "sqlite" {
		sql = "sqlite_up.sql"
	}
	up, _ := SqlBox.String(sql)
	requests := strings.Split(up, ";")
	for _, request := range requests {
		_, err := DbSession.Exec(request)
		if err != nil {
			utils.Log(2, err)
		}
	}
	//secret := NewSHA1Hash()
	//db.QueryRow("INSERT INTO core (secret, version) VALUES ($1, $2);", secret, VERSION).Scan()
	fmt.Println("Database Created")
	//SampleData()
}

func (c *DbConfig) Clean() *DbConfig {
	if os.Getenv("DB_PORT") != "" {
		if c.DbConn == "postgres" {
			c.DbHost = c.DbHost + ":" + os.Getenv("DB_PORT")
		}
	}
	return c
}
