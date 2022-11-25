package utils

import (
	"errors"
	"os"
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
