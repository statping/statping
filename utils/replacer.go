package utils

import (
	"bytes"
	"fmt"
	"text/template"
)

func ReplaceTemplate(tmpl string, data interface{}) string {
	buf := new(bytes.Buffer)
	var varStr string
	switch fmt.Sprintf("%T", data) {
	case "*services.Service":
		varStr = "Service"
	case "*failures.Failure":
		varStr = "Failure"
	default:
		varStr = "Object"
	}
	injectVars := make(map[string]interface{})
	injectVars[varStr] = data
	slackTemp, err := template.New("replacement").Parse(tmpl)
	if err != nil {
		Log.Error(err)
		return err.Error()
	}
	err = slackTemp.Execute(buf, injectVars)
	if err != nil {
		Log.Error(err)
		return err.Error()
	}
	return buf.String()
}
