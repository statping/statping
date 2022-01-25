// +build !windows

package utils

import (
	"errors"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"syscall"
)

func DirWritable(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, errors.New("path doesn't exist")
	}

	if !info.IsDir() {
		return false, errors.New("path isn't a directory")
	}

	var stat syscall.Stat_t
	if err = syscall.Stat(path, &stat); err != nil {
		return false, errors.New("unable to get stat")
	}

	if uint32(os.Geteuid()) == stat.Uid {
		if info.Mode().Perm()&(1<<7) != 0 {
			// owner matches and has write permissions
			return true, nil
		} else {
			return false, errors.New("owner doesn't have write permissions for this path")
		}
	}

	if uint32(os.Getegid()) == stat.Gid {
		if info.Mode().Perm()&(1<<4) != 0 {
			// group matches and has write permissions
			return true, nil
		} else {
			return false, errors.New("group doesn't have write permissions for this path")
		}
	}

	if info.Mode().Perm()&(1<<1) != 0 {
		// all users have write permissions
		return true, nil
	}

	return false, errors.New("user doesn't have write permissions for this path")
}

func Ping(address string, secondsTimeout int) (int64, error) {
	ping, err := exec.LookPath("ping")
	if err != nil {
		return 0, err
	}
	out, _, err := Command(ping, address, "-c", "1", "-W", strconv.Itoa(secondsTimeout))
	if err != nil {
		return 0, err
	}
	if strings.Contains(out, "Unknown host") {
		return 0, errors.New("unknown host")
	}
	if strings.Contains(out, "100.0% packet loss") {
		return 0, errors.New("destination host unreachable")
	}

	r := regexp.MustCompile(`time=(.*) ms`)
	strs := r.FindStringSubmatch(out)
	if len(strs) < 2 {
		return 0, errors.New("could not parse ping duration")
	}
	f, _ := strconv.ParseFloat(strs[1], 64)
	return int64(f * 1000), nil
}
