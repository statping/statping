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

func DeleteFile(file string) bool {
	err := os.Remove(file)
	if err != nil {
		Log(3, err)
		return false
	}
	return true
}
