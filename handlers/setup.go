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
	"errors"
	"github.com/hunterlong/statping/types/configs"
	"github.com/hunterlong/statping/types/core"
	"github.com/hunterlong/statping/types/null"
	"github.com/hunterlong/statping/utils"
	"net/http"
	"time"
)

func processSetupHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if core.App.Setup {
		sendErrorJson(errors.New("Statping has already been setup"), w, r)
		return
	}

	confgs, err := configs.LoadConfigForm(r)
	if err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	//sample, _ := strconv.ParseBool(r.PostForm.Get("sample_data"))

	log.WithFields(utils.ToFields(core.App, confgs)).Debugln("new configs posted")

	if err = configs.ConnectConfigs(confgs, true); err != nil {
		log.Errorln(err)
		if err := confgs.Delete(); err != nil {
			log.Errorln(err)
			sendErrorJson(err, w, r)
			return
		}
		sendErrorJson(err, w, r)
		return
	}

	if err := confgs.Save(utils.Directory); err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	if err = confgs.MigrateDatabase(); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	c := &core.Core{
		Name:        "Statping Sample Data",
		Description: "This data is only used to testing",
		//ApiKey:      apiKey.(string),
		//ApiSecret:   apiSecret.(string),
		Domain:    "http://localhost:8080",
		Version:   "test",
		CreatedAt: time.Now().UTC(),
		UseCdn:    null.NewNullBool(false),
		Footer:    null.NewNullString(""),
	}

	log.Infoln("Creating new Core")
	if err := c.Create(); err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	core.App = c

	//if sample {
	//	log.Infoln("Adding sample data into new database")
	//	if err = configs.TriggerSamples(); err != nil {
	//		log.Errorln(err)
	//		sendErrorJson(err, w, r)
	//		return
	//	}
	//}

	log.Infoln("Initializing new Statping instance")
	if err := core.InitApp(); err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	CacheStorage.Delete("/")
	resetCookies()
	time.Sleep(1 * time.Second)
	out := struct {
		Message string            `json:"message"`
		Config  *configs.DbConfig `json:"config"`
	}{
		"success",
		confgs,
	}
	returnJson(out, w, r)
}
