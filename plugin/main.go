// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package plugin

import (
	"github.com/hunterlong/statup/core/notifier"
	"github.com/jinzhu/gorm"
	"net/http"
)

//
//     STATUP PLUGIN INTERFACE
//
//            v0.1
//
//       https://statup.io
//
//
// An expandable plugin framework that will still
// work even if there's an update or addition.
//

type PluginObject struct{}

var (
	AllPlugins []*PluginObject
)

func Add(p Pluginer) *PluginObject {
	return &PluginObject{}
}

func (p *PluginObject) AddRoute(s string, i string, f http.HandlerFunc) {

}

type Pluginer interface {
	Select() *PluginObject
}

type Databaser interface {
	StatupDatabase(*gorm.DB)
}

type Router interface {
	AddRoute(string, string, http.HandlerFunc) error
}

type Notifier interface {
	notifier.Notifier
	notifier.BasicEvents
}

type AdvancedNotifier interface {
	notifier.Notifier
	notifier.BasicEvents
	notifier.UserEvents
	notifier.CoreEvents
	notifier.NotifierEvents
}

var (
	DB *gorm.DB
)

type Routing struct {
	URL     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

type Info struct {
	Name        string
	Description string
	Form        string
}

type Database *gorm.DB

type Plugin struct {
	Name        string
	Description string
}

type PluginDatabase interface {
	Database(gorm.DB)
	Update() error
}

type PluginInfo struct {
	i *Info
}

func SetDatabase(database *gorm.DB) {
	DB = database
}

func (p *PluginInfo) Form() string {
	return "okkokokkok"
}
