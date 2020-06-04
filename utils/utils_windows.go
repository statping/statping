package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
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

func Ping(address string, secondsTimeout int) error {
	ping, err := exec.LookPath("ping")
	if err != nil {
		return err
	}
	out, _, err := Command(ping, address, "-n 1", fmt.Sprintf("-w %v", secondsTimeout*1000))
	if err != nil {
		return err
	}
	if strings.Contains(out, "Destination Host Unreachable") {
		return errors.New("destination host unreachable")
	}
	return nil
}
