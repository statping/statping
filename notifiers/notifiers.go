package notifiers

import (
	"fmt"
	"github.com/hunterlong/statping/types/notifications"
	"github.com/hunterlong/statping/utils"
	"strings"
)

var (
	allowedVars = []string{"host", "username", "password", "port", "api_key", "api_secret", "var1", "var2"}
)

func checkNotifierForm(n notifications.Notifier) error {
	notifier := n.Select()
	for _, f := range notifier.Form {
		contains := contains(f.DbField, allowedVars)
		if !contains {
			return fmt.Errorf("the DbField '%v' is not allowed, allowed vars: %v", f.DbField, allowedVars)
		}
	}
	return nil
}

func contains(s string, arr []string) bool {
	for _, v := range arr {
		if strings.ToLower(s) == v {
			return true
		}
	}
	return false
}

// AddNotifier accept a Notifier interface to be added into the array
func AddNotifiers(notifiers ...notifications.Notifier) error {
	for _, n := range notifiers {
		if err := checkNotifierForm(n); err != nil {
			return err
		}
		notifications.AllCommunications = append(notifications.AllCommunications, n)
		if _, err := notifications.Init(n); err != nil {
			return err
		}
	}
	startAllNotifiers()
	return nil
}

// startAllNotifiers will start the go routine for each loaded notifier
func startAllNotifiers() {
	for _, comm := range notifications.AllCommunications {
		if utils.IsType(comm, new(notifications.Notifier)) {
			notify := comm.(notifications.Notifier)
			if notify.Select().Enabled.Bool {
				notify.Select().Close()
				notify.Select().Start()
				go notifications.Queue(notify)
			}
		}
	}
}

func AttachNotifiers() error {
	return AddNotifiers(
		Command,
		Discorder,
		Emailer,
		LineNotify,
		Mobile,
		Slacker,
		Telegram,
		Twilio,
		Webhook,
	)
}
