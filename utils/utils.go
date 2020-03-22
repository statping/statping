// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/statping/statping
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
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/ararog/timeago"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	// Directory returns the current path or the STATPING_DIR environment variable
	Directory   string
	disableLogs bool
)

// init will set the utils.Directory to the current running directory, or STATPING_DIR if it is set
func init() {
	defaultDir, err := os.Getwd()
	if err != nil {
		defaultDir = "."
	}

	Directory = Getenv("STATPING_DIR", defaultDir).(string)

	// check if logs are disabled
	disableLogs = Getenv("DISABLE_LOGS", false).(bool)
	if disableLogs {
		Log.Out = ioutil.Discard
	}

	Log.Debugln("current working directory: ", Directory)
	Log.AddHook(new(hook))
	Log.SetNoLock()
	checkVerboseMode()
}

func Getenv(key string, defaultValue interface{}) interface{} {
	if val, ok := os.LookupEnv(key); ok {
		if val != "" {
			switch d := defaultValue.(type) {

			case int, int64:
				return int(ToInt(val))

			case time.Duration:
				dur, err := time.ParseDuration(val)
				if err != nil {
					return d
				}
				return dur
			case bool:
				ok, err := strconv.ParseBool(val)
				if err != nil {
					return d
				}
				return ok
			default:
				return val
			}
		}
	}
	return defaultValue
}

func SliceConvert(g []*interface{}) []interface{} {
	var arr []interface{}
	for _, v := range g {
		arr = append(arr, v)
	}
	return arr
}

// ToInt converts a int to a string
func ToInt(s interface{}) int64 {
	switch v := s.(type) {
	case string:
		val, _ := strconv.Atoi(v)
		return int64(val)
	case []byte:
		val, _ := strconv.Atoi(string(v))
		return int64(val)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case int:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint:
		return int64(v)
	default:
		return 0
	}
}

// ConvertInterface will take all the keys/values from an interface and replace all %type.Key from a string
// Input:   {"name": "%service.Name", "domain": "%service.Domain"}
// Output:  {"name": "Google DNS", "domain": "8.8.8.8"}
func ConvertInterface(in string, obj interface{}) string {
	if reflect.ValueOf(obj).IsNil() {
		return in
	}
	s := reflect.ValueOf(obj).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		find := strings.Split(fmt.Sprintf("%s.%v", typeOfT, typeOfT.Field(i).Name), ".")
		find[1] = strings.ToLower(find[1])
		key := strings.Join(find[1:], ".")
		in = strings.ReplaceAll(in, fmt.Sprintf("%%%v", key), fmt.Sprintf("%v", f.Interface()))
	}
	return in
}

// ToString converts a int to a string
func ToString(s interface{}) string {
	switch v := s.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%v", v)
	case float32, float64:
		return fmt.Sprintf("%f", v)
	case []byte:
		return string(v)
	case bool:
		return fmt.Sprintf("%t", v)
	case time.Time:
		return v.Format("Monday January _2, 2006 at 03:04PM")
	case time.Duration:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
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
		Log.Debugf("file exist: %v (%v)", name, !os.IsNotExist(err))
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// DeleteFile will attempt to delete a file
//		DeleteFile("newfile.json")
func DeleteFile(file string) error {
	Log.Debugln("deleting file: " + file)
	return os.Remove(file)
}

// RenameDirectory will attempt rename a directory to a new name
func RenameDirectory(fromDir string, toDir string) error {
	Log.Debugln("renaming directory: " + fromDir + "to: " + toDir)
	return os.Rename(fromDir, toDir)
}

// DeleteDirectory will attempt to delete a directory and all contents inside
//		DeleteDirectory("assets")
func DeleteDirectory(directory string) error {
	Log.Debugln("removing directory: " + directory)
	return os.RemoveAll(directory)
}

// CreateDirectory will attempt to create a directory
//		CreateDirectory("assets")
func CreateDirectory(directory string) error {
	Log.Debugln("creating directory: " + directory)
	if err := os.Mkdir(directory, os.ModePerm); err != os.ErrExist {
		return err
	}
	return nil
}

// FolderExists will return true if the folder exists
func FolderExists(folder string) bool {
	if _, err := os.Stat(folder); os.IsExist(err) {
		return true
	}
	return false
}

func OpenFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	return string(data), err
}

// CopyFile will copy a file to a new directory
//		CopyFile("source.jpg", "/tmp/source.jpg")
func CopyFile(src, dst string) error {
	Log.Debugln(fmt.Sprintf("copying file: %v to %v", src, dst))
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

// IsType will return true if a variable can implement an interface
func IsType(n interface{}, obj interface{}) bool {
	one := reflect.TypeOf(n)
	two := reflect.ValueOf(obj).Elem()
	return one.Implements(two.Type())
}

// Command will run a terminal command with 'sh -c COMMAND' and return stdout and errOut as strings
//		in, out, err := Command("sass assets/scss assets/css/base.css")
func Command(name string, args ...string) (string, string, error) {
	Log.Debugln("running command: " + name + strings.Join(args, " "))
	testCmd := exec.Command(name, args...)
	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := testCmd.StdoutPipe()
	stderrIn, _ := testCmd.StderrPipe()
	err := testCmd.Start()
	if err != nil {
		return "", "", err
	}

	go func() {
		stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
	}()

	go func() {
		stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)
	}()

	err = testCmd.Wait()
	if err != nil {
		return string(stdout), string(stderr), err
	}

	if errStdout != nil || errStderr != nil {
		return string(stdout), string(stderr), errors.New("failed to capture stdout or stderr")
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
	err := ioutil.WriteFile(filename, data, os.ModePerm)
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
func HttpRequest(url, method string, content interface{}, headers []string, body io.Reader, timeout time.Duration, verifySSL bool) ([]byte, *http.Response, error) {
	var err error
	var req *http.Request
	t1 := time.Now()
	if req, err = http.NewRequest(method, url, body); err != nil {
		httpMetric.Errors++
		return nil, nil, err
	}
	req.Header.Set("User-Agent", "Statping")
	if content != nil {
		req.Header.Set("Content-Type", content.(string))
	}

	verifyHost := req.URL.Hostname()
	for _, h := range headers {
		keyVal := strings.SplitN(h, "=", 2)
		if len(keyVal) == 2 {
			if keyVal[0] != "" && keyVal[1] != "" {
				if strings.ToLower(keyVal[0]) == "host" {
					req.Host = strings.TrimSpace(keyVal[1])
					verifyHost = req.Host
				} else {
					req.Header.Set(keyVal[0], keyVal[1])
				}
			}
		}
	}
	var resp *http.Response

	dialer := &net.Dialer{
		Timeout:   timeout,
		KeepAlive: timeout,
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: !verifySSL,
			ServerName:         verifyHost,
		},
		DisableKeepAlives:     true,
		ResponseHeaderTimeout: timeout,
		TLSHandshakeTimeout:   timeout,
		Proxy:                 http.ProxyFromEnvironment,
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			// redirect all connections to host specified in url
			addr = strings.Split(req.URL.Host, ":")[0] + addr[strings.LastIndex(addr, ":"):]
			return dialer.DialContext(ctx, network, addr)
		},
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}

	if resp, err = client.Do(req); err != nil {
		httpMetric.Errors++
		return nil, resp, err
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	// record HTTP metrics
	t2 := time.Now().Sub(t1).Milliseconds()
	httpMetric.Requests++
	httpMetric.Milliseconds += t2 / httpMetric.Requests
	httpMetric.Bytes += int64(len(contents))

	return contents, resp, err
}

const (
	B  = 0x100
	N  = 0x1000
	BM = 0xff
)

func NewPerlin(alpha, beta float64, n int, seed int64) *Perlin {
	return NewPerlinRandSource(alpha, beta, n, rand.NewSource(seed))
}

// Perlin is the noise generator
type Perlin struct {
	alpha float64
	beta  float64
	n     int

	p  [B + B + 2]int
	g3 [B + B + 2][3]float64
	g2 [B + B + 2][2]float64
	g1 [B + B + 2]float64
}

func NewPerlinRandSource(alpha, beta float64, n int, source rand.Source) *Perlin {
	var p Perlin
	var i int

	p.alpha = alpha
	p.beta = beta
	p.n = n

	r := rand.New(source)

	for i = 0; i < B; i++ {
		p.p[i] = i
		p.g1[i] = float64((r.Int()%(B+B))-B) / B

		for j := 0; j < 2; j++ {
			p.g2[i][j] = float64((r.Int()%(B+B))-B) / B
		}

		normalize2(&p.g2[i])
	}

	for ; i > 0; i-- {
		k := p.p[i]
		j := r.Int() % B
		p.p[i] = p.p[j]
		p.p[j] = k
	}

	for i := 0; i < B+2; i++ {
		p.p[B+i] = p.p[i]
		p.g1[B+i] = p.g1[i]
		for j := 0; j < 2; j++ {
			p.g2[B+i][j] = p.g2[i][j]
		}
		for j := 0; j < 3; j++ {
			p.g3[B+i][j] = p.g3[i][j]
		}
	}

	return &p
}

func normalize2(v *[2]float64) {
	s := math.Sqrt(v[0]*v[0] + v[1]*v[1])
	v[0] = v[0] / s
	v[1] = v[1] / s
}

func (p *Perlin) Noise1D(x float64) float64 {
	var scale float64 = 1
	var sum float64
	px := x

	for i := 0; i < p.n; i++ {
		val := p.noise1(px)
		sum += val / scale
		scale *= p.alpha
		px *= p.beta
	}
	if sum < 0 {
		sum = sum * -1
	}
	return sum
}

func (p *Perlin) noise1(arg float64) float64 {
	var vec [1]float64
	vec[0] = arg

	t := vec[0] + N
	bx0 := int(t) & BM
	bx1 := (bx0 + 1) & BM
	rx0 := t - float64(int(t))
	rx1 := rx0 - 1.

	sx := sCurve(rx0)
	u := rx0 * p.g1[p.p[bx0]]
	v := rx1 * p.g1[p.p[bx1]]

	return lerp(sx, u, v)
}

func sCurve(t float64) float64 {
	return t * t * (3. - 2.*t)
}

func lerp(t, a, b float64) float64 {
	return a + t*(b-a)
}
