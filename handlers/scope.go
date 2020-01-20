// Statup
// Copyright (C) 2020.  Hunter Long and the project contributors
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

package handlers

import (
	"encoding/json"
	"reflect"
	"strings"
)

type scope struct {
	data  interface{}
	scope string
}

// MarshalJSON for Scopr
func (s scope) MarshalJSON() ([]byte, error) {
	svc := reflect.ValueOf(s.data)
	if svc.Kind() == reflect.Slice {
		alldata := make([]map[string]interface{}, 0)
		for i := 0; i < svc.Len(); i++ {
			objIndex := svc.Index(i)
			if objIndex.Kind() == reflect.Ptr {
				objIndex = objIndex.Elem()
			}
			alldata = append(alldata, SafeJson(objIndex.Interface(), s.scope))
		}
		return json.Marshal(alldata)
	}
	return json.Marshal(SafeJson(svc.Interface(), s.scope))
}

func SafeJson(input interface{}, scope string) map[string]interface{} {
	thisData := make(map[string]interface{})
	t := reflect.TypeOf(input)
	elem := reflect.ValueOf(input)
	d, _ := json.Marshal(input)

	var raw map[string]*json.RawMessage
	json.Unmarshal(d, &raw)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get("scope")
		tags := strings.Split(tag, ",")

		jTags := field.Tag.Get("json")
		jsonTag := strings.Split(jTags, ",")

		if len(jsonTag) == 0 {
			continue
		}

		if jsonTag[0] == "" || jsonTag[0] == "-" {
			continue
		}

		trueValue := elem.Field(i).Interface()

		if len(jsonTag) == 2 {
			if jsonTag[1] == "omitempty" && trueValue == "" {
				continue
			}
		}

		if tag == "" {
			thisData[jsonTag[0]] = trueValue
			continue
		}

		if forTag(tags, scope) {
			thisData[jsonTag[0]] = trueValue
		}
	}
	return thisData
}

func forTag(tags []string, scope string) bool {
	for _, v := range tags {
		if v == scope {
			return true
		}
	}
	return false
}
