// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/hunterlong/statup/source"
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
	currentMigration int64
)

type DbConfig types.DbConfig

func DbConnection(dbType string, retry bool, location string) error {
	var err error
	if dbType == "sqlite" {
		sqliteSettings = sqlite.ConnectionURL{
			Database: location + "/statup.db",
		}
		DbSession, err = sqlite.Open(sqliteSettings)
		if err != nil {
			return err
		}
	} else if dbType == "mysql" {
		if Configs.Port == "" {
			Configs.Port = "3306"
		}
		host := fmt.Sprintf("%v:%v", Configs.Host, Configs.Port)
		mysqlSettings = mysql.ConnectionURL{
			Database: Configs.Database,
			Host:     host,
			User:     Configs.User,
			Password: Configs.Password,
			Options:  map[string]string{"parseTime": "true", "charset": "utf8"},
		}
		DbSession, err = mysql.Open(mysqlSettings)
		if err != nil {
			if retry {
				utils.Log(1, fmt.Sprintf("Database connection to '%v' is not available, trying again in 5 seconds...", host))
				return waitForDb(dbType)
			} else {
				return err
			}
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
			if retry {
				utils.Log(1, fmt.Sprintf("Database connection to '%v' is not available, trying again in 5 seconds...", host))
				return waitForDb(dbType)
			} else {
				return err
			}
		}
	}
	err = DbSession.Ping()
	if err == nil {
		utils.Log(1, fmt.Sprintf("Database connection to '%v' was successful.", DbSession.Name()))
	}
	return err
}

func waitForDb(dbType string) error {
	time.Sleep(5 * time.Second)
	return DbConnection(dbType, true, ".")
}

func DatabaseMaintence() {
	for range time.Tick(60 * time.Minute) {
		utils.Log(1, "Checking for database records older than 7 days...")
		since := time.Now().AddDate(0, 0, -7)
		DeleteAllSince("failures", since)
		DeleteAllSince("hits", since)
	}
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
	err = DbConnection(Configs.Connection, false, c.Location)
	if err != nil {
		utils.Log(4, err)
		return err
	}
	DropDatabase()
	CreateDatabase()

	newCore := &types.Core{
		Name:        c.Project,
		Description: c.Description,
		Config:      "config.yml",
		ApiKey:      utils.NewSHA1Hash(9),
		ApiSecret:   utils.NewSHA1Hash(16),
		Domain:      c.Domain,
		MigrationId: time.Now().Unix(),
	}
	col := DbSession.Collection("core")
	_, err = col.Insert(newCore)
	if err == nil {
		CoreApp = &Core{Core: newCore}
	}

	CoreApp, err = SelectCore()
	if err != nil {
		utils.Log(4, err)
	}
	CoreApp.DbConnection = c.DbConn

	return err
}

func versionHigher(migrate int64) bool {
	if CoreApp.MigrationId < migrate {
		return true
	}
	return false
}

func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func RunDatabaseUpgrades() error {
	var err error
	currentMigration, err = SelectLastMigration()
	if err != nil {
		return err
	}
	utils.Log(1, fmt.Sprintf("Checking for Database Upgrades since #%v", currentMigration))
	upgrade, _ := source.SqlBox.String(CoreApp.DbConnection + "_upgrade.sql")
	// parse db version and upgrade file
	ups := strings.Split(upgrade, "=========================================== ")
	ups = reverseSlice(ups)
	var ran int
	var lastMigration int64
	for _, v := range ups {
		if len(v) == 0 {
			continue
		}
		vers := strings.Split(v, "\n")
		lastMigration = utils.StringInt(vers[0])
		data := vers[1:]

		//fmt.Printf("Checking Migration from v%v to v%v - %v\n", CoreApp.Version, version, versionHigher(version))
		if currentMigration >= lastMigration {
			continue
		}
		utils.Log(1, fmt.Sprintf("Migrating Database from #%v to #%v", currentMigration, lastMigration))
		for _, m := range data {
			if m == "" {
				continue
			}
			utils.Log(1, fmt.Sprintf("Running Query: %v", m))
			_, err := DbSession.Exec(db.Raw(m + ";"))
			ran++
			if err != nil {
				utils.Log(2, err)
				continue
			}
		}
		currentMigration = lastMigration
	}
	if ran > 0 {
		utils.Log(1, fmt.Sprintf("Database Upgraded %v queries ran, current #%v", ran, currentMigration))
		CoreApp, err = SelectCore()
		if err != nil {
			panic(err)
		}
		CoreApp.MigrationId = currentMigration
		UpdateCore(CoreApp)
	}
	return err
}

func DropDatabase() error {
	utils.Log(1, "Dropping Database Tables...")
	down, err := source.SqlBox.String("down.sql")
	if err != nil {
		return err
	}
	requests := strings.Split(down, ";")
	for _, request := range requests {
		_, err := DbSession.Exec(request)
		if err != nil {
			utils.Log(2, err)
		}
	}
	return err
}

func CreateDatabase() error {
	utils.Log(1, "Creating Database Tables...")
	sql := "postgres_up.sql"
	if CoreApp.DbConnection == "mysql" {
		sql = "mysql_up.sql"
	} else if CoreApp.DbConnection == "sqlite" {
		sql = "sqlite_up.sql"
	}
	up, err := source.SqlBox.String(sql)
	requests := strings.Split(up, ";")
	for _, request := range requests {
		_, err := DbSession.Exec(request)
		if err != nil {
			utils.Log(2, err)
		}
	}
	//secret := NewSHA1Hash()
	//db.QueryRow("INSERT INTO core (secret, version) VALUES ($1, $2);", secret, VERSION).Scan()
	utils.Log(1, "Database Created")
	//SampleData()
	return err
}

func (c *DbConfig) Clean() *DbConfig {
	if os.Getenv("DB_PORT") != "" {
		if c.DbConn == "postgres" {
			c.DbHost = c.DbHost + ":" + os.Getenv("DB_PORT")
		}
	}
	return c
}
