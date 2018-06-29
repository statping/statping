package main

import (
	"fmt"
	"github.com/hunterlong/statup/log"
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
			log.Send(3, err)
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
			log.Send(3, err)
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
			log.Send(3, err)
			return err
		}
	}
	//dbSession.SetLogging(true)
	dbServer = dbType
	OnLoad(dbSession)
	return err
}

func DatabaseMaintence() {
	defer DatabaseMaintence()
	since := time.Now().AddDate(0, 0, -7)
	DeleteAllSince("failures", since)
	DeleteAllSince("hits", since)
	time.Sleep(60 * time.Minute)
}

func DeleteAllSince(table string, date time.Time) {
	sql := fmt.Sprintf("DELETE FROM %v WHERE created_at < '%v';", table, date.Format("2006-01-02"))
	_, err := dbSession.Exec(db.Raw(sql))
	if err != nil {
		log.Send(2, err)
	}
}

func Backup() {

}
