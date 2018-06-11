package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"math/rand"
	"time"
	"github.com/hunterlong/statup/plugin"
)

func DbConnection() {
	var err error
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", configs.Host, configs.Port, configs.User, configs.Password, configs.Database)
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	plugin.SetDatabase(db)
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
