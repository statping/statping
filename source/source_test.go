package source

import (
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	dir           string
	requiredFiles = []string{
		"css/style.css",
		"css/style.css.gz",
		"css/main.css",
		"scss/main.scss",
		"scss/base.scss",
		"scss/forms.scss",
		"scss/layout.scss",
		"scss/mixin.scss",
		"scss/mobile.scss",
		"scss/variables.scss",
		"js/bundle.js",
		"js/main.chunk.js",
		"js/polyfill.chunk.js",
		"js/style.chunk.js",
		"banner.png",
		"favicon.ico",
		"robots.txt",
	}
)

func init() {
	dir = utils.Directory
	utils.InitLogs()
	Assets()
	utils.DeleteDirectory(dir + "/assets")
	dir = utils.Params.GetString("STATPING_DIR")
}

func assetFiles(t *testing.T, exist bool) {
	for _, f := range requiredFiles {
		if exist {
			assert.FileExists(t, dir+"/assets/"+f)
		} else {
			assert.NoFileExists(t, dir+"/assets/"+f)
		}
	}
}

func TestCore_UsingAssets(t *testing.T) {
	assert.False(t, UsingAssets(dir))
	assetFiles(t, false)
}

func TestCreateAssets(t *testing.T) {
	assert.Nil(t, CreateAllAssets(dir))
	assert.True(t, UsingAssets(dir))
	assert.Nil(t, CompileSASS())
	assetFiles(t, true)
}

func TestCopyAllToPublic(t *testing.T) {
	err := CopyAllToPublic(TmplBox)
	require.Nil(t, err)
	assetFiles(t, true)
}

func TestCompileSASS(t *testing.T) {
	err := CompileSASS()
	require.Nil(t, err)
	assert.True(t, UsingAssets(dir))
	assetFiles(t, true)
}

func TestSaveAndCompileAsset(t *testing.T) {
	scssData := "$bodycolor: #333; BODY { color: $bodycolor; }"

	err := SaveAsset([]byte(scssData), "scss/base.scss")
	require.Nil(t, err)
	assert.FileExists(t, dir+"/assets/scss/base.scss")

	asset := OpenAsset("scss/base.scss")
	assert.NotEmpty(t, asset)
	assert.Equal(t, scssData, asset)

	err = CompileSASS()
	require.Nil(t, err)
	assert.FileExists(t, dir+"/assets/css/main.css")

	themeCSS, err := utils.OpenFile(dir + "/assets/css/main.css")
	require.Nil(t, err)

	assert.Contains(t, themeCSS, `color: #333;`)
}

func TestOpenAsset(t *testing.T) {
	for _, f := range requiredFiles {
		asset := OpenAsset(f)
		assert.NotEmpty(t, asset)
	}
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
