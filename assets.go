package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/hunterlong/statup/log"
	"io/ioutil"
	"os"
	"os/exec"
)

var (
	useAssets bool
)

func CopyToPublic(box *rice.Box, folder, file string) {
	base, err := box.String(file)
	if err != nil {
		log.Send(2, err)
	}
	ioutil.WriteFile("assets/"+folder+"/"+file, []byte(base), 0644)
}

func MakePublicFolder(folder string) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err = os.MkdirAll(folder, 0755)
		if err != nil {
			log.Send(2, err)
		}
	}
}

func CompileSASS() {
	sassBin := os.Getenv("SASS")
	shell := os.Getenv("CMD_FILE")
	log.Send(1, fmt.Sprintf("Compiling SASS into /css/base.css..."))
	command := fmt.Sprintf("%v %v %v", sassBin, "assets/scss/base.scss", "assets/css/base.css")
	testCmd := exec.Command(shell, command)
	_, err := testCmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	log.Send(1, "SASS Compiling is complete!")
}

func hasAssets() bool {
	if _, err := os.Stat("assets"); err == nil {
		useAssets = true
		return true
	} else {
		assetEnv := os.Getenv("USE_ASSETS")
		if assetEnv == "true" {
			CreateAllAssets()
			useAssets = true
			return true
		}
	}
	return false
}

func SaveAsset(data, file string) {
	ioutil.WriteFile("assets/"+file, []byte(data), 0644)
}

func OpenAsset(file string) string {
	dat, err := ioutil.ReadFile("assets/" + file)
	log.Send(2, err)
	return string(dat)
}

func CreateAllAssets() {
	log.Send(1, "Creating folder 'assets' in current directory..")
	MakePublicFolder("assets")
	MakePublicFolder("assets/js")
	MakePublicFolder("assets/css")
	MakePublicFolder("assets/scss")
	MakePublicFolder("assets/emails")
	log.Send(1, "Inserting scss, css, emails, and javascript files into assets..")
	CopyToPublic(scssBox, "scss", "base.scss")
	CopyToPublic(scssBox, "scss", "variables.scss")
	CopyToPublic(emailBox, "emails", "error.html")
	CopyToPublic(emailBox, "emails", "failure.html")
	CopyToPublic(cssBox, "css", "bootstrap.min.css")
	CopyToPublic(jsBox, "js", "bootstrap.min.js")
	CopyToPublic(jsBox, "js", "Chart.bundle.min.js")
	CopyToPublic(jsBox, "js", "jquery-3.3.1.slim.min.js")
	CopyToPublic(jsBox, "js", "main.js")
	CopyToPublic(jsBox, "js", "setup.js")
	log.Send(1, "Compiling CSS from SCSS style...")
	CompileSASS()
	log.Send(1, "Statup assets have been inserted")
}
