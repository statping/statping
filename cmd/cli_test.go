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

package main

import (
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/utils"
	"github.com/rendon/testcli"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
	"time"
)

var (
	dir string
)

func init() {
	dir = utils.Directory
	core.SampleHits = 480
}

func TestStartServerCommand(t *testing.T) {
	t.SkipNow()
	os.Setenv("DB_CONN", "sqlite")
	cmd := helperCommand(nil, "")
	var got = make(chan string)
	commandAndSleep(cmd, time.Duration(60*time.Second), got)
	os.Unsetenv("DB_CONN")
	gg, _ := <-got
	assert.Contains(t, gg, "DB_CONN environment variable was found")
	assert.Contains(t, gg, "Core database does not exist, creating now!")
	assert.Contains(t, gg, "Starting monitoring process for 5 Services")
}

func TestVersionCommand(t *testing.T) {
	c := testcli.Command("statping", "version")
	c.Run()
	assert.True(t, c.StdoutContains(VERSION))
}

func TestHelpCommand(t *testing.T) {
	c := testcli.Command("statping", "help")
	c.Run()
	t.Log(c.Stdout())
	assert.True(t, c.StdoutContains("statping help               - Shows the user basic information about Statping"))
}

func TestStaticCommand(t *testing.T) {
	t.SkipNow()
	cmd := helperCommand(nil, "static")
	var got = make(chan string)
	commandAndSleep(cmd, time.Duration(10*time.Second), got)
	gg, _ := <-got
	t.Log(gg)
	assert.Contains(t, gg, "Exporting Static 'index.html' page...")
	assert.Contains(t, gg, "Exported Statping index page: 'index.html'")
	assert.FileExists(t, dir+"/index.html")
}

func TestExportCommand(t *testing.T) {
	t.SkipNow()
	cmd := helperCommand(nil, "export")
	var got = make(chan string)
	commandAndSleep(cmd, time.Duration(10*time.Second), got)
	gg, _ := <-got
	t.Log(gg)
	assert.FileExists(t, dir+"/statping-export.json")
}

func TestUpdateCommand(t *testing.T) {
	cmd := helperCommand(nil, "version")
	var got = make(chan string)
	commandAndSleep(cmd, time.Duration(15*time.Second), got)
	gg, _ := <-got
	t.Log(gg)
	assert.Contains(t, gg, VERSION)
}

func TestAssetsCommand(t *testing.T) {
	c := testcli.Command("statping", "assets")
	c.Run()
	t.Log(c.Stdout())
	t.Log("Directory for Assets: ", dir)
	time.Sleep(1 * time.Second)
	assert.FileExists(t, dir+"/assets/robots.txt")
	assert.FileExists(t, dir+"/assets/scss/base.scss")
}

func TestRunCommand(t *testing.T) {
	t.SkipNow()
	cmd := helperCommand(nil, "run")
	var got = make(chan string)
	commandAndSleep(cmd, time.Duration(15*time.Second), got)
	gg, _ := <-got
	t.Log(gg)
	assert.Contains(t, gg, "Running 1 time and saving to database...")
	assert.Contains(t, gg, "Check is complete.")
}

func TestEnvironmentVarsCommand(t *testing.T) {
	c := testcli.Command("statping", "env")
	c.Run()
	assert.True(t, c.StdoutContains("Statping Environment Variable"))
}

func TestVersionCLI(t *testing.T) {
	run := catchCLI([]string{"version"})
	assert.EqualError(t, run, "end")
}

func TestAssetsCLI(t *testing.T) {
	catchCLI([]string{"assets"})
	//assert.EqualError(t, run, "end")
	assert.FileExists(t, dir+"/assets/css/base.css")
	assert.FileExists(t, dir+"/assets/scss/base.scss")
}

func TestSassCLI(t *testing.T) {
	catchCLI([]string{"sass"})
	assert.FileExists(t, dir+"/assets/css/base.css")
}

func TestUpdateCLI(t *testing.T) {
	t.SkipNow()
	cmd := helperCommand(nil, "update")
	var got = make(chan string)
	commandAndSleep(cmd, time.Duration(15*time.Second), got)
	gg, _ := <-got
	t.Log(gg)
	assert.Contains(t, gg, "version")
}

func TestTestPackageCLI(t *testing.T) {
	t.SkipNow()
	run := catchCLI([]string{"test", "plugins"})
	assert.EqualError(t, run, "end")
}

func TestHelpCLI(t *testing.T) {
	run := catchCLI([]string{"help"})
	assert.EqualError(t, run, "end")
}

func TestRunOnceCLI(t *testing.T) {
	t.SkipNow()
	run := catchCLI([]string{"run"})
	assert.EqualError(t, run, "end")
}

func TestEnvCLI(t *testing.T) {
	run := catchCLI([]string{"env"})
	assert.Error(t, run)
}

func commandAndSleep(cmd *exec.Cmd, duration time.Duration, out chan<- string) {
	go func(out chan<- string) {
		runCommand(cmd, out)
	}(out)
	time.Sleep(duration)
	cmd.Process.Kill()
}

func helperCommand(envs []string, s ...string) *exec.Cmd {
	cmd := exec.Command("statping", s...)
	return cmd
}

func runCommand(c *exec.Cmd, out chan<- string) {
	bout, _ := c.CombinedOutput()
	out <- string(bout)
}
