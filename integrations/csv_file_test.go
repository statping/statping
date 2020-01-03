package integrations

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCsvFileIntegration(t *testing.T) {

	csvIntegrator.Fields[0].Value = "test_files/example_services.csv"

	t.Run("CSV File", func(t *testing.T) {
		path := csvIntegrator.Fields[0].Value
		assert.Equal(t, "test_files/example_services.csv", path)
	})

	t.Run("CSV Open File", func(t *testing.T) {
		err := csvIntegrator.Open()
		require.Nil(t, err)
	})

	t.Run("List Services from CSV File", func(t *testing.T) {
		services, err := csvIntegrator.List()
		require.Nil(t, err)
		assert.Equal(t, len(services), 1)
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

	t.Run("Close CSV", func(t *testing.T) {
		err := csvIntegrator.Close()
		require.Nil(t, err)
	})

}
