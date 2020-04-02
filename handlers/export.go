package handlers

import (
	"encoding/json"
	"github.com/statping/statping/types/checkins"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/groups"
	"github.com/statping/statping/types/messages"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/types/users"
)

// ExportChartsJs renders the charts for the index page

type ExportData struct {
	Core      *core.Core          `json:"core"`
	Services  []services.Service  `json:"services"`
	Messages  []*messages.Message `json:"messages"`
	Checkins  []*checkins.Checkin `json:"checkins"`
	Users     []*users.User       `json:"users"`
	Groups    []*groups.Group     `json:"groups"`
	Notifiers []core.AllNotifiers `json:"notifiers"`
}

// ExportSettings will export a JSON file containing all of the settings below:
// - Core
// - Notifiers
// - Checkins
// - Users
// - Services
// - Groups
// - Messages
func ExportSettings() ([]byte, error) {
	c, err := core.Select()
	if err != nil {
		return nil, err
	}
	data := ExportData{
		Core: c,
		//Notifiers: notifications.All(),
		Checkins: checkins.All(),
		Users:    users.All(),
		Services: services.AllInOrder(),
		Groups:   groups.All(),
		Messages: messages.All(),
	}
	export, err := json.Marshal(data)
	return export, err
}
