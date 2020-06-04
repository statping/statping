package main

import (
	"bytes"
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

var (
	dir string
)

func init() {
	utils.InitLogs()
}

func TestStatpingDirectory(t *testing.T) {
	dir := utils.Directory
	require.NotContains(t, dir, "/cmd")
	require.NotEmpty(t, dir)

	dir = utils.Params.GetString("STATPING_DIR")
	require.NotContains(t, dir, "/cmd")
	require.NotEmpty(t, dir)
}

func TestEnvCLI(t *testing.T) {
	cmd := rootCmd
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"env"})
	err := cmd.Execute()
	require.Nil(t, err)
	out, err := ioutil.ReadAll(b)
	require.Nil(t, err)
	assert.Contains(t, string(out), VERSION)
	assert.Contains(t, utils.Directory, string(out))
	assert.Contains(t, "SAMPLE_DATA=true", string(out))
}

func TestVersionCLI(t *testing.T) {
	cmd := rootCmd
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"version"})
	err := cmd.Execute()
	require.Nil(t, err)
	out, err := ioutil.ReadAll(b)
	require.Nil(t, err)
	assert.Contains(t, VERSION, string(out))
}

func TestAssetsCLI(t *testing.T) {
	cmd := rootCmd
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"assets"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	assert.Nil(t, err)
	assert.Contains(t, string(out), VERSION)
	assert.FileExists(t, utils.Directory+"/assets/css/main.css")
	assert.FileExists(t, utils.Directory+"/assets/css/style.css")
	assert.FileExists(t, utils.Directory+"/assets/css/vendor.css")
	assert.FileExists(t, utils.Directory+"/assets/scss/base.scss")
	assert.FileExists(t, utils.Directory+"/assets/scss/mobile.scss")
	assert.FileExists(t, utils.Directory+"/assets/scss/variables.scss")
}

func TestHelpCLI(t *testing.T) {
	cmd := rootCmd
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"help"})
	err := cmd.Execute()
	require.Nil(t, err)
	out, err := ioutil.ReadAll(b)
	require.Nil(t, err)
	assert.Contains(t, string(out), VERSION)
}

func TestResetCLI(t *testing.T) {
	err := utils.SaveFile(utils.Directory+"/statping.db", []byte("test data"))
	require.Nil(t, err)

	cmd := rootCmd
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"reset"})
	err = cmd.Execute()
	require.Nil(t, err)
	out, err := ioutil.ReadAll(b)
	require.Nil(t, err)
	assert.Contains(t, string(out), VERSION)

	assert.NoDirExists(t, utils.Directory+"/assets")
	assert.NoDirExists(t, utils.Directory+"/logs")
	assert.NoFileExists(t, utils.Directory+"/config.yml")
	assert.NoFileExists(t, utils.Directory+"/statping.db")
	assert.FileExists(t, utils.Directory+"/statping.db.backup")

	err = utils.DeleteFile(utils.Directory + "/statping.db.backup")
	require.Nil(t, err)
}
