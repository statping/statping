package main

import (
	"bytes"
	"fmt"
	"github.com/hunterlong/statup/log"
	"html/template"
	"io/ioutil"
	"strings"
)

func (f *Failure) ParseError() string {
	err := strings.Contains(f.Issue, "operation timed out")
	if err {
		return fmt.Sprintf("HTTP Request Timed Out")
	}
	err = strings.Contains(f.Issue, "x509: certificate is valid")
	if err {
		return fmt.Sprintf("SSL Certificate invalid")
	}
	err = strings.Contains(f.Issue, "no such host")
	if err {
		return fmt.Sprintf("Domain is offline or not found")
	}
	return f.Issue
}

func ExportIndexHTML() string {
	core.OfflineAssets = true
	out := index{*core, services}
	nav, _ := tmplBox.String("nav.html")
	footer, _ := tmplBox.String("footer.html")
	render, err := tmplBox.String("index.html")
	if err != nil {
		log.Send(3, err)
	}
	t := template.New("message")
	t.Funcs(template.FuncMap{
		"js": func(html string) template.JS {
			return template.JS(html)
		},
		"safe": func(html string) template.HTML {
			return template.HTML(html)
		},
		"VERSION": func() string {
			return VERSION
		},
		"underscore": func(html string) string {
			return UnderScoreString(html)
		},
	})
	t, _ = t.Parse(nav)
	t, _ = t.Parse(footer)
	t.Parse(render)

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, out); err != nil {
		log.Send(3, err)
	}
	result := tpl.String()

	fmt.Println(result)
	return result
}

func SaveFile(filename string, data []byte) error {
	err := ioutil.WriteFile(filename, data, 0644)
	return err
}
