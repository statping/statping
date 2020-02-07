package integrations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrations(t *testing.T) {

	t.Run("Collect Integrations", func(t *testing.T) {
		amount := len(Integrations)
		assert.Equal(t, 3, amount)
	})

}
