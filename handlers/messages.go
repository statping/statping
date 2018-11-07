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
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"time"
)

func messagesHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	messages, _ := core.SelectMessages()
	executeResponse(w, r, "messages.html", messages, nil)
}

func deleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	id := utils.StringInt(vars["id"])
	message, err := core.SelectMessage(id)
	if err != nil {
		http.Redirect(w, r, "/messages", http.StatusSeeOther)
		return
	}
	message.Delete()
	http.Redirect(w, r, "/messages", http.StatusSeeOther)
}

func viewMessageHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	id := utils.StringInt(vars["id"])
	message, err := core.SelectMessage(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	executeResponse(w, r, "message.html", message, nil)
}

func updateMessageHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	id := utils.StringInt(vars["id"])
	message, err := core.SelectMessage(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	r.ParseForm()

	title := r.PostForm.Get("title")
	description := r.PostForm.Get("description")
	notifyMethod := r.PostForm.Get("notify_method")
	notifyUsers := r.PostForm.Get("notify_users")
	startOn := r.PostForm.Get("start_on")
	endOn := r.PostForm.Get("end_on")
	notifyBefore := r.PostForm.Get("notify_before")
	serviceId := utils.StringInt(r.PostForm.Get("service_id"))

	start, err := time.Parse(utils.FlatpickrTime, startOn)
	if err != nil {
		utils.Log(3, err)
	}
	end, _ := time.Parse(utils.FlatpickrTime, endOn)
	before, _ := time.ParseDuration(notifyBefore)

	message.Title = title
	message.Description = description
	message.NotifyUsers = types.NewNullBool(notifyUsers == "on")
	message.NotifyMethod = notifyMethod
	message.StartOn = start.UTC()
	message.EndOn = end.UTC()
	message.NotifyBefore = before
	message.ServiceId = serviceId

	message.Update()
	executeResponse(w, r, "messages.html", message, "/messages")
}

func createMessageHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	r.ParseForm()

	title := r.PostForm.Get("title")
	description := r.PostForm.Get("description")
	notifyMethod := r.PostForm.Get("notify_method")
	notifyUsers := r.PostForm.Get("notify_users")
	startOn := r.PostForm.Get("start_on")
	endOn := r.PostForm.Get("end_on")
	notifyBefore := r.PostForm.Get("notify_before")
	serviceId := utils.StringInt(r.PostForm.Get("service_id"))

	start, _ := time.Parse(utils.FlatpickrTime, startOn)
	end, _ := time.Parse(utils.FlatpickrTime, endOn)
	before, _ := time.ParseDuration(notifyBefore)

	message := core.ReturnMessage(&types.Message{
		Title:        title,
		Description:  description,
		StartOn:      start.UTC(),
		EndOn:        end.UTC(),
		ServiceId:    serviceId,
		NotifyUsers:  types.NewNullBool(notifyUsers == "on"),
		NotifyMethod: notifyMethod,
		NotifyBefore: before,
	})
	_, err := message.Create()
	if err != nil {
		utils.Log(3, fmt.Sprintf("Error creating message %v", err))
	}
	messages, _ := core.SelectMessages()
	executeResponse(w, r, "messages.html", messages, "/messages")
}
