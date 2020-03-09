// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/statping/statping
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
	"fmt"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/types/users"
	"github.com/statping/statping/utils"
)

// OnSave will trigger a notifier when it has been saved - Notifier interface
func OnSave(method string) {
	for _, comm := range allNotifiers {
		if utils.IsType(comm, new(Notifier)) {
			notifier := comm.(Notifier)
			if notifier.Select().Method == method {
				notifier.OnSave()
			}
		}
	}
}

// OnFailure will be triggered when a service is failing - BasicEvents interface
func OnFailure(s *services.Service, f *failures.Failure) {
	if !s.AllowNotifications.Bool {
		return
	}

	// check if User wants to receive every Status Change
	if s.UpdateNotify.Bool {
		// send only if User hasn't been already notified about the Downtime
		if !s.UserNotified {
			s.UserNotified = true
			goto sendMessages
		} else {
			return
		}
	}

sendMessages:
	for _, comm := range allNotifiers {
		if utils.IsType(comm, new(BasicEvents)) && isEnabled(comm) && (s.Online || inLimits(comm)) {
			notifier := comm.(Notifier).Select()
			log.
				WithField("trigger", "OnFailure").
				WithFields(utils.ToFields(notifier, s)).Debugln(fmt.Sprintf("Sending [OnFailure] '%v' notification for service %v", notifier.Method, s.Name))
			comm.(BasicEvents).OnFailure(s, f)
			comm.Select().Hits.OnFailure++
		}
	}
}

// OnSuccess will be triggered when a service is successful - BasicEvents interface
func OnSuccess(s *services.Service) {
	if !s.AllowNotifications.Bool {
		return
	}

	// check if User wants to receive every Status Change
	if s.UpdateNotify.Bool && s.UserNotified {
		s.UserNotified = false
	}

	for _, comm := range allNotifiers {
		if utils.IsType(comm, new(BasicEvents)) && isEnabled(comm) && (!s.Online || inLimits(comm)) {
			notifier := comm.(Notifier).Select()
			log.
				WithField("trigger", "OnSuccess").
				WithFields(utils.ToFields(notifier, s)).Debugln(fmt.Sprintf("Sending [OnSuccess] '%v' notification for service %v", notifier.Method, s.Name))
			comm.(BasicEvents).OnSuccess(s)
			comm.Select().Hits.OnSuccess++
		}
	}
}

// OnNewService is triggered when a new service is created - ServiceEvents interface
func OnNewService(s *services.Service) {
	for _, comm := range allNotifiers {
		if utils.IsType(comm, new(ServiceEvents)) && isEnabled(comm) && inLimits(comm) {
			log.
				WithField("trigger", "OnNewService").
				Debugln(fmt.Sprintf("Sending new service notification for service %v", s.Name))
			comm.(ServiceEvents).OnNewService(s)
			comm.Select().Hits.OnNewService++
		}
	}
}

// OnUpdatedService is triggered when a service is updated - ServiceEvents interface
func OnUpdatedService(s *services.Service) {
	if !s.AllowNotifications.Bool {
		return
	}
	for _, comm := range allNotifiers {
		if utils.IsType(comm, new(ServiceEvents)) && isEnabled(comm) && inLimits(comm) {
			log.Debugln(fmt.Sprintf("Sending updated service notification for service %v", s.Name))
			comm.(ServiceEvents).OnUpdatedService(s)
			comm.Select().Hits.OnUpdatedService++
		}
	}
}

// OnDeletedService is triggered when a service is deleted - ServiceEvents interface
func OnDeletedService(s *services.Service) {
	if !s.AllowNotifications.Bool {
		return
	}
	for _, comm := range allNotifiers {
		if utils.IsType(comm, new(ServiceEvents)) && isEnabled(comm) && inLimits(comm) {
			log.Debugln(fmt.Sprintf("Sending deleted service notification for service %v", s.Name))
			comm.(ServiceEvents).OnDeletedService(s)
			comm.Select().Hits.OnDeletedService++
		}
	}
}

// OnNewUser is triggered when a new user is created - UserEvents interface
func OnNewUser(u *users.User) {
	for _, comm := range allNotifiers {
		if utils.IsType(comm, new(UserEvents)) && isEnabled(comm) && inLimits(comm) {
			log.Debugln(fmt.Sprintf("Sending new user notification for user %v", u.Username))
			comm.(UserEvents).OnNewUser(u)
			comm.Select().Hits.OnNewUser++
		}
	}
}

// OnUpdatedUser is triggered when a new user is updated - UserEvents interface
func OnUpdatedUser(u *users.User) {
	for _, comm := range allNotifiers {
		if utils.IsType(comm, new(UserEvents)) && isEnabled(comm) && inLimits(comm) {
			log.Debugln(fmt.Sprintf("Sending updated user notification for user %v", u.Username))
			comm.(UserEvents).OnUpdatedUser(u)
			comm.Select().Hits.OnUpdatedUser++
		}
	}
}

// OnDeletedUser is triggered when a new user is deleted - UserEvents interface
func OnDeletedUser(u *users.User) {
	for _, comm := range allNotifiers {
		if utils.IsType(comm, new(UserEvents)) && isEnabled(comm) && inLimits(comm) {
			log.Debugln(fmt.Sprintf("Sending deleted user notification for user %v", u.Username))
			comm.(UserEvents).OnDeletedUser(u)
			comm.Select().Hits.OnDeletedUser++
		}
	}
}

//// OnUpdatedCore is triggered when the CoreApp settings are saved - CoreEvents interface
//func OnUpdatedCore(c *core.Core) {
//	for _, comm := range allNotifiers {
//		if utils.IsType(comm, new(CoreEvents)) && isEnabled(comm) && inLimits(comm) {
//			log.Debugln(fmt.Sprintf("Sending updated core notification"))
//			comm.(CoreEvents).OnUpdatedCore(c)
//		}
//	}
//}
//
//// OnStart is triggered when the Statping service has started
//func OnStart(c *core.Core) {
//	for _, comm := range allNotifiers {
//		if utils.IsType(comm, new(CoreEvents)) && isEnabled(comm) && inLimits(comm) {
//			comm.(CoreEvents).OnUpdatedCore(c)
//		}
//	}
//}

// OnNewNotifier is triggered when a new notifier is loaded
func OnNewNotifier(n *Notification) {
	for _, comm := range allNotifiers {
		if utils.IsType(comm, new(NotifierEvents)) && isEnabled(comm) && inLimits(comm) {
			comm.(NotifierEvents).OnNewNotifier(n)
			comm.Select().Hits.OnNewNotifier++
		}
	}
}

// OnUpdatedNotifier is triggered when a notifier has been updated
func OnUpdatedNotifier(n *Notification) {
	for _, comm := range allNotifiers {
		if utils.IsType(comm, new(NotifierEvents)) && isEnabled(comm) && inLimits(comm) {
			log.Infoln(fmt.Sprintf("Sending updated notifier for %v", n.Id))
			comm.(NotifierEvents).OnUpdatedNotifier(n)
			comm.Select().Hits.OnUpdatedNotifier++
		}
	}
}
