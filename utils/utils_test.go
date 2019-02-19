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

package utils

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestCreateLog(t *testing.T) {
	err := createLog(Directory)
	assert.Nil(t, err)
}

func TestInitLogs(t *testing.T) {
	assert.Nil(t, InitLogs())
	assert.FileExists(t, Directory+"/logs/statup.log")
}

func TestDir(t *testing.T) {
	assert.Contains(t, Directory, "github.com/hunterlong/statping")
}

func TestCommand(t *testing.T) {
	t.SkipNow()
	in, out, err := Command("pwd")
	assert.Nil(t, err)
	assert.Contains(t, in, "statping")
	assert.Empty(t, out)
}

func TestLog(t *testing.T) {
	assert.Nil(t, Log(0, errors.New("this is a 0 level error")))
	assert.Nil(t, Log(1, errors.New("this is a 1 level error")))
	assert.Nil(t, Log(2, errors.New("this is a 2 level error")))
	assert.Nil(t, Log(3, errors.New("this is a 3 level error")))
	assert.Nil(t, Log(4, errors.New("this is a 4 level error")))
	assert.Nil(t, Log(5, errors.New("this is a 5 level error")))
}

func TestToInt(t *testing.T) {
	assert.Equal(t, int64(55), ToInt("55"))
	assert.Equal(t, int64(55), ToInt(55))
	assert.Equal(t, int64(55), ToInt(55.0))
	assert.Equal(t, int64(55), ToInt([]byte("55")))
}

func TestToString(t *testing.T) {
	assert.Equal(t, "55", ToString(55))
	assert.Equal(t, "55.000000", ToString(55.0))
	assert.Equal(t, "55", ToString([]byte("55")))
	dir, _ := time.ParseDuration("55s")
	assert.Equal(t, "55s", ToString(dir))
	assert.Equal(t, "true", ToString(true))
	assert.Equal(t, time.Now().Format("Monday January _2, 2006 at 03:04PM"), ToString(time.Now()))
}

func ExampleToString() {
	amount := 42
	fmt.Print(ToString(amount))
	// Output: 42
}

func TestSaveFile(t *testing.T) {
	assert.Nil(t, SaveFile(Directory+"/test.txt", []byte("testing saving a file")))
}

func TestFileExists(t *testing.T) {
	assert.True(t, FileExists(Directory+"/test.txt"))
	assert.False(t, FileExists(Directory+"fake.txt"))
}

func TestDeleteFile(t *testing.T) {
	assert.Nil(t, DeleteFile(Directory+"/test.txt"))
	assert.Error(t, DeleteFile(Directory+"/missingfilehere.txt"))
}

func TestFormatDuration(t *testing.T) {
	dur, _ := time.ParseDuration("158s")
	assert.Equal(t, "3 minutes", FormatDuration(dur))
	dur, _ = time.ParseDuration("-65s")
	assert.Equal(t, "1 minute", FormatDuration(dur))
	dur, _ = time.ParseDuration("3s")
	assert.Equal(t, "3 seconds", FormatDuration(dur))
	dur, _ = time.ParseDuration("48h")
	assert.Equal(t, "2 days", FormatDuration(dur))
	dur, _ = time.ParseDuration("12h")
	assert.Equal(t, "12 hours", FormatDuration(dur))
}

func ExampleDurationReadable() {
	dur, _ := time.ParseDuration("25m")
	readable := DurationReadable(dur)
	fmt.Print(readable)
	// Output: 25 minutes
}

func TestLogHTTP(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, Http(req))
}

func TestStringInt(t *testing.T) {
	assert.Equal(t, "1", ToString("1"))
}

func ExampleStringInt() {
	amount := "42"
	fmt.Print(ToString(amount))
	// Output: 42
}

func TestTimezone(t *testing.T) {
	zone := float32(-4.0)
	loc, _ := time.LoadLocation("America/Los_Angeles")
	timestamp := time.Date(2018, 1, 1, 10, 0, 0, 0, loc)
	timezone := Timezoner(timestamp, zone)
	assert.Equal(t, "2018-01-01 10:00:00 -0800 PST", timestamp.String())
	assert.Equal(t, "2018-01-01 18:00:00 +0000 UTC", timezone.UTC().String())
}

func TestTimestamp_Ago(t *testing.T) {
	now := Timestamp(time.Now())
	assert.Equal(t, "Just now", now.Ago())
}

func TestUnderScoreString(t *testing.T) {
	assert.Equal(t, "this_is_a_test", UnderScoreString("this is a test"))
}

func TestHashPassword(t *testing.T) {
	assert.Equal(t, 60, len(HashPassword("password123")))
}

func TestNewSHA1Hash(t *testing.T) {
	assert.NotEmpty(t, NewSHA1Hash(5))
}

func TestRandomString(t *testing.T) {
	assert.NotEmpty(t, RandomString(5))
}

func TestDeleteDirectory(t *testing.T) {
	assert.Nil(t, DeleteDirectory(Directory+"/logs"))
}
