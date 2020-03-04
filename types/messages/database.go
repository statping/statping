package messages

import "github.com/hunterlong/statping/database"

func DB() database.Database {
	return database.DB().Model(&Message{})
}

func Find(id int64) (*Message, error) {
	var user *Message
	db := DB().Where("id = ?", id).Find(&user)
	return user, db.Error()
}

func All() []*Message {
	var messages []*Message
	DB().Find(&messages)
	return messages
}

func (m *Message) Create() error {
	db := DB().Create(&m)
	return db.Error()
}

func (m *Message) Update() error {
	db := DB().Update(&m)
	return db.Error()
}

func (m *Message) Delete() error {
	db := DB().Delete(&m)
	return db.Error()
}
