package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/statping/statping/types/messages"
	"github.com/statping/statping/utils"
	"net/http"
)

func findMessage(r *http.Request) (*messages.Message, int64, error) {
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

func apiMessageGetHandler(r *http.Request) interface{} {
	message, id, err := findMessage(r)
	if err != nil {
		return fmt.Errorf("message #%d was not found", id)
	}
	return message
}

func apiMessageDeleteHandler(w http.ResponseWriter, r *http.Request) {
	message, id, err := findMessage(r)
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
	message, id, err := findMessage(r)
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
