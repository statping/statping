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
func (n Notification) LastSentDur() time.Duration {
	return time.Since(n.LastSent)
}

func (n *Notification) CanSend() bool {
	if !n.Enabled.Bool {
		return false
	}
	// the last sent notification was past 1 minute (limit per minute)
	if n.LastSent.Add(60 * time.Minute).Before(utils.Now()) {
		if n.LastSentCount != 0 {
			n.LastSentCount--
		}
	}

	// dont send if already beyond the notifier's limit
	if n.LastSentCount >= n.Limits {
		return false
	}

	return true
}

// GetValue returns the database value of a accept DbField value.
func (n *Notification) GetValue(dbField string) string {
	switch strings.ToLower(dbField) {
	case "host":
		return n.Host.String
	case "port":
		return fmt.Sprintf("%d", n.Port.Int64)
	case "username":
		return n.Username.String
	case "password":
		return n.Password.String
	case "var1":
		return n.Var1.String
	case "var2":
		return n.Var2.String
	case "api_key":
		return n.ApiKey.String
	case "api_secret":
		return n.ApiSecret.String
	case "limits":
		return utils.ToString(n.Limits)
	default:
		return ""
	}
}
