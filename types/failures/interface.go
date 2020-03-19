package failures

import (
	"fmt"
	"github.com/statping/statping/database"
	"time"
)

type ColumnIDInterfacer interface {
	FailuresColumnID() (string, int64)
}

type Failurer struct {
	db database.Database
}

func (f Failurer) Db() database.Database {
	return f.db
}

func (f Failurer) First() *Failure {
	var fail Failure
	f.db.Order("id ASC").Limit(1).Find(&fail)
	return &fail
}

func (f Failurer) Last() *Failure {
	var fail Failure
	f.db.Order("id DESC").Limit(1).Find(&fail)
	return &fail
}

func (f Failurer) List() []*Failure {
	var fails []*Failure
	f.db.Find(&fails)
	return fails
}

func (f Failurer) LastAmount(amount int) []*Failure {
	var fail []*Failure
	f.db.Order("id asc").Limit(amount).Find(&fail)
	return fail
}

func (f Failurer) Since(t time.Time) []*Failure {
	var fails []*Failure
	f.db.Since(t).Find(&fails)
	return fails
}

func (f Failurer) Count() int {
	var amount int
	f.db.Count(&amount)
	return amount
}

func (f Failurer) DeleteAll() error {
	q := f.db.Delete(&Failure{})
	return q.Error()
}

func AllFailures(obj ColumnIDInterfacer) Failurer {
	column, id := obj.FailuresColumnID()
	return Failurer{db.Where(fmt.Sprintf("%s = ?", column), id)}
}

func Since(t time.Time, obj ColumnIDInterfacer) Failurer {
	column, id := obj.FailuresColumnID()
	timestamp := db.FormatTime(t)
	return Failurer{db.Where(fmt.Sprintf("%s = ? AND created_at > ?", column), id, timestamp)}
}
