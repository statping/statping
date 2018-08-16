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

type ServiceInterface interface {
	AvgTime() float64
	Online24() float32
	ToService() *Service
	SmallText() string
	GraphData() string
	AvgUptime() string
	LimitedFailures() []*Failure
	TotalFailures() (uint64, error)
	TotalFailures24Hours() (uint64, error)
}

type FailureInterface interface {
	ToFailure() *Failure
	Ago() string
	ParseError() string
}
