package notifications

import (
	"fmt"
	"time"

	"github.com/gregdel/pushover"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
)

var (
	pushoverQueue []*types.PushoverNotification
	PushoverComm  *types.Communication
)

func PushoverRoutine() {
	for _, msg := range pushoverQueue {
		app := pushover.New(PushoverComm.Host)
		recipient := pushover.NewRecipient(msg.To)
		message := pushover.NewMessage(msg.Message)
		_, err := app.SendMessage(message, recipient)
		if err != nil {
			utils.Log(3, fmt.Sprintf("Issue sending Pushover notification: %v", err))
		}
		pushoverQueue = removePushover(pushoverQueue, msg)
	}
	time.Sleep(10 * time.Second)
	if PushoverComm.Enabled {
		PushoverRoutine()
	}
}

func removePushover(pushovers []*types.PushoverNotification, em *types.PushoverNotification) []*types.PushoverNotification {
	var newArr []*types.PushoverNotification
	for _, e := range pushovers {
		if e != em {
			newArr = append(newArr, e)
		}
	}
	return newArr
}

func SendPushover(notification *types.PushoverNotification) error {
	pushoverQueue = append(pushoverQueue, notification)
	return nil
}
