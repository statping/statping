package hits

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/metrics"
	"github.com/statping/statping/utils"
)

var log = utils.Log

var db database.Database

func SetDB(database database.Database) {
	db = database.Model(&Hit{})
}

func (h *Hit) AfterFind() {
	metrics.Query("hit", "find")
}

func (h *Hit) AfterUpdate() {
	metrics.Query("hit", "update")
}

func (h *Hit) AfterDelete() {
	metrics.Query("hit", "delete")
}

func (h *Hit) AfterCreate() {
	metrics.Query("hit", "create")
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
