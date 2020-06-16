package notifications

import (
	"github.com/sirupsen/logrus"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/utils"
	"time"
)

var (
	log = utils.Log.WithField("type", "notifier")
)

// Notification contains all the fields for a Statping Notifier.
type Notification struct {
	Id          int64         `gorm:"primary_key;column:id" json:"id"`
	Method      string        `gorm:"column:method" json:"method"`
	Host        string        `gorm:"not null;column:host" json:"host,omitempty"`
	Port        int           `gorm:"not null;column:port" json:"port,omitempty"`
	Username    string        `gorm:"not null;column:username" json:"username,omitempty"`
	Password    string        `gorm:"not null;column:password" json:"password,omitempty"`
	Var1        string        `gorm:"not null;column:var1" json:"var1,omitempty"`
	Var2        string        `gorm:"not null;column:var2" json:"var2,omitempty"`
	ApiKey      string        `gorm:"not null;column:api_key" json:"api_key,omitempty"`
	ApiSecret   string        `gorm:"not null;column:api_secret" json:"api_secret,omitempty"`
	Enabled     null.NullBool `gorm:"column:enabled;type:boolean;default:false" json:"enabled,omitempty"`
	Limits      int           `gorm:"not null;column:limits" json:"limits"`
	Removable   bool          `gorm:"column:removable" json:"removable"`
	SuccessData string        `gorm:"type:text;column:success_data" json:"success_data,omitempty"`
	FailureData string        `gorm:"type:text;column:failure_data" json:"failure_data,omitempty"`
	DataType    string        `gorm:"-" json:"data_type,omitempty"`
	RequestInfo string        `gorm:"-" json:"request_info,omitempty"`
	CreatedAt   time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time     `gorm:"column:updated_at" json:"updated_at"`
	Title       string        `gorm:"-" json:"title"`
	Description string        `gorm:"-" json:"description"`
	Author      string        `gorm:"-" json:"author"`
	AuthorUrl   string        `gorm:"-" json:"author_url"`
	Icon        string        `gorm:"-" json:"icon"`
	Delay       time.Duration `gorm:"-" json:"delay,string"`
	Running     chan bool     `gorm:"-" json:"-"`

	Form          []NotificationForm `gorm:"-" json:"form"`
	lastSent      time.Time          `gorm:"-" json:"-"`
	lastSentCount int                `gorm:"-" json:"-"`
}

func (n *Notification) Logger() *logrus.Logger {
	return log.WithField("notifier", n.Method).Logger
}

type RunFunc func(interface{}) error

// NotificationForm contains the HTML fields for each variable/input you want the notifier to accept.
type NotificationForm struct {
	Type        string   `json:"type"`        // the html input type (text, password, email)
	Title       string   `json:"title"`       // include a title for ease of use
	Placeholder string   `json:"placeholder"` // add a placeholder for the input
	DbField     string   `json:"field"`       // true variable key for input
	SmallText   string   `json:"small_text"`  // insert small text under a html input
	Required    bool     `json:"required"`    // require this input on the html form
	IsHidden    bool     `json:"hidden"`      // hide this form element from end user
	ListOptions []string `json:"list_options,omitempty"`
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

// NotificationOrder will reorder the services based on 'order_id' (Order)
type NotificationOrder []Notification

// Sort interface for resorting the Notifications in order
func (c NotificationOrder) Len() int           { return len(c) }
func (c NotificationOrder) Swap(i, j int)      { c[int64(i)], c[int64(j)] = c[int64(j)], c[int64(i)] }
func (c NotificationOrder) Less(i, j int) bool { return c[i].Id < c[j].Id }
