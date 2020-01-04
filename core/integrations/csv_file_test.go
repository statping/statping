package integrations

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestCsvFileIntegration(t *testing.T) {
	data, err := ioutil.ReadFile("../../source/tmpl/bulk_import.csv")
	require.Nil(t, err)

	t.Run("Set Field Value", func(t *testing.T) {
		formPost := map[string][]string{}
		formPost["input"] = []string{string(data)}
		_, err = SetFields(csvIntegrator, formPost)
		require.Nil(t, err)
	})

	t.Run("Get Field Value", func(t *testing.T) {
		value := Value(csvIntegrator, "input").(string)
		assert.Equal(t, string(data), value)
	})

	t.Run("List Services from CSV File", func(t *testing.T) {
		services, err := csvIntegrator.List()
		require.Nil(t, err)
		assert.Equal(t, 10, len(services))
	})

	t.Run("Confirm Services from CSV File", func(t *testing.T) {
		services, err := csvIntegrator.List()
		require.Nil(t, err)
		assert.Equal(t, "Bulk Upload", services[0].Name)
		assert.Equal(t, "http://google.com", services[0].Domain)
		assert.Equal(t, 60, services[0].Interval)
		for _, s := range services {
			t.Log(s)
		}
	})

}
