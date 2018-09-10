// +build test

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

package example

import (
	"fmt"
	"github.com/hunterlong/statup/notifiers"
	"github.com/hunterlong/statup/types"
	"sync"
)

var (
	exampler      *Example
	slackMessages []string
	messageLock   *sync.Mutex
)

type Example struct {
	*notifiers.Notification
}

// DEFINE YOUR NOTIFICATION HERE.
func init() {
	exampler = &Example{&notifiers.Notification{
		Id:     99999,
		Method: "slack",
		Host:   "https://webhooksurl.slack.com/***",
		Form: []notifiers.NotificationForm{{
			Type:        "text",
			Title:       "Incoming Webhook Url",
			Placeholder: "Insert your Slack webhook URL here.",
			DbField:     "Host",
		}}},
	}
	notifiers.AddNotifier(exampler)
	messageLock = new(sync.Mutex)
}

// Select Obj
func (u *Example) Select() *notifiers.Notification {
	return u.Notification
}

// WHEN NOTIFIER LOADS
func (u *Example) Init() error {
	err := u.Install()
	if err == nil {
		notifier, _ := notifiers.SelectNotification(u.Id)
		forms := u.Form
		u.Notification = notifier
		u.Form = forms
		if u.Enabled {
			go u.Run()
		}
	}

	return err
}

func (u *Example) Test() error {
	fmt.Println("Example notifier has been Tested!")
	return nil
}

// AFTER NOTIFIER LOADS, IF ENABLED, START A QUEUE PROCESS
func (u *Example) Run() error {
	if u.Enabled {
		u.Run()
	}
	return nil
}

// CUSTOM FUNCTION FO SENDING SLACK MESSAGES
func SendSlack(temp string, data interface{}) error {

	return nil
}

// ON SERVICE FAILURE, DO YOUR OWN FUNCTIONS
func (u *Example) OnFailure(s *types.Service) error {
	if u.Enabled {
		fmt.Println("Example notifier received a failing service event!")
	}
	return nil
}

// ON SERVICE SUCCESS, DO YOUR OWN FUNCTIONS
func (u *Example) OnSuccess(s *types.Service) error {
	if u.Enabled {
		fmt.Println("Example notifier received a successful service event!")
	}
	return nil
}

// ON SAVE OR UPDATE OF THE NOTIFIER FORM
func (u *Example) OnSave() error {
	fmt.Println("Example notifier was saved!")
	return nil
}

// ON SERVICE FAILURE, DO YOUR OWN FUNCTIONS
func (u *Example) Install() error {
	inDb := exampler.Notification.IsInDatabase()
	if !inDb {
		newNotifer, err := notifiers.InsertDatabase(u.Notification)
		if err != nil {
			return err
		}
		fmt.Println("Example notifier was installed!", newNotifer)
	}
	return nil
}
