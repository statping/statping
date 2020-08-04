package hits

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInit(t *testing.T) {
	err := utils.InitLogs()
	require.Nil(t, err)
	db, err := database.OpenTester()
	require.Nil(t, err)
	db.CreateTable(&Hit{}, &services.Service{})
	SetDB(db)
	services.SetDB(db)

	for i := 0; i <= 5; i++ {
		s := services.Example(true)
		assert.Nil(t, s.Create())
		assert.Len(t, s.AllHits().List(), 2)
	}

	require.Nil(t, Samples())
}
