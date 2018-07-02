package core

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/hunterlong/statup/utils"
	"io/ioutil"
	"os"
	"os/exec"
)

func CopyToPublic(box *rice.Box, folder, file string) {
	assetFolder := fmt.Sprintf("assets/%v/%v", folder, file)
	if folder == "" {
		assetFolder = fmt.Sprintf("assets/%v", file)
	}
	utils.Log(1, fmt.Sprintf("Copying %v to %v", file, assetFolder))
	base, err := box.String(file)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Failed to copy %v to %v, %v.", file, assetFolder, err))
	}
	err = ioutil.WriteFile(assetFolder, []byte(base), 0644)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Failed to write file %v to %v, %v.", file, assetFolder, err))
	}
}

func MakePublicFolder(folder string) {
	utils.Log(1, fmt.Sprintf("Creating folder '%v' in current directory...", folder))
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err = os.MkdirAll(folder, 0755)
		if err != nil {
			utils.Log(3, fmt.Sprintf("Failed to created %v directory, %v", folder, err))
		}
	}
}

func CompileSASS() error {
	sassBin := os.Getenv("SASS")
	shell := os.Getenv("CMD_FILE")
	utils.Log(1, fmt.Sprintf("Compiling SASS into /css/base.css..."))
	command := fmt.Sprintf("%v %v %v", sassBin, "assets/scss/base.scss", "assets/css/base.css")
	testCmd := exec.Command(shell, command)
	_, err := testCmd.Output()
	if err != nil {
		utils.Log(3, fmt.Sprintf("Failed to compile assets with SASS %v", err))
		utils.Log(3, fmt.Sprintf("SASS: %v | CMD_FILE:%v", sassBin, shell))
		return err
	}
	utils.Log(1, "SASS Compiling is complete!")
	return err
}

func HasAssets() bool {
	if _, err := os.Stat("assets"); err == nil {
		utils.Log(1, "Assets folder was found!")
		UsingAssets = true
		return true
	} else {
		assetEnv := os.Getenv("USE_ASSETS")
		if assetEnv == "true" {
			utils.Log(1, "Environment variable USE_ASSETS was found.")
			CreateAllAssets()
			UsingAssets = true
			return true
		}
	}
	return false
}

func SaveAsset(data, file string) {
	utils.Log(1, fmt.Sprintf("Saving %v into assets folder", file))
	err := ioutil.WriteFile("assets/"+file, []byte(data), 0644)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Failed to save %v, %v", file, err))
	}
}

func OpenAsset(file string) string {
	dat, err := ioutil.ReadFile("assets/" + file)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Failed to open %v, %v", file, err))
		return ""
	}
	return string(dat)
}

func CreateAllAssets() {
	utils.Log(1, "Dump Statup assets into current directory...")
	MakePublicFolder("assets")
	MakePublicFolder("assets/js")
	MakePublicFolder("assets/css")
	MakePublicFolder("assets/scss")
	MakePublicFolder("assets/emails")
	utils.Log(1, "Inserting scss, css, emails, and javascript files into assets..")
	CopyToPublic(ScssBox, "scss", "base.scss")
	CopyToPublic(ScssBox, "scss", "variables.scss")
	CopyToPublic(EmailBox, "emails", "message.html")
	CopyToPublic(EmailBox, "emails", "failure.html")
	CopyToPublic(CssBox, "css", "bootstrap.min.css")
	CopyToPublic(JsBox, "js", "bootstrap.min.js")
	CopyToPublic(JsBox, "js", "Chart.bundle.min.js")
	CopyToPublic(JsBox, "js", "jquery-3.3.1.slim.min.js")
	CopyToPublic(JsBox, "js", "main.js")
	CopyToPublic(JsBox, "js", "setup.js")
	CopyToPublic(JsBox, "js", "setup.js")
	CopyToPublic(TmplBox, "", "robots.txt")
	CopyToPublic(TmplBox, "", "favicon.ico")
	utils.Log(1, "Compiling CSS from SCSS style...")
	err := CompileSASS()
	if err != nil {
		CopyToPublic(CssBox, "css", "base.css")
		utils.Log(2, "Default 'base.css' was insert because SASS did not work.")
		return
	}
	utils.Log(1, "Statup assets have been inserted")
}
