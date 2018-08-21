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
	"net/http"
	"time"
	"upper.io/db.v3/lib/sqlbuilder"
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
	OnLoad(sqlbuilder.Database)
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
	Id        int       `db:"id,omitempty"`
	Service   int64     `db:"service"`
	Latency   float64   `db:"latency"`
	CreatedAt time.Time `db:"created_at"`
}

type Config struct {
	Connection string `yaml:"connection"`
	Host       string `yaml:"host"`
	Database   string `yaml:"database"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Port       string `yaml:"port"`
	Secret     string `yaml:"secret"`
}

type DbConfig struct {
	DbConn      string `yaml:"connection"`
	DbHost      string `yaml:"host"`
	DbUser      string `yaml:"user"`
	DbPass      string `yaml:"password"`
	DbData      string `yaml:"database"`
	DbPort      int    `yaml:"port"`
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
