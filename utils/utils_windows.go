package utils

import (
	"errors"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func DirWritable(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, errors.New("path doesn't exist")
	}

	if !info.IsDir() {
		return false, errors.New("path isn't a directory")
	}

	if info.Mode().Perm()&(1<<(uint(7))) == 0 {
		return false, errors.New("write permission bit is not set on this file for user")
	}

	return true, nil
}

func Ping(address string, secondsTimeout int) (int64, error) {
	ping, err := exec.LookPath("ping")
	if err != nil {
		return 0, err
	}
	out, _, err := Command(ping, address, "-n", "1", "-w", strconv.Itoa(secondsTimeout*1000))
	if err != nil {
		return 0, err
	}
	if strings.Contains(out, "Destination Host Unreachable") {
		return 0, errors.New("destination host unreachable")
	}
	r := regexp.MustCompile(`Average = (.*)ms`)
	strs := r.FindStringSubmatch(out)
	if len(strs) < 2 {
		return 0, errors.New("could not parse ping duration")
	}
	f, _ := strconv.ParseFloat(strs[1], 64)
	return int64(f * 1000), nil
}
