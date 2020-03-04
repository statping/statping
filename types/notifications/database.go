package notifications

import (
	"errors"
	"github.com/hunterlong/statping/database"
)

func DB() database.Database {
	return database.DB().Model(&Notification{})
}

func Find(method string) (Notifier, error) {
	for _, n := range AllCommunications {
		notif := n.Select()
		if notif.Method == method {
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

func Update(notifier Notifier) error {
	n := notifier.Select()
	n.ResetQueue()
	err := n.Update()
	if n.Enabled.Bool {
		n.Close()
		n.Start()
		go Queue(notifier)
	} else {
		n.Close()
	}
	return err
}

func (n *Notification) Update() error {
	return Update(n)
}

func (n *Notification) Delete() error {
	return nil
}
