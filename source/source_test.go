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

func assertFiles(t *testing.T, exist bool) {
	for _, f := range RequiredFiles {
		if exist {
			assert.FileExists(t, dir+"/assets/"+f)
		} else {
			assert.NoFileExists(t, dir+"/assets/"+f)
		}
	}
}

func TestCore_UsingAssets(t *testing.T) {
	assert.False(t, UsingAssets(dir))
	assertFiles(t, false)
}

func TestCreateAssets(t *testing.T) {
	assert.Nil(t, CreateAllAssets(dir))
	assert.True(t, UsingAssets(dir))
	assert.Nil(t, CompileSASS())
	assertFiles(t, true)
}

func TestCopyAllToPublic(t *testing.T) {
	err := CopyAllToPublic(TmplBox)
	require.Nil(t, err)
	assertFiles(t, true)
}

func TestCompileSASS(t *testing.T) {
	err := CompileSASS()
	require.Nil(t, err)
	assert.True(t, UsingAssets(dir))
	assertFiles(t, true)
}

func TestSaveAndCompileAsset(t *testing.T) {
	vars := OpenAsset("scss/variables.scss")
	vars += "$testingcolor: #b1b2b3;"

	err := SaveAsset([]byte(vars), "scss/variables.scss")
	require.Nil(t, err)
	assert.FileExists(t, dir+"/assets/scss/variables.scss")

	scssData := OpenAsset("scss/base.scss")
	scssData += "BODY { color: $testingcolor; }"
	err = SaveAsset([]byte(scssData), "scss/base.scss")
	require.Nil(t, err)
	assert.FileExists(t, dir+"/assets/scss/base.scss")

	asset := OpenAsset("scss/variables.scss")
	assert.NotEmpty(t, asset)
	assert.Equal(t, vars, asset)

	asset = OpenAsset("scss/base.scss")
	assert.NotEmpty(t, asset)
	assert.Equal(t, scssData, asset)

	err = CompileSASS()
	require.Nil(t, err)
	assertFiles(t, true)

	themeCSS, err := utils.OpenFile(dir + "/assets/css/index.css")
	require.Nil(t, err)

	assert.Contains(t, themeCSS, `color: #b1b2b3;`)
}

func TestOpenAsset(t *testing.T) {
	for _, f := range RequiredFiles {
		assert.FileExists(t, dir+"/assets/"+f)
		assert.NotEmpty(t, OpenAsset(f))
	}
}

func TestDeleteAssets(t *testing.T) {
	assert.True(t, UsingAssets(dir))
	assert.Nil(t, DeleteAllAssets(dir))
	assert.False(t, UsingAssets(dir))
	assertFiles(t, false)
}

func ExampleSaveAsset() {
	data := []byte("alert('helloooo')")
	SaveAsset(data, "js/test.js")
}

func ExampleOpenAsset() {
	OpenAsset("js/main.js")
}
