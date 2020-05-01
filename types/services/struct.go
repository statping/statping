package services

import (
	"github.com/statping/statping/types/checkins"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/utils"
	"time"
)

var (
	allServices map[int64]*Service
)

func init() {
	allServices = make(map[int64]*Service)
}

func Services() map[int64]*Service {
	return allServices
}

// Service is the main struct for Services
type Service struct {
	Id                  int64               `gorm:"primary_key;column:id" json:"id" yaml:"id"`
	Name                string              `gorm:"column:name" json:"name" yaml:"name"`
	Domain              string              `gorm:"column:domain" json:"domain" yaml:"domain" private:"true" scope:"user,admin"`
	Expected            null.NullString     `gorm:"column:expected" json:"expected" yaml:"expected" scope:"user,admin"`
	ExpectedStatus      int                 `gorm:"default:200;column:expected_status" json:"expected_status" yaml:"expected_status" scope:"user,admin"`
	Interval            int                 `gorm:"default:30;column:check_interval" json:"check_interval" yaml:"check_interval"`
	Type                string              `gorm:"column:check_type" json:"type" scope:"user,admin" yaml:"type"`
	Method              string              `gorm:"column:method" json:"method" scope:"user,admin" yaml:"method"`
	PostData            null.NullString     `gorm:"column:post_data" json:"post_data" scope:"user,admin" yaml:"post_data"`
	Port                int                 `gorm:"not null;column:port" json:"port" scope:"user,admin" yaml:"port"`
	Timeout             int                 `gorm:"default:30;column:timeout" json:"timeout" scope:"user,admin" yaml:"timeout"`
	Order               int                 `gorm:"default:0;column:order_id" json:"order_id" yaml:"order_id"`
	VerifySSL           null.NullBool       `gorm:"default:false;column:verify_ssl" json:"verify_ssl" scope:"user,admin" yaml:"verify_ssl"`
	Public              null.NullBool       `gorm:"default:true;column:public" json:"public" yaml:"public"`
	GroupId             int                 `gorm:"default:0;column:group_id" json:"group_id" yaml:"group_id"`
	Headers             null.NullString     `gorm:"column:headers" json:"headers" scope:"user,admin" yaml:"headers"`
	Permalink           null.NullString     `gorm:"column:permalink;unique;" json:"permalink" yaml:"permalink"`
	Redirect            null.NullBool       `gorm:"default:false;column:redirect" json:"redirect" scope:"user,admin" yaml:"redirect"`
	CreatedAt           time.Time           `gorm:"column:created_at" json:"created_at" yaml:"-"`
	UpdatedAt           time.Time           `gorm:"column:updated_at" json:"updated_at" yaml:"-"`
	Online              bool                `gorm:"-" json:"online" yaml:"-"`
	Latency             int64               `gorm:"-" json:"latency" yaml:"-"`
	PingTime            int64               `gorm:"-" json:"ping_time" yaml:"-"`
	Online24Hours       float32             `gorm:"-" json:"online_24_hours" yaml:"-"`
	Online7Days         float32             `gorm:"-" json:"online_7_days" yaml:"-"`
	AvgResponse         int64               `gorm:"-" json:"avg_response" yaml:"-"`
	FailuresLast24Hours int                 `gorm:"-" json:"failures_24_hours" yaml:"-"`
	Running             chan bool           `gorm:"-" json:"-" yaml:"-"`
	Checkpoint          time.Time           `gorm:"-" json:"-" yaml:"-"`
	SleepDuration       time.Duration       `gorm:"-" json:"-" yaml:"-"`
	LastResponse        string              `gorm:"-" json:"-" yaml:"-"`
	NotifyAfter         int64               `gorm:"column:notify_after" json:"notify_after" yaml:"notify_after" scope:"user,admin"`
	notifyAfterCount    int64               `gorm:"-" json:"-" yaml:"-"`
	AllowNotifications  null.NullBool       `gorm:"default:true;column:allow_notifications" json:"allow_notifications" yaml:"allow_notifications" scope:"user,admin"`
	UserNotified        bool                `gorm:"-" json:"-" yaml:"-"`                                                                                           // True if the User was already notified about a Downtime
	UpdateNotify        null.NullBool       `gorm:"default:true;column:notify_all_changes" json:"notify_all_changes" yaml:"notify_all_changes" scope:"user,admin"` // This Variable is a simple copy of `core.CoreApp.UpdateNotify.Bool`
	DownText            string              `gorm:"-" json:"-" yaml:"-"`                                                                                           // Contains the current generated Downtime Text
	SuccessNotified     bool                `gorm:"-" json:"-" yaml:"-"`                                                                                           // Is 'true' if the user has already be informed that the Services now again available
	LastStatusCode      int                 `gorm:"-" json:"status_code" yaml:"-"`
	Failures            []*failures.Failure `gorm:"-" json:"failures,omitempty" yaml:"-" scope:"user,admin"`
	AllCheckins         []*checkins.Checkin `gorm:"-" json:"checkins,omitempty" yaml:"-" scope:"user,admin"`
	LastLookupTime      int64               `gorm:"-" json:"-" yaml:"-"`
	LastLatency         int64               `gorm:"-" json:"-" yaml:"-"`
	LastCheck           time.Time           `gorm:"-" json:"-" yaml:"-"`
	LastOnline          time.Time           `gorm:"-" json:"last_success" yaml:"-"`
	LastOffline         time.Time           `gorm:"-" json:"last_error" yaml:"-"`
	Stats               *Stats              `gorm:"-" json:"stats,omitempty" yaml:"-"`

	SecondsOnline  int64 `gorm:"-" json:"-" yaml:"-"`
	SecondsOffline int64 `gorm:"-" json:"-" yaml:"-"`
}

type Stats struct {
	Failures int       `gorm:"-" json:"failures"`
	Hits     int       `gorm:"-" json:"hits"`
	FirstHit time.Time `gorm:"-" json:"first_hit"`
}

// BeforeCreate for Service will set CreatedAt to UTC
func (s *Service) BeforeCreate() (err error) {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = utils.Now()
		s.UpdatedAt = utils.Now()
	}
	return
}

// ServiceOrder will reorder the services based on 'order_id' (Order)
type ServiceOrder []Service

// Sort interface for resroting the Services in order
func (c ServiceOrder) Len() int           { return len(c) }
func (c ServiceOrder) Swap(i, j int)      { c[int64(i)], c[int64(j)] = c[int64(j)], c[int64(i)] }
func (c ServiceOrder) Less(i, j int) bool { return c[i].Order < c[j].Order }
