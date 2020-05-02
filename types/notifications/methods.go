package notifications

import (
	"fmt"
	"github.com/statping/statping/utils"
	"strings"
	"time"
)

func (n *Notification) Name() string {
	newName := strings.ToLower(n.Method)
	newName = strings.ReplaceAll(newName, " ", "_")
	return newName
}

// AfterFind for Notification will set the timezone
func (n *Notification) AfterFind() (err error) {
	n.CreatedAt = utils.Now()
	n.UpdatedAt = utils.Now()
	return
}

// AddQueue will add any type of interface (json, string, struct, etc) into the Notifiers queue
//func (n *Notification) AddQueue(uid string, msg interface{}) {
//	data := &QueueData{uid, msg}
//	n.Queue = append(n.Queue, data)
//	log.WithFields(utils.ToFields(data, n)).Debug(fmt.Sprintf("Notifier '%v' added new item (%v) to the queue. (%v queued)", n.Method, uid, len(n.Queue)))
//}

// CanTest returns true if the notifier implements the OnTest interface
//func (n *Notification) CanTest() bool {
//	return n.testable
//}

// LastSent returns a time.Duration of the last sent notification for the notifier
func (n *Notification) LastSent() time.Duration {
	since := time.Since(n.lastSent)
	return since
}

func (n *Notification) CanSend() bool {
	if !n.Enabled.Bool {
		return false
	}

	//fmt.Println("Last sent: ", n.lastSent.String())
	//fmt.Println("Last count: ", n.lastSentCount)
	//fmt.Println("Limits: ", n.Limits)
	//fmt.Println("Last sent before now: ", n.lastSent.Add(60*time.Second).Before(utils.Now()))

	// the last sent notification was past 1 minute (limit per minute)
	if n.lastSent.Add(60 * time.Minute).Before(utils.Now()) {
		if n.lastSentCount != 0 {
			n.lastSentCount--
		}
	}

	// dont send if already beyond the notifier's limit
	if n.lastSentCount >= n.Limits {
		return false
	}

	// action to do since notifier is able to send
	n.lastSentCount++
	n.lastSent = utils.Now()
	return true
}

// GetValue returns the database value of a accept DbField value.
func (n *Notification) GetValue(dbField string) string {
	switch strings.ToLower(dbField) {
	case "host":
		return n.Host
	case "port":
		return fmt.Sprintf("%d", n.Port)
	case "username":
		return n.Username
	case "password":
		return n.Password
	case "var1":
		return n.Var1
	case "var2":
		return n.Var2
	case "api_key":
		return n.ApiKey
	case "api_secret":
		return n.ApiSecret
	case "limits":
		return utils.ToString(int(n.Limits))
	default:
		return ""
	}
}

// start will start the go routine for the notifier queue
func (n *Notification) Start() {
	n.Running = make(chan bool)
}

// close will stop the go routine for queue
func (n *Notification) Close() {
	if n.IsRunning() {
		close(n.Running)
	}
}

// IsRunning will return true if the notifier is currently running a queue
func (n *Notification) IsRunning() bool {
	if n.Running == nil {
		return false
	}
	select {
	case <-n.Running:
		return false
	default:
		return true
	}
}

// Init accepts the Notifier interface to initialize the notifier
//func Init(n Notifier) (*Notification, error) {
//	if Exists(n.Select().Method) {
//		AllCommunications = append(AllCommunications, n)
//	} else {
//		_, err := insertDatabase(n)
//		if err != nil {
//			log.Errorln(err)
//			return nil, err
//		}
//		AllCommunications = append(AllCommunications, n)
//	}
//
//		notify, err := SelectNotification(n)
//		if err != nil {
//			return nil, errors.Wrap(err, "error selecting notification")
//		}
//
//		notify.CreatedAt = time.Now().UTC()
//		notify.UpdatedAt = time.Now().UTC()
//		if notify.Delay.Seconds() == 0 {
//			notify.Delay = 1 * time.Second
//		}
//		notify.testable = utils.IsType(n, new(Tester))
//		notify.Form = n.Select().Form
//
//		AllCommunications = append(AllCommunications, n)
//
//	return nil, err
//}
