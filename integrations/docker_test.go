package integrations

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDockerIntegration(t *testing.T) {

	t.Run("Docker Open containers", func(t *testing.T) {
		err := dockerIntegrator.Open()
		require.Nil(t, err)
	})

	t.Run("List Services from Docker", func(t *testing.T) {
		services, err := dockerIntegrator.List()
		require.Nil(t, err)
		assert.NotEqual(t, 0, len(services))
	})

	t.Run("Confirm Services from Docker", func(t *testing.T) {
		services, err := dockerIntegrator.List()
		require.Nil(t, err)
		for _, s := range services {
			t.Log(s)
		}
	})

	t.Run("Close Docker", func(t *testing.T) {
		err := dockerIntegrator.Close()
		require.Nil(t, err)
	})

}
