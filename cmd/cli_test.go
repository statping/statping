package main

import (
	"github.com/rendon/testcli"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	c := testcli.Command("statup", "version")
	c.Run()
	assert.True(t, c.StdoutContains("Statup v"+VERSION))
}

func TestHelpCommand(t *testing.T) {
	c := testcli.Command("statup", "help")
	c.Run()
	t.Log(c.Stdout())
	assert.True(t, c.StdoutContains("statup help               - Shows the user basic information about Statup"))
}

func TestExportCommand(t *testing.T) {
	t.SkipNow()
	c := testcli.Command("statup", "export")
	c.Run()
	t.Log(c.Stdout())
	assert.True(t, c.StdoutContains("Exporting Static 'index.html' page"))
	assert.True(t, fileExists(dir+"/cmd/index.html"))
}

func TestAssetsCommand(t *testing.T) {
	c := testcli.Command("statup", "assets")
	c.Run()
	t.Log(c.Stdout())
	t.Log("Directory for Assets: ", dir)
	assert.FileExists(t, dir+"/assets/robots.txt")
	assert.FileExists(t, dir+"/assets/js/main.js")
	assert.FileExists(t, dir+"/assets/scss/base.scss")
}

func TestVersionCLI(t *testing.T) {
	run := CatchCLI([]string{"statup", "version"})
	assert.Nil(t, run)
}

func TestAssetsCLI(t *testing.T) {
	run := CatchCLI([]string{"statup", "assets"})
	assert.Nil(t, run)
	assert.FileExists(t, dir+"/assets/css/base.css")
	assert.FileExists(t, dir+"/assets/scss/base.scss")
}

func TestSassCLI(t *testing.T) {
	run := CatchCLI([]string{"statup", "sass"})
	assert.Nil(t, run)
	assert.FileExists(t, dir+"/assets/css/base.css")
}

func TestUpdateCLI(t *testing.T) {
	run := CatchCLI([]string{"statup", "update"})
	assert.Nil(t, run)
}

func TestTestPackageCLI(t *testing.T) {
	run := CatchCLI([]string{"statup", "test", "plugins"})
	assert.Nil(t, run)
}

func TestHelpCLI(t *testing.T) {
	run := CatchCLI([]string{"statup", "help"})
	assert.Nil(t, run)
}

func TestRunOnceCLI(t *testing.T) {
	t.SkipNow()
	run := CatchCLI([]string{"statup", "run"})
	assert.Nil(t, run)
}

func TestEnvCLI(t *testing.T) {
	run := CatchCLI([]string{"statup", "env"})
	assert.Error(t, run)
}
