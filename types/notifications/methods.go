package notifications

import (
	"fmt"
	"github.com/statping/statping/utils"
	"strings"
	"time"
)

func (n Notification) Name() string {
	newName := strings.ToLower(n.Method)
	newName = strings.ReplaceAll(newName, " ", "_")
	return newName
}

// LastSent returns a time.Duration of the last sent notification for the notifier
func (n Notification) LastSent() time.Duration {
	return time.Since(n.lastSent)
}

func (n *Notification) CanSend() bool {
	if !n.Enabled.Bool {
		return false
	}

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
		return utils.ToString(n.Limits)
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
