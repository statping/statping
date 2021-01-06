package notifiers

import (
	"testing"

	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/services"
	"github.com/stretchr/testify/assert"
)

func TestReplaceTemplate(t *testing.T) {
	t.Parallel()
	temp := `{"id":{{.Service.Id}},"name":"{{.Service.Name}}"}`
	replaced := ReplaceTemplate(temp, replacer{Service: services.Example(true)})
	assert.Equal(t, `{"id":6283,"name":"Statping Example"}`, replaced)

	temp = `{"id":{{.Service.Id}},"name":"{{.Service.Name}}","failure":"{{.Failure.Issue}}"}`
	replaced = ReplaceTemplate(temp, replacer{Service: services.Example(false), Failure: failures.Example()})
	assert.Equal(t, `{"id":6283,"name":"Statping Example","failure":"Response did not response a 200 status code"}`, replaced)
}

func TestPushover_Select(t *testing.T) {
	tests := []struct {
		Value    string
		Expected string
	}{
		{
			"lowest",
			"-2",
		},
		{
			"low",
			"-1",
		},
		{
			"normal",
			"0",
		},
		{
			"high",
			"1",
		},
		{
			"emergency",
			"2",
		},
		{
			"",
			"0",
		},
	}

	for _, v := range tests {
		assert.Equal(t, v.Expected, priority(v.Value))
	}
}
