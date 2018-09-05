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

package types

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type PluginInfo struct {
	Info Info
	PluginActions
}

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

type PluginActions interface {
	GetInfo() Info
	GetForm() string
	OnLoad(db gorm.DB)
	SetInfo(map[string]interface{}) Info
	Routes() []Routing
	OnSave(map[string]interface{})
	OnFailure(map[string]interface{})
	OnSuccess(map[string]interface{})
	OnSettingsSaved(map[string]interface{})
	OnNewUser(map[string]interface{})
	OnNewService(map[string]interface{})
	OnUpdatedService(map[string]interface{})
	OnDeletedService(map[string]interface{})
	OnInstall(map[string]interface{})
	OnUninstall(map[string]interface{})
	OnBeforeRequest(map[string]interface{})
	OnAfterRequest(map[string]interface{})
	OnShutdown()
}

type AllNotifiers interface{}

type Hit struct {
	Id        int64     `gorm:"primary_key;column:id"`
	Service   int64     `gorm:"index;column:service"`
	Latency   float64   `gorm:"column:latency"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

type DbConfig struct {
	DbConn      string `yaml:"connection"`
	DbHost      string `yaml:"host"`
	DbUser      string `yaml:"user"`
	DbPass      string `yaml:"password"`
	DbData      string `yaml:"database"`
	DbPort      int    `yaml:"port"`
	ApiKey      string `yaml:"api_key"`
	ApiSecret   string `yaml:"api_secret"`
	Project     string `yaml:"-"`
	Description string `yaml:"-"`
	Domain      string `yaml:"-"`
	Username    string `yaml:"-"`
	Password    string `yaml:"-"`
	Email       string `yaml:"-"`
	Error       error  `yaml:"-"`
	Location    string `yaml:"location"`
}

type PluginRepos struct {
	Plugins []PluginJSON
}

type PluginJSON struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Repo        string `json:"repo"`
	Author      string `json:"author"`
	Namespace   string `json:"namespace"`
}
