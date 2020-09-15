package failures

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/metrics"
	"gorm.io/gorm"
)

var db *database.Database

func SetDB(dbz *database.Database) {
	db = database.Wrap(dbz.Model(&Failure{}))
}

func DB() *database.Database {
	return db
}

func (f *Failure) AfterFind(*gorm.DB) error {
	metrics.Query("failure", "find")
	return nil
}

func (f *Failure) AfterUpdate(*gorm.DB) error {
	metrics.Query("failure", "update")
	return nil
}

func (f *Failure) AfterDelete(*gorm.DB) error {
	metrics.Query("failure", "delete")
	return nil
}

func (f *Failure) AfterCreate(*gorm.DB) error {
	metrics.Query("failure", "create")
	return nil
}

func (f *Failure) Create() error {
	q := db.Create(f)
	return q.Error
}

func (f *Failure) Update() error {
	q := db.Save(f)
	return q.Error
}

func (f *Failure) Delete() error {
	q := db.Delete(f)
	return q.Error
}
