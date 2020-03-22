package database

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/statping/statping/utils"
	"strings"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	TIME_NANO  = "2006-01-02T15:04:05Z"
	TIME       = "2006-01-02 15:04:05"
	CHART_TIME = "2006-01-02T15:04:05.999999-07:00"
	TIME_DAY   = "2006-01-02"
)

var database Database

// Database is an interface which DB implements
type Database interface {
	Close() error
	DB() *sql.DB
	New() Database
	NewScope(value interface{}) *gorm.Scope
	CommonDB() gorm.SQLCommon
	Callback() *gorm.Callback
	SetLogger(l gorm.Logger)
	LogMode(enable bool) Database
	SingularTable(enable bool)
	Where(query interface{}, args ...interface{}) Database
	Or(query interface{}, args ...interface{}) Database
	Not(query interface{}, args ...interface{}) Database
	Limit(value int) Database
	Offset(value int) Database
	Order(value string, reorder ...bool) Database
	Select(query interface{}, args ...interface{}) Database
	Omit(columns ...string) Database
	Group(query string) Database
	Having(query string, values ...interface{}) Database
	Joins(query string, args ...interface{}) Database
	Scopes(funcs ...func(*gorm.DB) *gorm.DB) Database
	Unscoped() Database
	Attrs(attrs ...interface{}) Database
	Assign(attrs ...interface{}) Database
	First(out interface{}, where ...interface{}) Database
	Last(out interface{}, where ...interface{}) Database
	Find(out interface{}, where ...interface{}) Database
	Scan(dest interface{}) Database
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	ScanRows(rows *sql.Rows, result interface{}) error
	Pluck(column string, value interface{}) Database
	Count(value interface{}) Database
	Related(value interface{}, foreignKeys ...string) Database
	FirstOrInit(out interface{}, where ...interface{}) Database
	FirstOrCreate(out interface{}, where ...interface{}) Database
	Update(attrs ...interface{}) Database
	Updates(values interface{}, ignoreProtectedAttrs ...bool) Database
	UpdateColumn(attrs ...interface{}) Database
	UpdateColumns(values interface{}) Database
	Save(value interface{}) Database
	Create(value interface{}) Database
	Delete(value interface{}, where ...interface{}) Database
	Raw(sql string, values ...interface{}) Database
	Exec(sql string, values ...interface{}) Database
	Model(value interface{}) Database
	Table(name string) Database
	Debug() Database
	Begin() Database
	Commit() Database
	Rollback() Database
	NewRecord(value interface{}) bool
	RecordNotFound() bool
	CreateTable(values ...interface{}) Database
	DropTable(values ...interface{}) Database
	DropTableIfExists(values ...interface{}) Database
	HasTable(value interface{}) bool
	AutoMigrate(values ...interface{}) Database
	ModifyColumn(column string, typ string) Database
	DropColumn(column string) Database
	AddIndex(indexName string, column ...string) Database
	AddUniqueIndex(indexName string, column ...string) Database
	RemoveIndex(indexName string) Database
	AddForeignKey(field string, dest string, onDelete string, onUpdate string) Database
	Association(column string) *gorm.Association
	Preload(column string, conditions ...interface{}) Database
	Set(name string, value interface{}) Database
	InstantSet(name string, value interface{}) Database
	Get(name string) (value interface{}, ok bool)
	SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface)
	AddError(err error) error
	GetErrors() (errors []error)

	// extra
	Error() error
	RowsAffected() int64

	Since(time.Time) Database
	Between(time.Time, time.Time) Database

	SelectByTime(time.Duration) string
	MultipleSelects(args ...string) Database

	FormatTime(t time.Time) string
	ParseTime(t string) (time.Time, error)
	DbType() string
}

func (it *Db) DbType() string {
	return it.Database.Dialect().GetName()
}

func Close(db Database) error {
	if db == nil {
		return nil
	}
	return db.Close()
}

func LogMode(db Database, b bool) Database {
	return db.LogMode(b)
}

func Begin(db Database, model interface{}) Database {
	if all, ok := model.(string); ok {
		if all == "migration" {
			return db.Begin()
		}
	}
	return db.Model(model).Begin()
}

func Available(db Database) bool {
	if db == nil {
		return false
	}
	if err := db.DB().Ping(); err != nil {
		return false
	}
	return true
}

func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
	return db.Where("service = ?", 1000)
}

func (it *Db) MultipleSelects(args ...string) Database {
	joined := strings.Join(args, ", ")
	return it.Select(joined)
}

type Db struct {
	Database *gorm.DB
	Type     string
}

// Openw is a drop-in replacement for Open()
func Openw(dialect string, args ...interface{}) (db Database, err error) {
	gorm.NowFunc = func() time.Time {
		return time.Now().UTC()
	}
	gormdb, err := gorm.Open(dialect, args...)
	if err != nil {
		return nil, err
	}
	database = Wrap(gormdb)
	return database, err
}

func OpenTester() (Database, error) {
	testDB := utils.Getenv("TEST_DB", "sqlite3").(string)
	var dbParamsstring string
	switch testDB {
	case "mysql":
		dbParamsstring = fmt.Sprintf("root:password123@tcp(localhost:3306)/statping?charset=utf8&parseTime=True&loc=UTC&time_zone=%%27UTC%%27")
	case "postgres":
		dbParamsstring = fmt.Sprintf("host=localhost port=5432 user=root dbname=statping password=password123 timezone=UTC")
	default:
		dbParamsstring = fmt.Sprintf("file:%s?mode=memory&cache=shared", utils.RandomString(12))
	}
	fmt.Println(testDB, dbParamsstring)
	newDb, err := Openw(testDB, dbParamsstring)
	if err != nil {
		return nil, err
	}
	newDb.DB().SetMaxOpenConns(1)
	return newDb, err
}

// Wrap wraps gorm.DB in an interface
func Wrap(db *gorm.DB) Database {
	return &Db{
		Database: db,
		Type:     db.Dialect().GetName(),
	}
}

func (it *Db) Close() error {
	return it.Database.Close()
}

func (it *Db) DB() *sql.DB {
	return it.Database.DB()
}

func (it *Db) New() Database {
	return Wrap(it.Database.New())
}

func (it *Db) NewScope(value interface{}) *gorm.Scope {
	return it.Database.NewScope(value)
}

func (it *Db) CommonDB() gorm.SQLCommon {
	return it.Database.CommonDB()
}

func (it *Db) Callback() *gorm.Callback {
	return it.Database.Callback()
}

func (it *Db) SetLogger(log gorm.Logger) {
	it.Database.SetLogger(log)
}

func (it *Db) LogMode(enable bool) Database {
	return Wrap(it.Database.LogMode(enable))
}

func (it *Db) SingularTable(enable bool) {
	it.Database.SingularTable(enable)
}

func (it *Db) Where(query interface{}, args ...interface{}) Database {
	return Wrap(it.Database.Where(query, args...))
}

func (it *Db) Or(query interface{}, args ...interface{}) Database {
	return Wrap(it.Database.Or(query, args...))
}

func (it *Db) Not(query interface{}, args ...interface{}) Database {
	return Wrap(it.Database.Not(query, args...))
}

func (it *Db) Limit(value int) Database {
	return Wrap(it.Database.Limit(value))
}

func (it *Db) Offset(value int) Database {
	return Wrap(it.Database.Offset(value))
}

func (it *Db) Order(value string, reorder ...bool) Database {
	return Wrap(it.Database.Order(value, reorder...))
}

func (it *Db) Select(query interface{}, args ...interface{}) Database {
	return Wrap(it.Database.Select(query, args...))
}

func (it *Db) Omit(columns ...string) Database {
	return Wrap(it.Database.Omit(columns...))
}

func (it *Db) Group(query string) Database {
	return Wrap(it.Database.Group(query))
}

func (it *Db) Having(query string, values ...interface{}) Database {
	return Wrap(it.Database.Having(query, values...))
}

func (it *Db) Joins(query string, args ...interface{}) Database {
	return Wrap(it.Database.Joins(query, args...))
}

func (it *Db) Scopes(funcs ...func(*gorm.DB) *gorm.DB) Database {
	return Wrap(it.Database.Scopes(funcs...))
}

func (it *Db) Unscoped() Database {
	return Wrap(it.Database.Unscoped())
}

func (it *Db) Attrs(attrs ...interface{}) Database {
	return Wrap(it.Database.Attrs(attrs...))
}

func (it *Db) Assign(attrs ...interface{}) Database {
	return Wrap(it.Database.Assign(attrs...))
}

func (it *Db) First(out interface{}, where ...interface{}) Database {
	return Wrap(it.Database.First(out, where...))
}

func (it *Db) Last(out interface{}, where ...interface{}) Database {
	return Wrap(it.Database.Last(out, where...))
}

func (it *Db) Find(out interface{}, where ...interface{}) Database {
	return Wrap(it.Database.Find(out, where...))
}

func (it *Db) Scan(dest interface{}) Database {
	return Wrap(it.Database.Scan(dest))
}

func (it *Db) Row() *sql.Row {
	return it.Database.Row()
}

func (it *Db) Rows() (*sql.Rows, error) {
	return it.Database.Rows()
}

func (it *Db) ScanRows(rows *sql.Rows, result interface{}) error {
	return it.Database.ScanRows(rows, result)
}

func (it *Db) Pluck(column string, value interface{}) Database {
	return Wrap(it.Database.Pluck(column, value))
}

func (it *Db) Count(value interface{}) Database {
	return Wrap(it.Database.Count(value))
}

func (it *Db) Related(value interface{}, foreignKeys ...string) Database {
	return Wrap(it.Database.Related(value, foreignKeys...))
}

func (it *Db) FirstOrInit(out interface{}, where ...interface{}) Database {
	return Wrap(it.Database.FirstOrInit(out, where...))
}

func (it *Db) FirstOrCreate(out interface{}, where ...interface{}) Database {
	return Wrap(it.Database.FirstOrCreate(out, where...))
}

func (it *Db) Update(attrs ...interface{}) Database {
	return Wrap(it.Database.Update(attrs...))
}

func (it *Db) Updates(values interface{}, ignoreProtectedAttrs ...bool) Database {
	return Wrap(it.Database.Updates(values, ignoreProtectedAttrs...))
}

func (it *Db) UpdateColumn(attrs ...interface{}) Database {
	return Wrap(it.Database.UpdateColumn(attrs...))
}

func (it *Db) UpdateColumns(values interface{}) Database {
	return Wrap(it.Database.UpdateColumns(values))
}

func (it *Db) Save(value interface{}) Database {
	return Wrap(it.Database.Save(value))
}

func (it *Db) Create(value interface{}) Database {
	return Wrap(it.Database.Create(value))
}

func (it *Db) Delete(value interface{}, where ...interface{}) Database {
	return Wrap(it.Database.Delete(value, where...))
}

func (it *Db) Raw(sql string, values ...interface{}) Database {
	return Wrap(it.Database.Raw(sql, values...))
}

func (it *Db) Exec(sql string, values ...interface{}) Database {
	return Wrap(it.Database.Exec(sql, values...))
}

func (it *Db) Model(value interface{}) Database {
	return Wrap(it.Database.Model(value))
}

func (it *Db) Table(name string) Database {
	return Wrap(it.Database.Table(name))
}

func (it *Db) Debug() Database {
	return Wrap(it.Database.Debug())
}

func (it *Db) Begin() Database {
	return Wrap(it.Database.Begin())
}

func (it *Db) Commit() Database {
	return Wrap(it.Database.Commit())
}

func (it *Db) Rollback() Database {
	return Wrap(it.Database.Rollback())
}

func (it *Db) NewRecord(value interface{}) bool {
	return it.Database.NewRecord(value)
}

func (it *Db) RecordNotFound() bool {
	return it.Database.RecordNotFound()
}

func (it *Db) CreateTable(values ...interface{}) Database {
	return Wrap(it.Database.CreateTable(values...))
}

func (it *Db) DropTable(values ...interface{}) Database {
	return Wrap(it.Database.DropTable(values...))
}

func (it *Db) DropTableIfExists(values ...interface{}) Database {
	return Wrap(it.Database.DropTableIfExists(values...))
}

func (it *Db) HasTable(value interface{}) bool {
	return it.Database.HasTable(value)
}

func (it *Db) AutoMigrate(values ...interface{}) Database {
	return Wrap(it.Database.AutoMigrate(values...))
}

func (it *Db) ModifyColumn(column string, typ string) Database {
	return Wrap(it.Database.ModifyColumn(column, typ))
}

func (it *Db) DropColumn(column string) Database {
	return Wrap(it.Database.DropColumn(column))
}

func (it *Db) AddIndex(indexName string, columns ...string) Database {
	return Wrap(it.Database.AddIndex(indexName, columns...))
}

func (it *Db) AddUniqueIndex(indexName string, columns ...string) Database {
	return Wrap(it.Database.AddUniqueIndex(indexName, columns...))
}

func (it *Db) RemoveIndex(indexName string) Database {
	return Wrap(it.Database.RemoveIndex(indexName))
}

func (it *Db) Association(column string) *gorm.Association {
	return it.Database.Association(column)
}

func (it *Db) Preload(column string, conditions ...interface{}) Database {
	return Wrap(it.Database.Preload(column, conditions...))
}

func (it *Db) Set(name string, value interface{}) Database {
	return Wrap(it.Database.Set(name, value))
}

func (it *Db) InstantSet(name string, value interface{}) Database {
	return Wrap(it.Database.InstantSet(name, value))
}

func (it *Db) Get(name string) (interface{}, bool) {
	return it.Database.Get(name)
}

func (it *Db) SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface) {
	it.Database.SetJoinTableHandler(source, column, handler)
}

func (it *Db) AddForeignKey(field string, dest string, onDelete string, onUpdate string) Database {
	return Wrap(it.Database.AddForeignKey(field, dest, onDelete, onUpdate))
}

func (it *Db) AddError(err error) error {
	return it.Database.AddError(err)
}

func (it *Db) GetErrors() (errors []error) {
	return it.Database.GetErrors()
}

func (it *Db) RowsAffected() int64 {
	return it.Database.RowsAffected
}

func (it *Db) Error() error {
	return it.Database.Error
}

func (it *Db) Since(ago time.Time) Database {
	return it.Where("created_at > ?", it.FormatTime(ago))
}

func (it *Db) Between(t1 time.Time, t2 time.Time) Database {
	return it.Where("created_at BETWEEN ? AND ?", it.FormatTime(t1), it.FormatTime(t2))
}

type TimeValue struct {
	Timeframe string `json:"timeframe"`
	Amount    int64  `json:"amount"`
}
