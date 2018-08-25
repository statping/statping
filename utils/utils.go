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

package utils

import (
	"github.com/ararog/timeago"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	Directory string
)

func init() {
	if os.Getenv("STATUP_DIR") != "" {
		Directory = os.Getenv("STATUP_DIR")
	} else {
		Directory = dir()
	}
}

func StringInt(s string) int64 {
	num, _ := strconv.Atoi(s)
	return int64(num)
}

func IntString(s int) string {
	return strconv.Itoa(s)
}

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

func (t Timestamp) Ago() string {
	got, _ := timeago.TimeAgoWithTime(time.Now(), time.Time(t))
	return got
}

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

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func DeleteFile(file string) error {
	err := os.Remove(file)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDirectory(directory string) error {
	return os.RemoveAll(directory)
}
