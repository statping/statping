package main

import (
	"fmt"
	"strings"
	"time"
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
		if configs.Port == "" {
			configs.Port = "3306"
		}
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
		if configs.Port == "" {
			configs.Port = "5432"
		}
		host := fmt.Sprintf("%v:%v", configs.Host, configs.Port)
		postgresSettings = postgresql.ConnectionURL{
			Database: configs.Database,
			Host:     host,
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
	OnLoad(dbSession)
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

	for i := 0; i < 100; i++ {
		s1.Check()
		s2.Check()
		s3.Check()
		s4.Check()
		time.Sleep(250 * time.Millisecond)
	}

	return nil
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
