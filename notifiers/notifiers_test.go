package notifiers

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestReplaceTemplate(t *testing.T) {
	temp := `{"id":{{.Service.Id}},"name":"{{.Service.Name}}"}`
	replaced := ReplaceTemplate(temp, replacer{Service: exampleService})
	assert.Equal(t, `{"id":1,"name":"Statping"}`, replaced)

	temp = `{"id":{{.Service.Id}},"name":"{{.Service.Name}}","failure":"{{.Failure.Issue}}"}`
	replaced = ReplaceTemplate(temp, replacer{Service: exampleService, Failure: exampleFailure})
	assert.Equal(t, `{"id":1,"name":"Statping","failure":"HTTP returned a 500 status code"}`, replaced)
}
