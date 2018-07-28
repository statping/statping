package core

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/hunterlong/statup/utils"
	"io/ioutil"
	"os"
	"os/exec"
)

func RenderBoxes() {
	SqlBox = rice.MustFindBox("../source/sql")
	CssBox = rice.MustFindBox("../source/css")
	ScssBox = rice.MustFindBox("../source/scss")
	JsBox = rice.MustFindBox("../source/js")
	TmplBox = rice.MustFindBox("../source/tmpl")
}

func CopyToPublic(box *rice.Box, folder, file string) {
	assetFolder := fmt.Sprintf("%v/%v", folder, file)
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

func CompileSASS(folder string) error {
	sassBin := os.Getenv("SASS")
	shell := os.Getenv("CMD_FILE")

	scssFile := fmt.Sprintf("%v/%v", folder, "assets/scss/base.scss")
	baseFile := fmt.Sprintf("%v/%v", folder, "assets/css/base.css")

	utils.Log(1, fmt.Sprintf("Compiling SASS %v into %v", scssFile, baseFile))
	command := fmt.Sprintf("%v %v %v", sassBin, scssFile, baseFile)
	testCmd := exec.Command(shell, command)
	_, err := testCmd.Output()
	if err != nil {
		utils.Log(3, fmt.Sprintf("Failed to compile assets with SASS %v", err))
		utils.Log(3, fmt.Sprintf("%v %v %v", sassBin, scssFile, baseFile))
		return err
	}
	utils.Log(1, "SASS Compiling is complete!")
	return err
}

func HasAssets(folder string) bool {
	if _, err := os.Stat(folder + "/assets"); err == nil {
		utils.Log(1, "Assets folder was found!")
		UsingAssets = true
		return true
	} else {
		assetEnv := os.Getenv("USE_ASSETS")
		if assetEnv == "true" {
			utils.Log(1, "Environment variable USE_ASSETS was found.")
			CreateAllAssets(folder)
			err := CompileSASS(folder)
			if err != nil {
				CopyToPublic(CssBox, folder+"/css", "base.css")
				utils.Log(2, "Default 'base.css' was insert because SASS did not work.")
				return true
			}
			UsingAssets = true
			return true
		}
	}
	return false
}

func SaveAsset(data, folder, file string) {
	utils.Log(1, fmt.Sprintf("Saving %v/%v into assets folder", folder, file))
	err := ioutil.WriteFile(folder+"/assets/"+file, []byte(data), 0644)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Failed to save %v/%v, %v", folder, file, err))
	}
}

func OpenAsset(folder, file string) string {
	dat, err := ioutil.ReadFile(folder + "/assets/" + file)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Failed to open %v, %v", file, err))
		return ""
	}
	return string(dat)
}

func CreateAllAssets(folder string) error {
	utils.Log(1, fmt.Sprintf("Dump Statup assets into %v/assets", folder))
	MakePublicFolder(folder + "/assets")
	MakePublicFolder(folder + "/assets/js")
	MakePublicFolder(folder + "/assets/css")
	MakePublicFolder(folder + "/assets/scss")
	MakePublicFolder(folder + "/assets/emails")
	utils.Log(1, "Inserting scss, css, emails, and javascript files into assets..")
	CopyToPublic(ScssBox, folder+"/assets/scss", "base.scss")
	CopyToPublic(ScssBox, folder+"/assets/scss", "variables.scss")
	CopyToPublic(CssBox, folder+"/assets/css", "bootstrap.min.css")
	CopyToPublic(JsBox, folder+"/assets/js", "bootstrap.min.js")
	CopyToPublic(JsBox, folder+"/assets/js", "Chart.bundle.min.js")
	CopyToPublic(JsBox, folder+"/assets/js", "jquery-3.3.1.slim.min.js")
	CopyToPublic(JsBox, folder+"/assets/js", "main.js")
	CopyToPublic(JsBox, folder+"/assets/js", "setup.js")
	CopyToPublic(TmplBox, folder+"/assets/", "robots.txt")
	CopyToPublic(TmplBox, folder+"/assets/", "favicon.ico")
	utils.Log(1, "Compiling CSS from SCSS style...")
	err := utils.Log(1, "Statup assets have been inserted")
	return err
}

func DeleteAllAssets(folder string) error {
	err := os.RemoveAll(folder + "/assets")
	if err != nil {
		utils.Log(1, fmt.Sprintf("There was an issue deleting Statup Assets, %v", err))
		return err
	}
	utils.Log(1, "Statup assets have been deleted")
	return err
}
