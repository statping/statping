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

package notifier

import (
	"errors"
	"fmt"
	"github.com/hunterlong/statping/source"
	"github.com/hunterlong/statping/types"
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
	AddNotifiers(example)
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
func (n *ExampleNotifier) OnSuccess(s *types.Service) {
	msg := fmt.Sprintf("received a count trigger for service: %v\n", s.Name)
	n.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnFailure is a required basic event for the Notifier interface
func (n *ExampleNotifier) OnFailure(s *types.Service, f *types.Failure) {
	msg := fmt.Sprintf("received a failure trigger for service: %v\n", s.Name)
	n.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnTest is a option testing event for the Notifier interface
func (n *ExampleNotifier) OnTest() error {
	fmt.Printf("received a test trigger with form data: %v\n", n.Host)
	return nil
}

// OnNewService is a option event for new services
func (n *ExampleNotifier) OnNewService(s *types.Service) {
	msg := fmt.Sprintf("received a new service trigger for service: %v\n", s.Name)
	n.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnUpdatedService is a option event for updated services
func (n *ExampleNotifier) OnUpdatedService(s *types.Service) {
	msg := fmt.Sprintf("received a update service trigger for service: %v\n", s.Name)
	n.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnDeletedService is a option event for deleted services
func (n *ExampleNotifier) OnDeletedService(s *types.Service) {
	msg := fmt.Sprintf("received a delete service trigger for service: %v\n", s.Name)
	n.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnNewUser is a option event for new users
func (n *ExampleNotifier) OnNewUser(s *types.User) {
	msg := fmt.Sprintf("received a new user trigger for user: %v\n", s.Username)
	n.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnUpdatedUser is a option event for updated users
func (n *ExampleNotifier) OnUpdatedUser(s *types.User) {
	msg := fmt.Sprintf("received a updated user trigger for user: %v\n", s.Username)
	n.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnDeletedUser is a option event for deleted users
func (n *ExampleNotifier) OnDeletedUser(s *types.User) {
	msg := fmt.Sprintf("received a deleted user trigger for user: %v\n", s.Username)
	n.AddQueue(fmt.Sprintf("service_%v", s.Id), msg)
}

// OnUpdatedCore is a option event when the settings are updated
func (n *ExampleNotifier) OnUpdatedCore(s *types.Core) {
	msg := fmt.Sprintf("received a updated core trigger for core: %v\n", s.Name)
	n.AddQueue("core", msg)
}

// OnStart is triggered when statup has been started
func (n *ExampleNotifier) OnStart(s *types.Core) {
	msg := fmt.Sprintf("received a trigger on Statping boot: %v\n", s.Name)
	n.AddQueue(fmt.Sprintf("core"), msg)
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

// Create a new notifier that includes a form for the end user to insert their own values
func ExampleNotification() {
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
		}, {
			Type:        "text",
			Title:       "API Key",
			Placeholder: "Include some type of API key here",
			DbField:     "api_key",
		}},
	}}

	// AddNotifier accepts a Notifier to load into the Statping Notification system
	err := AddNotifiers(example)
	fmt.Println(err)
	// Output: <nil>
}

// Add a Notifier to the AddQueue function to insert it into the system
func ExampleAddNotifier() {
	err := AddNotifiers(example)
	fmt.Println(err)
	// Output: <nil>
}

// OnSuccess will be triggered everytime a service is online
func ExampleNotification_OnSuccess() {
	msg := fmt.Sprintf("this is a successful message as a string passing into AddQueue function")
	example.AddQueue("example", msg)
	fmt.Println(len(example.Queue))
	// Output:
	// 1
}

// Add a new message into the queue OnSuccess
func ExampleOnSuccess() {
	msg := fmt.Sprintf("received a count trigger for service: %v\n", service.Name)
	example.AddQueue("example", msg)
}

// Add a new message into the queue OnFailure
func ExampleOnFailure() {
	msg := fmt.Sprintf("received a failing service: %v\n", service.Name)
	example.AddQueue("example", msg)
}

// OnTest allows your notifier to be testable
func ExampleOnTest() {
	err := example.OnTest()
	fmt.Print(err)
	// Output <nil>
}

// Implement the Test interface to give your notifier testing abilities
func ExampleNotification_CanTest() {
	testable := example.CanTest()
	fmt.Print(testable)
	// Output: true
}

// Add any type of interface to the AddQueue function to be ran in the queue
func ExampleNotification_AddQueue() {
	msg := fmt.Sprintf("this is a failing message as a string passing into AddQueue function")
	example.AddQueue("example", msg)
	queue := example.Queue
	fmt.Printf("Example has %v items in the queue", len(queue))
	// Output:
	// Example has 2 items in the queue
}

// The Send method will run the main functionality of your notifier
func ExampleNotification_Send() {
	msg := "this can be any type of interface"
	example.Send(msg)
	queue := example.Queue
	fmt.Printf("Example has %v items in the queue", len(queue))
	// Output:
	// i received this string: this can be any type of interface
	// Example has 2 items in the queue
}

// LastSent will return the time.Duration of the last sent message
func ExampleNotification_LastSent() {
	last := example.LastSent()
	fmt.Printf("Last message was sent %v seconds ago", last.Seconds())
	// Output: Last message was sent 0 seconds ago
}

// Logs will return a slice of previously sent items from your notifier
func ExampleNotification_Logs() {
	logs := example.Logs()
	fmt.Printf("Example has %v items in the log", len(logs))
	// Output: Example has 0 items in the log
}

// SentLastMinute will return he amount of notifications sent in last 1 minute
func ExampleNotification_SentLastMinute() {
	lastMinute := example.SentLastMinute()
	fmt.Printf("%v notifications sent in the last minute", lastMinute)
	// Output: 0 notifications sent in the last minute
}

// SentLastHour will return he amount of notifications sent in last 1 hour
func ExampleNotification_SentLastHour() {
	lastHour := example.SentLastHour()
	fmt.Printf("%v notifications sent in the last hour", lastHour)
	// Output: 0 notifications sent in the last hour
}

// SentLastHour will return he amount of notifications sent in last 1 hour
func ExampleNotification_WithinLimits() {
	ok, err := example.WithinLimits()
	if err != nil {
		panic(err)
	}
	if ok {
		fmt.Printf("Example notifier is still within its sending limits")
	}
	// Output: Example notifier is still within its sending limits
}
