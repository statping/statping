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

// Notifier interface
func OnSave(method string) {
	for _, comm := range AllCommunications {
		if IsType(comm, "Notifier") {
			notifier := comm.(Notifier).Select()
			if notifier.Method == method {
				comm.(Notifier).OnSave()
			}
		}
	}
}

// BasicEvents interface
func OnFailure(s *types.Service, f *types.Failure) {
	for _, comm := range AllCommunications {
		if IsType(comm, "BasicEvents") {
			comm.(BasicEvents).OnFailure(s, f)
		}
	}
}

// BasicEvents interface
func OnSuccess(s *types.Service) {
	for _, comm := range AllCommunications {
		if IsType(comm, "BasicEvents") {
			comm.(BasicEvents).OnSuccess(s)
		}
	}
}

// ServiceEvents interface
func OnNewService(s *types.Service) {
	for _, comm := range AllCommunications {
		if IsType(comm, "ServiceEvents") {
			comm.(ServiceEvents).OnNewService(s)
		}
	}
}

// ServiceEvents interface
func OnUpdatedService(s *types.Service) {
	for _, comm := range AllCommunications {
		if IsType(comm, "ServiceEvents") {
			comm.(ServiceEvents).OnUpdatedService(s)
		}
	}
}

// ServiceEvents interface
func OnDeletedService(s *types.Service) {
	for _, comm := range AllCommunications {
		if IsType(comm, "ServiceEvents") {
			comm.(ServiceEvents).OnDeletedService(s)
		}
	}
}

// UserEvents interface
func OnNewUser(u *types.User) {
	for _, comm := range AllCommunications {
		if IsType(comm, "UserEvents") {
			comm.(UserEvents).OnNewUser(u)
		}
	}
}

// UserEvents interface
func OnUpdatedUser(u *types.User) {
	for _, comm := range AllCommunications {
		if IsType(comm, "UserEvents") {
			comm.(UserEvents).OnUpdatedUser(u)
		}
	}
}

// UserEvents interface
func OnDeletedUser(u *types.User) {
	for _, comm := range AllCommunications {
		if IsType(comm, "UserEvents") {
			comm.(UserEvents).OnDeletedUser(u)
		}
	}
}

// CoreEvents interface
func OnUpdatedCore(c *types.Core) {
	for _, comm := range AllCommunications {
		if IsType(comm, "CoreEvents") {
			comm.(CoreEvents).OnUpdatedCore(c)
		}
	}
}

// NotifierEvents interface
func OnNewNotifier(n *Notification) {
	for _, comm := range AllCommunications {
		if IsType(comm, "NotifierEvents") {
			comm.(NotifierEvents).OnNewNotifier(n)
		}
	}
}

// NotifierEvents interface
func OnUpdatedNotifier(n *Notification) {
	for _, comm := range AllCommunications {
		if IsType(comm, "NotifierEvents") {
			comm.(NotifierEvents).OnUpdatedNotifier(n)
		}
	}
}
