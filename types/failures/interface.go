package failures

import (
	"fmt"
	"github.com/hunterlong/statping/database"
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

func (f Failurer) List() []*Failure {
	var fails []*Failure
	f.db.Find(&fails)
	return fails
}

func (f Failurer) Count() int {
	var amount int
	f.db.Count(&amount)
	return amount
}

func (f Failurer) Last(amount int) []*Failure {
	var fails []*Failure
	f.db.Limit(amount).Find(&fails)
	return fails
}

func (f Failurer) Since(t time.Time) []*Failure {
	var fails []*Failure
	f.db.Since(t).Find(&fails)
	return fails
}

func AllFailures(obj ColumnIDInterfacer) Failurer {
	column, id := obj.FailuresColumnID()
	return Failurer{DB().Where(fmt.Sprintf("%s = ?", column), id)}
}

func FailuresSince(t time.Time, obj ColumnIDInterfacer) Failurer {
	column, id := obj.FailuresColumnID()
	timestamp := DB().FormatTime(t)
	return Failurer{DB().Where(fmt.Sprintf("%s = ? AND created_at > ?", column), id, timestamp)}
}
