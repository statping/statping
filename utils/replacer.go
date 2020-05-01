package utils

import (
	"bytes"
	"text/template"
)

func ReplaceTemplate(tmpl string, data interface{}) string {
	buf := new(bytes.Buffer)

	tmp, err := template.New("replacement").Parse(tmpl)
	if err != nil {
		Log.Error(err)
		return err.Error()
	}

	err = tmp.Execute(buf, data)
	if err != nil {
		Log.Error(err)
		return err.Error()
	}

	return buf.String()
}
