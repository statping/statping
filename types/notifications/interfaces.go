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

package notifications

import (
	"github.com/hunterlong/statping/types/failures"
	"github.com/hunterlong/statping/types/services"
	"github.com/hunterlong/statping/types/users"
)

// Notifier interface is required to create a new Notifier
type Notifier interface {
	OnSave() error          // OnSave is triggered when the notifier is saved
	Send(interface{}) error // OnSave is triggered when the notifier is saved
	Select() *Notification  // Select returns the *Notification for a notifier
}

// BasicEvents includes the most minimal events, failing and successful service triggers
type BasicEvents interface {
	OnSuccess(*services.Service)                    // OnSuccess is triggered when a service is successful
	OnFailure(*services.Service, *failures.Failure) // OnFailure is triggered when a service is failing
}

// Tester interface will include a function to Test users settings before saving
type Tester interface {
	OnTest() error
}

// ServiceEvents are events for Services
type ServiceEvents interface {
	OnNewService(*services.Service)
	OnUpdatedService(*services.Service)
	OnDeletedService(*services.Service)
}

// UserEvents are events for Users
type UserEvents interface {
	OnNewUser(*users.User)
	OnUpdatedUser(*users.User)
	OnDeletedUser(*users.User)
}

// CoreEvents are events for the main Core app
//type CoreEvents interface {
//	OnUpdatedCore(*core.Core)
//	OnStart(*core.Core)
//}

// NotifierEvents are events for other Notifiers
type NotifierEvents interface {
	OnNewNotifier(*Notification)
	OnUpdatedNotifier(*Notification)
}

// HTTPRouter interface will allow your notifier to accept http GET/POST requests
type HTTPRouter interface {
	OnGET() error
	OnPOST() error
}
