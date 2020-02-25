package database

import "github.com/hunterlong/statping/types"

type MessageObj struct {
	*types.Message
	o *Object
}

func Message(id int64) (*MessageObj, error) {
	var message types.Message
	query := database.Messages().Where("id = ?", id)
	finder := query.Find(&message)
	return &MessageObj{Message: &message, o: wrapObject(id, &message, query)}, finder.Error()
}

func AllMessages() []*types.Message {
	var messages []*types.Message
	database.Messages().Find(&messages)
	return messages
}

func (m *MessageObj) object() *Object {
	return m.o
}
