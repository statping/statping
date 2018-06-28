package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
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
		fmt.Println(err)
	}
	ioutil.WriteFile("assets/"+folder+"/"+file, []byte(base), 0644)
}

func MakePublicFolder(folder string) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err = os.MkdirAll(folder, 0755)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func CompileSASS() {
	cmdBin := os.Getenv("CMD_FILE")
	fmt.Println("Compiling SASS into /css/base.css...")
	testCmd := exec.Command(cmdBin, "sass assets/scss/base.scss assets/css/base.css")
	_, err := testCmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("SASS Compiling is complete!")
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
	dat, _ := ioutil.ReadFile("assets/" + file)
	return string(dat)
}

func CreateAllAssets() {
	fmt.Println("Creating folder 'assets' in current directory..")
	MakePublicFolder("assets")
	MakePublicFolder("assets/js")
	MakePublicFolder("assets/css")
	MakePublicFolder("assets/scss")
	MakePublicFolder("assets/emails")
	fmt.Println("Inserting scss, css, emails, and javascript files into assets..")
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
	fmt.Println("Compiling CSS from SCSS style...")
	CompileSASS()
	fmt.Println("Statup assets have been inserted")
}
