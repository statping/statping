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
	AllCommunications []types.AllNotifiers
	db                *gorm.DB
)

type Notification struct {
	Id        int64              `gorm:"primary_key;column:id" json:"id"`
	Method    string             `gorm:"column:method" json:"method"`
	Host      string             `gorm:"not null;column:host" json:"-"`
	Port      int                `gorm:"not null;column:port" json:"-"`
	Username  string             `gorm:"not null;column:username" json:"-"`
	Password  string             `gorm:"not null;column:password" json:"-"`
	Var1      string             `gorm:"not null;column:var1" json:"-"`
	Var2      string             `gorm:"not null;column:var2" json:"-"`
	ApiKey    string             `gorm:"not null;column:api_key" json:"-"`
	ApiSecret string             `gorm:"not null;column:api_secret" json:"-"`
	Enabled   bool               `gorm:"column:enabled;type:boolean;default:false" json:"enabled"`
	Limits    int                `gorm:"not null;column:limits" json:"-"`
	Removable bool               `gorm:"column:removable" json:"-"`
	CreatedAt time.Time          `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time          `gorm:"column:updated_at" json:"updated_at"`
	Form      []NotificationForm `gorm:"-" json:"-"`
	Routine   chan struct{}      `gorm:"-" json:"-"`
	logs      []*NotificationLog `gorm:"-" json:"-"`
}

type NotificationForm struct {
	Type        string
	Title       string
	Placeholder string
	DbField     string
}

type NotificationLog struct {
	Message   string
	Time      utils.Timestamp
	Timestamp time.Time
}

// db will return the notifier database column/record
func (n *Notification) db() *gorm.DB {
	return db.Model(&Notification{}).Where("method = ?", n.Method).Find(n)
}

// SetDB is called by core to inject the database for a notifier to use
func SetDB(d *gorm.DB) {
	db = d
}

// AddNotifier accept a Notifier interface to be added into the array
func AddNotifier(c Notifier) error {
	if notifier, ok := c.(Notifier); ok {
		err := checkNotifierForm(notifier)
		if err != nil {
			return err
		}
		AllCommunications = append(AllCommunications, notifier)
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
		//n.Test()
	}
	return notifiers
}

func (n *Notification) Select() *Notification {
	return n
}

// Log will record a new notification into memory and will show the logs on the settings page
func (n *Notification) Log(msg string) {
	log := &NotificationLog{
		Message:   msg,
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
func (n *Notification) isInDatabase() bool {
	inDb := n.db().RecordNotFound()
	return !inDb
}

// SelectNotification returns the Notification struct from the database
func SelectNotification(method string) (*Notification, error) {
	var notifier Notification
	err := db.Model(&Notification{}).Where("method = ?", method).Scan(&notifier)
	return &notifier, err.Error
}

// Update will update the notification into the database
func (n *Notification) Update() (*Notification, error) {
	err := n.db().Update(n)
	return n, err.Error
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
func SelectNotifier(method string) (*Notification, error) {
	for _, comm := range AllCommunications {
		n, ok := comm.(Notifier)
		if !ok {
			return nil, errors.New(fmt.Sprintf("incorrect notification type: %v", reflect.TypeOf(n).String()))
		}
		notifier := n.Select()
		if notifier.Method == method {
			return notifier, nil
		}
	}
	return nil, nil
}

// CanSend will return true if notifier has not passed its Limits within the last hour
func (f *Notification) CanSend() bool {
	if f.SentLastHour() >= f.Limits {
		return false
	}
	return true
}

// Init accepts the Notifier interface to initialize the notifier
func Init(n Notifier) (*Notification, error) {
	err := install(n)
	var notify *Notification
	if err == nil {
		notify, _ = SelectNotification(n.Select().Method)
		notify.Form = n.Select().Form
	}
	return notify, err
}

// install will check the database for the notification, if its not inserted it will insert a new record for it
func install(n Notifier) error {
	inDb := n.Select().isInDatabase()
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
func (f *Notification) LastSent() time.Duration {
	if len(f.logs) == 0 {
		return time.Duration(0)
	}
	last := f.Logs()[0]
	since := time.Since(last.Timestamp)
	return since
}

// SentLastHour returns the amount of sent notifications within the last hour
func (f *Notification) SentLastHour() int {
	sent := 0
	hourAgo := time.Now().Add(-1 * time.Hour)
	for _, v := range f.Logs() {
		lastTime := time.Time(v.Time)
		if lastTime.After(hourAgo) {
			sent++
		}
	}
	return sent
}

// Limit returns the limits on how many notifications can be sent in 1 hour
func (f *Notification) Limit() int64 {
	return utils.StringInt(f.GetValue("limits"))
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
		return utils.IntString(int(n.Limits))
	}
	return ""
}

// isType will return true if a variable can implement an interface
func isType(n interface{}, obj interface{}) bool {
	objOne := reflect.TypeOf(n)
	obj2 := reflect.TypeOf(obj)
	return objOne.String() == obj2.String()
}

// isEnabled returns true if the notifier is enabled
func isEnabled(n interface{}) bool {
	notify := n.(Notifier).Select()
	return notify.Enabled
}

func UniqueStrings(elements []string) []string {
	result := []string{}

	for i := 0; i < len(elements); i++ {
		// Scan slice for a previous element of the same value.
		exists := false
		for v := 0; v < i; v++ {
			if elements[v] == elements[i] {
				exists = true
				break
			}
		}
		// If no previous element exists, append this one.
		if !exists {
			result = append(result, elements[i])
		}
	}
	return result
}
