package services

import (
	"github.com/hunterlong/statping/types/checkins"
	"github.com/hunterlong/statping/types/failures"
	"github.com/hunterlong/statping/types/null"
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
	Id                  int64               `gorm:"primary_key;column:id" json:"id"`
	Name                string              `gorm:"column:name" json:"name"`
	Domain              string              `gorm:"column:domain" json:"domain" private:"true" scope:"user,admin"`
	Expected            null.NullString     `gorm:"column:expected" json:"expected" scope:"user,admin"`
	ExpectedStatus      int                 `gorm:"default:200;column:expected_status" json:"expected_status" scope:"user,admin"`
	Interval            int                 `gorm:"default:30;column:check_interval" json:"check_interval" scope:"user,admin"`
	Type                string              `gorm:"column:check_type" json:"type" scope:"user,admin"`
	Method              string              `gorm:"column:method" json:"method" scope:"user,admin"`
	PostData            null.NullString     `gorm:"column:post_data" json:"post_data" scope:"user,admin"`
	Port                int                 `gorm:"not null;column:port" json:"port" scope:"user,admin"`
	Timeout             int                 `gorm:"default:30;column:timeout" json:"timeout" scope:"user,admin"`
	Order               int                 `gorm:"default:0;column:order_id" json:"order_id"`
	VerifySSL           null.NullBool       `gorm:"default:false;column:verify_ssl" json:"verify_ssl" scope:"user,admin"`
	Public              null.NullBool       `gorm:"default:true;column:public" json:"public"`
	GroupId             int                 `gorm:"default:0;column:group_id" json:"group_id"`
	Headers             null.NullString     `gorm:"column:headers" json:"headers" scope:"user,admin"`
	Permalink           null.NullString     `gorm:"column:permalink" json:"permalink"`
	CreatedAt           time.Time           `gorm:"column:created_at" json:"created_at"`
	UpdatedAt           time.Time           `gorm:"column:updated_at" json:"updated_at"`
	Online              bool                `gorm:"-" json:"online"`
	Latency             float64             `gorm:"-" json:"latency"`
	PingTime            float64             `gorm:"-" json:"ping_time"`
	Online24Hours       float32             `gorm:"-" json:"online_24_hours"`
	Online7Days         float32             `gorm:"-" json:"online_7_days"`
	AvgResponse         float64             `gorm:"-" json:"avg_response"`
	FailuresLast24Hours int                 `gorm:"-" json:"failures_24_hours"`
	Running             chan bool           `gorm:"-" json:"-"`
	Checkpoint          time.Time           `gorm:"-" json:"-"`
	SleepDuration       time.Duration       `gorm:"-" json:"-"`
	LastResponse        string              `gorm:"-" json:"-"`
	AllowNotifications  null.NullBool       `gorm:"default:true;column:allow_notifications" json:"allow_notifications" scope:"user,admin"`
	UserNotified        bool                `gorm:"-" json:"-"`                                                                          // True if the User was already notified about a Downtime
	UpdateNotify        null.NullBool       `gorm:"default:true;column:notify_all_changes" json:"notify_all_changes" scope:"user,admin"` // This Variable is a simple copy of `core.CoreApp.UpdateNotify.Bool`
	DownText            string              `gorm:"-" json:"-"`                                                                          // Contains the current generated Downtime Text
	SuccessNotified     bool                `gorm:"-" json:"-"`                                                                          // Is 'true' if the user has already be informed that the Services now again available
	LastStatusCode      int                 `gorm:"-" json:"status_code"`
	LastOnline          time.Time           `gorm:"-" json:"last_success"`
	Failures            []*failures.Failure `gorm:"-" json:"failures,omitempty" scope:"user,admin"`
	AllCheckins         []*checkins.Checkin `gorm:"-" json:"checkins,omitempty" scope:"user,admin"`
	Stats               *Stats              `gorm:"-" json:"stats,omitempty"`
}

type Stats struct {
	Failures int `gorm:"-" json:"failures"`
	Hits     int `gorm:"-" json:"hits"`
}

// BeforeCreate for Service will set CreatedAt to UTC
func (s *Service) BeforeCreate() (err error) {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now().UTC()
		s.UpdatedAt = time.Now().UTC()
	}
	return
}

// ServiceOrder will reorder the services based on 'order_id' (Order)
type ServiceOrder map[int64]*Service

// Sort interface for resroting the Services in order
func (c ServiceOrder) Len() int      { return len(c) }
func (c ServiceOrder) Swap(i, j int) { c[int64(i)], c[int64(j)] = c[int64(j)], c[int64(i)] }
func (c ServiceOrder) Less(i, j int) bool {
	if c[int64(i)] == nil {
		return false
	}
	if c[int64(j)] == nil {
		return false
	}
	return c[int64(i)].Order < c[int64(j)].Order
}
