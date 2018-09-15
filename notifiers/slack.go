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
	"bytes"
	"fmt"
	"github.com/hunterlong/statup/core/notifier"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"text/template"
	"time"
)

const (
	SLACK_ID         = 2
	SLACK_METHOD     = "slack"
	FAILING_TEMPLATE = `{ "attachments": [ { "fallback": "Service {{.Service.Name}} - is currently failing", "text": "<{{.Service.Domain}}|{{.Service.Name}}> - Your Statup service '{{.Service.Name}}' has just received a Failure notification with a HTTP Status code of {{.Service.LastStatusCode}}.", "fields": [ { "title": "Expected", "value": "{{.Service.Expected}}", "short": true }, { "title": "Status Code", "value": "{{.Service.LastStatusCode}}", "short": true } ], "color": "#FF0000", "thumb_url": "https://statup.io", "footer": "Statup", "footer_icon": "https://img.cjx.io/statuplogo32.png" } ] }`
	SUCCESS_TEMPLATE = `{ "attachments": [ { "fallback": "Service {{.Service.Name}} - is now back online", "text": "<{{.Service.Domain}}|{{.Service.Name}}> - Your Statup service '{{.Service.Name}}' has just received a Failure notification.", "fields": [ { "title": "Issue", "value": "Awesome Project", "short": true }, { "title": "Status Code", "value": "{{.Service.LastStatusCode}}", "short": true } ], "color": "#00FF00", "thumb_url": "https://statup.io", "footer": "Statup", "footer_icon": "https://img.cjx.io/statuplogo32.png" } ] }`
	TEST_TEMPLATE    = `{"text":"{{.}}"}`
)

type Slack struct {
	*notifier.Notification
}

var slacker = &Slack{&notifier.Notification{
	Method:      SLACK_METHOD,
	Title:       "Slack",
	Description: "Send notifications to your Slack channel when a service is offline. Insert your Incoming Webhook URL for your channel to receive notifications. Based on the <a href=\"https://api.slack.com/incoming-webhooks\">Slack API</a>.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(10 * time.Second),
	Host:        "https://webhooksurl.slack.com/***",
	CanTest:     true,
	Form: []notifier.NotificationForm{{
		Type:        "text",
		Title:       "Incoming Webhook Url",
		Placeholder: "Insert your Slack webhook URL here.",
		SmallText:   "Incoming Webhook URL from <a href=\"https://api.slack.com/apps\" target=\"_blank\">Slack Apps</a>",
		DbField:     "Host",
	}}},
}

func sendSlack(temp string, data interface{}) error {
	buf := new(bytes.Buffer)
	slackTemp, _ := template.New("slack").Parse(temp)
	err := slackTemp.Execute(buf, data)
	if err != nil {
		return err
	}
	slacker.AddQueue(buf.String())
	return nil
}

type slackMessage struct {
	Service  *types.Service
	Template string
	Time     int64
}

// DEFINE YOUR NOTIFICATION HERE.
func init() {
	err := notifier.AddNotifier(slacker)
	if err != nil {
		panic(err)
	}
}

func (u *Slack) Send(msg interface{}) error {
	message := msg.(string)
	client := new(http.Client)
	_, err := client.Post(u.Host, "application/json", bytes.NewBuffer([]byte(message)))
	if err != nil {
		return err
	}
	return nil
}

func (u *Slack) Select() *notifier.Notification {
	return u.Notification
}

func (u *Slack) OnTest(n notifier.Notification) (bool, error) {
	utils.Log(1, "Slack notifier loaded")
	msg := fmt.Sprintf("You're Statup Slack Notifier is working correctly!")
	err := sendSlack(TEST_TEMPLATE, msg)
	return true, err
}

// ON SERVICE FAILURE, DO YOUR OWN FUNCTIONS
func (u *Slack) OnFailure(s *types.Service, f *types.Failure) {
	message := slackMessage{
		Service:  s,
		Template: FAILURE,
		Time:     time.Now().Unix(),
	}
	sendSlack(FAILING_TEMPLATE, message)
}

// ON SERVICE SUCCESS, DO YOUR OWN FUNCTIONS
func (u *Slack) OnSuccess(s *types.Service) {

}

// ON SAVE OR UPDATE OF THE NOTIFIER FORM
func (u *Slack) OnSave() error {
	message := fmt.Sprintf("Notification %v is receiving updated information.", u.Method)
	u.AddQueue(message)
	return nil
}
