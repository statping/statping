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

package core

import (
	"fmt"
	"github.com/hunterlong/statping/types"
	"time"
)

type Message struct {
	*types.Message
}

// SelectServiceMessages returns all messages for a service
func SelectServiceMessages(id int64) []*Message {
	var message []*Message
	messagesDb().Where("service = ?", id).Limit(10).Find(&message)
	return message
}

// ReturnMessage will convert *types.Message to *core.Message
func ReturnMessage(m *types.Message) *Message {
	return &Message{m}
}

// SelectMessages returns all messages
func SelectMessages() ([]*Message, error) {
	var messages []*Message
	db := messagesDb().Find(&messages).Order("id desc")
	return messages, db.Error
}

// SelectMessage returns a Message based on the ID passed
func SelectMessage(id int64) (*Message, error) {
	var message Message
	db := messagesDb().Where("id = ?", id).Find(&message)
	return &message, db.Error
}

func (m *Message) Service() *Service {
	if m.ServiceId == 0 {
		return nil
	}
	return SelectService(m.ServiceId)
}

// Create will create a Message and insert it into the database
func (m *Message) Create() (int64, error) {
	m.CreatedAt = time.Now().UTC()
	db := messagesDb().Create(m)
	if db.Error != nil {
		log.Errorln(fmt.Sprintf("Failed to create message %v #%v: %v", m.Title, m.Id, db.Error))
		return 0, db.Error
	}
	return m.Id, nil
}

// Delete will delete a Message from database
func (m *Message) Delete() error {
	db := messagesDb().Delete(m)
	return db.Error
}

// Update will update a Message in the database
func (m *Message) Update() (*Message, error) {
	db := messagesDb().Update(m)
	if db.Error != nil {
		log.Errorln(fmt.Sprintf("Failed to update message %v #%v: %v", m.Title, m.Id, db.Error))
		return nil, db.Error
	}
	return m, nil
}
