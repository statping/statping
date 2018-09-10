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

package notifiers

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
	Collections       *gorm.DB
	Logs              []*NotificationLog
)

type Notification struct {
	Id        int64              `gorm:"primary_key column:id" json:"id"`
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
	Notifier
}

type NotificationForm struct {
	Type        string
	Title       string
	Placeholder string
	DbField     string
}

type NotificationLog struct {
	Notifier *Notification
	Message  string
	Time     utils.Timestamp
}

func AddNotifier(c interface{}) error {
	if _, ok := c.(Notifier); ok {
		AllCommunications = append(AllCommunications, c)
	} else {
		return errors.New("notifier does not have the required methods")
	}
	return nil
}

func Load() []types.AllNotifiers {
	var notifiers []types.AllNotifiers
	for _, comm := range AllCommunications {
		n := comm.(Notifier)
		n.Init()
		notifiers = append(notifiers, n)
		n.Test()
	}
	return notifiers
}

func (n *Notification) Select() *Notification {
	return n
}

func (n *Notification) Log(msg string) {
	log := &NotificationLog{
		Notifier: n,
		Message:  msg,
		Time:     utils.Timestamp(time.Now()),
	}
	Logs = append(Logs, log)
}

func (n *Notification) Logs() []*NotificationLog {
	var logs []*NotificationLog
	for _, v := range Logs {
		if v.Notifier.Id == n.Id {
			logs = append(logs, v)
		}
	}
	return reverseLogs(logs)
}

func reverseLogs(input []*NotificationLog) []*NotificationLog {
	if len(input) == 0 {
		return input
	}
	return append(reverseLogs(input[1:]), input[0])
}

func (n *Notification) IsInDatabase() bool {
	return !Collections.Find(n).RecordNotFound()
}

func SelectNotification(id int64) (*Notification, error) {
	var notifier Notification
	err := Collections.Find(&notifier, id)
	return &notifier, err.Error
}

func (n *Notification) Update() (*Notification, error) {
	err := Collections.Update(n)
	return n, err.Error
}

func InsertDatabase(n *Notification) (int64, error) {
	n.CreatedAt = time.Now()
	n.Limits = 3
	db := Collections.Create(n)
	if db.Error != nil {
		return 0, db.Error
	}
	return n.Id, db.Error
}

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

func (f *Notification) CanSend() bool {
	if f.SentLastHour() >= f.Limits {
		return false
	}
	return true
}

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

func (f *Notification) LimitValue() int64 {
	return utils.StringInt(f.GetValue("limits"))
}

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

func IsType(n interface{}, obj string) bool {
	objOne := reflect.TypeOf(n)
	return objOne.String() == obj
}

func uniqueStrings(elements []string) []string {
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
