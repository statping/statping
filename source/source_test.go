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

package source

import (
	"github.com/hunterlong/statping/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	dir string
)

func init() {
	dir = utils.Directory
	utils.InitLogs()
	Assets()
	utils.DeleteDirectory(dir + "/assets")
}

func TestCore_UsingAssets(t *testing.T) {
	assert.False(t, UsingAssets(dir))
}

func TestCreateAssets(t *testing.T) {
	assert.Nil(t, CreateAllAssets(dir))
	assert.True(t, UsingAssets(dir))
	assert.FileExists(t, dir+"/assets/css/base.css")
	assert.FileExists(t, dir+"/assets/scss/base.scss")
}

func TestCompileSASS(t *testing.T) {
	CompileSASS(dir)
	assert.True(t, UsingAssets(dir))
}

func TestSaveAsset(t *testing.T) {
	data := []byte("BODY { color: black; }")
	asset := SaveAsset(data, dir, "scss/theme.scss")
	assert.Nil(t, asset)
	assert.FileExists(t, dir+"/assets/scss/theme.scss")
}

func TestOpenAsset(t *testing.T) {
	asset := OpenAsset(dir, "scss/theme.scss")
	assert.NotEmpty(t, asset)
}

func TestDeleteAssets(t *testing.T) {
	assert.Nil(t, DeleteAllAssets(dir))
	assert.False(t, UsingAssets(dir))
}

func TestCopyToPluginFailed(t *testing.T) {
	assert.Nil(t, DeleteAllAssets(dir))
	assert.False(t, UsingAssets(dir))
}

func ExampleSaveAsset() {
	data := []byte("alert('helloooo')")
	SaveAsset(data, "js", "test.js")
}

func ExampleOpenAsset() {
	OpenAsset("js", "main.js")
}
