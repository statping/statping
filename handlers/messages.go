package handlers

import (
	"github.com/gorilla/mux"
	"github.com/statping/statping/types/errors"
	"github.com/statping/statping/types/messages"
	"github.com/statping/statping/utils"
	"net/http"
)

func findMessage(r *http.Request) (*messages.Message, int64, error) {
	vars := mux.Vars(r)
	if utils.NotNumber(vars["id"]) {
		return nil, 0, errors.NotNumber
	}
	id := utils.ToInt(vars["id"])
	message, err := messages.Find(id)
	if err != nil {
		return nil, id, err
	}
	return message, id, nil
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

func apiMessageGetHandler(r *http.Request) interface{} {
	message, _, err := findMessage(r)
	if err != nil {
		return err
	}
	return message
}

func apiMessageDeleteHandler(w http.ResponseWriter, r *http.Request) {
	message, _, err := findMessage(r)
	if err != nil {
		sendErrorJson(err, w, r)
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
	message, _, err := findMessage(r)
	if err != nil {
		sendErrorJson(err, w, r)
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
