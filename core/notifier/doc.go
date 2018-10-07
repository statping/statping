// Package notifier contains the main functionality for the Statup Notification system
//
// Example Notifier
//
// Below is an example of a Notifier with multiple Form values to custom your inputs.
//
//		type ExampleNotifier struct {
//			*Notification
//		}
//
//		var example = &ExampleNotifier{&Notification{
//			Method:      METHOD,
//			Host:        "http://exmaplehost.com",
//			Title:       "Example",
//			Description: "Example Notifier",
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
// Loading the Notifier into the Statup Notification system with the following
//
//		func init() {
//			AddNotifier(example)
//		}
//
// More info on: https://github.com/hunterlong/statup/wiki/Notifiers
package notifier
