// +build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

const wikiUrl = "https://github.com/statping/statping.wiki"

var vue = `<template>
<div class="col-12">
<div class="row mb-4">
	{{ range .Categories }}
	<div class="col">
	<h4 class="h4 mb-2">{{ .String }}</h4>
		{{ range .Pages }}
			<a @click.prevent='tab="{{.String}}"' class="d-block mb-1 text-link" href="#">{{.String}}</a>
		{{end}}
	</div>
	{{end}}
</div>

<div class="col-12" v-if='tab === "Home"'>
	<div v-pre>
		{{html .Home.Data}}
	</div>
</div>
{{ range .Pages }}
	<div class="col-12" v-if='tab === "{{.String}}"'>
		<h1 class="h1 mt-5 mb-5 text-muted">{{ .String }}</h1>
		<span class="spacer"></span>
		<div v-pre>
				{{html .Data}}
		</div>
	</div>
{{end}}

<div class="col-12 shadow-md mt-5">
	<div class="text-dim" v-pre>
		{{html .Footer.Data}}
	</div>
</div>

<div class="text-center small text-dim" v-pre>
Automatically generated from Statping's Wiki on {{.CreatedAt}}
</div>

</div>
</template>

<script>
export default {
  name: 'Help',
	data () {
		  return {
			  tab: "Home",
		  }
	  }
}
</script>

<style scoped>
IMG {
	max-width: 80%;
	alignment: center;
	display: block;
}
</style>
`

var temp *template.Template

type Category struct {
	String string
	Pages  []*Page
}

type Page struct {
	String string
	Data   string
}

type Render struct {
	Categories []*Category
	Pages      []*Page
	Home       *Page
	Footer     *Page
	CreatedAt  time.Time
}

func main() {
	fmt.Println("RUNNING: ./source/generate_help.go")
	fmt.Println("\n\nGenerating Help.vue from Statping's Wiki")
	fmt.Println("Cloning ", wikiUrl)
	cmd := exec.Command("git", "clone", wikiUrl)
	cmd.Start()
	cmd.Wait()

	fmt.Println("Generating Help view from Wiki")

	d, _ := ioutil.ReadFile("statping.wiki/_Sidebar.md")

	var cats []*Category
	var pages []*Page
	scanner := bufio.NewScanner(strings.NewReader(string(d)))
	var thisCategory *Category
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}
		if txt[0:1] == "#" {
			newCate := &Category{
				String: txt[2:len(txt)],
			}
			if txt[2:len(txt)] == "Contact" || txt[2:len(txt)] == "Badges" {
				continue
			}
			thisCategory = newCate
			cats = append(cats, newCate)
		}
		if txt[0:2] == "[[" {
			file := "statping.wiki/" + txt[2:len(txt)-2] + ".md"
			file = strings.ReplaceAll(file, " ", "-")
			page := &Page{
				String: txt[2 : len(txt)-2],
				Data:   open(file),
			}
			pages = append(pages, page)
			thisCategory.Pages = append(thisCategory.Pages, page)
		}
	}

	home := &Page{
		String: "Home",
		Data:   open("statping.wiki/Home.md"),
	}

	footer := &Page{
		String: "Footer",
		Data:   open("statping.wiki/_Footer.md"),
	}

	w := bytes.NewBufferString("")
	temp = template.New("wiki")
	temp.Funcs(template.FuncMap{
		"html": func(val string) template.HTML {
			return template.HTML(val)
		},
		"fake": func(val string) template.HTML {
			return template.HTML(`{{` + val + `}}`)
		},
	})
	temp, _ = temp.Parse(vue)
	temp.ExecuteTemplate(w, "wiki", Render{Categories: cats, Pages: pages, Home: home, Footer: footer, CreatedAt: time.Now().UTC()})

	fmt.Println("Saving wiki page to: ./frontend/src/pages/Home.vue")
	ioutil.WriteFile("../frontend/src/pages/Help.vue", w.Bytes(), os.FileMode(0755))

	fmt.Println("Deleting statping wiki repo")
	os.RemoveAll("statping.wiki")
}

func open(filename string) string {
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	d, _ := ioutil.ReadFile(filename)
	output := markdown.ToHTML(d, nil, renderer)
	return string(output)
}
