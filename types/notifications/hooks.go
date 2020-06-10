package notifications

import "github.com/statping/statping/utils"

// AfterFind for Notification will set the timezone
func (n *Notification) AfterFind() (err error) {
	n.CreatedAt = utils.Now()
	n.UpdatedAt = utils.Now()
	return
}
