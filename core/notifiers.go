package core

import (
	"github.com/hunterlong/statping/core/notifier"
	"github.com/hunterlong/statping/notifiers"
)

// AttachNotifiers will attach all the notifier's into the system
func AttachNotifiers() error {
	return notifier.AddNotifiers(
		notifiers.Command,
		notifiers.Discorder,
		notifiers.Emailer,
		notifiers.LineNotify,
		notifiers.Mobile,
		notifiers.Slacker,
		notifiers.Telegram,
		notifiers.Twilio,
		notifiers.Webhook,
	)
}
