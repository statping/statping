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
	"time"
)

type User struct {
	Id            int64     `gorm:"primary_key;column:id" json:"id"`
	Username      string    `gorm:"type:varchar(100);unique;column:username;" json:"username"`
	Password      string    `gorm:"column:password" json:"-"`
	Email         string    `gorm:"type:varchar(100);unique;column:email" json:"-"`
	ApiKey        string    `gorm:"column:api_key" json:"api_key"`
	ApiSecret     string    `gorm:"column:api_secret" json:"-"`
	Admin         bool      `gorm:"column:administrator" json:"admin"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserInterface `gorm:"-" json:"-"`
}

type UserInterface interface {
	// Database functions
	Create() (int64, error)
	Update() error
	Delete() error
}
