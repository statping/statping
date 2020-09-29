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

func (n *Notification) Values() Values {
	return Values{
		Host:      n.Host.String,
		Port:      n.Port.Int64,
		Username:  n.Username.String,
		Password:  n.Password.String,
		Var1:      n.Var1.String,
		Var2:      n.Var2.String,
		ApiKey:    n.ApiKey.String,
		ApiSecret: n.ApiSecret.String,
	}
}

func All() []*Notification {
	var n []*Notification
	q := db.Find(&n)
	if q.Error() != nil {
		return nil
	}
	return n
}

func Find(method string) (*Notification, error) {
	var n Notification
	q := db.Where("method = ?", method).Find(&n)
	if q.Error() != nil {
		return nil, q.Error()
	}
	return &n, nil
}

func (n *Notification) Create() error {
	var p Notification
	q := db.Where("method = ?", n.Method).Find(&p)
	if q.RecordNotFound() {
		log.Infof("Notifier '%s' was not found, adding into database...\n", n.Method)
		if err := db.Create(n).Error(); err != nil {
			return err
		}
		return nil
	}
	if p.FailureData.String == "" {
		p.FailureData = n.FailureData
	}
	if p.SuccessData.String == "" {
		p.SuccessData = n.SuccessData
	}
	if err := p.Update(); err != nil {
		return err
	}
	return nil
}

func (n *Notification) UpdateFields(notif *Notification) *Notification {
	if notif == nil {
		return n
	}
	n.Id = notif.Id
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
