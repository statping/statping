package main

import (
	"database/sql"
	"fmt"
	"github.com/hunterlong/statup/plugin"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
)

func DbConnection(dbType string) error {
	var err error
	var dbInfo string
	if dbType=="sqlite3" {
		dbInfo = "./statup.db"
	} else if dbType=="mysql" {
		dbInfo = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", configs.User, configs.Password, configs.Host, configs.Port, configs.Database)
	} else {
		dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", configs.Host, configs.Port, configs.User, configs.Password, configs.Database)
	}
	db, err = sql.Open(dbType, dbInfo)
	if err != nil {
		return err
	}

	//stmt, err := db.Prepare("CREATE database statup;")
	//if err != nil {
	//	panic(err)
	//}
	//stmt.Exec()

	plugin.SetDatabase(db)
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
	db.QueryRow(upgrade).Scan()
}

func DropDatabase() {
	fmt.Println("Dropping Tables...")
	down, _ := sqlBox.String("down.sql")
	db.QueryRow(down).Scan()
}

func CreateDatabase() {
	fmt.Println("Creating Tables...")
	VERSION = "1.1.1"
	up, _ := sqlBox.String("up.sql")
	db.QueryRow(up).Scan()
	//secret := NewSHA1Hash()
	//db.QueryRow("INSERT INTO core (secret, version) VALUES ($1, $2);", secret, VERSION).Scan()
	fmt.Println("Database Created")
	//SampleData()
}

func SampleData() {
	i := 0
	for i < 300 {
		ran := rand.Float32()
		latency := fmt.Sprintf("%0.2f", ran)
		date := time.Now().AddDate(0, 0, i)
		db.QueryRow("INSERT INTO hits (service, latency, created_at) VALUES (1, $1, $2);", latency, date).Scan()
		i++
	}
}
