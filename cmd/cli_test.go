// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"github.com/rendon/testcli"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfirmVersion(t *testing.T) {
	t.SkipNow()
	assert.NotEmpty(t, VERSION)
}

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
	assert.True(t, fileExists(dir+"/index.html"))
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
	assert.EqualError(t, run, "end")
}

func TestAssetsCLI(t *testing.T) {
	t.SkipNow()
	run := CatchCLI([]string{"statup", "assets"})
	assert.EqualError(t, run, "end")
	assert.FileExists(t, dir+"/assets/css/base.css")
	assert.FileExists(t, dir+"/assets/scss/base.scss")
}

func TestSassCLI(t *testing.T) {
	run := CatchCLI([]string{"statup", "sass"})
	assert.EqualError(t, run, "end")
	assert.FileExists(t, dir+"/assets/css/base.css")
}

func TestUpdateCLI(t *testing.T) {
	t.SkipNow()
	run := CatchCLI([]string{"statup", "update"})
	assert.EqualError(t, run, "end")
}

func TestTestPackageCLI(t *testing.T) {
	run := CatchCLI([]string{"statup", "test", "plugins"})
	assert.EqualError(t, run, "end")
}

func TestHelpCLI(t *testing.T) {
	run := CatchCLI([]string{"statup", "help"})
	assert.EqualError(t, run, "end")
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
