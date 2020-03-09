package notifiers

import (
	"fmt"
	"github.com/google/martian/log"
	"github.com/hunterlong/statping/types/notifications"
	"github.com/hunterlong/statping/utils"
	"strings"
)

var (
	allowedVars = []string{"host", "username", "password", "port", "api_key", "api_secret", "var1", "var2"}
)

func checkNotifierForm(n *notifications.Notification) error {
	for _, f := range n.Form {
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
	log.Infof("Initiating %d Notifiers\n", len(notifiers))

	for _, n := range notifiers {
		notif := n.Select()
		log.Infof("Initiating %s Notifier\n", notif.Method)

		if err := checkNotifierForm(notif); err != nil {
			log.Errorf(err.Error())
			return err
		}

		log.Infof("Creating %s Notifier\n", notif.Method)
		if err := notif.Create(); err != nil {
			return err
		}

		notifications.Append(notif)

		if notif.Enabled.Bool {
			notif.Close()
			notif.Start()
			go notifications.Queue(notif)
		}

	}
	return nil
}

// startAllNotifiers will start the go routine for each loaded notifier
func startAllNotifiers() {
	for _, notify := range notifications.All() {
		n := notify.Select()
		log.Infof("Initiating %s Notifier\n", n.Method)
		if utils.IsType(notify, new(notifications.Notifier)) {
			if n.Enabled.Bool {
				n.Close()
				n.Start()
				go notifications.Queue(notify)
			}
		}
	}
}

func Migrate() error {
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
