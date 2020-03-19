package notifiers

import (
	"fmt"
	"github.com/statping/statping/types/services"
)

func InitNotifiers() {
	Add(
		slacker,
		Command,
		Discorder,
		email,
		LineNotify,
		Telegram,
		Twilio,
		Webhook,
		Mobile,
	)
}

func Add(notifs ...services.ServiceNotifier) {
	for _, n := range notifs {
		services.AddNotifier(n)
		if err := n.Select().Create(); err != nil {
			fmt.Println(err)
		}
	}
}
