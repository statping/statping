// Statup
// Copyright (C) 2020.  Hunter Long and the project contributors
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

package types

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"time"
)

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
	QuerySearch(*http.Request) Database

	Since(time.Time) Database
	Between(time.Time, time.Time) Database
	Hits() ([]*Hit, error)
	ToChart() ([]*DateScan, error)

	Failurer
}

type Failurer interface {
	Failures(id int64) Database
	Fails() ([]*Failure, error)
}

func (it *database) Failures(id int64) Database {
	return it.Model(&Failure{}).Where("service = ?", id).Not("method = 'checkin'").Order("id desc")
}

func (it *database) Fails() ([]*Failure, error) {
	var fails []*Failure
	err := it.Find(&fails)
	return fails, err.Error()
}

type database struct {
	w *gorm.DB
}

// Openw is a drop-in replacement for Open()
func Openw(dialect string, args ...interface{}) (db Database, err error) {
	gormdb, err := gorm.Open(dialect, args...)
	return Wrap(gormdb), err
}

// Wrap wraps gorm.DB in an interface
func Wrap(db *gorm.DB) Database {
	return &database{db}
}

func (it *database) Close() error {
	return it.w.Close()
}

func (it *database) DB() *sql.DB {
	return it.w.DB()
}

func (it *database) New() Database {
	return Wrap(it.w.New())
}

func (it *database) NewScope(value interface{}) *gorm.Scope {
	return it.w.NewScope(value)
}

func (it *database) CommonDB() gorm.SQLCommon {
	return it.w.CommonDB()
}

func (it *database) Callback() *gorm.Callback {
	return it.w.Callback()
}

func (it *database) SetLogger(log gorm.Logger) {
	it.w.SetLogger(log)
}

func (it *database) LogMode(enable bool) Database {
	return Wrap(it.w.LogMode(enable))
}

func (it *database) SingularTable(enable bool) {
	it.w.SingularTable(enable)
}

func (it *database) Where(query interface{}, args ...interface{}) Database {
	return Wrap(it.w.Where(query, args...))
}

func (it *database) Or(query interface{}, args ...interface{}) Database {
	return Wrap(it.w.Or(query, args...))
}

func (it *database) Not(query interface{}, args ...interface{}) Database {
	return Wrap(it.w.Not(query, args...))
}

func (it *database) Limit(value int) Database {
	return Wrap(it.w.Limit(value))
}

func (it *database) Offset(value int) Database {
	return Wrap(it.w.Offset(value))
}

func (it *database) Order(value string, reorder ...bool) Database {
	return Wrap(it.w.Order(value, reorder...))
}

func (it *database) Select(query interface{}, args ...interface{}) Database {
	return Wrap(it.w.Select(query, args...))
}

func (it *database) Omit(columns ...string) Database {
	return Wrap(it.w.Omit(columns...))
}

func (it *database) Group(query string) Database {
	return Wrap(it.w.Group(query))
}

func (it *database) Having(query string, values ...interface{}) Database {
	return Wrap(it.w.Having(query, values...))
}

func (it *database) Joins(query string, args ...interface{}) Database {
	return Wrap(it.w.Joins(query, args...))
}

func (it *database) Scopes(funcs ...func(*gorm.DB) *gorm.DB) Database {
	return Wrap(it.w.Scopes(funcs...))
}

func (it *database) Unscoped() Database {
	return Wrap(it.w.Unscoped())
}

func (it *database) Attrs(attrs ...interface{}) Database {
	return Wrap(it.w.Attrs(attrs...))
}

func (it *database) Assign(attrs ...interface{}) Database {
	return Wrap(it.w.Assign(attrs...))
}

func (it *database) First(out interface{}, where ...interface{}) Database {
	return Wrap(it.w.First(out, where...))
}

func (it *database) Last(out interface{}, where ...interface{}) Database {
	return Wrap(it.w.Last(out, where...))
}

func (it *database) Find(out interface{}, where ...interface{}) Database {
	return Wrap(it.w.Find(out, where...))
}

func (it *database) Scan(dest interface{}) Database {
	return Wrap(it.w.Scan(dest))
}

func (it *database) Row() *sql.Row {
	return it.w.Row()
}

func (it *database) Rows() (*sql.Rows, error) {
	return it.w.Rows()
}

func (it *database) ScanRows(rows *sql.Rows, result interface{}) error {
	return it.w.ScanRows(rows, result)
}

func (it *database) Pluck(column string, value interface{}) Database {
	return Wrap(it.w.Pluck(column, value))
}

func (it *database) Count(value interface{}) Database {
	return Wrap(it.w.Count(value))
}

func (it *database) Related(value interface{}, foreignKeys ...string) Database {
	return Wrap(it.w.Related(value, foreignKeys...))
}

func (it *database) FirstOrInit(out interface{}, where ...interface{}) Database {
	return Wrap(it.w.FirstOrInit(out, where...))
}

func (it *database) FirstOrCreate(out interface{}, where ...interface{}) Database {
	return Wrap(it.w.FirstOrCreate(out, where...))
}

func (it *database) Update(attrs ...interface{}) Database {
	return Wrap(it.w.Update(attrs...))
}

func (it *database) Updates(values interface{}, ignoreProtectedAttrs ...bool) Database {
	return Wrap(it.w.Updates(values, ignoreProtectedAttrs...))
}

func (it *database) UpdateColumn(attrs ...interface{}) Database {
	return Wrap(it.w.UpdateColumn(attrs...))
}

func (it *database) UpdateColumns(values interface{}) Database {
	return Wrap(it.w.UpdateColumns(values))
}

func (it *database) Save(value interface{}) Database {
	return Wrap(it.w.Save(value))
}

func (it *database) Create(value interface{}) Database {
	return Wrap(it.w.Create(value))
}

func (it *database) Delete(value interface{}, where ...interface{}) Database {
	return Wrap(it.w.Delete(value, where...))
}

func (it *database) Raw(sql string, values ...interface{}) Database {
	return Wrap(it.w.Raw(sql, values...))
}

func (it *database) Exec(sql string, values ...interface{}) Database {
	return Wrap(it.w.Exec(sql, values...))
}

func (it *database) Model(value interface{}) Database {
	return Wrap(it.w.Model(value))
}

func (it *database) Table(name string) Database {
	return Wrap(it.w.Table(name))
}

func (it *database) Debug() Database {
	return Wrap(it.w.Debug())
}

func (it *database) Begin() Database {
	return Wrap(it.w.Begin())
}

func (it *database) Commit() Database {
	return Wrap(it.w.Commit())
}

func (it *database) Rollback() Database {
	return Wrap(it.w.Rollback())
}

func (it *database) NewRecord(value interface{}) bool {
	return it.w.NewRecord(value)
}

func (it *database) RecordNotFound() bool {
	return it.w.RecordNotFound()
}

func (it *database) CreateTable(values ...interface{}) Database {
	return Wrap(it.w.CreateTable(values...))
}

func (it *database) DropTable(values ...interface{}) Database {
	return Wrap(it.w.DropTable(values...))
}

func (it *database) DropTableIfExists(values ...interface{}) Database {
	return Wrap(it.w.DropTableIfExists(values...))
}

func (it *database) HasTable(value interface{}) bool {
	return it.w.HasTable(value)
}

func (it *database) AutoMigrate(values ...interface{}) Database {
	return Wrap(it.w.AutoMigrate(values...))
}

func (it *database) ModifyColumn(column string, typ string) Database {
	return Wrap(it.w.ModifyColumn(column, typ))
}

func (it *database) DropColumn(column string) Database {
	return Wrap(it.w.DropColumn(column))
}

func (it *database) AddIndex(indexName string, columns ...string) Database {
	return Wrap(it.w.AddIndex(indexName, columns...))
}

func (it *database) AddUniqueIndex(indexName string, columns ...string) Database {
	return Wrap(it.w.AddUniqueIndex(indexName, columns...))
}

func (it *database) RemoveIndex(indexName string) Database {
	return Wrap(it.w.RemoveIndex(indexName))
}

func (it *database) Association(column string) *gorm.Association {
	return it.w.Association(column)
}

func (it *database) Preload(column string, conditions ...interface{}) Database {
	return Wrap(it.w.Preload(column, conditions...))
}

func (it *database) Set(name string, value interface{}) Database {
	return Wrap(it.w.Set(name, value))
}

func (it *database) InstantSet(name string, value interface{}) Database {
	return Wrap(it.w.InstantSet(name, value))
}

func (it *database) Get(name string) (interface{}, bool) {
	return it.w.Get(name)
}

func (it *database) SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface) {
	it.w.SetJoinTableHandler(source, column, handler)
}

func (it *database) AddForeignKey(field string, dest string, onDelete string, onUpdate string) Database {
	return Wrap(it.w.AddForeignKey(field, dest, onDelete, onUpdate))
}

func (it *database) AddError(err error) error {
	return it.w.AddError(err)
}

func (it *database) GetErrors() (errors []error) {
	return it.w.GetErrors()
}

func (it *database) RowsAffected() int64 {
	return it.w.RowsAffected
}

func (it *database) Error() error {
	return it.w.Error
}

func (it *database) Hits() ([]*Hit, error) {
	var hits []*Hit
	err := it.Find(&hits)
	return hits, err.Error()
}

func (it *database) Since(ago time.Time) Database {
	return it.Where("created_at > ?", ago.UTC().Format(TIME))
}

func (it *database) Between(t1 time.Time, t2 time.Time) Database {
	return it.Where("created_at BETWEEN ? AND ?", t1.UTC().Format(TIME), t2.UTC().Format(TIME))
}

// DateScan struct is for creating the charts.js graph JSON array
type DateScan struct {
	CreatedAt string `json:"x,omitempty"`
	Value     int64  `json:"y"`
}

func (it *database) ToChart() ([]*DateScan, error) {
	rows, err := it.w.Rows()
	if err != nil {
		return nil, err
	}
	var data []*DateScan
	for rows.Next() {
		gd := new(DateScan)
		var createdAt string
		var value float64
		if err := rows.Scan(&createdAt, &value); err != nil {
			return nil, err
		}
		createdTime, err := time.Parse(TIME, createdAt)
		if err != nil {
			return nil, err
		}
		gd.CreatedAt = createdTime.UTC().String()
		gd.Value = int64(value * 1000)
		data = append(data, gd)
	}
	return data, err
}

func (it *database) QuerySearch(r *http.Request) Database {
	if r == nil {
		return it
	}
	db := it.w
	start := defaultField(r, "start")
	end := defaultField(r, "end")
	limit := defaultField(r, "limit")
	offset := defaultField(r, "offset")
	params := &Params{
		Start:  start,
		End:    end,
		Limit:  limit,
		Offset: offset,
	}
	if params.Start != nil && params.End != nil {
		db = db.Where("created_at BETWEEN ? AND ?", time.Unix(*params.Start, 0).Format(TIME), time.Unix(*params.End, 0).UTC().Format(TIME))
	} else if params.Start != nil && params.End == nil {
		db = db.Where("created_at > ?", time.Unix(*params.Start, 0).UTC().Format(TIME))
	}
	if params.Limit != nil {
		db = db.Limit(*params.Limit)
	} else {
		db = db.Limit(10000)
	}
	if params.Offset != nil {
		db = db.Offset(*params.Offset)
	} else {
		db = db.Offset(0)
	}
	return Wrap(db)
}

type Params struct {
	Start  *int64
	End    *int64
	Limit  *int64
	Offset *int64
}

func defaultField(r *http.Request, key string) *int64 {
	r.ParseForm()
	val := r.Form.Get(key)
	if val == "" {
		return nil
	}

	gg, _ := strconv.Atoi(val)
	num := int64(gg)
	return &num
}
