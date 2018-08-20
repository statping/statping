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
	"io/ioutil"
)

func ExportIndexHTML() string {
	source.Assets()
	CoreApp.UseCdn = true
	//out := index{*CoreApp, CoreApp.DbServices}
	nav, _ := source.TmplBox.String("nav.html")
	footer, _ := source.TmplBox.String("footer.html")
	render, err := source.TmplBox.String("index.html")
	if err != nil {
		utils.Log(3, err)
	}
	t := template.New("message")
	t.Funcs(template.FuncMap{
		"js": func(html string) template.JS {
			return template.JS(html)
		},
		"safe": func(html string) template.HTML {
			return template.HTML(html)
		},
		"VERSION": func() string {
			return VERSION
		},
		"CoreApp": func() *Core {
			return CoreApp
		},
		"USE_CDN": func() bool {
			return CoreApp.UseCdn
		},
		"underscore": func(html string) string {
			return utils.UnderScoreString(html)
		},
	})
	t, _ = t.Parse(nav)
	t, _ = t.Parse(footer)
	t.Parse(render)
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, nil); err != nil {
		utils.Log(3, err)
	}
	result := tpl.String()
	return result
}

func SaveFile(filename string, data []byte) error {
	err := ioutil.WriteFile(filename, data, 0644)
	return err
}
