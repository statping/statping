package database

import (
	"github.com/hunterlong/statping/types"
	"time"
)

type FailureObj struct {
	o *Object
}

type Failurer interface {
	Model() []*types.Failure
}

func (f *FailureObj) Model() []*types.Failure {
	return f.All()
}

func (f *FailureObj) All() []*types.Failure {
	var fails []*types.Failure
	f.o.db.Find(&fails)
	return fails
}

func AllFailures() int {
	var amount int
	database.Failures().Count(&amount)
	return amount
}

func (f *FailureObj) DeleteAll() error {
	query := database.Exec(`DELETE FROM failures WHERE service = ?`, f.o.Id)
	return query.Error()
}

func (f *FailureObj) Last(amount int) *types.Failure {
	var fail types.Failure
	f.o.db.Limit(amount).Last(&fail)
	return &fail
}

func (f *FailureObj) Count() int {
	var amount int
	f.o.db.Count(&amount)
	return amount
}

func (f *FailureObj) Since(t time.Time) []*types.Failure {
	var fails []*types.Failure
	f.o.db.Since(t).Find(&fails)
	return fails
}

func (f *FailureObj) object() *Object {
	return f.o
}
