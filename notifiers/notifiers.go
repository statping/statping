package notifiers

import (
	"github.com/hunterlong/statup/utils"
	"time"
	"upper.io/db.v3"
)

var (
	AllCommunications []*Notification
	Collections       db.Collection
)

func add(c *Notification) {
	AllCommunications = append(AllCommunications, c)
}

func Load() {
	utils.Log(1, "Loading notifiers")
	for _, comm := range AllCommunications {
		comm.Init()
	}
}

type Notification struct {
	Id        int64     `db:"id,omitempty" json:"id"`
	Method    string    `db:"method" json:"method"`
	Host      string    `db:"host" json:"-"`
	Port      int       `db:"port" json:"-"`
	Username  string    `db:"username" json:"-"`
	Password  string    `db:"password" json:"-"`
	Var1      string    `db:"var1" json:"-"`
	Var2      string    `db:"var2" json:"-"`
	ApiKey    string    `db:"api_key" json:"-"`
	ApiSecret string    `db:"api_secret" json:"-"`
	Enabled   bool      `db:"enabled" json:"enabled"`
	Limits    int64     `db:"limits" json:"-"`
	Removable bool      `db:"removable" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	Form      []NotificationForm
	Routine   chan struct{}
}

type Notifier interface {
	Init() error
	Run() error
	OnFailure() error
	OnSuccess() error
}

type NotificationForm struct {
	Type        string
	Title       string
	Placeholder string
}

func OnFailure() {
	for _, comm := range AllCommunications {
		comm.OnFailure()
	}
}

func OnSuccess() {
	for _, comm := range AllCommunications {
		comm.OnSuccess()
	}
}

func uniqueMessages(arr []string, v string) []string {
	var newArray []string
	for _, i := range arr {
		if i != v {
			newArray = append(newArray, v)
		}
	}
	return newArray
}
