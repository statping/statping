package failures

import "github.com/hunterlong/statping/database"

func DB() database.Database {
	return database.DB().Model(&Failure{})
}

func Find(id int64) (*Failure, error) {
	var failure Failure
	db := DB().Where("id = ?", id).Find(&failure)
	return &failure, db.Error()
}

func All() []*Failure {
	var failures []*Failure
	DB().Find(&failures)
	return failures
}

func (f *Failure) Create() error {
	db := DB().Create(f)
	return db.Error()
}

func (f *Failure) Update() error {
	db := DB().Update(f)
	return db.Error()
}

func (f *Failure) Delete() error {
	db := DB().Delete(f)
	return db.Error()
}
