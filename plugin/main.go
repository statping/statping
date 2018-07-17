package plugin

import (
	"github.com/hunterlong/statup/types"
	"upper.io/db.v3/lib/sqlbuilder"
)

//
//     STATUP PLUGIN INTERFACE
//
//            v0.1
//
//       https://statup.io
//
//
// An expandable plugin framework that will still
// work even if there's an update or addition.
//

var (
	DB sqlbuilder.Database
)

type PluginInfo struct {
	i *types.Info
}

func SetDatabase(database sqlbuilder.Database) {
	DB = database
}

func (p *PluginInfo) Form() string {
	return "okkokokkok"
}
