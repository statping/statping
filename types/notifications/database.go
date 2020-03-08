package notifications

import (
	"errors"
	"github.com/hunterlong/statping/database"
)

func DB() database.Database {
	return database.DB().Model(&Notification{})
}

func Find(name string) (Notifier, error) {
	for _, n := range AllCommunications {
		notif := n.Select()
		if notif.Name() == name || notif.Method == name {
			return n, nil
		}
	}
	return nil, errors.New("notifier not found")
}

func All() []*Notification {
	var notifiers []*Notification
	DB().Find(&notifiers)
	return notifiers
}

func (n *Notification) Create() error {
	db := DB().FirstOrCreate(&n)
	return db.Error()
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
	err := DB().Update(n)
	return err.Error()
}

func (n *Notification) Delete() error {
	return nil
}
