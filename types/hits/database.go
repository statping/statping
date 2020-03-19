package hits

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/utils"
)

var log = utils.Log

var db database.Database

func SetDB(database database.Database) {
	db = database.Model(&Hit{})
}

func Find(id int64) (*Hit, error) {
	var group Hit
	q := db.Where("id = ?", id).Find(&group)
	return &group, q.Error()
}

func All() []*Hit {
	var hits []*Hit
	db.Find(&hits)
	return hits
}

func (h *Hit) Create() error {
	q := db.Create(h)
	return q.Error()
}

func (h *Hit) Update() error {
	q := db.Update(h)
	return q.Error()
}

func (h *Hit) Delete() error {
	q := db.Delete(h)
	return q.Error()
}
