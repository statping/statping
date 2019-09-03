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

const (
	TIME_NANO  = "2006-01-02T15:04:05Z"
	TIME       = "2006-01-02 15:04:05"
	CHART_TIME = "2006-01-02T15:04:05.999999-07:00"
	TIME_DAY   = "2006-01-02"
)

var (
	NOW = func() time.Time { return time.Now() }()
	//HOUR_1_AGO  = time.Now().Add(-1 * time.Hour)
	//HOUR_24_AGO = time.Now().Add(-24 * time.Hour)
	//HOUR_72_AGO = time.Now().Add(-72 * time.Hour)
	//DAY_7_AGO   = NOW.AddDate(0, 0, -7)
	//MONTH_1_AGO = NOW.AddDate(0, -1, 0)
	//YEAR_1_AGO  = NOW.AddDate(-1, 0, 0)
)
