package database

import (
	"fmt"
	"github.com/statping/statping/types/metrics"
	"github.com/statping/statping/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Database struct {
	*gorm.DB
}

func Close() {

}

func Open(dialect, dsn string) (*Database, error) {
	var dia gorm.Dialector
	config := &gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc:                utils.Now,
	}
	switch dialect {
	case "sqlite", "sqlite3":
		dia = sqlite.Open(dsn)
	case "postgres":
		dia = postgres.Open(dsn)
	case "mysql":
		dia = mysql.Open(dsn)
	}
	conn, err := gorm.Open(dia, config)
	if err != nil {
		return nil, err
	}
	return Wrap(conn), nil
}

func (d *Database) ChunkSize() int {
	switch d.DB.Dialector.Name() {
	case "mysql":
		return 3000
	case "postgres":
		return 3000
	default:
		return 100
	}
}

func (d *Database) MultipleSelects(args ...string) *Database {
	joined := strings.Join(args, ", ")
	return Wrap(d.DB.Select(joined))
}

func (d *Database) Status() int {
	switch d.DB.Error {
	case gorm.ErrRecordNotFound:
		return 404
	default:
		return 500
	}
}

func (d *Database) Close() error {
	return nil
}

func Routine(db *Database) {
	for {
		if db == nil {
			time.Sleep(5 * time.Second)
			continue
		}
		stats, err := db.DB.DB()
		if err != nil {
			log.Errorln(err)
			time.Sleep(5 * time.Second)
			continue
		}

		metrics.CollectDatabase(stats.Stats())
		time.Sleep(5 * time.Second)
	}
}

func OpenTester() (*Database, error) {
	testDB := utils.Params.GetString("DB_CONN")
	var dbString string

	switch testDB {
	case "mysql":
		dbString = fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=UTC&time_zone=%%27UTC%%27",
			utils.Params.GetString("DB_HOST"),
			utils.Params.GetString("DB_PASS"),
			utils.Params.GetString("DB_HOST"),
			utils.Params.GetInt("DB_PORT"),
			utils.Params.GetString("DB_DATABASE"),
		)
	case "postgres":
		dbString = fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s sslmode=disable timezone=UTC",
			utils.Params.GetString("DB_HOST"),
			utils.Params.GetInt("DB_PORT"),
			utils.Params.GetString("DB_USER"),
			utils.Params.GetString("DB_DATABASE"),
			utils.Params.GetString("DB_PASS"))
	default:
		dbString = fmt.Sprintf("file:%s?mode=memory&cache=shared", utils.RandomString(12))
	}
	newDb, err := Open(testDB, dbString)
	if err != nil {
		return nil, err
	}
	db, err := newDb.DB.DB()
	if err != nil {
		return nil, err
	}
	if testDB == "sqlite3" || testDB == "sqlite" {
		db.SetMaxOpenConns(1)
	}
	return newDb, err
}

// Wrap wraps gorm.DB in an interface
func Wrap(db *gorm.DB) *Database {
	return &Database{
		DB: db,
	}
}
