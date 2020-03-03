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
	"github.com/stretchr/testify/require"
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
	assert.Nil(t, CompileSASS(DefaultScss...))
	assert.FileExists(t, dir+"/assets/css/main.css")
	assert.FileExists(t, dir+"/assets/css/style.css")
	assert.FileExists(t, dir+"/assets/css/vendor.css")
	assert.FileExists(t, dir+"/assets/scss/base.scss")
	assert.FileExists(t, dir+"/assets/scss/mobile.scss")
	assert.FileExists(t, dir+"/assets/scss/variables.scss")
}

//func TestCopyAllToPublic(t *testing.T) {
//	err := CopyAllToPublic(TmplBox)
//	require.Nil(t, err)
//}

func TestCompileSASS(t *testing.T) {
	err := CompileSASS(DefaultScss...)
	require.Nil(t, err)
	assert.True(t, UsingAssets(dir))
}

func TestSaveAndCompileAsset(t *testing.T) {
	scssData := "$bodycolor: #333; BODY { color: $bodycolor; }"

	err := SaveAsset([]byte(scssData), "scss/theme.scss")
	require.Nil(t, err)
	assert.FileExists(t, dir+"/assets/scss/theme.scss")

	asset := OpenAsset("scss/theme.scss")
	assert.NotEmpty(t, asset)
	assert.Equal(t, scssData, asset)

	err = CompileSASS("scss/theme.scss")
	require.Nil(t, err)
	assert.FileExists(t, dir+"/assets/css/theme.css")

	themeCSS, err := utils.OpenFile(dir + "/assets/css/theme.css")
	require.Nil(t, err)

	assert.Contains(t, themeCSS, `color: #333;`)
}

func TestOpenAsset(t *testing.T) {
	asset := OpenAsset("scss/theme.scss")
	assert.NotEmpty(t, asset)
}

func TestDeleteAssets(t *testing.T) {
	assert.True(t, UsingAssets(dir))
	assert.Nil(t, DeleteAllAssets(dir))
	assert.False(t, UsingAssets(dir))
}

func ExampleSaveAsset() {
	data := []byte("alert('helloooo')")
	SaveAsset(data, "js/test.js")
}

func ExampleOpenAsset() {
	OpenAsset("js/main.js")
}
