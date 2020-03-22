// Package main for building the Statping CLI binary application. This package
// connects to all the other packages to make a runnable binary for multiple
// operating system.
//
// Compile Assets
//
// Before building, you must compile the Statping Assets with Rice, to install rice run the command below:
//		go get github.com/GeertJohan/go.rice
//		go get github.com/GeertJohan/go.rice/rice
//
// Once you have rice install, you can run the following command to build all assets inside the source directory.
//		cd source && rice embed-go
//
// Build Statping Binary
//
// To build the statup binary for your local environment, run the command below:
//		go build -o statup ./cmd
//
// Build All Binary Arch's
//
// To build Statping for Mac, Windows, Linux, and ARM devices, you can run xgo to build for all. xgo is an awesome
// golang package that requires Docker. https://github.com/crazy-max/xgo
//		docker pull crazy-max/xgo
//		build-all
//
// More info on: https://github.com/statping/statping
package main
