// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/statping/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
