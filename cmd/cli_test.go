package main

import (
	"bytes"
	"github.com/statping/statping/source"
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

var (
	dir string
)

func init() {
	utils.InitLogs()
}

func TestStatpingDirectory(t *testing.T) {
	dir = utils.Params.GetString("STATPING_DIR")
	require.NotContains(t, dir, "/cmd")
	require.NotEmpty(t, dir)
}

func TestEnvCLI(t *testing.T) {
	os.Setenv("API_SECRET", "demoapisecret123")
	os.Setenv("SASS", "/usr/local/bin/sass")

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
	assert.Contains(t, "API_SECRET=demoapisecret123", string(out))
	assert.Contains(t, "STATPING_DIR="+dir, string(out))
	assert.Contains(t, "SASS=/usr/local/bin/sass", string(out))

	os.Unsetenv("API_SECRET")
	os.Unsetenv("SASS")
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
	err := cmd.Execute()
	require.Nil(t, err)
	out, err := ioutil.ReadAll(b)
	assert.Nil(t, err)
	assert.Contains(t, string(out), VERSION)
	for _, f := range source.RequiredFiles {
		assert.FileExists(t, utils.Directory+"/assets/"+f)
	}
}

func TestUpdateCLI(t *testing.T) {
	cmd := rootCmd
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"update"})
	err := cmd.Execute()
	require.Nil(t, err)
	out, err := ioutil.ReadAll(b)
	require.Nil(t, err)
	assert.Contains(t, string(out), VERSION)
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
