package database

import (
	"github.com/hunterlong/statping/types"
)

type failures struct {
	DB Database
}

func (f *failures) All() []*types.Failure {
	var fails []*types.Failure
	f.DB = f.DB.Find(&fails)
	return fails
}

func (f *failures) Last(amount int) *types.Failure {
	var fail types.Failure
	f.DB = f.DB.Limit(amount).Find(&fail)
	return &fail
}

func (f *failures) Count() int {
	var amount int
	f.DB = f.DB.Count(&amount)
	return amount
}

func (f *failures) Find(data interface{}) error {
	q := f.Find(&data)
	return q
}
