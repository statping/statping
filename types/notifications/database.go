package notifications

import (
	"github.com/statping/statping/database"
)

var (
	db database.Database
)

func SetDB(database database.Database) {
	db = database.Model(&Notification{})
}

func Find(method string) (*Notification, error) {
	var notification Notification
	q := db.Where("method = ?", method).Find(&notification)
	return &notification, q.Error()
}

func (n *Notification) Create() error {
	q := db.Where("method = ?", n.Method).Find(n)
	if q.RecordNotFound() {
		if err := db.Create(n).Error(); err != nil {
			return err
		}
	}
	return nil
}

func (n *Notification) Update() error {
	n.ResetQueue()
	if n.Enabled.Bool {
		n.Close()
		n.Start()
	} else {
		n.Close()
	}
	err := db.Update(n)
	return err.Error()
}

func loadAll() []*Notification {
	var notifications []*Notification
	db.Find(&notifications)
	return notifications
}
