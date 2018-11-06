// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
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

import (
	"errors"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/hunterlong/statup/utils"
	"gopkg.in/russross/blackfriday.v2"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

var (
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
	helpSrc, err := TmplBox.Bytes("help.md")
	if err != nil {
		utils.Log(4, err)
		return "error generating markdown"
	}
	output := blackfriday.Run(helpSrc)
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

	utils.Log(1, fmt.Sprintf("Compiling SASS %v into %v", scssFile, baseFile))
	command := fmt.Sprintf("%v %v %v", sassBin, scssFile, baseFile)

	utils.Log(1, fmt.Sprintf("Command: sh -c %v", command))

	testCmd := exec.Command("sh", "-c", command)

	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := testCmd.StdoutPipe()
	stderrIn, _ := testCmd.StderrPipe()
	testCmd.Start()

	go func() {
		stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
	}()

	go func() {
		stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)
	}()

	err := testCmd.Wait()
	if err != nil {
		utils.Log(3, err)
		return err
	}

	if errStdout != nil || errStderr != nil {
		utils.Log(3, fmt.Sprintf("Failed to compile assets with SASS %v", err))
		return errors.New("failed to capture stdout or stderr")
	}

	if err != nil {
		utils.Log(3, fmt.Sprintf("Failed to compile assets with SASS %v", err))
		utils.Log(3, fmt.Sprintf("bash -c %v %v %v", sassBin, scssFile, baseFile))
		return err
	}

	outStr, errStr := string(stdout), string(stderr)
	utils.Log(1, fmt.Sprintf("out: %v | error: %v", outStr, errStr))
	utils.Log(1, "SASS Compiling is complete!")
	return err
}

// UsingAssets returns true if the '/assets' folder is found in the directory
func UsingAssets(folder string) bool {
	if _, err := os.Stat(folder + "/assets"); err == nil {
		return true
	} else {
		if os.Getenv("USE_ASSETS") == "true" {
			utils.Log(1, "Environment variable USE_ASSETS was found.")
			CreateAllAssets(folder)
			err := CompileSASS(folder)
			if err != nil {
				CopyToPublic(CssBox, folder+"/css", "base.css")
				utils.Log(2, "Default 'base.css' was insert because SASS did not work.")
				return true
			}
			return true
		}
	}
	return false
}

// SaveAsset will save an asset to the '/assets/' folder.
func SaveAsset(data []byte, folder, file string) error {
	utils.Log(1, fmt.Sprintf("Saving %v/assets/%v into assets folder", folder, file))
	err := ioutil.WriteFile(folder+"/assets/"+file, data, 0744)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Failed to save %v/%v, %v", folder, file, err))
		return err
	}
	return nil
}

// OpenAsset returns a file's contents as a string
func OpenAsset(folder, file string) string {
	dat, err := ioutil.ReadFile(folder + "/assets/" + file)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Failed to open %v, %v", file, err))
		return ""
	}
	return string(dat)
}

// CreateAllAssets will dump HTML, CSS, SCSS, and JS assets into the '/assets' directory
func CreateAllAssets(folder string) error {
	utils.Log(1, fmt.Sprintf("Dump Statup assets into %v/assets", folder))
	MakePublicFolder(folder + "/assets")
	MakePublicFolder(folder + "/assets/js")
	MakePublicFolder(folder + "/assets/css")
	MakePublicFolder(folder + "/assets/scss")
	MakePublicFolder(folder + "/assets/font")
	utils.Log(1, "Inserting scss, css, and javascript files into assets folder")
	CopyAllToPublic(FontBox, "font")
	CopyAllToPublic(ScssBox, "scss")
	CopyAllToPublic(CssBox, "css")
	CopyAllToPublic(JsBox, "js")
	CopyToPublic(FontBox, folder+"/assets/font", "all.css")
	CopyToPublic(TmplBox, folder+"/assets", "robots.txt")
	CopyToPublic(TmplBox, folder+"/assets", "statup.png")
	CopyToPublic(TmplBox, folder+"/assets", "favicon.ico")
	utils.Log(1, "Compiling CSS from SCSS style...")
	err := CompileSASS(utils.Directory)
	utils.Log(1, "Statup assets have been inserted")
	return err
}

// DeleteAllAssets will delete the '/assets' folder
func DeleteAllAssets(folder string) error {
	err := os.RemoveAll(folder + "/assets")
	if err != nil {
		utils.Log(1, fmt.Sprintf("There was an issue deleting Statup Assets, %v", err))
		return err
	}
	utils.Log(1, "Statup assets have been deleted")
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
			MakePublicFolder(folder)
			return nil
		}
		file, err := box.Bytes(path)
		if err != nil {
			return nil
		}
		filePath := fmt.Sprintf("%v%v", folder, path)
		SaveAsset(file, utils.Directory, filePath)
		return nil
	})
	return err
}

// CopyToPublic will create a file from a rice Box to the '/assets' directory
func CopyToPublic(box *rice.Box, folder, file string) error {
	assetFolder := fmt.Sprintf("%v/%v", folder, file)
	utils.Log(1, fmt.Sprintf("Copying %v to %v", file, assetFolder))
	base, err := box.String(file)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Failed to copy %v to %v, %v.", file, assetFolder, err))
		return err
	}
	err = ioutil.WriteFile(assetFolder, []byte(base), 0744)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Failed to write file %v to %v, %v.", file, assetFolder, err))
		return err
	}
	return nil
}

// MakePublicFolder will create a new folder
func MakePublicFolder(folder string) error {
	utils.Log(1, fmt.Sprintf("Creating folder '%v'", folder))
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err = os.MkdirAll(folder, 0777)
		if err != nil {
			utils.Log(3, fmt.Sprintf("Failed to created %v directory, %v", folder, err))
			return err
		}
	}
	return nil
}

// copyAndCapture captures the response from a terminal command
func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}
