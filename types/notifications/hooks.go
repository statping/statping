package notifications

import (
	"github.com/statping/statping/types/metrics"
	"gorm.io/gorm"
)

func (n *Notification) AfterFind(*gorm.DB) error {
	metrics.Query("notifier", "find")
	return nil
}

func (n *Notification) AfterCreate(*gorm.DB) error {
	metrics.Query("notifier", "create")
	return nil
}

func (n *Notification) AfterUpdate(*gorm.DB) error {
	metrics.Query("notifier", "update")
	return nil
}

func (n *Notification) AfterDelete(*gorm.DB) error {
	metrics.Query("notifier", "delete")
	return nil
}
