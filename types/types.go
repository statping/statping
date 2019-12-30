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

package types

import (
	"time"
)

// Hit struct is a 'successful' ping or web response entry for a service.
type Hit struct {
	Id        int64     `gorm:"primary_key;column:id" json:"id"`
	Service   int64     `gorm:"column:service" json:"-"`
	Latency   float64   `gorm:"column:latency" json:"latency"`
	PingTime  float64   `gorm:"column:ping_time" json:"ping_time"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

// BeforeCreate for Hit will set CreatedAt to UTC
func (h *Hit) BeforeCreate() (err error) {
	if h.CreatedAt.IsZero() {
		h.CreatedAt = time.Now().UTC()
	}
	return
}

// DbConfig struct is used for the database connection and creates the 'config.yml' file
type DbConfig struct {
	DbConn      string `yaml:"connection" json:"connection"`
	DbHost      string `yaml:"host" json:"-"`
	DbUser      string `yaml:"user" json:"-"`
	DbPass      string `yaml:"password" json:"-"`
	DbData      string `yaml:"database" json:"-"`
	DbPort      int64  `yaml:"port" json:"-"`
	ApiKey      string `yaml:"api_key" json:"-"`
	ApiSecret   string `yaml:"api_secret" json:"-"`
	Project     string `yaml:"-" json:"-"`
	Description string `yaml:"-" json:"-"`
	Domain      string `yaml:"-" json:"-"`
	Username    string `yaml:"-" json:"-"`
	Password    string `yaml:"-" json:"-"`
	Email       string `yaml:"-" json:"-"`
	Error       error  `yaml:"-" json:"-"`
	Location    string `yaml:"location" json:"-"`
	SqlFile     string `yaml:"sqlfile,omitempty" json:"-"`
	LocalIP     string `yaml:"-" json:"-"`
}
