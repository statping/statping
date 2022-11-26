//go:build !windows
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
