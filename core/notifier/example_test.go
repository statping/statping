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
	"fmt"
	"github.com/hunterlong/statup/types"
)

type Example struct {
	*Notification
}

const (
	EXAMPLE_METHOD = "example"
)

var example = &Example{&Notification{
	Method: EXAMPLE_METHOD,
	Host:   "http://exmaplehost.com",
	Form: []NotificationForm{{
		Type:        "text",
		Title:       "Host",
		Placeholder: "Insert your Host here.",
		DbField:     "host",
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
	}}},
}

// REQUIRED init() will install/load the notifier
func init() {
	AddNotifier(example)
}

// REQUIRED
func (n *Example) Run() error {
	return nil
}

// REQUIRED
func (n *Example) OnSave() error {
	return nil
}

// REQUIRED
func (n *Example) Test() error {
	return nil
}

// REQUIRED
func (n *Example) Select() *Notification {
	return n.Notification
}

// REQUIRED - BASIC EVENT
func (n *Example) OnSuccess(s *types.Service) {
	saySomething("service is is online!")
}

// REQUIRED - BASIC EVENT
func (n *Example) OnFailure(s *types.Service, f *types.Failure) {
	saySomething("service is failing!")
}

// Example function to do something awesome or not...
func saySomething(text ...interface{}) {
	fmt.Println(text)
}

// OPTIONAL
func (n *Example) OnNewService(s *types.Service) {

}

// OPTIONAL
func (n *Example) OnUpdatedService(s *types.Service) {

}

// OPTIONAL
func (n *Example) OnDeletedService(s *types.Service) {

}

// OPTIONAL
func (n *Example) OnNewUser(s *types.User) {

}

// OPTIONAL
func (n *Example) OnUpdatedUser(s *types.User) {

}

// OPTIONAL
func (n *Example) OnDeletedUser(s *types.User) {

}

// OPTIONAL
func (n *Example) OnUpdatedCore(s *types.Core) {

}

// OPTIONAL
func (n *Example) OnNewNotifier(s *Notification) {

}

// OPTIONAL
func (n *Example) OnUpdatedNotifier(s *Notification) {

}
