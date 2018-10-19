// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package notifier

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"github.com/jinzhu/gorm"
	"reflect"
	"strings"
	"time"
)

var (
	// AllCommunications holds all the loaded notifiers
	AllCommunications []types.AllNotifiers
	// db holds the Statup database connection
	db *gorm.DB
)

// Notification contains all the fields for a Statup Notifier.
type Notification struct {
	Id          int64              `gorm:"primary_key;column:id" json:"id"`
	Method      string             `gorm:"column:method" json:"method"`
	Host        string             `gorm:"not null;column:host" json:"-"`
	Port        int                `gorm:"not null;column:port" json:"-"`
	Username    string             `gorm:"not null;column:username" json:"-"`
	Password    string             `gorm:"not null;column:password" json:"-"`
	Var1        string             `gorm:"not null;column:var1" json:"-"`
	Var2        string             `gorm:"not null;column:var2" json:"-"`
	ApiKey      string             `gorm:"not null;column:api_key" json:"-"`
	ApiSecret   string             `gorm:"not null;column:api_secret" json:"-"`
	Enabled     bool               `gorm:"column:enabled;type:boolean;default:false" json:"enabled"`
	Limits      int                `gorm:"not null;column:limits" json:"-"`
	Removable   bool               `gorm:"column:removable" json:"-"`
	CreatedAt   time.Time          `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time          `gorm:"column:updated_at" json:"updated_at"`
	Form        []NotificationForm `gorm:"-" json:"-"`
	logs        []*NotificationLog `gorm:"-" json:"-"`
	Title       string             `gorm:"-" json:"-"`
	Description string             `gorm:"-" json:"-"`
	Author      string             `gorm:"-" json:"-"`
	AuthorUrl   string             `gorm:"-" json:"-"`
	Delay       time.Duration      `gorm:"-" json:"-"`
	Queue       []interface{}      `gorm:"-" json:"-"`
	Running     chan bool          `gorm:"-" json:"-"`
	Online      bool               `gorm:"-" json:"-"`
	testable    bool               `gorm:"-" json:"-"`
}

// NotificationForm contains the HTML fields for each variable/input you want the notifier to accept.
type NotificationForm struct {
	Type        string // the html input type (text, password, email)
	Title       string // include a title for ease of use
	Placeholder string // add a placeholder for the input
	DbField     string // true variable key for input
	SmallText   string // insert small text under a html input
	Required    bool   // require this input on the html form
}

// NotificationLog contains the normalized message from previously sent notifications
type NotificationLog struct {
	Message   string
	Time      utils.Timestamp
	Timestamp time.Time
}

// AddQueue will add any type of interface (json, string, struct, etc) into the Notifiers queue
func (n *Notification) AddQueue(msg interface{}) {
	n.Queue = append(n.Queue, msg)
}

// CanTest returns true if the notifier implements the OnTest interface
func (n *Notification) CanTest() bool {
	return n.testable
}

// db will return the notifier database column/record
func modelDb(n *Notification) *gorm.DB {
	return db.Model(&Notification{}).Where("method = ?", n.Method).Find(n)
}

// SetDB is called by core to inject the database for a notifier to use
func SetDB(d *gorm.DB) {
	db = d
}

// asNotification accepts a Notifier and returns a Notification struct
func asNotification(n Notifier) *Notification {
	return n.Select()
}

// AddNotifier accept a Notifier interface to be added into the array
func AddNotifier(n Notifier) error {
	if isType(n, new(Notifier)) {
		err := checkNotifierForm(n)
		if err != nil {
			return err
		}
		AllCommunications = append(AllCommunications, n)
	} else {
		return errors.New("notifier does not have the required methods")
	}
	return nil
}

// Load is called by core to add all the notifier into memory
func Load() []types.AllNotifiers {
	var notifiers []types.AllNotifiers
	for _, comm := range AllCommunications {
		n := comm.(Notifier)
		Init(n)
		notifiers = append(notifiers, n)
	}
	startAllNotifiers()
	return notifiers
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

// removeQueue will remove a specific notification and return the new one
func (n *Notification) removeQueue(msg interface{}) interface{} {
	var newArr []interface{}
	for _, q := range n.Queue {
		if q != msg {
			newArr = append(newArr, q)
		}
	}
	n.Queue = newArr
	return newArr
}

// Log will record a new notification into memory and will show the logs on the settings page
func (n *Notification) makeLog(msg interface{}) {
	log := &NotificationLog{
		Message:   normalizeType(msg),
		Time:      utils.Timestamp(time.Now()),
		Timestamp: time.Now(),
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

// isInDatabase returns true if the notifier has already been installed
func isInDatabase(n *Notification) bool {
	inDb := modelDb(n).RecordNotFound()
	return !inDb
}

// SelectNotification returns the Notification struct from the database
func SelectNotification(n Notifier) (*Notification, error) {
	notifier := n.Select()
	err := db.Model(&Notification{}).Where("method = ?", notifier.Method).Scan(&notifier)
	return notifier, err.Error
}

// Update will update the notification into the database
func Update(n Notifier, notif *Notification) (*Notification, error) {
	err := db.Model(&Notification{}).Update(notif)
	if notif.Enabled {
		notif.close()
		notif.start()
		go Queue(n)
	}
	return notif, err.Error
}

// insertDatabase will create a new record into the database for the notifier
func insertDatabase(n *Notification) (int64, error) {
	n.Limits = 3
	query := db.Create(n)
	if query.Error != nil {
		return 0, query.Error
	}
	return n.Id, query.Error
}

// SelectNotifier returns the Notification struct from the database
func SelectNotifier(method string) (*Notification, Notifier, error) {
	for _, comm := range AllCommunications {
		n, ok := comm.(Notifier)
		if !ok {
			return nil, nil, fmt.Errorf("incorrect notification type: %v", reflect.TypeOf(n).String())
		}
		notifier := n.Select()
		if notifier.Method == method {
			return notifier, comm.(Notifier), nil
		}
	}
	return nil, nil, nil
}

// Init accepts the Notifier interface to initialize the notifier
func Init(n Notifier) (*Notification, error) {
	err := install(n)
	var notify *Notification
	if err == nil {
		notify, _ = SelectNotification(n)
		if notify.Delay.Seconds() == 0 {
			notify.Delay = time.Duration(60 * time.Second)
		}
		notify.testable = isType(n, new(Tester))
		notify.Form = n.Select().Form
	}
	return notify, err
}

// startAllNotifiers will start the go routine for each loaded notifier
func startAllNotifiers() {
	for _, comm := range AllCommunications {
		if isType(comm, new(Notifier)) {
			notify := comm.(Notifier)
			if notify.Select().Enabled {
				notify.Select().close()
				notify.Select().start()
				go Queue(notify)
			}
		}
	}
}

// Queue is the FIFO go routine to send notifications when objects are triggered
func Queue(n Notifier) {
	notification := n.Select()
	rateLimit := notification.Delay

CheckNotifier:
	for {
		select {
		case <-notification.Running:
			break CheckNotifier
		case <-time.After(rateLimit):
			notification = n.Select()
			if len(notification.Queue) > 0 {
				ok, _ := notification.WithinLimits()
				if ok {
					msg := notification.Queue[0]
					err := n.Send(msg)
					if err != nil {
						utils.Log(2, fmt.Sprintf("notifier %v had an error: %v", notification.Method, err))
					}
					notification.makeLog(msg)
					notification.Queue = notification.Queue[1:]
					rateLimit = notification.Delay
				}
			}
		}
		continue
	}
}

// install will check the database for the notification, if its not inserted it will insert a new record for it
func install(n Notifier) error {
	inDb := isInDatabase(n.Select())
	if !inDb {
		_, err := insertDatabase(n.Select())
		if err != nil {
			utils.Log(3, err)
			return err
		}
	}
	return nil
}

// LastSent returns a time.Duration of the last sent notification for the notifier
func (n *Notification) LastSent() time.Duration {
	if len(n.logs) == 0 {
		return time.Duration(0)
	}
	last := n.Logs()[0]
	since := time.Since(last.Timestamp)
	return since
}

// SentLastHour returns the total amount of notifications sent in last 1 hour
func (n *Notification) SentLastHour() int {
	since := time.Now().Add(-1 * time.Hour)
	return n.SentLast(since)
}

// SentLastMinute returns the total amount of notifications sent in last 1 minute
func (n *Notification) SentLastMinute() int {
	since := time.Now().Add(-1 * time.Minute)
	return n.SentLast(since)
}

// SentLast accept a time.Time and returns the amount of sent notifications within your time to current
func (n *Notification) SentLast(since time.Time) int {
	sent := 0
	for _, v := range n.Logs() {
		lastTime := time.Time(v.Time)
		if lastTime.After(since) {
			sent++
		}
	}
	return sent
}

// GetValue returns the database value of a accept DbField value.
func (n *Notification) GetValue(dbField string) string {
	dbField = strings.ToLower(dbField)
	switch dbField {
	case "host":
		return n.Host
	case "port":
		return fmt.Sprintf("%v", n.Port)
	case "username":
		return n.Username
	case "password":
		if n.Password != "" {
			return "##########"
		}
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
	}
	return ""
}

// isType will return true if a variable can implement an interface
func isType(n interface{}, obj interface{}) bool {
	one := reflect.TypeOf(n)
	two := reflect.ValueOf(obj).Elem()
	return one.Implements(two.Type())
}

// isEnabled returns true if the notifier is enabled
func isEnabled(n interface{}) bool {
	notifier, _ := SelectNotification(n.(Notifier))
	return notifier.Enabled
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
	if n.Delay.Seconds() == 0 {
		n.Delay = time.Duration(500 * time.Millisecond)
	}
	if n.LastSent().Seconds() == 0 {
		return true, nil
	}
	if n.Delay.Seconds() >= n.LastSent().Seconds() {
		return false, fmt.Errorf("notifiers delay (%v) is greater than last message sent (%v)", n.Delay.Seconds(), n.LastSent().Seconds())
	}
	return true, nil
}

// ResetQueue will clear the notifiers Queue
func (n *Notification) ResetQueue() {
	n.Queue = nil
}

// start will start the go routine for the notifier queue
func (n *Notification) start() {
	n.Running = make(chan bool)
}

// close will stop the go routine for queue
func (n *Notification) close() {
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
