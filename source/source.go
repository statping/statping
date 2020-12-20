package source

//go:generate go run generate_help.go
//go:generate go run generate_version.go
//go:generate go run generate_languages.go

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/pkg/errors"
	"github.com/statping/statping/utils"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	log           = utils.Log.WithField("type", "source")
	TmplBox       *rice.Box // HTML and other small files from the 'source/dist' directory, this will be loaded into '/assets'
	RequiredFiles = []string{
		"css/style.css",
		"css/style.css.gz",
		"css/index.css",
		"scss/index.scss",
		"scss/base.scss",
		"scss/forms.scss",
		"scss/layout.scss",
		"scss/mixin.scss",
		"scss/mobile.scss",
		"scss/variables.scss",
		"robots.txt",
	}
)

// Assets will load the Rice boxes containing the CSS, SCSS, and HTML files.
func Assets() error {
	if utils.Params.GetBool("DISABLE_HTTP") {
		return nil
	}
	var err error
	TmplBox, err = rice.FindBox("dist")
	if err != nil {
		return err
	}

	return err
}

func scssRendered(name string) string {
	spl := strings.Split(name, "/")
	path := spl[:len(spl)-2]
	file := spl[len(spl)-1]
	splFile := strings.Split(file, ".")
	return filepath.Join(strings.Join(path, "/"), "css", splFile[len(splFile)-2]+".css")
}

// CompileSASS will attempt to compile the SASS files into CSS
func CompileSASS() error {
	sassBin := utils.Params.GetString("SASS")
	if sassBin == "" {
		bin, err := exec.LookPath("sass")
		if err != nil {
			log.Warnf("could not find sass executable in PATH: %s", err)
			return err
		}
		sassBin = bin
	}

	scssFile := filepath.Join(utils.Params.GetString("STATPING_DIR"), "assets", "scss", "index.scss")
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

	log.Infoln("SASS Compiling is complete!")
	return nil
}

// UsingAssets returns true if the '/assets' folder is found in the directory
func UsingAssets(folder string) bool {
	if _, err := os.Stat(folder + "/assets"); err == nil {
		return true
	} else {
		if utils.Params.GetBool("USE_ASSETS") {
			log.Infoln("Environment variable USE_ASSETS was found.")
			if err := CreateAllAssets(folder); err != nil {
				log.Warnln(err)
			}
			if err := CompileSASS(); err != nil {
				//CopyToPublic(CssBox, folder+"/css", "base.css")
				log.Warn(errors.Wrap(err, "Default 'base.css' was insert because SASS did not work."))
				return true
			}
			return true
		}
	}
	return false
}

// SaveAsset will save an asset to the '/assets/' folder.
func SaveAsset(data []byte, path string) error {
	path = filepath.Join(utils.Directory, "assets", path)
	err := utils.SaveFile(path, data)
	if err != nil {
		log.Errorln(fmt.Sprintf("Failed to save %v, %v", path, err))
		return err
	}
	return nil
}

// OpenAsset returns a file's contents as a string
func OpenAsset(path string) string {
	path = filepath.Join(utils.Directory, "assets", path)
	data, err := utils.OpenFile(path)
	if err != nil {
		log.Errorln(fmt.Sprintf("Failed to open %s, %v", path, err))
		return ""
	}
	return data
}

// CreateAllAssets will dump HTML, CSS, and SCSS assets into the '/assets' directory
func CreateAllAssets(folder string) error {
	log.Infoln(fmt.Sprintf("Dump Statping assets into %s/assets", folder))
	fp := filepath.Join

	if err := MakePublicFolder(fp(folder, "/assets")); err != nil {
		return err
	}
	if err := MakePublicFolder(fp(folder, "assets", "css")); err != nil {
		return err
	}
	if err := MakePublicFolder(fp(folder, "assets", "scss")); err != nil {
		return err
	}
	log.Infoln("Inserting scss, and css files into assets folder")

	if err := CopyAllToPublic(TmplBox); err != nil {
		log.Errorln(err)
		return errors.Wrap(err, "copying all to public")
	}

	if err := CopyToPublic(TmplBox, "", "robots.txt"); err != nil {
		return err
	}
	log.Infoln("Compiling CSS from SCSS style...")
	if err := CompileSASS(); err != nil {
		return err
	}
	log.Infoln("Statping assets have been inserted")
	return nil
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

	return box.Walk("/", func(path string, info os.FileInfo, err error) error {
		if info.Name() == "" {
			return nil
		}
		if exclude[info.Name()] {
			return nil
		}
		if strings.Contains(path, "/js") {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		file, err := box.Bytes(path)
		if err != nil {
			return err
		}
		return SaveAsset(file, path)
	})
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
	log.Infoln(fmt.Sprintf("Creating folder '%s'", folder))
	if !utils.FolderExists(folder) {
		err := utils.CreateDirectory(folder)
		if err != nil {
			log.Errorln(fmt.Sprintf("Failed to created %s directory, %v", folder, err))
			return err
		}
	}
	return nil
}
