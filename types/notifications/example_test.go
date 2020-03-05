// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package notifications

import (
	"errors"
	"fmt"
	"github.com/hunterlong/statping/source"
	"github.com/hunterlong/statping/types/failures"
	"github.com/hunterlong/statping/types/services"
	"github.com/hunterlong/statping/types/users"
	"github.com/hunterlong/statping/utils"
	"time"
)

// ExampleNotifier is an example on how to use the Statping notifier struct
type ExampleNotifier struct {
	*Notification
}

// example is a example variable for a example notifier
var example = &ExampleNotifier{&Notification{
	Method:      METHOD,
	Host:        "http://exmaplehost.com",
	Title:       "Example",
	Description: "Example Notifier",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(3 * time.Second),
	Limits:      7,
	Form: []NotificationForm{{
		Type:        "text",
		Title:       "Host",
		Placeholder: "Insert your Host here.",
		DbField:     "host",
		SmallText:   "this is where you would put the host",
	}, {
		Type:        "text",
		Title:       "Username",
		Placeholder: "Insert your Username here.",
		DbField:     "username",
	}, {
		Type:        "password",
		Title:       "Password",
		Placeholder: "Insert your Password here.",
		DbField:     "password",
	}, {
		Type:        "number",
		Title:       "Port",
		Placeholder: "Insert your Port here.",
		DbField:     "port",
	}, {
		Type:        "text",
		Title:       "API Key",
		Placeholder: "Insert your API Key here",
		DbField:     "api_key",
	}, {
		Type:        "text",
		Title:       "API Secret",
		Placeholder: "Insert your API Secret here",
		DbField:     "api_secret",
	}, {
		Type:        "text",
		Title:       "Var 1",
		Placeholder: "Insert your Var1 here",
		DbField:     "var1",
	}, {
		Type:        "text",
		Title:       "Var2",
		Placeholder: "Var2 goes here",
		DbField:     "var2",
	}},
}}

// init will be ran when Statping is loaded, AddNotifier will add the notifier instance to the system
func init() {
	dir = utils.Directory
	source.Assets()
	utils.InitLogs()
	injectDatabase()
}

// Send is the main function to hold your notifier functionality
func (n *ExampleNotifier) Send(msg interface{}) error {
	message := msg.(string)
	fmt.Printf("i received this string: %v\n", message)
	return nil
}

// Select is a required basic event for the Notifier interface
func (n *ExampleNotifier) Select() *Notification {
	return n.Notification
}

// OnSave is a required basic event for the Notifier interface
func (n *ExampleNotifier) OnSave() error {
	msg := fmt.Sprintf("received on save trigger")
	n.AddQueue("onsave", msg)
	return errors.New("onsave triggered")
}

// OnSuccess is a required basic event for the Notifier interface
func (n *ExampleNotifier) OnSuccess(s *services.Service) {
	msg := fmt.Sprintf("received a count trigger for service: %v\n", s.Name)
	n.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnFailure is a required basic event for the Notifier interface
func (n *ExampleNotifier) OnFailure(s *services.Service, f *failures.Failure) {
	msg := fmt.Sprintf("received a failure trigger for service: %v\n", s.Name)
	n.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnTest is a option testing event for the Notifier interface
func (n *ExampleNotifier) OnTest() error {
	fmt.Printf("received a test trigger with form data: %v\n", n.Host)
	return nil
}

// OnNewService is a option event for new services
func (n *ExampleNotifier) OnNewService(s *services.Service) {
	msg := fmt.Sprintf("received a new service trigger for service: %v\n", s.Name)
	n.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnUpdatedService is a option event for updated services
func (n *ExampleNotifier) OnUpdatedService(s *services.Service) {
	msg := fmt.Sprintf("received a update service trigger for service: %v\n", s.Name)
	n.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnDeletedService is a option event for deleted services
func (n *ExampleNotifier) OnDeletedService(s *services.Service) {
	msg := fmt.Sprintf("received a delete service trigger for service: %v\n", s.Name)
	n.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnNewUser is a option event for new users
func (n *ExampleNotifier) OnNewUser(s *users.User) {
	msg := fmt.Sprintf("received a new user trigger for user: %v\n", s.Username)
	n.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnUpdatedUser is a option event for updated users
func (n *ExampleNotifier) OnUpdatedUser(s *users.User) {
	msg := fmt.Sprintf("received a updated user trigger for user: %v\n", s.Username)
	n.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnDeletedUser is a option event for deleted users
func (n *ExampleNotifier) OnDeletedUser(s *users.User) {
	msg := fmt.Sprintf("received a deleted user trigger for user: %v\n", s.Username)
	n.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnNewNotifier is triggered when a new notifier has initialized
func (n *ExampleNotifier) OnNewNotifier(s *Notification) {
	msg := fmt.Sprintf("received a new notifier trigger for notifier: %v\n", s.Method)
	n.AddQueue(fmt.Sprintf("notifier_%v", s.Id), msg)
}

// OnUpdatedNotifier is triggered when a notifier has been updated
func (n *ExampleNotifier) OnUpdatedNotifier(s *Notification) {
	msg := fmt.Sprintf("received a update notifier trigger for notifier: %v\n", s.Method)
	n.AddQueue(fmt.Sprintf("notifier_%v", s.Id), msg)
}
