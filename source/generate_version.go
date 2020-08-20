// +build ignore

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const replace = `this\.version = "[0-9]\.[0-9]{2}\.[0-9]{2}";`
const replaceCommit = `this\.commit = \"[a-z0-9]{40}\"\;`

func main() {
	fmt.Println("RUNNING: ./source/generate_version.go")
	version, _ := ioutil.ReadFile("../version.txt")
	apiJsFile, _ := ioutil.ReadFile("../frontend/src/API.js")

	w := bytes.NewBuffer(nil)
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Stdout = w
	cmd.Run()
	gitCommit := strings.TrimSpace(w.String())

	fmt.Println("git commit: ", gitCommit)

	replaceWith := `this.version = "` + strings.TrimSpace(string(version)) + `";`
	replaceCommitWith := `this.commit = "` + gitCommit + `";`

	vRex := regexp.MustCompile(replace)
	newApiFile := vRex.ReplaceAllString(string(apiJsFile), replaceWith)
	cRex := regexp.MustCompile(replaceCommit)
	newApiFile = cRex.ReplaceAllString(newApiFile, replaceCommitWith)

	fmt.Printf("Setting version %s to frontend/src/API.js\n", string(version))
	ioutil.WriteFile("../frontend/src/API.js", []byte(newApiFile), os.FileMode(0755))
}
