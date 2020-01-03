// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package integrations

import (
	"errors"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
)

var (
	Integrations []types.Integrator
	log          = utils.Log.WithField("type", "integration")
)

func init() {
	Integrations = append(Integrations,
		csvIntegrator,
		dockerIntegrator,
		traefikIntegrator,
	)
}

func Value(intg types.Integrator, fieldName string) interface{} {
	for _, v := range intg.Get().Fields {
		if fieldName == v.Name {
			return v.Value
		}
	}
	return nil
}

func SetFields(intg types.Integrator, data map[string][]string) (*types.Integration, error) {
	i := intg.Get()
	for _, v := range i.Fields {
		if data[v.Name] != nil {
			v.Value = data[v.Name][0]
		}
	}
	return i, nil
}

func Find(name string) (types.Integrator, error) {
	for _, i := range Integrations {
		obj := i.Get()
		if obj.ShortName == name {
			return i, nil
		}
	}
	return nil, errors.New(name + " not found")
}
