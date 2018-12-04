// Package notifier contains the main functionality for the Statping Notification system
//
// Example Notifier
//
// Below is an example of a Notifier with multiple Form values to custom your inputs. Place your notifier go file
// into the /notifiers/ directory and follow the example below.
//
//		type ExampleNotifier struct {
//			*Notification
//		}
//
//		var example = &ExampleNotifier{&Notification{
//			Method:      "example",
//			Title:       "Example Notifier",
//			Description: "This is an example of a notifier for Statping!",
//			Author:      "Hunter Long",
//			AuthorUrl:   "https://github.com/hunterlong",
//			Delay:       time.Duration(3 * time.Second),
//			Limits:      7,
//			Form: []NotificationForm{{
//				Type:        "text",
//				Title:       "Host",
//				Placeholder: "Insert your Host here.",
//				DbField:     "host",
//				SmallText:   "this is where you would put the host",
//			}, {
//				Type:        "text",
//				Title:       "Username",
//				Placeholder: "Insert your Username here.",
//				DbField:     "username",
//			}, {
//				Type:        "password",
//				Title:       "Password",
//				Placeholder: "Insert your Password here.",
//				DbField:     "password",
//			}, {
//				Type:        "number",
//				Title:       "Port",
//				Placeholder: "Insert your Port here.",
//				DbField:     "port",
//			}, {
//				Type:        "text",
//				Title:       "API Key",
//				Placeholder: "Insert your API Key here",
//				DbField:     "api_key",
//			}, {
//				Type:        "text",
//				Title:       "API Secret",
//				Placeholder: "Insert your API Secret here",
//				DbField:     "api_secret",
//			}, {
//				Type:        "text",
//				Title:       "Var 1",
//				Placeholder: "Insert your Var1 here",
//				DbField:     "var1",
//			}, {
//				Type:        "text",
//				Title:       "Var2",
//				Placeholder: "Var2 goes here",
//				DbField:     "var2",
//			}},
//		}}
//
// Load the Notifier
//
// Include the init() function with AddNotifier and your notification struct. This is ran on start of Statping
// and will automatically create a new row in the database so the end user can save their own values.
//
//		func init() {
//			AddNotifier(example)
//		}
//
// Required Methods for Notifier Interface
//
// Below are the required methods to have your notifier implement the Notifier interface. The Send method
// will be where you would include the logic for your notification.
//
// 		// REQUIRED
//		func (n *ExampleNotifier) Send(msg interface{}) error {
//			message := msg.(string)
//			fmt.Printf("i received this string: %v\n", message)
//			return nil
//		}
//
//		// REQUIRED
//		func (n *ExampleNotifier) Select() *Notification {
//			return n.Notification
//		}
//
//		// REQUIRED
//		func (n *ExampleNotifier) OnSave() error {
//			msg := fmt.Sprintf("received on save trigger")
//			n.AddQueue(msg)
//			return errors.New("onsave triggered")
//		}
//
// Basic Events for Notifier
//
// You must include OnSuccess and OnFailure methods for your notifier. Anytime a service is online or offline
// these methods will be ran with the service corresponding to it.
//
// 		// REQUIRED - BASIC EVENT
//		func (n *ExampleNotifier) OnSuccess(s *types.Service) {
//			msg := fmt.Sprintf("received a count trigger for service: %v\n", s.Name)
//			n.AddQueue(msg)
//		}
//
//		// REQUIRED - BASIC EVENT
//		func (n *ExampleNotifier) OnFailure(s *types.Service, f *types.Failure) {
//			msg := fmt.Sprintf("received a failure trigger for service: %v\n", s.Name)
//			n.AddQueue(msg)
//		}
//
// Additional Events
//
// You can implement your notifier to different types of events that are triggered. Checkout the wiki to
// see more details and examples of how to build your own notifier.
//
// More info on: https://github.com/hunterlong/statping/wiki/Notifiers
package notifier
