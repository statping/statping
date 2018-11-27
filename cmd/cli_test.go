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
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/source"
	"github.com/hunterlong/statup/utils"
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
}

func TestStartServerCommand(t *testing.T) {
	os.Setenv("DB_CONN", "sqlite")
	cmd := helperCommand(nil, "")
	var got = make(chan string)
	commandAndSleep(cmd, time.Duration(8*time.Second), got)
	os.Unsetenv("DB_CONN")
	gg, _ := <-got
	assert.Contains(t, gg, "DB_CONN environment variable was found")
	assert.Contains(t, gg, "Core database does not exist, creating now!")
	assert.Contains(t, gg, "Starting monitoring process for 5 Services")
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
	cmd := helperCommand(nil, "static")
	var got = make(chan string)
	commandAndSleep(cmd, time.Duration(10*time.Second), got)
	gg, _ := <-got
	t.Log(gg)
	assert.Contains(t, gg, "Exporting Static 'index.html' page...")
	assert.Contains(t, gg, "Exported Statup index page: 'index.html'")
	assert.True(t, fileExists(dir+"/index.html"))
}

func TestUpdateCommand(t *testing.T) {
	cmd := helperCommand(nil, "version")
	var got = make(chan string)
	commandAndSleep(cmd, time.Duration(10*time.Second), got)
	gg, _ := <-got
	t.Log(gg)
	assert.Contains(t, gg, "Statup")
}

func TestAssetsCommand(t *testing.T) {
	c := testcli.Command("statup", "assets")
	c.Run()
	t.Log(c.Stdout())
	t.Log("Directory for Assets: ", dir)
	assert.FileExists(t, dir+"/assets/robots.txt")
	assert.FileExists(t, dir+"/assets/scss/base.scss")
}

func TestRunCommand(t *testing.T) {
	cmd := helperCommand(nil, "run")
	var got = make(chan string)
	commandAndSleep(cmd, time.Duration(5*time.Second), got)
	gg, _ := <-got
	t.Log(gg)
	assert.Contains(t, gg, "Running 1 time and saving to database...")
	assert.Contains(t, gg, "Check is complete.")
}

func TestEnvironmentVarsCommand(t *testing.T) {
	c := testcli.Command("statup", "env")
	c.Run()
	assert.True(t, c.StdoutContains("Statup Environment Variable"))
}

func TestVersionCLI(t *testing.T) {
	run := catchCLI([]string{"version"})
	assert.EqualError(t, run, "end")
}

func TestAssetsCLI(t *testing.T) {
	run := catchCLI([]string{"assets"})
	assert.EqualError(t, run, "end")
	assert.FileExists(t, dir+"/assets/css/base.css")
	assert.FileExists(t, dir+"/assets/scss/base.scss")
}

func TestSassCLI(t *testing.T) {
	run := catchCLI([]string{"sass"})
	assert.EqualError(t, run, "end")
	assert.FileExists(t, dir+"/assets/css/base.css")
}

func TestUpdateCLI(t *testing.T) {
	t.SkipNow()
	run := catchCLI([]string{"update"})
	assert.EqualError(t, run, "end")
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
	run := catchCLI([]string{"run"})
	assert.EqualError(t, run, "end")
}

func TestEnvCLI(t *testing.T) {
	run := catchCLI([]string{"env"})
	assert.Error(t, run)
	Clean()
}

func commandAndSleep(cmd *exec.Cmd, duration time.Duration, out chan<- string) {
	go func(out chan<- string) {
		runCommand(cmd, out)
	}(out)
	time.Sleep(duration)
	cmd.Process.Kill()
}

func helperCommand(envs []string, s ...string) *exec.Cmd {
	cmd := exec.Command("statup", s...)
	return cmd
}

func runCommand(c *exec.Cmd, out chan<- string) {
	bout, _ := c.CombinedOutput()
	out <- string(bout)
}

func fileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func Clean() {
	utils.DeleteFile(dir + "/config.yml")
	utils.DeleteFile(dir + "/statup.db")
	utils.DeleteDirectory(dir + "/assets")
	utils.DeleteDirectory(dir + "/logs")
	core.CoreApp = core.NewCore()
	source.Assets()
	//core.CloseDB()
	os.Unsetenv("DB_CONN")
	time.Sleep(2 * time.Second)
}
