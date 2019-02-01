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

// Message is for creating Announcements, Alerts and other messages for the end users
type Message struct {
	Id                int64     `gorm:"primary_key;column:id" json:"id"`
	Title             string    `gorm:"column:title" json:"title"`
	Description       string    `gorm:"column:description" json:"description"`
	StartOn           time.Time `gorm:"column:start_on" json:"start_on"`
	EndOn             time.Time `gorm:"column:end_on" json:"end_on"`
	ServiceId         int64     `gorm:"index;column:service" json:"service"`
	NotifyUsers       NullBool  `gorm:"column:notify_users" json:"notify_users"`
	NotifyMethod      string    `gorm:"column:notify_method" json:"notify_method"`
	NotifyBefore      NullInt64 `gorm:"column:notify_before" json:"notify_before"`
	NotifyBeforeScale string    `gorm:"column:notify_before_scale" json:"notify_before_scale"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at" json:"updated_at"`
}

// BeforeCreate for Message will set CreatedAt to UTC
func (u *Message) BeforeCreate() (err error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now().UTC()
		u.UpdatedAt = time.Now().UTC()
	}
	return
}
