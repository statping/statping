package notifiers

import (
	"fmt"
	"github.com/google/martian/log"
	"github.com/statping/statping/notifiers/senders"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"strings"
)

var (
	allowedVars = []string{"host", "username", "password", "port", "api_key", "api_secret", "var1", "var2"}
)

func SendEvent(data ...interface{}) {
	d1 := data[0]
	service, ok := d1.(*services.Service)
	if !ok {
		return
	}
	d2 := data[1]
	if d2 == nil {
		OnSuccess(service)
	}
	fail, ok := d2.(*failures.Failure)
	if !ok {
		return
	}
	OnFailure(service, fail)
}

func checkNotifierForm(n *Notification) error {
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
func AddNotifiers(notifiers ...Notifier) error {
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

		if notif.Enabled.Bool {
			notif.Close()
			notif.Start()
			go Queue(notif)
		}

	}
	return nil
}

// startAllNotifiers will start the go routine for each loaded notifier
func startAllNotifiers() {
	for _, notify := range All() {
		n := notify.Select()
		log.Infof("Initiating %s Notifier\n", n.Method)
		if utils.IsType(notify, new(Notifier)) {
			if n.Enabled.Bool {
				n.Close()
				n.Start()
				go Queue(notify)
			}
		}
	}
}

func Migrate() error {
	return AddNotifiers(
		senders.Command,
		senders.Discorder,
		senders.Emailer,
		senders.LineNotify,
		senders.Mobile,
		senders.Slacker,
		senders.Telegram,
		senders.Twilio,
		senders.Webhook,
	)
}
