package utils

import (
	"bytes"
	"text/template"
)

func ReplaceTemplate(tmpl string, data interface{}) string {
	buf := new(bytes.Buffer)

	slackTemp, err := template.New("replacement").Parse(tmpl)
	if err != nil {
		Log.Error(err)
		return err.Error()
	}

	err = slackTemp.Execute(buf, data)
	if err != nil {
		Log.Error(err)
		return err.Error()
	}

	return buf.String()
}
