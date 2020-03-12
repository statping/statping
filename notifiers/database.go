package notifiers

import (
	"errors"
	"github.com/statping/statping/database"
)

var db database.Database

func SetDB(database database.Database) {
	db = database.Model(&Notification{})
}

func appendList(n Notifier) {
	allNotifiers = append(allNotifiers, n)
}

func Find(name string) (*Notification, error) {
	for _, n := range allNotifiers {
		notif := n.(*Notification)
		if notif.Method == name {
			return notif, nil
		}
	}
	return nil, errors.New("notifier not found")
}

func All() []Notifier {
	return allNotifiers
}

func (n *Notification) Create() error {
	q := db.Where("method = ?", n.Method).Find(n)
	if q.RecordNotFound() {
		if err := db.Create(n).Error(); err != nil {
			return err
		}
	}
	appendList(n)
	return nil
}

func (n *Notification) Update() error {
	n.ResetQueue()
	if n.Enabled.Bool {
		n.Close()
		n.Start()
		go Queue(n)
	} else {
		n.Close()
	}
	err := db.Update(n)
	return err.Error()
}

func (n *Notification) Delete() error {
	return nil
}
