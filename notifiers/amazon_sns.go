package notifiers

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/statping/statping/types/null"

	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/notifier"
	"github.com/statping/statping/types/services"
)

var _ notifier.Notifier = (*amazonSNS)(nil)

type amazonSNS struct {
	*notifications.Notification
}

func (g *amazonSNS) Select() *notifications.Notification {
	return g.Notification
}

func (g *amazonSNS) Valid(values notifications.Values) error {
	return nil
}

var AmazonSNS = &amazonSNS{&notifications.Notification{
	Method:      "amazon_sns",
	Title:       "Amazon SNS",
	Description: "Use amazonSNS to receive push notifications. Add your amazonSNS URL and App Token to receive notifications.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Icon:        "fab fa-amazon",
	Delay:       5 * time.Second,
	Limits:      60,
	SuccessData: null.NewNullString(`{{.Service.Name}} is back online and was down for {{.Service.Downtime.Human}}`),
	FailureData: null.NewNullString(`{{.Service.Name}} is offline and has been down for {{.Service.Downtime.Human}}`),
	DataType:    "html",
	Form: []notifications.NotificationForm{{
		Type:        "text",
		Title:       "AWS Access Token",
		DbField:     "api_key",
		Placeholder: "AKPMED5XUXSEU3O5AB6M",
		Required:    true,
	}, {
		Type:        "text",
		Title:       "AWS Secret Key",
		DbField:     "api_secret",
		Placeholder: "39eAZODxEosHRgzLx173ttX9sCtJVOE8rzElRE9B",
		Required:    true,
	}, {
		Type:        "text",
		Title:       "Region",
		SmallText:   "Amazon Region for SNS",
		DbField:     "var1",
		Placeholder: "us-west-2",
		Required:    true,
	}, {
		Type:        "text",
		Title:       "SNS Topic ARN",
		SmallText:   "The ARN of the Topic",
		DbField:     "Host",
		Placeholder: "arn:aws:sns:us-west-2:123456789012:YourTopic",
		Required:    true,
	}}},
}

func valToAttr(val interface{}) *sns.MessageAttributeValue {
	dataType := "String"
	switch val.(type) {
	case string, bool:
		dataType = "String"
	case int, int64, uint, uint64, uint32:
		dataType = "Number"
	}
	return &sns.MessageAttributeValue{
		DataType:    aws.String(dataType),
		StringValue: aws.String(fmt.Sprintf("%v", val)),
	}
}

func messageAttributesSNS(s services.Service, f failures.Failure) map[string]*sns.MessageAttributeValue {
	attr := make(map[string]*sns.MessageAttributeValue)
	attr["service_id"] = valToAttr(s.Id)
	attr["online"] = valToAttr(s.Online)
	attr["downtime_milliseconds"] = valToAttr(s.Downtime().Milliseconds())
	if f.Id != 0 {
		attr["failure_issue"] = valToAttr(f.Issue)
		attr["failure_reason"] = valToAttr(f.Reason)
		attr["failure_status_code"] = valToAttr(f.ErrorCode)
		attr["failure_ping"] = valToAttr(f.PingTime)
	}
	return attr
}

// Send will send a HTTP Post to the amazonSNS API. It accepts type: string
func (g *amazonSNS) sendMessage(msg string, s services.Service, f failures.Failure) (string, error) {
	creds := credentials.NewStaticCredentials(g.ApiKey.String, g.ApiSecret.String, "")
	c := aws.NewConfig()
	c.Credentials = creds
	c.Region = aws.String(g.Var1.String)
	sess, err := session.NewSession(c)
	if err != nil {
		return "", err
	}

	client := sns.New(sess)
	input := &sns.PublishInput{
		Message:           aws.String(msg),
		TopicArn:          aws.String(g.Host.String),
		MessageAttributes: messageAttributesSNS(s, f),
	}

	result, err := client.Publish(input)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}

// OnFailure will trigger failing service
func (g *amazonSNS) OnFailure(s services.Service, f failures.Failure) (string, error) {
	msg := ReplaceVars(g.FailureData.String, s, f)
	return g.sendMessage(msg, s, f)
}

// OnSuccess will trigger successful service
func (g *amazonSNS) OnSuccess(s services.Service) (string, error) {
	msg := ReplaceVars(g.SuccessData.String, s, failures.Failure{})
	return g.sendMessage(msg, s, failures.Failure{})
}

// OnTest will test the amazonSNS notifier
func (g *amazonSNS) OnTest() (string, error) {
	s := services.Example(true)
	f := failures.Example()
	msg := ReplaceVars(`This is a test SNS notification from Statping. Service: {{.Service.Name}} - Downtime: {{.Service.Downtime.Human}}`, s, f)
	return g.sendMessage(msg, s, f)
}

// OnSave will trigger when this notifier is saved
func (g *amazonSNS) OnSave() (string, error) {
	return "", nil
}
