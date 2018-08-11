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
}

func TestCore_UsingAssets(t *testing.T) {
	assert.False(t, UsingAssets)
}

func TestHasAssets(t *testing.T) {
	assert.False(t, HasAssets(dir))
}

func TestCreateAssets(t *testing.T) {
	assert.Nil(t, CreateAllAssets(dir))
	assert.True(t, HasAssets(dir))
}

func TestCompileSASS(t *testing.T) {
	t.SkipNow()
	os.Setenv("SASS", "sass")
	os.Setenv("CMD_FILE", dir+"/cmd.sh")
	assert.Nil(t, CompileSASS(dir))
	assert.True(t, HasAssets(dir))
}

func TestDeleteAssets(t *testing.T) {
	assert.Nil(t, DeleteAllAssets(dir))
	assert.False(t, HasAssets(dir))
}
