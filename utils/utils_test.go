package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestCreateLog(t *testing.T) {
	err := createLog(Directory)
	assert.Nil(t, err)
}

func TestReplaceValue(t *testing.T) {
	assert.Equal(t, true, replaceVal(true))
	assert.Equal(t, 42, replaceVal(42))
	assert.Equal(t, "hello world", replaceVal("hello world"))
	assert.Equal(t, "5s", replaceVal(time.Duration(5*time.Second)))
}

func TestInitLogs(t *testing.T) {
	assert.Nil(t, InitLogs())
	Log.Infoln("this is a test")
	assert.FileExists(t, Directory+"/logs/statping.log")
}

func TestDir(t *testing.T) {
	assert.Contains(t, Directory, "statping/statping")
}

func TestCommand(t *testing.T) {
	t.SkipNow()
	_, out, err := Command("/bin/echo", "\"statping testing\"")
	assert.Nil(t, err)
	assert.Contains(t, out, "statping")
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
	assert.Equal(t, Now().Format("Monday January _2, 2006 at 03:04PM"), ToString(Now()))
}

func ExampleToString() {
	amount := 42
	fmt.Print(ToString(amount))
	// Output: 42
}

func TestSaveFile(t *testing.T) {
	assert.Nil(t, SaveFile(Directory+"/test.txt", []byte("testing saving a file")))
}

func TestOpenFile(t *testing.T) {
	f, err := OpenFile(Directory + "/test.txt")
	require.Nil(t, err)
	assert.Equal(t, "testing saving a file", f)
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
	assert.NotNil(t, req)
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

func TestHashPassword(t *testing.T) {
	assert.Equal(t, 60, len(HashPassword("password123")))
}

func TestNewSHA1Hash(t *testing.T) {
	hash := NewSHA256Hash()
	assert.NotEmpty(t, hash)
	assert.Len(t, hash, 64)
	assert.Len(t, NewSHA256Hash(), 64)
	assert.NotEqual(t, hash, NewSHA256Hash())
}

func TestRandomString(t *testing.T) {
	assert.NotEmpty(t, RandomString(5))
}

func TestDeleteDirectory(t *testing.T) {
	assert.Nil(t, DeleteDirectory(Directory+"/logs"))
}

func TestRenameDirectory(t *testing.T) {
	assert.Nil(t, CreateDirectory(Directory+"/example"))
	require.DirExists(t, Directory+"/example")
	assert.Nil(t, RenameDirectory(Directory+"/example", Directory+"/renamed_example"))
	require.DirExists(t, Directory+"/renamed_example")
	assert.Nil(t, os.RemoveAll(Directory+"/renamed_example"))
}

func TestHttpRequest(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "/")
		assert.Equal(t, req.Header["Aaa"], []string{"bbbb="})
		assert.Equal(t, req.Header["Ccc"], []string{"ddd"})
		// Send response to be tested
		rw.Write([]byte(`OK`))
	}))
	// Close the server when test finishes
	defer server.Close()

	body, resp, err := HttpRequest(server.URL, "GET", "application/json", []string{"aaa=bbbb=", "ccc=ddd"}, nil, 2*time.Second, false)

	assert.Nil(t, err)
	assert.Equal(t, []byte("OK"), body)
	assert.Equal(t, resp.StatusCode, 200)
}

func TestConfigLoad(t *testing.T) {
	InitCLI()
	setDefaults()

	s := Params.GetString
	b := Params.GetBool

	Params.Set("DB_CONN", "sqlite")

	assert.Equal(t, "sqlite", s("DB_CONN"))
	assert.Equal(t, Directory, s("STATPING_DIR"))
	assert.True(t, b("SAMPLE_DATA"))
	assert.True(t, b("ALLOW_REPORTS"))
}
