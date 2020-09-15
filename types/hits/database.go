package hits

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/metrics"
	"github.com/statping/statping/utils"
	"gorm.io/gorm"
	"time"
)

var log = utils.Log

var db *database.Database

func SetDB(dbz *database.Database) {
	db = database.Wrap(dbz.Model(&Hit{}))
}

// BeforeCreate for Hit will set CreatedAt to UTC
func (h *Hit) BeforeCreate(*gorm.DB) error {
	if h.CreatedAt.IsZero() {
		h.CreatedAt = time.Now().UTC()
	}
	return nil
}

func (h *Hit) AfterFind(*gorm.DB) error {
	metrics.Query("hit", "find")
	return nil
}

func (h *Hit) AfterUpdate(*gorm.DB) error {
	metrics.Query("hit", "update")
	return nil
}

func (h *Hit) AfterDelete(*gorm.DB) error {
	metrics.Query("hit", "delete")
	return nil
}

func (h *Hit) AfterCreate(*gorm.DB) error {
	metrics.Query("hit", "create")
	return nil
}

func (h *Hit) Create() error {
	q := db.Create(h)
	return q.Error
}

func (h *Hit) Update() error {
	q := db.Save(h)
	return q.Error
}

func (h *Hit) Delete() error {
	q := db.Delete(h)
	return q.Error
}
