package source

import (
	"github.com/statping/statping/utils"
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
	dir = utils.Params.GetString("STATPING_DIR")
}

func TestCore_UsingAssets(t *testing.T) {
	assert.False(t, UsingAssets(dir))
}

func TestCreateAssets(t *testing.T) {
	CreateAllAssets(dir)
	assert.True(t, UsingAssets(dir))
	CompileSASS(DefaultScss...)
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
	CompileSASS(DefaultScss...)
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
