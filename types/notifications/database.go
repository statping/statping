package notifications

import (
	"errors"
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
	if &notification == nil {
		return nil, errors.New("cannot find notifier")
	}
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

func (n *Notification) UpdateFields(notif *Notification) *Notification {
	n.Enabled = notif.Enabled
	n.Host = notif.Host
	n.Port = notif.Port
	n.Username = notif.Username
	n.Password = notif.Password
	n.ApiKey = notif.ApiKey
	n.ApiSecret = notif.ApiSecret
	n.Var1 = notif.Var1
	n.Var2 = notif.Var2
	return n
}

func (n *Notification) Update() error {
	if err := db.Update(n); err.Error() != nil {
		return err.Error()
	}
	n.ResetQueue()
	if n.Enabled.Bool {
		n.Close()
		n.Start()
	} else {
		n.Close()
	}
	return nil
}

func loadAll() []*Notification {
	var notifications []*Notification
	db.Find(&notifications)
	return notifications
}
