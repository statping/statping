package notifiers

import (
	"github.com/magiconair/properties/assert"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/services"
	"testing"
)

func TestReplaceTemplate(t *testing.T) {
	temp := `{"id":{{.Service.Id}},"name":"{{.Service.Name}}"}`
	replaced := ReplaceTemplate(temp, replacer{Service: services.Example(true)})
	assert.Equal(t, `{"id":6283,"name":"Statping"}`, replaced)

	temp = `{"id":{{.Service.Id}},"name":"{{.Service.Name}}","failure":"{{.Failure.Issue}}"}`
	replaced = ReplaceTemplate(temp, replacer{Service: services.Example(false), Failure: failures.Example()})
	assert.Equal(t, `{"id":6283,"name":"Statping","failure":"HTTP returned a 500 status code"}`, replaced)
}
