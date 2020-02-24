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
	"fmt"
	"time"
)

const (
	TIME_NANO  = "2006-01-02T15:04:05Z"
	TIME       = "2006-01-02 15:04:05"
	CHART_TIME = "2006-01-02T15:04:05.999999-07:00"
	TIME_DAY   = "2006-01-02"
)

var (
	Second = time.Second
	Minute = time.Minute
	Hour   = time.Hour
	Day    = Hour * 24
	Week   = Day * 7
	Month  = Week * 4
	Year   = Day * 365
)

func FixedTime(t time.Time, d time.Duration) string {
	switch d {
	case Month:
		month := fmt.Sprintf("%v", int(t.Month()))
		if int(t.Month()) < 10 {
			month = fmt.Sprintf("0%v", int(t.Month()))
		}
		return fmt.Sprintf("%v-%v-01T00:00:00Z", t.Year(), month)
	case Year:
		return fmt.Sprintf("%v-01-01T00:00:00Z", t.Year())
	default:
		return t.Format(durationStr(d))
	}
}

func durationStr(d time.Duration) string {
	switch d {
	case Second:
		return "2006-01-02T15:04:05Z"
	case Minute:
		return "2006-01-02T15:04:00Z"
	case Hour:
		return "2006-01-02T15:00:00Z"
	case Day:
		return "2006-01-02T00:00:00Z"
	default:
		return "2006-01-02T00:00:00Z"
	}
}
