package services

import (
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
)

var (
	allNotifiers = make(map[string]ServiceNotifier)
)

func AllNotifiers() map[string]ServiceNotifier {
	return allNotifiers
}

func ReturnNotifier(method string) ServiceNotifier {
	return allNotifiers[method]
}

func FindNotifier(method string) *notifications.Notification {
	n := allNotifiers[method]
	if n != nil {
		return n.Select()
	}
	return nil
}

type ServiceNotifier interface {
	OnSuccess(Service) (string, error)                   // OnSuccess is triggered when a service is successful
	OnFailure(Service, failures.Failure) (string, error) // OnFailure is triggered when a service is failing
	OnTest() (string, error)                             // OnTest is triggered for testing
	OnSave() (string, error)                             // OnSave is triggered for testing
	Select() *notifications.Notification                 // OnTest is triggered for testing
}
