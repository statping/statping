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
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/ararog/timeago"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	// Directory returns the current path or the STATPING_DIR environment variable
	Directory string
)

// init will set the utils.Directory to the current running directory, or STATPING_DIR if it is set
func init() {
	if os.Getenv("STATPING_DIR") != "" {
		Directory = os.Getenv("STATPING_DIR")
	} else {
		Directory = dir()
	}
}

// ToInt converts a int to a string
func ToInt(s interface{}) int64 {
	switch v := s.(type) {
	case string:
		val, _ := strconv.Atoi(v)
		return int64(val)
	default:
		return 0
	}
}

// ToString converts a int to a string
func ToString(s interface{}) string {
	switch v := s.(type) {
	case int, int32, int64:
		return fmt.Sprintf("%v", v)
	case float32, float64:
		return fmt.Sprintf("%f", v)
	case []byte:
		return string(v)
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// dir returns the current working directory
func dir() string {
	dir, err := os.Getwd()
	if err != nil {
		return "."
	}
	return dir
}

type Timestamp time.Time
type Timestamper interface {
	Ago() string
}

// Ago returns a human readable timestamp based on the Timestamp (time.Time) interface
func (t Timestamp) Ago() string {
	got, _ := timeago.TimeAgoWithTime(time.Now(), time.Time(t))
	return got
}

// UnderScoreString will return a string that replaces spaces and other characters to underscores
//		UnderScoreString("Example String")
//		// example_string
func UnderScoreString(str string) string {

	// convert every letter to lower case
	newStr := strings.ToLower(str)

	// convert all spaces/tab to underscore
	regExp := regexp.MustCompile("[[:space:][:blank:]]")
	newStrByte := regExp.ReplaceAll([]byte(newStr), []byte("_"))

	regExp = regexp.MustCompile("`[^a-z0-9]`i")
	newStrByte = regExp.ReplaceAll(newStrByte, []byte("_"))

	regExp = regexp.MustCompile("[!/']")
	newStrByte = regExp.ReplaceAll(newStrByte, []byte("_"))

	// and remove underscore from beginning and ending

	newStr = strings.TrimPrefix(string(newStrByte), "_")
	newStr = strings.TrimSuffix(newStr, "_")

	return newStr
}

// FileExists returns true if a file exists
//		exists := FileExists("assets/css/base.css")
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// DeleteFile will attempt to delete a file
//		DeleteFile("newfile.json")
func DeleteFile(file string) error {
	Log(1, "deleting file: "+file)
	err := os.Remove(file)
	if err != nil {
		return err
	}
	return nil
}

// DeleteDirectory will attempt to delete a directory and all contents inside
//		DeleteDirectory("assets")
func DeleteDirectory(directory string) error {
	return os.RemoveAll(directory)
}

// Command will run a terminal command with 'sh -c COMMAND' and return stdout and errOut as strings
//		in, out, err := Command("sass assets/scss assets/css/base.css")
func Command(cmd string) (string, string, error) {
	Log(1, "running command: "+cmd)
	testCmd := exec.Command("sh", "-c", cmd)
	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := testCmd.StdoutPipe()
	stderrIn, _ := testCmd.StderrPipe()
	testCmd.Start()

	go func() {
		stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
	}()

	go func() {
		stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)
	}()

	err := testCmd.Wait()
	if err != nil {
		return "", "", err
	}

	if errStdout != nil || errStderr != nil {
		return "", "", errors.New("failed to capture stdout or stderr")
	}

	outStr, errStr := string(stdout), string(stderr)
	return outStr, errStr, err
}

// copyAndCapture will read a terminal command into bytes
func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}

// DurationReadable will return a time.Duration into a human readable string
// // t := time.Duration(5 * time.Minute)
// // DurationReadable(t)
// // returns: 5 minutes
func DurationReadable(d time.Duration) string {
	if d.Hours() >= 1 {
		return fmt.Sprintf("%0.0f hours", d.Hours())
	} else if d.Minutes() >= 1 {
		return fmt.Sprintf("%0.0f minutes", d.Minutes())
	} else if d.Seconds() >= 1 {
		return fmt.Sprintf("%0.0f seconds", d.Seconds())
	}
	return d.String()
}

// SaveFile will create a new file with data inside it
//		SaveFile("newfile.json", []byte('{"data": "success"}')
func SaveFile(filename string, data []byte) error {
	err := ioutil.WriteFile(filename, data, 0644)
	return err
}

// HttpRequest is a global function to send a HTTP request
// // url - The URL for HTTP request
// // method - GET, POST, DELETE, PATCH
// // content - The HTTP request content type (text/plain, application/json, or nil)
// // headers - An array of Headers to be sent (KEY=VALUE) []string{"Authentication=12345", ...}
// // body - The body or form data to send with HTTP request
// // timeout - Specific duration to timeout on. time.Duration(30 * time.Seconds)
// // You can use a HTTP Proxy if you HTTP_PROXY environment variable
func HttpRequest(url, method string, content interface{}, headers []string, body io.Reader, timeout time.Duration) ([]byte, *http.Response, error) {
	var err error
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		DisableKeepAlives:     true,
		ResponseHeaderTimeout: timeout,
		TLSHandshakeTimeout:   timeout,
		Proxy:                 http.ProxyFromEnvironment,
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}
	var req *http.Request
	if req, err = http.NewRequest(method, url, body); err != nil {
		return nil, nil, err
	}
	req.Header.Set("User-Agent", "Statping")
	if content != nil {
		req.Header.Set("Content-Type", content.(string))
	}
	for _, h := range headers {
		keyVal := strings.Split(h, "=")
		if len(keyVal) == 2 {
			if keyVal[0] != "" && keyVal[1] != "" {
				req.Header.Set(keyVal[0], keyVal[1])
			}
		}
	}
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	return contents, resp, err
}
