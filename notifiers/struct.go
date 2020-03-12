// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/statping/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package notifiers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"reflect"
	"time"
)

var (
	// db holds the Statping database connection
	log          = utils.Log.WithField("type", "notifier")
	allNotifiers []Notifier
)

// Notification contains all the fields for a Statping Notifier.
type Notification struct {
	Id          int64              `gorm:"primary_key;column:id" json:"id"`
	Method      string             `gorm:"column:method" json:"method"`
	Host        string             `gorm:"not null;column:host" json:"host,omitempty"`
	Port        int                `gorm:"not null;column:port" json:"port,omitempty"`
	Username    string             `gorm:"not null;column:username" json:"username,omitempty"`
	Password    string             `gorm:"not null;column:password" json:"password,omitempty"`
	Var1        string             `gorm:"not null;column:var1" json:"var1,omitempty"`
	Var2        string             `gorm:"not null;column:var2" json:"var2,omitempty"`
	ApiKey      string             `gorm:"not null;column:api_key" json:"api_key,omitempty"`
	ApiSecret   string             `gorm:"not null;column:api_secret" json:"api_secret,omitempty"`
	Enabled     null.NullBool      `gorm:"column:enabled;type:boolean;default:false" json:"enabled"`
	Limits      int                `gorm:"not null;column:limits" json:"limits"`
	Removable   bool               `gorm:"column:removable" json:"removeable"`
	CreatedAt   time.Time          `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time          `gorm:"column:updated_at" json:"updated_at"`
	Form        []NotificationForm `gorm:"-" json:"form"`
	logs        []*NotificationLog `gorm:"-" json:"logs"`
	Title       string             `gorm:"-" json:"title"`
	Description string             `gorm:"-" json:"description"`
	Author      string             `gorm:"-" json:"author"`
	AuthorUrl   string             `gorm:"-" json:"author_url"`
	Icon        string             `gorm:"-" json:"icon"`
	Delay       time.Duration      `gorm:"-" json:"delay,string"`
	Queue       []*QueueData       `gorm:"-" json:"-"`
	Running     chan bool          `gorm:"-" json:"-"`
	testable    bool               `gorm:"-" json:"testable"`

	Hits notificationHits
	Notifier
}

type notificationHits struct {
	OnSuccess         int64 `gorm:"-" json:"-"`
	OnFailure         int64 `gorm:"-" json:"-"`
	OnSave            int64 `gorm:"-" json:"-"`
	OnNewService      int64 `gorm:"-" json:"-"`
	OnUpdatedService  int64 `gorm:"-" json:"-"`
	OnDeletedService  int64 `gorm:"-" json:"-"`
	OnNewUser         int64 `gorm:"-" json:"-"`
	OnUpdatedUser     int64 `gorm:"-" json:"-"`
	OnDeletedUser     int64 `gorm:"-" json:"-"`
	OnNewNotifier     int64 `gorm:"-" json:"-"`
	OnUpdatedNotifier int64 `gorm:"-" json:"-"`
}

// QueueData is the struct for the messaging queue with service
type QueueData struct {
	Id   string
	Data interface{}
}

// NotificationForm contains the HTML fields for each variable/input you want the notifier to accept.
type NotificationForm struct {
	Type        string `json:"type"`        // the html input type (text, password, email)
	Title       string `json:"title"`       // include a title for ease of use
	Placeholder string `json:"placeholder"` // add a placeholder for the input
	DbField     string `json:"field"`       // true variable key for input
	SmallText   string `json:"small_text"`  // insert small text under a html input
	Required    bool   `json:"required"`    // require this input on the html form
	IsHidden    bool   `json:"hidden"`      // hide this form element from end user
	IsList      bool   `json:"list"`        // make this form element a comma separated list
	IsSwitch    bool   `json:"switch"`      // make the notifier a boolean true/false switch
}

// NotificationLog contains the normalized message from previously sent notifications
type NotificationLog struct {
	Message   string          `json:"message"`
	Time      utils.Timestamp `json:"time"`
	Timestamp time.Time       `json:"timestamp"`
}

// normalizeType will accept multiple interfaces and converts it into a string for logging
func normalizeType(ty interface{}) string {
	switch v := ty.(type) {
	case int, int32, int64:
		return fmt.Sprintf("%v", v)
	case float32, float64:
		return fmt.Sprintf("%v", v)
	case string:
		return v
	case []byte:
		return string(v)
	case []string:
		return fmt.Sprintf("%v", v)
	case interface{}, map[string]interface{}:
		j, _ := json.Marshal(v)
		return string(j)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// Log will record a new notification into memory and will show the logs on the settings page
func (n *Notification) makeLog(msg interface{}) {
	log := &NotificationLog{
		Message:   normalizeType(msg),
		Time:      utils.Timestamp(utils.Now()),
		Timestamp: utils.Now(),
	}
	n.logs = append(n.logs, log)
}

// Logs returns an array of the notifiers logs
func (n *Notification) Logs() []*NotificationLog {
	return reverseLogs(n.logs)
}

// reverseLogs will reverse the notifier's logs to be time desc
func reverseLogs(input []*NotificationLog) []*NotificationLog {
	if len(input) == 0 {
		return input
	}
	return append(reverseLogs(input[1:]), input[0])
}

// SelectNotification returns the Notification struct from the database
func SelectNotification(n Notifier) (*Notification, error) {
	notifier := n.Select()
	err := db.Where("method = ?", notifier.Method).Find(&notifier)
	return notifier, err.Error()
}

// SelectNotifier returns the Notification struct from the database
func SelectNotifier(method string) (*Notification, Notifier, error) {
	for _, comm := range allNotifiers {
		n, ok := comm.(Notifier)
		if !ok {
			return nil, nil, fmt.Errorf("incorrect notification type: %v", reflect.TypeOf(n).String())
		}
		notifier := n.Select()
		if notifier.Method == method {
			return notifier, comm.(Notifier), nil
		}
	}
	return nil, nil, errors.New("cannot find notifier")
}

// Queue is the FIFO go routine to send notifications when objects are triggered
func Queue(notifer Notifier) {
	n := notifer.(*Notification)
	rateLimit := n.Delay

CheckNotifier:
	for {
		select {
		case <-n.Running:
			break CheckNotifier
		case <-time.After(rateLimit):
			n := notifer.(*Notification)
			fmt.Printf("checking %s %d\n", n.Method, len(n.Queue))
			if len(n.Queue) > 0 {
				ok, _ := n.WithinLimits()
				if ok {
					msg := n.Queue[0]
					err := notifer.Send(msg.Data)
					if err != nil {
						log.WithFields(utils.ToFields(n, msg)).Error(fmt.Sprintf("Notifier '%v' had an error: %v", n.Method, err))
					} else {
						log.WithFields(utils.ToFields(n, msg)).Debug(fmt.Sprintf("Notifier '%v' sent outgoing message (%v) %v left in queue.", n.Method, msg.Id, len(n.Queue)))
					}
					n.makeLog(msg.Data)
					if len(n.Queue) > 1 {
						n.Queue = n.Queue[1:]
					} else {
						n.Queue = nil
					}
					rateLimit = n.Delay
				}
			}
		}
		continue
	}
}

// install will check the database for the notification, if its not inserted it will insert a new record for it
//func install(n Notifier) error {
//	log.WithFields(utils.ToFields(n)).
//		Debugln(fmt.Sprintf("Checking if notifier '%v' is installed", n.Select().Method))
//
//	if Exists(n.Select().Method) {
//		AllCommunications = append(AllCommunications, n)
//	} else {
//		_, err := insertDatabase(n)
//		if err != nil {
//			log.Errorln(err)
//			return err
//		}
//		AllCommunications = append(AllCommunications, n)
//	}
//	return nil
//}

// isEnabled returns true if the notifier is enabled
func isEnabled(n interface{}) bool {
	notifier := n.(Notifier).Select()
	return notifier.Enabled.Bool
}

// inLimits will return true if the notifier is within sending limits
func inLimits(n interface{}) bool {
	notifier := n.(Notifier).Select()
	ok, _ := notifier.WithinLimits()
	return ok
}

// WithinLimits returns true if the notifier is within its sending limits
func (n *Notification) WithinLimits() (bool, error) {
	if n.SentLastMinute() == 0 {
		return true, nil
	}
	if n.SentLastMinute() >= n.Limits {
		return false, fmt.Errorf("notifier sent %v out of %v in last minute", n.SentLastMinute(), n.Limits)
	}
	if n.LastSent().Seconds() == 0 {
		return true, nil
	}
	if n.Delay.Seconds() >= n.LastSent().Seconds() {
		return false, fmt.Errorf("notifiers delay (%v) is greater than last message sent (%v)", n.Delay.Seconds(), n.LastSent().Seconds())
	}
	return true, nil
}

// ExampleService can be used for the OnTest() method for notifiers
var ExampleService = &services.Service{
	Id:             1,
	Name:           "Interpol - All The Rage Back Home",
	Domain:         "https://www.youtube.com/watch?v=-u6DvRyyKGU",
	ExpectedStatus: 200,
	Interval:       30,
	Type:           "http",
	Method:         "GET",
	Timeout:        20,
	LastStatusCode: 404,
	Expected:       null.NewNullString("test example"),
	LastResponse:   "<html>this is an example response</html>",
	CreatedAt:      utils.Now().Add(-24 * time.Hour),
}
