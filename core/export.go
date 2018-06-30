package core

import (
	"bytes"
	"github.com/hunterlong/statup/utils"
	"html/template"
	"io/ioutil"
)

func ExportIndexHTML() string {
	CoreApp.OfflineAssets = true
	//out := index{*CoreApp, CoreApp.Services}
	nav, _ := TmplBox.String("nav.html")
	footer, _ := TmplBox.String("footer.html")
	render, err := TmplBox.String("index.html")
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
			return "version here"
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
