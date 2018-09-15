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

type Example struct {
	*Notification
}

var example = &Example{&Notification{
	Method:      METHOD,
	Host:        "http://exmaplehost.com",
	Title:       "Example",
	Description: "Example Notifier",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(5 * time.Second),
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
func (n *Example) Send(msg interface{}) error {
	message := msg.(string)
	fmt.Printf("i received this string: %v\n", message)
	return nil
}

// REQUIRED
func (n *Example) Select() *Notification {
	return n.Notification
}

// REQUIRED
func (n *Example) OnSave() error {
	msg := fmt.Sprintf("received on save trigger")
	n.AddQueue(msg)
	return errors.New("onsave triggered")
}

// REQUIRED
func (n *Example) Test() error {
	msg := fmt.Sprintf("received a test trigger\n")
	n.AddQueue(msg)
	return errors.New("test triggered")
}

// REQUIRED - BASIC EVENT
func (n *Example) OnSuccess(s *types.Service) {
	msg := fmt.Sprintf("received a count trigger for service: %v\n", s.Name)
	n.AddQueue(msg)
}

// REQUIRED - BASIC EVENT
func (n *Example) OnFailure(s *types.Service, f *types.Failure) {
	msg := fmt.Sprintf("received a failure trigger for service: %v\n", s.Name)
	n.AddQueue(msg)
}

// OPTIONAL
func (n *Example) OnNewService(s *types.Service) {
	msg := fmt.Sprintf("received a new service trigger for service: %v\n", s.Name)
	n.AddQueue(msg)
}

// OPTIONAL
func (n *Example) OnUpdatedService(s *types.Service) {
	msg := fmt.Sprintf("received a update service trigger for service: %v\n", s.Name)
	n.AddQueue(msg)
}

// OPTIONAL
func (n *Example) OnDeletedService(s *types.Service) {
	msg := fmt.Sprintf("received a delete service trigger for service: %v\n", s.Name)
	n.AddQueue(msg)
}

// OPTIONAL
func (n *Example) OnNewUser(s *types.User) {
	msg := fmt.Sprintf("received a new user trigger for user: %v\n", s.Username)
	n.AddQueue(msg)
}

// OPTIONAL
func (n *Example) OnUpdatedUser(s *types.User) {
	msg := fmt.Sprintf("received a updated user trigger for user: %v\n", s.Username)
	n.AddQueue(msg)
}

// OPTIONAL
func (n *Example) OnDeletedUser(s *types.User) {
	msg := fmt.Sprintf("received a deleted user trigger for user: %v\n", s.Username)
	n.AddQueue(msg)
}

// OPTIONAL
func (n *Example) OnUpdatedCore(s *types.Core) {
	msg := fmt.Sprintf("received a updated core trigger for core: %v\n", s.Name)
	n.AddQueue(msg)
}

// OPTIONAL
func (n *Example) OnNewNotifier(s *Notification) {
	msg := fmt.Sprintf("received a new notifier trigger for notifier: %v\n", s.Method)
	n.AddQueue(msg)
}

// OPTIONAL
func (n *Example) OnUpdatedNotifier(s *Notification) {
	msg := fmt.Sprintf("received a update notifier trigger for notifier: %v\n", s.Method)
	n.AddQueue(msg)
}
