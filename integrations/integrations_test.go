package integrations

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIntegrations(t *testing.T) {

	t.Run("Collect Integrations", func(t *testing.T) {
		amount := len(Integrations)
		assert.Equal(t, 2, amount)
	})

	t.Run("Close All Integrations", func(t *testing.T) {
		closedAll := CloseAll()
		require.Nil(t, closedAll)
	})

}
