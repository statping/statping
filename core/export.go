// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"encoding/json"
	"github.com/hunterlong/statping/types"
)

// ExportChartsJs renders the charts for the index page

type ExportData struct {
	Core      *types.Core              `json:"core"`
	Services  []types.ServiceInterface `json:"services"`
	Messages  []*Message               `json:"messages"`
	Checkins  []*Checkin               `json:"checkins"`
	Users     []*User                  `json:"users"`
	Groups    []*Group                 `json:"groups"`
	Notifiers []types.AllNotifiers     `json:"notifiers"`
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
	users, err := SelectAllUsers()
	messages, err := SelectMessages()
	if err != nil {
		return nil, err
	}
	data := ExportData{
		Core:      CoreApp.Core,
		Notifiers: CoreApp.Notifications,
		Checkins:  AllCheckins(),
		Users:     users,
		Services:  CoreApp.Services,
		Groups:    SelectGroups(true, true),
		Messages:  messages,
	}
	export, err := json.Marshal(data)
	return export, err
}
