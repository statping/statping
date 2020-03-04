package messages

import "github.com/hunterlong/statping/database"

func Find(id int64) (*Message, error) {
	var user *Message
	db := database.DB().Model(&Message{}).Where("id = ?", id).Find(&user)
	return user, db.Error()
}

func All() []*Message {
	var messages []*Message
	database.DB().Model(&Message{}).Find(&messages)
	return messages
}

func (m *Message) Create() error {
	db := database.DB().Create(&m)
	return db.Error()
}

func (m *Message) Update() error {
	db := database.DB().Update(&m)
	return db.Error()
}

func (m *Message) Delete() error {
	db := database.DB().Delete(&m)
	return db.Error()
}
