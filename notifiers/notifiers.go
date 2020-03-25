package notifiers

import (
	"github.com/prometheus/common/log"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"time"
)

func InitNotifiers() {
	Add(
		slacker,
		Command,
		Discorder,
		email,
		LineNotify,
		Telegram,
		Twilio,
		Webhook,
		Mobile,
	)
}

func Add(notifs ...services.ServiceNotifier) {
	for _, n := range notifs {
		services.AddNotifier(n)
		if err := n.Select().Create(); err != nil {
			log.Error(err)
		}
	}
}

func ToMap(srv *services.Service, f *failures.Failure) map[string]interface{} {
	m := make(map[string]interface{})
	m["Service"] = srv
	m["Failure"] = f
	return m
}

func ReplaceVars(input string, s *services.Service, f *failures.Failure) string {
	input = utils.ReplaceTemplate(input, s)
	if f != nil {
		input = utils.ReplaceTemplate(input, f)
	}
	return input
}

var exampleService = &services.Service{
	Id:                  1,
	Name:                "Statping",
	Domain:              "https://statping.com",
	Expected:            null.NewNullString("a better response"),
	ExpectedStatus:      200,
	Interval:            60,
	Type:                "http",
	Method:              "get",
	Timeout:             10,
	Order:               2,
	VerifySSL:           null.NewNullBool(true),
	Public:              null.NewNullBool(true),
	GroupId:             0,
	Permalink:           null.NewNullString("statping"),
	Online:              true,
	Latency:             324399,
	PingTime:            18399,
	Online24Hours:       99.2,
	Online7Days:         99.8,
	AvgResponse:         300233,
	FailuresLast24Hours: 4,
	Checkpoint:          utils.Now().Add(-10 * time.Minute),
	SleepDuration:       55,
	LastResponse:        "returning from a response",
	AllowNotifications:  null.NewNullBool(true),
	UserNotified:        false,
	UpdateNotify:        null.NewNullBool(true),
	SuccessNotified:     false,
	LastStatusCode:      200,
	LastLookupTime:      5233,
	LastLatency:         270233,
	LastCheck:           utils.Now().Add(-15 * time.Second),
	LastOnline:          utils.Now().Add(-15 * time.Second),
	LastOffline:         utils.Now().Add(-10 * time.Minute),
	SecondsOnline:       4500,
	SecondsOffline:      300,
}

var exampleFailure = &failures.Failure{
	Id:        1,
	Issue:     "HTTP returned a 500 status code",
	ErrorCode: 500,
	Service:   1,
	PingTime:  43203,
	CreatedAt: utils.Now().Add(-10 * time.Minute),
}
