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

package core

import (
	"github.com/hunterlong/statping/types"
)

type Failure struct{}

const (
	limitedFailures = 32
	limitedHits     = 32
)

// Delete will remove a Failure record from the database
func (f *Failure) Delete() error {
	db := Database(&types.Failure{}).Delete(f)
	return db.Error()
}
