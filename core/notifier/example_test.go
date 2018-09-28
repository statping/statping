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

package notifier

import (
	"errors"
	"fmt"
	"github.com/hunterlong/statup/types"
	"time"
)

type ExampleNotifier struct {
	*Notification
}

var example = &ExampleNotifier{&Notification{
	Method:      METHOD,
	Host:        "http://exmaplehost.com",
	Title:       "Example",
	Description: "Example Notifier",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(200 * time.Millisecond),
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

// REQUIRED init() will install/load the notifier
func init() {
	AddNotifier(example)
}

// REQUIRED
func (n *ExampleNotifier) Send(msg interface{}) error {
	message := msg.(string)
	fmt.Printf("i received this string: %v\n", message)
	return nil
}

// REQUIRED
func (n *ExampleNotifier) Select() *Notification {
	return n.Notification
}

// REQUIRED
func (n *ExampleNotifier) OnSave() error {
	msg := fmt.Sprintf("received on save trigger")
	n.AddQueue(msg)
	return errors.New("onsave triggered")
}

// REQUIRED - BASIC EVENT
func (n *ExampleNotifier) OnSuccess(s *types.Service) {
	msg := fmt.Sprintf("received a count trigger for service: %v\n", s.Name)
	n.AddQueue(msg)
}

// REQUIRED - BASIC EVENT
func (n *ExampleNotifier) OnFailure(s *types.Service, f *types.Failure) {
	msg := fmt.Sprintf("received a failure trigger for service: %v\n", s.Name)
	n.AddQueue(msg)
}

// OPTIONAL Test function before user saves
func (n *ExampleNotifier) OnTest() error {
	fmt.Printf("received a test trigger with form data: %v\n", n.Host)
	return nil
}

// OPTIONAL
func (n *ExampleNotifier) OnNewService(s *types.Service) {
	msg := fmt.Sprintf("received a new service trigger for service: %v\n", s.Name)
	n.AddQueue(msg)
}

// OPTIONAL
func (n *ExampleNotifier) OnUpdatedService(s *types.Service) {
	msg := fmt.Sprintf("received a update service trigger for service: %v\n", s.Name)
	n.AddQueue(msg)
}

// OPTIONAL
func (n *ExampleNotifier) OnDeletedService(s *types.Service) {
	msg := fmt.Sprintf("received a delete service trigger for service: %v\n", s.Name)
	n.AddQueue(msg)
}

// OPTIONAL
func (n *ExampleNotifier) OnNewUser(s *types.User) {
	msg := fmt.Sprintf("received a new user trigger for user: %v\n", s.Username)
	n.AddQueue(msg)
}

// OPTIONAL
func (n *ExampleNotifier) OnUpdatedUser(s *types.User) {
	msg := fmt.Sprintf("received a updated user trigger for user: %v\n", s.Username)
	n.AddQueue(msg)
}

// OPTIONAL
func (n *ExampleNotifier) OnDeletedUser(s *types.User) {
	msg := fmt.Sprintf("received a deleted user trigger for user: %v\n", s.Username)
	n.AddQueue(msg)
}

// OPTIONAL
func (n *ExampleNotifier) OnUpdatedCore(s *types.Core) {
	msg := fmt.Sprintf("received a updated core trigger for core: %v\n", s.Name)
	n.AddQueue(msg)
}

// OPTIONAL
func (n *ExampleNotifier) OnStart(s *types.Core) {
	msg := fmt.Sprintf("received a trigger on Statup boot: %v\n", s.Name)
	n.AddQueue(msg)
}

// OPTIONAL
func (n *ExampleNotifier) OnNewNotifier(s *Notification) {
	msg := fmt.Sprintf("received a new notifier trigger for notifier: %v\n", s.Method)
	n.AddQueue(msg)
}

// OPTIONAL
func (n *ExampleNotifier) OnUpdatedNotifier(s *Notification) {
	msg := fmt.Sprintf("received a update notifier trigger for notifier: %v\n", s.Method)
	n.AddQueue(msg)
}

// Create a new notifier that includes a form for the end user to insert their own values
func Example() {
	// Create a new variable for your Notifier
	example = &ExampleNotifier{&Notification{
		Method:      "Example",
		Title:       "Example Notifier",
		Description: "Example Notifier can hold many different types of fields for a customized look.",
		Author:      "Hunter Long",
		AuthorUrl:   "https://github.com/hunterlong",
		Delay:       time.Duration(1500 * time.Millisecond),
		Limits:      7,
		Form: []NotificationForm{{
			Type:        "text",
			Title:       "Host",
			Placeholder: "Insert your Host here.",
			DbField:     "host",
			SmallText:   "you can also use SmallText to insert some helpful hints under this input",
		}},
	}}

	// AddNotifier accepts a notifier to load into the Statup Notification system
	AddNotifier(example)
}

// Add any type of interface to the AddQueue function when a service is successful
func Example_onSuccess() {
	msg := fmt.Sprintf("this is a successful message as a string passing into AddQueue function")
	example.AddQueue(msg)
}

// Add any type of interface to the AddQueue function when a service is successful
func Example_onFailure() {
	msg := fmt.Sprintf("this is a failing message as a string passing into AddQueue function")
	example.AddQueue(msg)
}

// The Send method will run the main functionality of your notifier
func Example_send() {
	// example.Send(msg interface{})
	for i := 0; i <= 10; i++ {
		fmt.Printf("do something awesome rather than a loop %v\n", i)
	}
}
