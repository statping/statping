package failures

import "github.com/statping/statping/database"

var db database.Database

func SetDB(database database.Database) {
	db = database.Model(&Failure{})
}

func DB() database.Database {
	return db
}

func All() []*Failure {
	var failures []*Failure
	db.Find(&failures)
	return failures
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
