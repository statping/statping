package source

import (
	"github.com/hunterlong/statup/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	dir string
)

func init() {
	dir = utils.Directory
	utils.InitLogs()
	Assets()
	os.RemoveAll(dir + "/cmd/assets")
}

func TestCore_UsingAssets(t *testing.T) {
	assert.False(t, UsingAssets)
}

func TestCreateAssets(t *testing.T) {
	assert.Nil(t, CreateAllAssets(dir))
	assert.True(t, HasAssets(dir))
	assert.FileExists(t, "../assets/css/base.css")
	assert.FileExists(t, "../assets/scss/base.scss")
}

func TestCompileSASS(t *testing.T) {
	assert.Nil(t, CompileSASS(dir))
	assert.True(t, HasAssets(dir))
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
	assert.False(t, HasAssets(dir))
}

func TestCopyToPluginFailed(t *testing.T) {

	assert.Nil(t, DeleteAllAssets(dir))
	assert.False(t, HasAssets(dir))
}
