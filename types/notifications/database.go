package notifications

import (
	"errors"
	"github.com/statping/statping/database"
)

var db database.Database

func SetDB(database database.Database) {
	db = database.Model(&Notification{})
}

func Append(n Notifier) {
	allNotifiers = append(allNotifiers, n)
}

func Find(name string) (*Notification, error) {
	for _, n := range allNotifiers {
		notif := n.Select()
		if notif.Name() == name || notif.Method == name {
			return notif, nil
		}
	}
	return nil, errors.New("notifier not found")
}

func All() []Notifier {
	return allNotifiers
}

func (n *Notification) Create() error {
	var notif Notification
	if db.Where("method = ?", n.Method).Find(&notif).RecordNotFound() {
		Append(n)
		return db.Create(n).Error()
	}
	Append(n)
	return nil
}

func (n *Notification) Update() error {
	n.ResetQueue()
	if n.Enabled.Bool {
		n.Close()
		n.Start()
		go Queue(Notifier(n))
	} else {
		n.Close()
	}
	err := db.Update(n)
	return err.Error()
}

func (n *Notification) Delete() error {
	return nil
}
