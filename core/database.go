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
			Options:  map[string]string{"parseTime": "true", "charset": "utf8"},
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

	CoreApp, err = SelectCore()
	CoreApp.DbConnection = c.DbConn

	return err
}

func versionSplit(v string) (int64, int64, int64) {
	currSplit := strings.Split(v, ".")
	if len(currSplit) < 2 {
		return 9999, 9999, 9999
	}
	var major, mid, minor string
	if len(currSplit) == 3 {
		major = currSplit[0]
		mid = currSplit[1]
		minor = currSplit[2]
		return utils.StringInt(major), utils.StringInt(mid), utils.StringInt(minor)
	}
	major = currSplit[0]
	mid = currSplit[1]
	return utils.StringInt(major), utils.StringInt(mid), 0
}

func versionHigher(migrate string) bool {
	cM, cMi, cMn := versionSplit(CoreApp.Version)
	mM, mMi, mMn := versionSplit(migrate)
	if mM > cM {
		return true
	}
	if mMi > cMi {
		return true
	}
	if mMn > cMn {
		return true
	}
	return false
}

func RunDatabaseUpgrades() error {
	var err error
	utils.Log(1, fmt.Sprintf("Checking Database Upgrades from v%v in '%v_upgrade.sql'...", CoreApp.Version, CoreApp.DbConnection))
	upgrade, _ := SqlBox.String(CoreApp.DbConnection + "_upgrade.sql")
	// parse db version and upgrade file
	ups := strings.Split(upgrade, "=========================================== ")
	var ran int
	for _, v := range ups {
		if len(v) == 0 {
			continue
		}
		vers := strings.Split(v, "\n")
		version := vers[0]
		data := vers[1:]
		//fmt.Printf("Checking Migration from v%v to v%v - %v\n", CoreApp.Version, version, versionHigher(version))
		if !versionHigher(version) {
			//fmt.Printf("Already up-to-date with v%v\n", version)
			continue
		}
		fmt.Printf("Migration Database from v%v to v%v\n", CoreApp.Version, version)
		for _, m := range data {
			if m == "" {
				continue
			}
			fmt.Printf("Running Migration: %v\n", m)
			_, err := DbSession.Exec(db.Raw(m + ";"))
			if err != nil {
				utils.Log(2, err)
				continue
			}
			ran++
			CoreApp.Version = m
		}
	}
	CoreApp.Update()
	CoreApp, err = SelectCore()
	if ran > 0 {
		utils.Log(1, fmt.Sprintf("Database Upgraded, %v query ran", ran))
	} else {
		utils.Log(1, fmt.Sprintf("Database is already up-to-date, latest v%v", CoreApp.Version))
	}
	return err
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
	if CoreApp.DbConnection == "mysql" {
		sql = "mysql_up.sql"
	} else if CoreApp.DbConnection == "sqlite" {
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
