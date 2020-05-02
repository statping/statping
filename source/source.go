package source

//go:generate go run generate_wiki.go

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/pkg/errors"
	"github.com/russross/blackfriday/v2"
	"github.com/statping/statping/utils"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	log         = utils.Log.WithField("type", "source")
	TmplBox     *rice.Box // HTML and other small files from the 'source/tmpl' directory, this will be loaded into '/assets'
	DefaultScss = []string{"scss/main.scss", "scss/base.scss", "scss/mobile.scss"}
)

// Assets will load the Rice boxes containing the CSS, SCSS, JS, and HTML files.
func Assets() error {
	var err error
	TmplBox, err = rice.FindBox("dist")
	return err
}

// HelpMarkdown will return the Markdown of help.md into HTML
func HelpMarkdown() string {
	output := blackfriday.Run(CompiledWiki)
	return string(output)
}

func scssRendered(name string) string {
	spl := strings.Split(name, "/")
	path := spl[:len(spl)-2]
	file := spl[len(spl)-1]
	splFile := strings.Split(file, ".")
	return fmt.Sprintf("%s/css/%s.css", strings.Join(path, "/"), splFile[len(splFile)-2])
}

// CompileSASS will attempt to compile the SASS files into CSS
func CompileSASS(files ...string) error {
	sassBin := utils.Params.GetString("SASS")
	if sassBin == "" {
		bin, err := exec.LookPath("sass")
		if err != nil {
			log.Warnf("could not find sass executable in PATH: %s", err)
			return err
		}
		sassBin = bin
	}

	for _, file := range files {
		scssFile := fmt.Sprintf("%v/assets/%v", utils.Params.GetString("STATPING_DIR"), file)

		log.Infoln(fmt.Sprintf("Compiling SASS %v into %v", scssFile, scssRendered(scssFile)))

		stdout, stderr, err := utils.Command(sassBin, scssFile, scssRendered(scssFile))

		if err != nil {
			log.Errorln(fmt.Sprintf("Failed to compile assets with SASS %v", err))
			log.Errorln(fmt.Sprintf("%s %s %s", sassBin, scssFile, scssRendered(scssFile)))
			return errors.Wrapf(err, "failed to compile assets, %s %s %s", err, stdout, stderr)
		}

		if stdout != "" || stderr != "" {
			log.Infoln(fmt.Sprintf("out: %v | error: %v", stdout, stderr))
		}
	}
	log.Infoln("SASS Compiling is complete!")
	return nil
}

// UsingAssets returns true if the '/assets' folder is found in the directory
func UsingAssets(folder string) bool {
	if _, err := os.Stat(folder + "/assets"); err == nil {
		return true
	} else {
		useAssets := utils.Params.GetBool("USE_ASSETS")

		if useAssets {
			log.Infoln("Environment variable USE_ASSETS was found.")
			if err := CreateAllAssets(folder); err != nil {
				log.Warnln(err)
			}
			err := CompileSASS(DefaultScss...)
			if err != nil {
				//CopyToPublic(CssBox, folder+"/css", "base.css")
				log.Warnln("Default 'base.css' was insert because SASS did not work.")
				return true
			}
			return true
		}
	}
	return false
}

// SaveAsset will save an asset to the '/assets/' folder.
func SaveAsset(data []byte, path string) error {
	path = fmt.Sprintf("%s/assets/%s", utils.Directory, path)
	err := utils.SaveFile(path, data)
	if err != nil {
		log.Errorln(fmt.Sprintf("Failed to save %v, %v", path, err))
		return err
	}
	return nil
}

// OpenAsset returns a file's contents as a string
func OpenAsset(path string) string {
	path = fmt.Sprintf("%s/assets/%s", utils.Directory, path)
	data, err := utils.OpenFile(path)
	if err != nil {
		log.Errorln(fmt.Sprintf("Failed to open %v, %v", path, err))
		return ""
	}
	return data
}

// CreateAllAssets will dump HTML, CSS, SCSS, and JS assets into the '/assets' directory
func CreateAllAssets(folder string) error {
	log.Infoln(fmt.Sprintf("Dump Statping assets into %v/assets", folder))
	fp := filepath.Join

	if err := MakePublicFolder(fp(folder, "/assets")); err != nil {
		return err
	}
	if err := MakePublicFolder(fp(folder, "assets", "js")); err != nil {
		return err
	}
	if err := MakePublicFolder(fp(folder, "assets", "css")); err != nil {
		return err
	}
	if err := MakePublicFolder(fp(folder, "assets", "scss")); err != nil {
		return err
	}
	log.Infoln("Inserting scss, css, and javascript files into assets folder")

	if err := CopyAllToPublic(TmplBox); err != nil {
		log.Errorln(err)
		return errors.Wrap(err, "copying all to public")
	}

	CopyToPublic(TmplBox, "", "robots.txt")
	CopyToPublic(TmplBox, "", "banner.png")
	CopyToPublic(TmplBox, "", "favicon.ico")
	log.Infoln("Compiling CSS from SCSS style...")
	err := CompileSASS(DefaultScss...)
	log.Infoln("Statping assets have been inserted")
	return err
}

// DeleteAllAssets will delete the '/assets' folder
func DeleteAllAssets(folder string) error {
	err := utils.DeleteDirectory(folder + "/assets")
	if err != nil {
		log.Infoln(fmt.Sprintf("There was an issue deleting Statping Assets, %v", err))
		return err
	}
	log.Infoln("Statping assets have been deleted")
	return err
}

// CopyAllToPublic will copy all the files in a rice box into a local folder
func CopyAllToPublic(box *rice.Box) error {

	exclude := map[string]bool{
		"base.gohtml": true,
		"index.html":  true,
	}

	err := box.Walk("/", func(path string, info os.FileInfo, err error) error {
		if info.Name() == "" {
			return nil
		}
		if exclude[info.Name()] {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		utils.Log.Infoln(path)
		file, err := box.Bytes(path)
		if err != nil {
			return err
		}
		return SaveAsset(file, path)
	})
	return err
}

// CopyToPublic will create a file from a rice Box to the '/assets' directory
func CopyToPublic(box *rice.Box, path, file string) error {
	assetPath := fmt.Sprintf("%v/assets/%v/%v", utils.Directory, path, file)
	if path == "" {
		assetPath = fmt.Sprintf("%v/assets/%v", utils.Directory, file)
	}
	log.Infoln(fmt.Sprintf("Copying %v to %v", file, assetPath))
	base, err := box.String(file)
	if err != nil {
		log.Errorln(fmt.Sprintf("Failed to copy %v to %v, %v.", file, assetPath, err))
		return err
	}
	err = utils.SaveFile(assetPath, []byte(base))
	if err != nil {
		log.Errorln(fmt.Sprintf("Failed to write file %v to %v, %v.", file, assetPath, err))
		return err
	}
	return nil
}

// MakePublicFolder will create a new folder
func MakePublicFolder(folder string) error {
	log.Infoln(fmt.Sprintf("Creating folder '%v'", folder))
	if !utils.FolderExists(folder) {
		err := utils.CreateDirectory(folder)
		if err != nil {
			log.Errorln(fmt.Sprintf("Failed to created %v directory, %v", folder, err))
			return err
		}
	}
	return nil
}
