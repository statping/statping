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
	"fmt"
	"strings"
)

var (
	allowedVars = []string{"host", "username", "password", "port", "api_key", "api_secret", "var1", "var2"}
)

func checkNotifierForm(n Notifier) error {
	notifier := n.Select()
	for _, f := range notifier.Form {
		contains := contains(f.DbField, allowedVars)
		if !contains {
			return fmt.Errorf("the DbField '%v' is not allowed, allowed vars: %v", f.DbField, allowedVars)
		}
	}
	return nil
}

func contains(s string, arr []string) bool {
	for _, v := range arr {
		if strings.ToLower(s) == v {
			return true
		}
	}
	return false
}
