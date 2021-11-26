package services

import (
	"github.com/statping-ng/statping-ng/types/failures"
	"github.com/statping-ng/statping-ng/types/notifications"
	"github.com/statping-ng/statping-ng/utils"
)

func AddNotifier(n ServiceNotifier) {
	notif := n.Select()
	allNotifiers[notif.Method] = n
}

func UpdateNotifiers() {
	for _, n := range notifications.All() {
		notifier := allNotifiers[n.Method]
		notifier.Select().UpdateFields(n)
	}
}

func sendSuccess(s *Service) {
	if !s.AllowNotifications.Bool {
		return
	}

	s.notifyAfterCount = 0

	if s.prevOnline == s.Online {
		return
	}
	s.prevOnline = true

	for _, n := range allNotifiers {
		notif := n.Select()
		if notif.CanSend() {
			log.Infof("Sending notification to: %s!", notif.Method)
			out, err := n.OnSuccess(*s)
			if err != nil {
				notif.Logger().Errorln(err)
				logMessage(notif.Method, "", err, false, s.Id)
				return
			}
			logMessage(notif.Method, out, nil, true, s.Id)
			notif.LastSentCount++
			notif.LastSent = utils.Now()
		}
	}
}

func sendFailure(s *Service, f *failures.Failure) {
	if !s.AllowNotifications.Bool {
		return
	}

	if s.prevOnline == s.Online && !s.UpdateNotify.Bool {
		return
	}

	if s.NotifyAfter != 0 {
		if s.NotifyAfter > s.notifyAfterCount {
			s.notifyAfterCount++
			return
		}
	}

	s.prevOnline = false

	for _, n := range allNotifiers {
		notif := n.Select()
		if notif.CanSend() {
			log.Infof("Sending Failure notification to: %s!", notif.Method)
			out, err := n.OnFailure(*s, *f)
			if err != nil {
				notif.Logger().WithField("failure", f.Issue).Errorln(err)
				logMessage(notif.Method, "", err, false, s.Id)
			}
			logMessage(notif.Method, out, nil, false, s.Id)

			notif.LastSentCount++
			notif.LastSent = utils.Now()
		}
	}
}

func logMessage(method string, msg string, error error, onSuccesss bool, serviceId int64) {
	notif := FindNotifier(method)
	l := &notifications.NotificationLog{
		Message:   msg,
		Error:     error,
		Success:   onSuccesss,
		Service:   serviceId,
		CreatedAt: utils.Now(),
	}
	notif.Logs = append(notif.Logs, l)
	if len(notif.Logs) > 32 {
		notif.Logs = notif.Logs[1:]
	}
}
