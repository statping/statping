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
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-yaml/yaml"
	"github.com/hunterlong/statup/notifiers"
	"github.com/hunterlong/statup/source"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strings"
	"time"
)

var (
	DbSession        *gorm.DB
	currentMigration int64
)

func failuresDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "TEST" {
		return DbSession.Model(&types.Failure{}).Debug()
	}
	return DbSession.Model(&types.Failure{})
}

func (s *Service) allHits() *gorm.DB {
	var hits []*Hit
	return servicesDB().Find(s).Related(&hits)
}

func hitsDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "TEST" {
		return DbSession.Model(&types.Hit{}).Debug()
	}
	return DbSession.Model(&types.Hit{})
}

func servicesDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "TEST" {
		return DbSession.Model(&types.Service{}).Debug()
	}
	return DbSession.Model(&types.Service{})
}

func coreDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "TEST" {
		return DbSession.Table("core").Debug()
	}
	return DbSession.Table("core")
}

func usersDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "TEST" {
		return DbSession.Model(&types.User{}).Debug()
	}
	return DbSession.Model(&types.User{})
}

func commDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "TEST" {
		return DbSession.Table("communication").Model(&notifiers.Notification{}).Debug()
	}
	return DbSession.Table("communication").Model(&notifiers.Notification{})
}

func checkinDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "TEST" {
		return DbSession.Model(&types.Checkin{}).Debug()
	}
	return DbSession.Model(&types.Checkin{})
}

type DbConfig struct {
	*types.DbConfig
}

func (db *DbConfig) Close() error {
	return DbSession.Close()
}

func (db *DbConfig) InsertCore() (*Core, error) {
	CoreApp = &Core{Core: &types.Core{
		Name:        db.Project,
		Description: db.Description,
		Config:      "config.yml",
		ApiKey:      utils.NewSHA1Hash(9),
		ApiSecret:   utils.NewSHA1Hash(16),
		Domain:      db.Domain,
		MigrationId: time.Now().Unix(),
	}}
	CoreApp.DbConnection = db.DbConn
	query := coreDB().Create(&CoreApp)
	return CoreApp, query.Error
}

func (db *DbConfig) Connect(retry bool, location string) error {
	var err error
	if DbSession != nil {
		DbSession = nil
	}
	switch Configs.DbConn {
	case "sqlite":
		DbSession, err = gorm.Open("sqlite3", utils.Directory+"/statup.db")
		if err != nil {
			return err
		}
	case "mysql":
		if Configs.DbPort == 0 {
			Configs.DbPort = 3306
		}
		host := fmt.Sprintf("%v:%v", Configs.DbHost, Configs.DbPort)
		conn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local", Configs.DbUser, Configs.DbPass, host, Configs.DbData)
		DbSession, err = gorm.Open("mysql", conn)
		DbSession.DB().SetConnMaxLifetime(time.Minute * 5)
		DbSession.DB().SetMaxIdleConns(0)
		DbSession.DB().SetMaxOpenConns(5)
		if err != nil {
			if retry {
				utils.Log(1, fmt.Sprintf("Database connection to '%v' is not available, trying again in 5 seconds...", host))
				return db.waitForDb()
			} else {
				return err
			}
		}
	case "postgres":
		if Configs.DbPort == 0 {
			Configs.DbPort = 5432
		}
		conn := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", Configs.DbHost, Configs.DbPort, Configs.DbUser, Configs.DbData, Configs.DbPass)
		DbSession, err = gorm.Open("postgres", conn)
		if err != nil {
			if retry {
				utils.Log(1, fmt.Sprintf("Database connection to '%v' is not available, trying again in 5 seconds...", Configs.DbHost))
				return db.waitForDb()
			} else {
				fmt.Println("ERROR:", err)
				return err
			}
		}
	case "mssql":
		if Configs.DbPort == 0 {
			Configs.DbPort = 1433
		}
		host := fmt.Sprintf("%v:%v", Configs.DbHost, Configs.DbPort)
		conn := fmt.Sprintf("sqlserver://%v:%v@%v?database=%v", Configs.DbUser, Configs.DbPass, host, Configs.DbData)
		DbSession, err = gorm.Open("mssql", conn)
		if err != nil {
			if retry {
				utils.Log(1, fmt.Sprintf("Database connection to '%v' is not available, trying again in 5 seconds...", host))
				return db.waitForDb()
			} else {
				return err
			}
		}
	}
	err = DbSession.DB().Ping()
	if err == nil {
		utils.Log(1, fmt.Sprintf("Database connection to '%v' was successful.", Configs.DbData))
	}
	return err
}

func (db *DbConfig) waitForDb() error {
	time.Sleep(5 * time.Second)
	return db.Connect(true, utils.Directory)
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
	db := DbSession.Raw(sql)
	defer db.Close()
	if db.Error != nil {
		utils.Log(2, db.Error)
	}
}

func (c *DbConfig) Update() error {
	var err error
	config, err := os.Create(utils.Directory + "/config.yml")
	if err != nil {
		utils.Log(4, err)
		return err
	}
	data, err := yaml.Marshal(c.DbConfig)
	if err != nil {
		utils.Log(3, err)
		return err
	}
	config.WriteString(string(data))
	config.Close()
	return err
}

func (c *DbConfig) Save() (*DbConfig, error) {
	var err error
	config, err := os.Create(utils.Directory + "/config.yml")
	if err != nil {
		utils.Log(4, err)
		return nil, err
	}
	c.ApiKey = utils.NewSHA1Hash(16)
	c.ApiSecret = utils.NewSHA1Hash(16)
	data, err := yaml.Marshal(c.DbConfig)
	if err != nil {
		utils.Log(3, err)
		return nil, err
	}
	config.WriteString(string(data))
	defer config.Close()
	return c, err
}

func (c *DbConfig) CreateCore() *Core {
	newCore := &types.Core{
		Name:        c.Project,
		Description: c.Description,
		Config:      "config.yml",
		ApiKey:      c.ApiKey,
		ApiSecret:   c.ApiSecret,
		Domain:      c.Domain,
		MigrationId: time.Now().Unix(),
	}
	db := coreDB().Create(&newCore)
	if db.Error == nil {
		CoreApp = &Core{Core: newCore}
	}
	CoreApp, err := SelectCore()
	if err != nil {
		utils.Log(4, err)
	}
	return CoreApp
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
			db := DbSession.Raw(m)
			ran++
			if db.Error != nil {
				utils.Log(2, db.Error)
				continue
			}
		}
		currentMigration = lastMigration
	}
	if ran > 0 {
		utils.Log(1, fmt.Sprintf("Database Upgraded %v queries ran, current #%v", ran, currentMigration))
		CoreApp, err = SelectCore()
		if err != nil {
			return err
		}
		CoreApp.MigrationId = currentMigration
		UpdateCore(CoreApp)
	}
	return err
}

func (db *DbConfig) SeedSchema() (string, string, error) {
	utils.Log(1, "Seeding Schema Database with Dummy Data...")
	dir := utils.Directory
	var cmd string
	switch db.DbConn {
	case "sqlite":
		cmd = fmt.Sprintf("cat %v/source/sql/sqlite_up.sql | sqlite3 %v/statup.db", dir, dir)
	case "mysql":
		cmd = fmt.Sprintf("mysql -h %v -P %v -u %v --password=%v %v < %v/source/sql/mysql_up.sql", Configs.DbHost, Configs.DbPort, Configs.DbUser, Configs.DbPass, Configs.DbData, dir)
	case "postgres":
		cmd = fmt.Sprintf("PGPASSWORD=%v psql -U %v -h %v -d %v -1 -f %v/source/sql/postgres_up.sql", db.DbPass, db.DbUser, db.DbHost, db.DbData, dir)
	}
	out, outErr, err := utils.Command(cmd)
	if err != nil {
		return out, outErr, err
	}
	return out, outErr, err
}

func (db *DbConfig) SeedDatabase() (string, string, error) {
	utils.Log(1, "Seeding Database with Dummy Data...")
	dir := utils.Directory
	var cmd string
	switch db.DbConn {
	case "sqlite":
		cmd = fmt.Sprintf("cat %v/dev/sqlite_seed.sql | sqlite3 %v/statup.db", dir, dir)
	case "mysql":
		cmd = fmt.Sprintf("mysql -h %v -P %v -u %v --password=%v %v < %v/dev/mysql_seed.sql", Configs.DbHost, Configs.DbPort, Configs.DbUser, Configs.DbPass, Configs.DbData, dir)
	case "postgres":
		cmd = fmt.Sprintf("PGPASSWORD=%v psql -U %v -h %v -d %v -1 -f %v/dev/postgres_seed.sql", db.DbPass, db.DbUser, db.DbHost, db.DbData, dir)
	}
	out, outErr, err := utils.Command(cmd)
	return out, outErr, err
}

func (db *DbConfig) DropDatabase() error {
	utils.Log(1, "Dropping Database Tables...")
	err := DbSession.DropTableIfExists("checkins")
	err = DbSession.DropTableIfExists("communication")
	err = DbSession.DropTableIfExists("core")
	err = DbSession.DropTableIfExists("failures")
	err = DbSession.DropTableIfExists("hits")
	err = DbSession.DropTableIfExists("services")
	err = DbSession.DropTableIfExists("users")
	return err.Error
}

func (db *DbConfig) CreateDatabase() error {
	utils.Log(1, "Creating Database Tables...")
	err := DbSession.CreateTable(&types.Checkin{})
	err = DbSession.Table("communication").CreateTable(&notifiers.Notification{})
	err = DbSession.Table("core").CreateTable(&types.Core{})
	err = DbSession.CreateTable(&types.Failure{})
	err = DbSession.CreateTable(&types.Hit{})
	err = DbSession.CreateTable(&types.Service{})
	err = DbSession.CreateTable(&types.User{})
	utils.Log(1, "Statup Database Created")
	return err.Error
}

func (db *DbConfig) MigrateDatabase() error {
	utils.Log(1, "Migrating Database Tables...")
	err := DbSession.AutoMigrate(&types.Checkin{})
	err = DbSession.Table("communication").AutoMigrate(&notifiers.Notification{})
	err = DbSession.Table("core").AutoMigrate(&types.Core{})
	err = DbSession.AutoMigrate(&types.Failure{})
	err = DbSession.AutoMigrate(&types.Hit{})
	err = DbSession.AutoMigrate(&types.Service{})
	err = DbSession.AutoMigrate(&types.User{})
	utils.Log(1, "Statup Database Migrated")
	return err.Error
}

func (c *DbConfig) Clean() *DbConfig {
	if os.Getenv("DB_PORT") != "" {
		if c.DbConn == "postgres" {
			c.DbHost = c.DbHost + ":" + os.Getenv("DB_PORT")
		}
	}
	return c
}
