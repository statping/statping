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
	var n Notification
	q := db.Where("method = ?", method).Find(&n)
	if &n == nil {
		return nil, errors.New("cannot find notifier")
	}
	n.UpdateFields(&n)
	return &n, q.Error()
}

func (n *Notification) Create() error {
	var p Notification
	q := db.Where("method = ?", n.Method).Find(&p)
	if q.RecordNotFound() {
		if err := db.Create(n).Error(); err != nil {
			return err
		}
		return nil
	}
	if p.FailureData == "" {
		p.FailureData = n.FailureData
	}
	if p.SuccessData == "" {
		p.SuccessData = n.SuccessData
	}
	if err := p.Update(); err != nil {
		return err
	}
	return nil
}

func (n *Notification) UpdateFields(notif *Notification) *Notification {
	n.Limits = notif.Limits
	n.Enabled = notif.Enabled
	n.Host = notif.Host
	n.Port = notif.Port
	n.Username = notif.Username
	n.Password = notif.Password
	n.ApiKey = notif.ApiKey
	n.ApiSecret = notif.ApiSecret
	n.Var1 = notif.Var1
	n.Var2 = notif.Var2
	n.SuccessData = notif.SuccessData
	n.FailureData = notif.FailureData
	return n
}

func (n *Notification) Update() error {
	if err := db.Update(n).Error(); err != nil {
		return err
	}
	return nil
}
