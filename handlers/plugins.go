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

package handlers

import (
	"net/http"
	"strings"
)

type PluginSelect struct {
	Plugin string
	Form   string
	Params map[string]interface{}
}

func pluginSavedHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//vars := mux.Vars(router)
	//plug := SelectPlugin(vars["name"])
	data := make(map[string]string)
	for k, v := range r.PostForm {
		data[k] = strings.Join(v, "")
	}
	//plug.OnSave(structs.Map(data))
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}

func pluginsDownloadHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(router)
	//name := vars["name"]
	//DownloadPlugin(name)
	//core.LoadConfig(utils.Directory)
	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}
