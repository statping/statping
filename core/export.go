package core

import (
	"bytes"
	"github.com/hunterlong/statup/source"
	"github.com/hunterlong/statup/utils"
	"html/template"
	"io/ioutil"
)

func ExportIndexHTML() string {
	CoreApp.UseCdn = true
	//out := index{*CoreApp, CoreApp.Services}
	nav, _ := source.TmplBox.String("nav.html")
	footer, _ := source.TmplBox.String("footer.html")
	render, err := source.TmplBox.String("index.html")
	if err != nil {
		utils.Log(3, err)
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
			return utils.UnderScoreString(html)
		},
	})
	t, _ = t.Parse(nav)
	t, _ = t.Parse(footer)
	t.Parse(render)
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, nil); err != nil {
		utils.Log(3, err)
	}
	result := tpl.String()
	return result
}

func SaveFile(filename string, data []byte) error {
	err := ioutil.WriteFile(filename, data, 0644)
	return err
}
