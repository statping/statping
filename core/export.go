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
	"bytes"
	"github.com/hunterlong/statup/source"
	"github.com/hunterlong/statup/utils"
	"html/template"
)

// ExportChartsJs renders the charts for the index page
func ExportChartsJs() string {
	render, err := source.JsBox.String("charts.js")
	if err != nil {
		utils.Log(4, err)
	}
	t := template.New("charts")
	t.Funcs(template.FuncMap{
		"safe": func(html string) template.HTML {
			return template.HTML(html)
		},
	})
	t.Parse(render)
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, CoreApp.Services); err != nil {
		utils.Log(3, err)
	}
	result := tpl.String()
	return result
}
