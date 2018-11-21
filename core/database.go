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
	"github.com/hunterlong/statup/core/notifier"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"time"
)

var (
	// DbSession stores the Statup database session
	DbSession *gorm.DB
	DbModels  []interface{}
)

func init() {
	DbModels = []interface{}{&types.Service{}, &types.User{}, &types.Hit{}, &types.Failure{}, &types.Message{}, &types.Checkin{}, &types.CheckinHit{}, &notifier.Notification{}}
}

// DbConfig stores the config.yml file for the statup configuration
type DbConfig types.DbConfig

// failuresDB returns the 'failures' database column
func failuresDB() *gorm.DB {
	return DbSession.Model(&types.Failure{})
}

// hitsDB returns the 'hits' database column
func hitsDB() *gorm.DB {
	return DbSession.Model(&types.Hit{})
}

// servicesDB returns the 'services' database column
func servicesDB() *gorm.DB {
	return DbSession.Model(&types.Service{})
}

// coreDB returns the single column 'core'
func coreDB() *gorm.DB {
	return DbSession.Table("core").Model(&CoreApp)
}

// usersDB returns the 'users' database column
func usersDB() *gorm.DB {
	return DbSession.Model(&types.User{})
}

// checkinDB returns the Checkin records for a service
func checkinDB() *gorm.DB {
	return DbSession.Table("checkins").Model(&types.Checkin{})
}

// checkinHitsDB returns the Checkin Hits records for a service
func checkinHitsDB() *gorm.DB {
	return DbSession.Model(&types.CheckinHit{})
}

// messagesDb returns the Checkin records for a service
func messagesDb() *gorm.DB {
	return DbSession.Model(&types.Message{})
}

// HitsBetween returns the gorm database query for a collection of service hits between a time range
func (s *Service) HitsBetween(t1, t2 time.Time, group string, column string) *gorm.DB {
	selector := Dbtimestamp(group, column)
	if Configs.DbConn == "postgres" {
		timeQuery := fmt.Sprintf("service = %v AND created_at BETWEEN '%v.000000' AND '%v.000000'", s.Id, t1.UTC().Format(types.POSTGRES_TIME), t2.UTC().Format(types.POSTGRES_TIME))
		return DbSession.Model(&types.Hit{}).Select(selector).Where(timeQuery)
	} else {
		return DbSession.Model(&types.Hit{}).Select(selector).Where("service = ? AND created_at BETWEEN ? AND ?", s.Id, t1.UTC().Format(types.TIME_DAY), t2.UTC().Format(types.TIME_DAY))
	}
}

// CloseDB will close the database connection if available
func CloseDB() {
	if DbSession != nil {
		DbSession.DB().Close()
	}
}

// Close shutsdown the database connection
func (db *DbConfig) Close() error {
	return DbSession.DB().Close()
}

// AfterFind for Core will set the timezone
func (c *Core) AfterFind() (err error) {
	c.CreatedAt = utils.Timezoner(c.CreatedAt, CoreApp.Timezone)
	c.UpdatedAt = utils.Timezoner(c.UpdatedAt, CoreApp.Timezone)
	return
}

// AfterFind for Service will set the timezone
func (s *Service) AfterFind() (err error) {
	s.CreatedAt = utils.Timezoner(s.CreatedAt, CoreApp.Timezone)
	s.UpdatedAt = utils.Timezoner(s.UpdatedAt, CoreApp.Timezone)
	return
}

// AfterFind for Hit will set the timezone
func (h *Hit) AfterFind() (err error) {
	h.CreatedAt = utils.Timezoner(h.CreatedAt, CoreApp.Timezone)
	return
}

// AfterFind for failure will set the timezone
func (f *failure) AfterFind() (err error) {
	f.CreatedAt = utils.Timezoner(f.CreatedAt, CoreApp.Timezone)
	return
}

// AfterFind for USer will set the timezone
func (u *User) AfterFind() (err error) {
	u.CreatedAt = utils.Timezoner(u.CreatedAt, CoreApp.Timezone)
	u.UpdatedAt = utils.Timezoner(u.UpdatedAt, CoreApp.Timezone)
	return
}

// AfterFind for Checkin will set the timezone
func (c *Checkin) AfterFind() (err error) {
	c.CreatedAt = utils.Timezoner(c.CreatedAt, CoreApp.Timezone)
	c.UpdatedAt = utils.Timezoner(c.UpdatedAt, CoreApp.Timezone)
	return
}

// AfterFind for checkinHit will set the timezone
func (c *CheckinHit) AfterFind() (err error) {
	c.CreatedAt = utils.Timezoner(c.CreatedAt, CoreApp.Timezone)
	return
}

// AfterFind for Message will set the timezone
func (u *Message) AfterFind() (err error) {
	u.CreatedAt = utils.Timezoner(u.CreatedAt, CoreApp.Timezone)
	u.UpdatedAt = utils.Timezoner(u.UpdatedAt, CoreApp.Timezone)
	u.StartOn = utils.Timezoner(u.StartOn, CoreApp.Timezone)
	u.EndOn = utils.Timezoner(u.EndOn, CoreApp.Timezone)
	return
}

// BeforeCreate for Hit will set CreatedAt to UTC
func (h *Hit) BeforeCreate() (err error) {
	if h.CreatedAt.IsZero() {
		h.CreatedAt = time.Now().UTC()
	}
	return
}

// BeforeCreate for failure will set CreatedAt to UTC
func (f *failure) BeforeCreate() (err error) {
	if f.CreatedAt.IsZero() {
		f.CreatedAt = time.Now().UTC()
	}
	return
}

// BeforeCreate for User will set CreatedAt to UTC
func (u *User) BeforeCreate() (err error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now().UTC()
		u.UpdatedAt = time.Now().UTC()
	}
	return
}

// BeforeCreate for Message will set CreatedAt to UTC
func (u *Message) BeforeCreate() (err error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now().UTC()
		u.UpdatedAt = time.Now().UTC()
	}
	return
}

// BeforeCreate for Service will set CreatedAt to UTC
func (s *Service) BeforeCreate() (err error) {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now().UTC()
		s.UpdatedAt = time.Now().UTC()
	}
	return
}

// BeforeCreate for Checkin will set CreatedAt to UTC
func (c *Checkin) BeforeCreate() (err error) {
	if c.CreatedAt.IsZero() {
		c.CreatedAt = time.Now().UTC()
		c.UpdatedAt = time.Now().UTC()
	}
	return
}

// BeforeCreate for checkinHit will set CreatedAt to UTC
func (c *CheckinHit) BeforeCreate() (err error) {
	if c.CreatedAt.IsZero() {
		c.CreatedAt = time.Now().UTC()
	}
	return
}

// InsertCore create the single row for the Core settings in Statup
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

// Connect will attempt to connect to the sqlite, postgres, or mysql database
func (db *DbConfig) Connect(retry bool, location string) error {
	if DbSession != nil {
		return nil
	}
	var conn, dbType string
	var err error
	dbType = Configs.DbConn
	if Configs.DbPort == 0 {
		Configs.DbPort = DefaultPort(dbType)
	}
	switch dbType {
	case "sqlite":
		conn = location + "/statup.db"
		dbType = "sqlite3"
	case "mysql":
		host := fmt.Sprintf("%v:%v", Configs.DbHost, Configs.DbPort)
		conn = fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=UTC", Configs.DbUser, Configs.DbPass, host, Configs.DbData)
	case "postgres":
		conn = fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v timezone=UTC sslmode=disable", Configs.DbHost, Configs.DbPort, Configs.DbUser, Configs.DbData, Configs.DbPass)
	case "mssql":
		host := fmt.Sprintf("%v:%v", Configs.DbHost, Configs.DbPort)
		conn = fmt.Sprintf("sqlserver://%v:%v@%v?database=%v", Configs.DbUser, Configs.DbPass, host, Configs.DbData)
	}
	dbSession, err := gorm.Open(dbType, conn)
	if err != nil {
		if retry {
			utils.Log(1, fmt.Sprintf("Database connection to '%v' is not available, trying again in 5 seconds...", conn))
			return db.waitForDb()
		} else {
			return err
		}
	}
	err = dbSession.DB().Ping()
	if err == nil {
		DbSession = dbSession
		utils.Log(1, fmt.Sprintf("Database %v connection '%v@%v:%v' at %v was successful.", dbType, Configs.DbUser, Configs.DbHost, Configs.DbPort, Configs.DbData))
	}
	return err
}

// waitForDb will sleep for 5 seconds and try to connect to the database again
func (db *DbConfig) waitForDb() error {
	time.Sleep(5 * time.Second)
	return db.Connect(true, utils.Directory)
}

// DatabaseMaintence will automatically delete old records from 'failures' and 'hits'
// this function is currently set to delete records 7+ days old every 60 minutes
func DatabaseMaintence() {
	for range time.Tick(60 * time.Minute) {
		utils.Log(1, "Checking for database records older than 3 months...")
		since := time.Now().AddDate(0, -3, 0).UTC()
		DeleteAllSince("failures", since)
		DeleteAllSince("hits", since)
	}
}

// DeleteAllSince will delete a specific table's records based on a time.
func DeleteAllSince(table string, date time.Time) {
	sql := fmt.Sprintf("DELETE FROM %v WHERE created_at < '%v';", table, date.Format("2006-01-02"))
	db := DbSession.Exec(sql)
	if db.Error != nil {
		utils.Log(2, db.Error)
	}
}

// Update will save the config.yml file
func (db *DbConfig) Update() error {
	var err error
	config, err := os.Create(utils.Directory + "/config.yml")
	if err != nil {
		utils.Log(4, err)
		return err
	}
	data, err := yaml.Marshal(db)
	if err != nil {
		utils.Log(3, err)
		return err
	}
	config.WriteString(string(data))
	config.Close()
	return err
}

// Save will initially create the config.yml file
func (db *DbConfig) Save() (*DbConfig, error) {
	var err error
	config, err := os.Create(utils.Directory + "/config.yml")
	if err != nil {
		utils.Log(4, err)
		return nil, err
	}
	db.ApiKey = utils.NewSHA1Hash(16)
	db.ApiSecret = utils.NewSHA1Hash(16)
	data, err := yaml.Marshal(db)
	if err != nil {
		utils.Log(3, err)
		return nil, err
	}
	config.WriteString(string(data))
	defer config.Close()
	return db, err
}

// CreateCore will initialize the global variable 'CoreApp". This global variable contains most of Statup app.
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

// DropDatabase will DROP each table Statup created
func (db *DbConfig) DropDatabase() error {
	utils.Log(1, "Dropping Database Tables...")
	err := DbSession.DropTableIfExists("checkins")
	err = DbSession.DropTableIfExists("checkin_hits")
	err = DbSession.DropTableIfExists("notifications")
	err = DbSession.DropTableIfExists("core")
	err = DbSession.DropTableIfExists("failures")
	err = DbSession.DropTableIfExists("hits")
	err = DbSession.DropTableIfExists("services")
	err = DbSession.DropTableIfExists("users")
	err = DbSession.DropTableIfExists("messages")
	return err.Error
}

// CreateDatabase will CREATE TABLES for each of the Statup elements
func (db *DbConfig) CreateDatabase() error {
	var err error
	utils.Log(1, "Creating Database Tables...")
	for _, table := range DbModels {
		if err := DbSession.CreateTable(table); err.Error != nil {
			return err.Error
		}
	}
	if err := DbSession.Table("core").CreateTable(&types.Core{}); err.Error != nil {
		return err.Error
	}
	utils.Log(1, "Statup Database Created")
	return err
}

// MigrateDatabase will migrate the database structure to current version.
// This function will NOT remove previous records, tables or columns from the database.
// If this function has an issue, it will ROLLBACK to the previous state.
func (db *DbConfig) MigrateDatabase() error {
	utils.Log(1, "Migrating Database Tables...")
	tx := DbSession.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}
	for _, table := range DbModels {
		tx = tx.AutoMigrate(table)
	}
	if err := tx.Table("core").AutoMigrate(&types.Core{}); err.Error != nil {
		tx.Rollback()
		utils.Log(3, fmt.Sprintf("Statup Database could not be migrated: %v", tx.Error))
		return tx.Error
	}
	utils.Log(1, "Statup Database Migrated")
	return tx.Commit().Error
}
