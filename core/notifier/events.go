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

package notifier

import "github.com/hunterlong/statup/types"

// OnSave will trigger a notifier when it has been saved - Notifier interface
func OnSave(method string) {
	for _, comm := range AllCommunications {
		if isType(comm, new(Notifier)) {
			notifier := comm.(Notifier).Select()
			if notifier.Method == method {
				comm.(Notifier).OnSave()
			}
		}
	}
}

// OnFailure will be triggered when a service is failing - BasicEvents interface
func OnFailure(s *types.Service, f *types.Failure) {
	for _, comm := range AllCommunications {
		if isType(comm, new(BasicEvents)) && isEnabled(comm) {
			comm.(BasicEvents).OnFailure(s, f)
		}
	}
}

// OnSuccess will be triggered when a service is successful - BasicEvents interface
func OnSuccess(s *types.Service) {
	for _, comm := range AllCommunications {
		if isType(comm, new(BasicEvents)) && isEnabled(comm) {
			comm.(BasicEvents).OnSuccess(s)
		}
	}
}

// OnNewService is triggered when a new service is created - ServiceEvents interface
func OnNewService(s *types.Service) {
	for _, comm := range AllCommunications {
		if isType(comm, new(ServiceEvents)) && isEnabled(comm) {
			comm.(ServiceEvents).OnNewService(s)
		}
	}
}

// OnUpdatedService is triggered when a service is updated - ServiceEvents interface
func OnUpdatedService(s *types.Service) {
	for _, comm := range AllCommunications {
		if isType(comm, new(ServiceEvents)) && isEnabled(comm) {
			comm.(ServiceEvents).OnUpdatedService(s)
		}
	}
}

// OnDeletedService is triggered when a service is deleted - ServiceEvents interface
func OnDeletedService(s *types.Service) {
	for _, comm := range AllCommunications {
		if isType(comm, new(ServiceEvents)) && isEnabled(comm) {
			comm.(ServiceEvents).OnDeletedService(s)
		}
	}
}

// OnNewUser is triggered when a new user is created - UserEvents interface
func OnNewUser(u *types.User) {
	for _, comm := range AllCommunications {
		if isType(comm, new(UserEvents)) && isEnabled(comm) {
			comm.(UserEvents).OnNewUser(u)
		}
	}
}

// OnUpdatedUser is triggered when a new user is updated - UserEvents interface
func OnUpdatedUser(u *types.User) {
	for _, comm := range AllCommunications {
		if isType(comm, new(UserEvents)) && isEnabled(comm) {
			comm.(UserEvents).OnUpdatedUser(u)
		}
	}
}

// OnDeletedUser is triggered when a new user is deleted - UserEvents interface
func OnDeletedUser(u *types.User) {
	for _, comm := range AllCommunications {
		if isType(comm, new(UserEvents)) && isEnabled(comm) {
			comm.(UserEvents).OnDeletedUser(u)
		}
	}
}

// OnUpdatedCore is triggered when the CoreApp settings are saved - CoreEvents interface
func OnUpdatedCore(c *types.Core) {
	for _, comm := range AllCommunications {
		if isType(comm, new(CoreEvents)) && isEnabled(comm) {
			comm.(CoreEvents).OnUpdatedCore(c)
		}
	}
}

// NotifierEvents interface
func OnNewNotifier(n *Notification) {
	for _, comm := range AllCommunications {
		if isType(comm, new(NotifierEvents)) && isEnabled(comm) {
			comm.(NotifierEvents).OnNewNotifier(n)
		}
	}
}

// NotifierEvents interface
func OnUpdatedNotifier(n *Notification) {
	for _, comm := range AllCommunications {
		if isType(comm, new(NotifierEvents)) && isEnabled(comm) {
			comm.(NotifierEvents).OnUpdatedNotifier(n)
		}
	}
}
