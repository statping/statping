package integrations

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDockerIntegration(t *testing.T) {

	t.Run("Set Field Value", func(t *testing.T) {
		formPost := map[string][]string{}
		formPost["path"] = []string{"unix:///var/run/docker.sock"}
		formPost["version"] = []string{"1.25"}
		_, err := SetFields(csvIntegrator, formPost)
		require.Nil(t, err)
	})

	t.Run("Get Field Value", func(t *testing.T) {
		path := Value(dockerIntegrator, "path").(string)
		version := Value(dockerIntegrator, "version").(string)
		assert.Equal(t, "unix:///var/run/docker.sock", path)
		assert.Equal(t, "1.25", version)
	})

	t.Run("List Services from Docker", func(t *testing.T) {
		services, err := dockerIntegrator.List()
		require.Nil(t, err)
		assert.Equal(t, 0, len(services))
	})

	t.Run("Confirm Services from Docker", func(t *testing.T) {
		services, err := dockerIntegrator.List()
		require.Nil(t, err)
		for _, s := range services {
			t.Log(s)
		}
	})

}
