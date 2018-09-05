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
	"fmt"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"github.com/jinzhu/gorm"
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
	Host      string             `gorm:"column:host" json:"-"`
	Port      int                `gorm:"column:port" json:"-"`
	Username  string             `gorm:"column:username" json:"-"`
	Password  string             `gorm:"column:password" json:"-"`
	Var1      string             `gorm:"column:var1" json:"-"`
	Var2      string             `gorm:"column:var2" json:"-"`
	ApiKey    string             `gorm:"column:api_key" json:"-"`
	ApiSecret string             `gorm:"column:api_secret" json:"-"`
	Enabled   bool               `gorm:"column:enabled" json:"enabled"`
	Limits    int                `gorm:"column:limits" json:"-"`
	Removable bool               `gorm:"column:removable" json:"-"`
	CreatedAt time.Time          `gorm:"column:created_at" json:"created_at"`
	Form      []NotificationForm `gorm:"-" json:"-"`
	Routine   chan struct{}      `gorm:"-" json:"-"`
}

type Notifier interface {
	Init() error
	Install() error
	Run() error
	OnFailure(*types.Service) error
	OnSuccess(*types.Service) error
	Select() *Notification
	Test() error
}

type NotificationForm struct {
	Id          int64
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

func add(c interface{}) {
	AllCommunications = append(AllCommunications, c)
}

func Load() []types.AllNotifiers {
	utils.Log(1, "Loading notifiers")
	var notifiers []types.AllNotifiers
	for _, comm := range AllCommunications {
		n := comm.(Notifier)
		n.Init()
		notifiers = append(notifiers, n)
		n.Test()
	}
	return notifiers
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
	n.CreatedAt = time.Now()
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

func SelectNotifier(id int64) Notifier {
	var notifier Notifier
	for _, n := range AllCommunications {
		notif := n.(Notifier)
		n := notif.Select()
		if n.Id == id {
			return notif
		}
	}
	return notifier
}

func (f Notification) CanSend() bool {
	if f.SentLastHour() >= f.Limits {
		return false
	}
	return true
}

func (f Notification) SentLastHour() int {
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

func (f NotificationForm) Value() string {
	noti := SelectNotifier(f.Id)
	return noti.Select().GetValue(f.DbField)
}

func (f Notification) LimitValue() int64 {
	notifier, _ := SelectNotification(f.Id)
	return utils.StringInt(notifier.GetValue("limits"))
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

func OnFailure(s *types.Service) {
	for _, comm := range AllCommunications {
		n := comm.(Notifier)
		n.OnFailure(s)
	}
}

func OnSuccess(s *types.Service) {
	for _, comm := range AllCommunications {
		n := comm.(Notifier)
		n.OnSuccess(s)
	}
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
