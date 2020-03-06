package hits

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/utils"
)

var log = utils.Log

func DB() database.Database {
	return database.DB().Model(&Hit{})
}

func Find(id int64) (*Hit, error) {
	var group Hit
	db := DB().Where("id = ?", id).Find(&group)
	return &group, db.Error()
}

func All() []*Hit {
	var hits []*Hit
	DB().Find(&hits)
	return hits
}

func (h *Hit) Create() error {
	db := DB().Create(h)
	return db.Error()
}

func (h *Hit) Update() error {
	db := DB().Update(h)
	return db.Error()
}

func (h *Hit) Delete() error {
	db := DB().Delete(h)
	return db.Error()
}
