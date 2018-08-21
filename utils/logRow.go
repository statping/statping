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

package utils

import (
	"fmt"
	"time"
)

type LogRow struct {
	Date time.Time
	Line interface{}
}

func NewLogRow(line interface{}) (logRow *LogRow) {
	logRow = new(LogRow)
	logRow.Date = time.Now()
	logRow.Line = line
	return
}

func (o *LogRow) LineAsString() string {
	switch v := o.Line.(type) {
	case string:
		return v
	case error:
		return v.Error()
	case []byte:
		return string(v)
	}
	return ""
}

func (o *LogRow) FormatForHtml() string {
	return fmt.Sprintf("%s: %s", o.Date.Format("2006-01-02 15:04:05"), o.LineAsString())
}
