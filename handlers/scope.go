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

// TODO: make a better way to parse
func (s scope) MarshalJSON() ([]byte, error) {
	svc := reflect.ValueOf(s.data)
	if svc.Kind() == reflect.Slice {
		alldata := make([]map[string]interface{}, svc.Len())
		for i := 0; i < svc.Len(); i++ {
			objIndex := svc.Index(i)
			alldata[i] = SafeJson(objIndex, s.scope)
		}
		return json.Marshal(alldata)
	}
	return json.Marshal(SafeJson(svc, s.scope))
}

// TODO: make a better way to parse
func SafeJson(val reflect.Value, scope string) map[string]interface{} {
	thisData := make(map[string]interface{})
	if val.Kind() == reflect.Interface && !val.IsNil() {
		elm := val.Elem()
		if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
			val = elm
		}
	}
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tagVal := typeField.Tag

		tag := tagVal.Get("scope")
		tags := strings.Split(tag, ",")

		jTags := tagVal.Get("json")
		jsonTag := strings.Split(jTags, ",")

		if len(jsonTag) == 0 {
			continue
		}

		if jsonTag[0] == "" || jsonTag[0] == "-" {
			continue
		}

		if len(jsonTag) == 2 {
			if jsonTag[1] == "omitempty" && valueField.Interface() == "" {
				continue
			}
		}

		if tag == "" {
			thisData[jsonTag[0]] = valueField.Interface()
			continue
		}

		if forTag(tags, scope) {
			thisData[jsonTag[0]] = valueField.Interface()
		}
	}
	return thisData
}

// TODO: make a better way to parse
func forTag(tags []string, scope string) bool {
	for _, v := range tags {
		if v == scope {
			return true
		}
	}
	return false
}
