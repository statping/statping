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

// Failure is a failed attempt to check a service. Any a service does not meet the expected requirements,
// a new Failure will be inserted into database.
type Failure struct {
	Id        int64     `gorm:"primary_key;column:id" json:"id"`
	Issue     string    `gorm:"column:issue" json:"issue"`
	Method    string    `gorm:"column:method" json:"method,omitempty"`
	MethodId  int64     `gorm:"column:method_id" json:"method_id,omitempty"`
	ErrorCode int       `gorm:"column:error_code" json:"error_code"`
	Service   int64     `gorm:"index;column:service" json:"-"`
	Checkin   int64     `gorm:"index;column:checkin" json:"-"`
	PingTime  float64   `gorm:"column:ping_time"  json:"ping"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

type FailureInterface interface {
	Select() *Failure
	Ago() string        // Ago returns a human readable timestamp
	ParseError() string // ParseError returns a human readable error for a service failure
}

// BeforeCreate for Failure will set CreatedAt to UTC
func (f *Failure) BeforeCreate() (err error) {
	if f.CreatedAt.IsZero() {
		f.CreatedAt = time.Now().UTC()
	}
	return
}

type FailSort []FailureInterface

func (s FailSort) Len() int {
	return len(s)
}
func (s FailSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s FailSort) Less(i, j int) bool {
	return s[i].Select().Id < s[j].Select().Id
}
