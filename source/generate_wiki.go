// +build ignore

package main

import (
	"github.com/statping/statping/utils"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

const (
	wikiUrl = "http://assets.statping.com/wiki/"
)

func replaceDash(s string) string {
	return strings.Replace(s, "-", " ", -1)
}

func main() {
	var compiled string

	utils.InitLogs()
	utils.Command("git clone https://github.com/statping/statping.wiki.git")

	pages := []string{"Types-of-Monitoring", "Features", "Start-Statping", "Linux", "Mac", "Windows", "AWS-EC2", "Docker", "Mobile-App", "Heroku", "API", "Makefile",
		"Notifiers", "Notifier-Events", "Notifier-Example", "Prometheus-Exporter", "SSL", "Config-with-.env-File", "Static-Export", "Statping-Plugins", "Statuper", "Build-and-Test", "Contributing", "PGP-Signature", "Testing", "Deployment"}
	newPages := map[string]string{}

	for k, v := range pages {
		compiled += "<a class=\"scrollclick\" href=\"#\" data-id=\"page_" + utils.ToString(k) + "\">" + replaceDash(v) + "</a><br>"
	}

	for k, v := range pages {
		sc, _ := ioutil.ReadFile("statping.wiki/" + v + ".md")
		newPages[v] = string(sc)
		compiled += "\n\n<div class=\"mt-5\" id=\"page_" + utils.ToString(k) + "\"><h1>" + replaceDash(v) + "</h1></div>\n" + string(sc)
	}

	utils.DeleteDirectory("./statping.wiki")
	utils.DeleteDirectory("./logs")

	f, err := os.Create("../frontend/src/pages/Help.vue")
	die(err)
	defer f.Close()

	packageTemplate.Execute(f, struct {
		Timestamp time.Time
		URL       string
		Compiled  string
		Pages     map[string]string
	}{
		Timestamp: utils.Now(),
		URL:       "ok",
		Compiled:  compiled,
		Pages:     newPages,
	})
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var packageTemplate = template.Must(template.New("").Parse(`<template>
    <div class="col-12 bg-white p-4" v-html={{printf "%q" .Compiled}}></div>
</template>

<script>
export default {
  name: 'Help',
}
</script>`))
