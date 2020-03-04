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
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statping/types/messages"
	"github.com/hunterlong/statping/utils"
	"net/http"
)

func getMessageByID(r *http.Request) (*messages.Message, int64, error) {
	vars := mux.Vars(r)
	num := utils.ToInt(vars["id"])
	message, err := messages.Find(num)
	if err != nil {
		return nil, num, err
	}
	return message, num, nil
}

func apiAllMessagesHandler(r *http.Request) interface{} {
	msgs := messages.All()
	return msgs
}

func apiMessageCreateHandler(w http.ResponseWriter, r *http.Request) {
	var message *messages.Message
	if err := DecodeJSON(r, &message); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	if err := message.Create(); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(message, "create", w, r)
}

func apiMessageGetHandler(w http.ResponseWriter, r *http.Request) {
	message, id, err := getMessageByID(r)
	if err != nil {
		sendErrorJson(fmt.Errorf("message #%d was not found", id), w, r)
		return
	}
	returnJson(message, w, r)
}

func apiMessageDeleteHandler(w http.ResponseWriter, r *http.Request) {
	message, id, err := getMessageByID(r)
	if err != nil {
		sendErrorJson(fmt.Errorf("message #%d was not found", id), w, r)
		return
	}
	err = message.Delete()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(message, "delete", w, r)
}

func apiMessageUpdateHandler(w http.ResponseWriter, r *http.Request) {
	message, id, err := getMessageByID(r)
	if err != nil {
		sendErrorJson(fmt.Errorf("message #%d was not found", id), w, r)
		return
	}
	if err := DecodeJSON(r, &message); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	if err := message.Update(); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(message, "update", w, r)
}
