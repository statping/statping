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

package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"net/http"
)

func ApiIndexHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(core.CoreApp)
}

func ApiRenewHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	var err error
	core.CoreApp.ApiKey = utils.NewSHA1Hash(40)
	core.CoreApp.ApiSecret = utils.NewSHA1Hash(40)
	core.CoreApp, err = core.UpdateCore(core.CoreApp)
	if err != nil {
		utils.Log(3, err)
	}
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}

func ApiCheckinHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	checkin := core.FindCheckin(vars["api"])
	//checkin.Receivehit()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(checkin)
}

func ApiServiceHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))
	json.NewEncoder(w).Encode(service)
}

func ApiServiceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	serv := core.SelectService(utils.StringInt(vars["id"]))
	var s *types.Service
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&s)
	serv.Update()
	json.NewEncoder(w).Encode(s)
}

func ApiAllServicesHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	services, _ := core.CoreApp.SelectAllServices()
	json.NewEncoder(w).Encode(services)
}

func ApiUserHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	user, _ := core.SelectUser(utils.StringInt(vars["id"]))
	json.NewEncoder(w).Encode(user)
}

func ApiAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	users, _ := core.SelectAllUsers()
	json.NewEncoder(w).Encode(users)
}

func isAPIAuthorized(r *http.Request) bool {
	if IsAuthenticated(r) {
		return true
	}
	if isAuthorized(r) {
		return true
	}
	return false
}
