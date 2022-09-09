package services

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/razorpay/statping/types/checkins"
	"github.com/razorpay/statping/types/downtimes"
	"github.com/razorpay/statping/types/failures"
	"github.com/razorpay/statping/types/incidents"
	"github.com/razorpay/statping/types/messages"
	"github.com/razorpay/statping/types/null"
	"time"
)

type ServiceWithDowntime struct {
	Service
	Downtime *downtimes.Downtime `gorm:"-" json:"downtime" yaml:"-"`
}

// Service is the main struct for Services
type Service struct {
	Id                  int64                 `gorm:"primary_key;column:id" json:"id" yaml:"id"`
	Name                string                `gorm:"column:name" json:"name" yaml:"name"`
	Description         string                `gorm:"column:description" json:"description" yaml:"-"`
	DisplayName         string                `json:"display_name"`
	Domain              string                `gorm:"column:domain" json:"domain" yaml:"domain" private:"true" scope:"user,admin"`
	Expected            null.NullString       `gorm:"column:expected" json:"expected" yaml:"expected" scope:"user,admin"`
	ExpectedStatus      int                   `gorm:"default:200;column:expected_status" json:"expected_status" yaml:"expected_status" scope:"user,admin"`
	Interval            int                   `gorm:"default:30;column:check_interval" json:"check_interval" yaml:"check_interval"`
	Type                string                `gorm:"column:check_type" json:"type" yaml:"type"`
	Method              string                `gorm:"column:method" json:"method" scope:"user,admin" yaml:"method"`
	PostData            null.NullString       `gorm:"column:post_data" json:"post_data" scope:"user,admin" yaml:"post_data"`
	Port                int                   `gorm:"not null;column:port" json:"port" scope:"user,admin" yaml:"port"`
	Timeout             int                   `gorm:"default:30;column:timeout" json:"timeout" scope:"user,admin" yaml:"timeout"`
	Order               int                   `gorm:"default:0;column:order_id" json:"order_id" yaml:"order_id"`
	Ftc                 int                   `gorm:"default:3;column:ftc" json:"ftc" yaml:"ftc"`
	VerifySSL           null.NullBool         `gorm:"default:false;column:verify_ssl" json:"verify_ssl" scope:"user,admin" yaml:"verify_ssl"`
	GrpcHealthCheck     null.NullBool         `gorm:"default:false;column:grpc_health_check" json:"grpc_health_check" scope:"user,admin" yaml:"grpc_health_check"`
	Public              null.NullBool         `gorm:"default:true;column:public" json:"public" yaml:"public"`
	GroupId             int                   `gorm:"default:0;column:group_id" json:"group_id" yaml:"group_id"`
	TLSCert             null.NullString       `gorm:"column:tls_cert" json:"tls_cert" scope:"user,admin" yaml:"tls_cert"`
	TLSCertKey          null.NullString       `gorm:"column:tls_cert_key" json:"tls_cert_key" scope:"user,admin" yaml:"tls_cert_key"`
	TLSCertRoot         null.NullString       `gorm:"column:tls_cert_root" json:"tls_cert_root" scope:"user,admin" yaml:"tls_cert_root"`
	Headers             null.NullString       `gorm:"column:headers" json:"headers" scope:"user,admin" yaml:"headers"`
	Permalink           null.NullString       `gorm:"column:permalink" json:"permalink" yaml:"permalink"`
	Redirect            null.NullBool         `gorm:"default:false;column:redirect" json:"redirect" scope:"user,admin" yaml:"redirect"`
	SubServicesDetails  SubServicesDetail     `gorm:"column:sub_services_detail;type:json;DEFAULT:null" json:"sub_services_detail" yaml:"sub_services_detail"`
	CreatedAt           time.Time             `gorm:"column:created_at" json:"created_at" yaml:"-"`
	UpdatedAt           time.Time             `gorm:"column:updated_at" json:"updated_at" yaml:"-"`
	Online              bool                  `gorm:"column:online" json:"online" yaml:"-"`
	Latency             int64                 `gorm:"-" json:"latency" yaml:"-"`
	PingTime            int64                 `gorm:"-" json:"ping_time" yaml:"-"`
	Online24Hours       float32               `gorm:"-" json:"online_24_hours" yaml:"-"`
	Online7Days         float32               `gorm:"-" json:"online_7_days" yaml:"-"`
	AvgResponse         int64                 `gorm:"-" json:"avg_response" yaml:"-"`
	FailuresLast24Hours int                   `gorm:"-" json:"failures_24_hours" yaml:"-"`
	Running             chan bool             `gorm:"-" json:"-" yaml:"-"`
	Checkpoint          time.Time             `gorm:"-" json:"-" yaml:"-"`
	SleepDuration       time.Duration         `gorm:"-" json:"-" yaml:"-"`
	LastResponse        string                `gorm:"-" json:"-" yaml:"-"`
	NotifyAfter         int64                 `gorm:"column:notify_after" json:"notify_after" yaml:"notify_after" scope:"user,admin"`
	AllowNotifications  null.NullBool         `gorm:"default:true;column:allow_notifications" json:"allow_notifications" yaml:"allow_notifications" scope:"user,admin"`
	UpdateNotify        null.NullBool         `gorm:"default:true;column:notify_all_changes" json:"notify_all_changes" yaml:"notify_all_changes" scope:"user,admin"` // This Variable is a simple copy of `core.CoreApp.UpdateNotify.Bool`
	DownText            string                `gorm:"-" json:"-" yaml:"-"`                                                                                           // Contains the current generated Downtime Text 	// Is 'true' if the user has already be informed that the Services now again available // Is 'true' if the user has already be informed that the Services now again available
	LastStatusCode      int                   `gorm:"-" json:"status_code" yaml:"-"`
	LastLookupTime      int64                 `gorm:"-" json:"-" yaml:"-"`
	LastLatency         int64                 `gorm:"-" json:"-" yaml:"-"`
	LastCheck           time.Time             `gorm:"column:last_check" json:"last_check" yaml:"-"`
	LastOnline          time.Time             `gorm:"column:last_success" json:"last_success" yaml:"-"`
	LastOffline         time.Time             `gorm:"column:last_error" json:"last_error" yaml:"-"`
	Stats               *Stats                `gorm:"-" json:"stats,omitempty" yaml:"-"`
	Messages            []*messages.Message   `gorm:"foreignkey:service;association_foreignkey:id" json:"messages,omitempty" yaml:"messages"`
	Incidents           []*incidents.Incident `gorm:"foreignkey:service;association_foreignkey:id" json:"incidents,omitempty" yaml:"incidents"`
	Checkins            []*checkins.Checkin   `gorm:"foreignkey:service;association_foreignkey:id" json:"checkins,omitempty" yaml:"-" scope:"user,admin"`
	Failures            []*failures.Failure   `gorm:"-" json:"failures,omitempty" yaml:"-" scope:"user,admin"`
	LastProcessingTime  time.Time             `gorm:"column:last_processing_time" json:"-"`

	notifyAfterCount int64 `gorm:"column:notify_after_count" yaml:"-"`
	prevOnline       bool  `gorm:"column:prev_online" yaml:"-"`

	FailureCounter  int    `gorm:"column:failure_counter" json:"-" yaml:"-"`
	CurrentDowntime int64  `gorm:"column:current_downtime" json:"-" yaml:"-"`
	LastFailureType string `gorm:"-" json:"-" yaml:"-"`
	ManualDowntime  bool   `gorm:"default:false;column:manual_downtime" json:"manual_downtime"`
}

// ServiceOrder will reorder the services based on 'order_id' (Order)
type ServiceOrder []Service

// Sort interface for resroting the Services in order
func (c ServiceOrder) Len() int           { return len(c) }
func (c ServiceOrder) Swap(i, j int)      { c[int64(i)], c[int64(j)] = c[int64(j)], c[int64(i)] }
func (c ServiceOrder) Less(i, j int) bool { return c[i].Order < c[j].Order }

type Stats struct {
	Failures int       `gorm:"-" json:"failures"`
	Hits     int       `gorm:"-" json:"hits"`
	FirstHit time.Time `gorm:"-" json:"first_hit"`
}

type ser struct {
	Time      time.Time
	Online    bool
	SubStatus string
}

type UptimeSeries struct {
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Uptime   int64     `json:"uptime"`
	Downtime int64     `json:"downtime"`
	Series   []series  `json:"series"`
}

type BlockSeries struct {
	Start    time.Time `json:"start, omitempty"`
	End      time.Time `json:"end, omitempty"`
	Uptime   int64     `json:"uptime, omitempty"`
	Downtime int64     `json:"downtime, omitempty"`
	Series   *[]Block  `json:"series, omitempty"`
}

type ByTime []ser

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Less(i, j int) bool { return a[i].Time.Before(a[j].Time) }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (s *Service) GetFtc() int {
	if s.Ftc < FAILURE_THRESHOLD {
		return FAILURE_THRESHOLD
	} else {
		return s.Ftc
	}
}

type Block struct {
	Timeframe string      `json:"timeframe"`
	Status    string      `json:"status"`
	Downtimes *[]Downtime `json:"downtimes"`
}

type Downtime struct {
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	Duration  int64     `json:"duration"`
	SubStatus string    `json:"sub_status"`
}

type series struct {
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	Duration  int64     `json:"duration"`
	Online    bool      `json:"online"`
	SubStatus string    `json:"sub_status"`
}

type SubServicesDetail map[int64]SubService

type SubService struct {
	DisplayName    string `json:"display_name" yaml:"display_name"`
	DependencyType string `json:"dependency_type" yaml:"dependency_type"`
	Public         bool   `json:"public" yaml:"public"`
}

func (j SubServicesDetail) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *SubServicesDetail) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}
