package main

import (
	"fmt"
	"github.com/hunterlong/statping/source"
	"github.com/hunterlong/statping/utils"
	"net"
)

// UsingAssets will return true if /assets folder is present
func UsingAssets() bool {
	return source.UsingAssets(utils.Directory)
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "http://localhost"
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return fmt.Sprintf("http://%v", ipnet.IP.String())
			}
		}
	}
	return "http://localhost"
}
