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

package notifiers

import "github.com/hunterlong/statup/types"

// Notifier interface is required to create a new Notifier
type Notifier interface {
	Init() error
	Install() error
	Run() error
	OnSave() error
	Test() error
	Select() *Notification
}

// BasicEvents includes the basic events, failing and successful service triggers
type BasicEvents interface {
	OnSuccess(*types.Service)
	OnFailure(*types.Service, *types.Failure)
}

// ServiceEvents are events for Services
type ServiceEvents interface {
	OnNewService(*types.Service)
	OnUpdatedService(*types.Service)
	OnDeletedService(*types.Service)
}

// UserEvents are events for Users
type UserEvents interface {
	OnNewUser(*types.User)
	OnUpdatedUser(*types.User)
	OnDeletedUser(*types.User)
}

// CoreEvents are events for the main Core app
type CoreEvents interface {
	OnUpdatedCore(*types.Core)
}

// NotifierEvents are events for other Notifiers
type NotifierEvents interface {
	OnNewNotifier(*Notification)
	OnUpdatedNotifier(*Notification)
}
