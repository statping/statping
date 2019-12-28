// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package source

//go:generate go run generate_wiki.go

import (
	"errors"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/hunterlong/statping/utils"
	"github.com/russross/blackfriday/v2"
	"io/ioutil"
	"os"
)

var (
	log = utils.Log.WithField("type", "source")
	CssBox  *rice.Box // CSS files from the 'source/css' directory, this will be loaded into '/assets/css'
	ScssBox *rice.Box // SCSS files from the 'source/scss' directory, this will be loaded into '/assets/scss'
	JsBox   *rice.Box // JS files from the 'source/js' directory, this will be loaded into '/assets/js'
	TmplBox *rice.Box // HTML and other small files from the 'source/tmpl' directory, this will be loaded into '/assets'
	FontBox *rice.Box // HTML and other small files from the 'source/tmpl' directory, this will be loaded into '/assets'
)

// Assets will load the Rice boxes containing the CSS, SCSS, JS, and HTML files.
func Assets() {
	CssBox = rice.MustFindBox("css")
	ScssBox = rice.MustFindBox("scss")
	JsBox = rice.MustFindBox("js")
	TmplBox = rice.MustFindBox("tmpl")
	FontBox = rice.MustFindBox("font")
}

// HelpMarkdown will return the Markdown of help.md into HTML
func HelpMarkdown() string {
	output := blackfriday.Run(CompiledWiki)
	return string(output)
}

// CompileSASS will attempt to compile the SASS files into CSS
func CompileSASS(folder string) error {
	sassBin := os.Getenv("SASS")
	if sassBin == "" {
		sassBin = "sass"
	}

	scssFile := fmt.Sprintf("%v/%v", folder, "assets/scss/base.scss")
	baseFile := fmt.Sprintf("%v/%v", folder, "assets/css/base.css")

	log.Infoln(fmt.Sprintf("Compiling SASS %v into %v", scssFile, baseFile))
	command := fmt.Sprintf("%v %v %v", sassBin, scssFile, baseFile)

	stdout, stderr, err := utils.Command(command)

	if stdout != "" || stderr != "" {
		log.Errorln(fmt.Sprintf("Failed to compile assets with SASS %v", err))
		return errors.New("failed to capture stdout or stderr")
	}

	if err != nil {
		log.Errorln(fmt.Sprintf("Failed to compile assets with SASS %v", err))
		log.Errorln(fmt.Sprintf("sh -c %v", command))
		return err
	}

	log.Infoln(fmt.Sprintf("out: %v | error: %v", stdout, stderr))
	log.Infoln("SASS Compiling is complete!")
	return nil
}

// UsingAssets returns true if the '/assets' folder is found in the directory
func UsingAssets(folder string) bool {
	if _, err := os.Stat(folder + "/assets"); err == nil {
		return true
	} else {
		if os.Getenv("USE_ASSETS") == "true" {
			log.Infoln("Environment variable USE_ASSETS was found.")
			CreateAllAssets(folder)
			err := CompileSASS(folder)
			if err != nil {
				CopyToPublic(CssBox, folder+"/css", "base.css")
				log.Warnln("Default 'base.css' was insert because SASS did not work.")
				return true
			}
			return true
		}
	}
	return false
}

// SaveAsset will save an asset to the '/assets/' folder.
func SaveAsset(data []byte, folder, file string) error {
	location := folder + "/assets/" + file
	log.Infoln(fmt.Sprintf("Saving %v", location))
	err := utils.SaveFile(location, data)
	if err != nil {
		log.Errorln(fmt.Sprintf("Failed to save %v, %v", location, err))
		return err
	}
	return nil
}

// OpenAsset returns a file's contents as a string
func OpenAsset(folder, file string) string {
	dat, err := ioutil.ReadFile(folder + "/assets/" + file)
	if err != nil {
		log.Errorln(fmt.Sprintf("Failed to open %v, %v", file, err))
		return ""
	}
	return string(dat)
}

// CreateAllAssets will dump HTML, CSS, SCSS, and JS assets into the '/assets' directory
func CreateAllAssets(folder string) error {
	log.Infoln(fmt.Sprintf("Dump Statping assets into %v/assets", folder))
	MakePublicFolder(folder + "/assets")
	MakePublicFolder(folder + "/assets/js")
	MakePublicFolder(folder + "/assets/css")
	MakePublicFolder(folder + "/assets/scss")
	MakePublicFolder(folder + "/assets/font")
	MakePublicFolder(folder + "/assets/files")
	log.Infoln("Inserting scss, css, and javascript files into assets folder")
	CopyAllToPublic(ScssBox, "scss")
	CopyAllToPublic(FontBox, "font")
	CopyAllToPublic(CssBox, "css")
	CopyAllToPublic(JsBox, "js")
	CopyToPublic(FontBox, folder+"/assets/font", "all.css")
	CopyToPublic(TmplBox, folder+"/assets", "robots.txt")
	CopyToPublic(TmplBox, folder+"/assets", "banner.png")
	CopyToPublic(TmplBox, folder+"/assets", "favicon.ico")
	CopyToPublic(TmplBox, folder+"/assets/files", "swagger.json")
	CopyToPublic(TmplBox, folder+"/assets/files", "postman.json")
	CopyToPublic(TmplBox, folder+"/assets/files", "grafana.json")
	log.Infoln("Compiling CSS from SCSS style...")
	err := CompileSASS(utils.Directory)
	log.Infoln("Statping assets have been inserted")
	return err
}

// DeleteAllAssets will delete the '/assets' folder
func DeleteAllAssets(folder string) error {
	err := os.RemoveAll(folder + "/assets")
	if err != nil {
		log.Infoln(fmt.Sprintf("There was an issue deleting Statping Assets, %v", err))
		return err
	}
	log.Infoln("Statping assets have been deleted")
	return err
}

// CopyAllToPublic will copy all the files in a rice box into a local folder
func CopyAllToPublic(box *rice.Box, folder string) error {
	err := box.Walk("/", func(path string, info os.FileInfo, err error) error {
		if info.Name() == "" {
			return nil
		}
		if info.IsDir() {
			folder := fmt.Sprintf("%v/assets/%v/%v", utils.Directory, folder, info.Name())
			return MakePublicFolder(folder)
		}
		file, err := box.Bytes(path)
		if err != nil {
			return err
		}
		filePath := fmt.Sprintf("%v/%v", folder, path)
		return SaveAsset(file, utils.Directory, filePath)
	})
	return err
}

// CopyToPublic will create a file from a rice Box to the '/assets' directory
func CopyToPublic(box *rice.Box, folder, file string) error {
	assetFolder := fmt.Sprintf("%v/%v", folder, file)
	log.Infoln(fmt.Sprintf("Copying %v to %v", file, assetFolder))
	base, err := box.String(file)
	if err != nil {
		log.Errorln(fmt.Sprintf("Failed to copy %v to %v, %v.", file, assetFolder, err))
		return err
	}
	err = ioutil.WriteFile(assetFolder, []byte(base), 0744)
	if err != nil {
		log.Errorln(fmt.Sprintf("Failed to write file %v to %v, %v.", file, assetFolder, err))
		return err
	}
	return nil
}

// MakePublicFolder will create a new folder
func MakePublicFolder(folder string) error {
	log.Infoln(fmt.Sprintf("Creating folder '%v'", folder))
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err = os.MkdirAll(folder, 0777)
		if err != nil {
			log.Errorln(fmt.Sprintf("Failed to created %v directory, %v", folder, err))
			return err
		}
	}
	return nil
}
