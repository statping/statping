package failures

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/metrics"
)

var db database.Database

func SetDB(database database.Database) {
	db = database.Model(&Failure{})
}

func DB() database.Database {
	return db
}

func (f *Failure) AfterFind() {
	metrics.Query("failure", "find")
}

func (f *Failure) AfterUpdate() {
	metrics.Query("failure", "update")
}

func (f *Failure) AfterDelete() {
	metrics.Query("failure", "delete")
}

func (f *Failure) AfterCreate() {
	metrics.Query("failure", "create")
}

func (f *Failure) Create() error {
	q := db.Create(f)
	return q.Error()
}

func (f *Failure) Update() error {
	q := db.Update(f)
	return q.Error()
}

func (f *Failure) Delete() error {
	q := db.Delete(f)
	return q.Error()
}
