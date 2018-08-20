// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"github.com/fatih/structs"
	"github.com/hunterlong/statup/types"
	"upper.io/db.v3/lib/sqlbuilder"
)

func OnLoad(db sqlbuilder.Database) {
	for _, p := range CoreApp.AllPlugins {
		p.OnLoad(db)
	}
}

func OnSuccess(s *Service) {
	for _, p := range CoreApp.AllPlugins {
		p.OnSuccess(structs.Map(s))
	}
	//notifiers.OnSuccess(s)
	// TODO convert notifiers to correct type
}

func OnFailure(s *Service, f *types.Failure) {
	for _, p := range CoreApp.AllPlugins {
		p.OnFailure(structs.Map(s))
	}
	//notifiers.OnFailure(s)
	// TODO convert notifiers to correct type
}

func OnSettingsSaved(c *types.Core) {
	for _, p := range CoreApp.AllPlugins {
		p.OnSettingsSaved(structs.Map(c))
	}
}

func OnNewUser(u *User) {
	for _, p := range CoreApp.AllPlugins {
		p.OnNewUser(structs.Map(u))
	}
}

func OnNewService(s *Service) {
	for _, p := range CoreApp.AllPlugins {
		p.OnNewService(structs.Map(s))
	}
}

func OnDeletedService(s *Service) {
	for _, p := range CoreApp.AllPlugins {
		p.OnDeletedService(structs.Map(s))
	}
}

func OnUpdateService(s *Service) {
	for _, p := range CoreApp.AllPlugins {
		p.OnUpdatedService(structs.Map(s))
	}
}

func SelectPlugin(name string) types.PluginActions {
	for _, p := range CoreApp.AllPlugins {
		if p.GetInfo().Name == name {
			return p
		}
	}
	return types.PluginInfo{}
}
