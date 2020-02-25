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

func (m *MessageObj) object() *Object {
	return m.o
}
