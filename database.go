package main

import (
	"fmt"
	"github.com/hunterlong/statup/plugin"
	"strings"
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
	dbSession        sqlbuilder.Database
)

func DbConnection(dbType string) error {
	var err error
	if dbType == "sqlite" {
		sqliteSettings = sqlite.ConnectionURL{
			Database: "statup.db",
		}
		dbSession, err = sqlite.Open(sqliteSettings)
		if err != nil {
			return err
		}
	} else if dbType == "mysql" {
		mysqlSettings = mysql.ConnectionURL{
			Database: configs.Database,
			Host:     configs.Host,
			User:     configs.User,
			Password: configs.Password,
		}
		dbSession, err = mysql.Open(mysqlSettings)
		if err != nil {
			return err
		}
	} else {
		postgresSettings = postgresql.ConnectionURL{
			Database: configs.Database,
			Host:     configs.Host,
			User:     configs.User,
			Password: configs.Password,
		}
		dbSession, err = postgresql.Open(postgresSettings)
		if err != nil {
			return err
		}
	}
	//dbSession.SetLogging(true)
	dbServer = dbType
	plugin.SetDatabase(dbSession)
	return err
}

func UpgradeDatabase() {
	fmt.Println("New Version:     ", core.Version)
	fmt.Println("Current Version: ", VERSION)
	if VERSION == core.Version {
		fmt.Println("Database already up to date")
		return
	}
	fmt.Println("Upgrading Database...")
	upgrade, _ := sqlBox.String("upgrade.sql")
	requests := strings.Split(upgrade, ";")
	for _, request := range requests {
		_, err := db.Exec(request)
		if err != nil {
			fmt.Println(err)
		}
	}
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
	admin := &User{
		Username: "admin",
		Password: "admin",
		Email:    "admin@admin.com",
	}
	s1.Create()
	s2.Create()
	s3.Create()
	s4.Create()
	admin.Create()

	for i := 0; i < 20; i++ {
		s1.Check()
		s2.Check()
		s3.Check()
		s4.Check()
	}

	return nil
}

func CreateDatabase() {
	fmt.Println("Creating Tables...")
	VERSION = "1.1.1"
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
