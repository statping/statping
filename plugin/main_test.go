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

package plugin

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	database *gorm.DB
	example  *PluginObject
)

func (p *PluginObject) StatupDatabase(db *gorm.DB) {
	database = db
}

func (p *PluginObject) Select() *PluginObject {
	return p
}

func setupHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func TestAdd(t *testing.T) {
	err := Add(example)
	assert.NotNil(t, err)
}

func TestSelect(t *testing.T) {
	err := example.Select()
	assert.Nil(t, err)
}

func TestAddRoute(t *testing.T) {
	example.AddRoute("/plugin_example", "GET", setupHandler)
}
