package notifiers

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/null"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var form1 = NotificationForm{
	Type:     "text",
	Title:    "Example Input",
	DbField:  "Host",
	Required: true,
	IsHidden: false,
	IsList:   false,
	IsSwitch: false,
}

var form2 = NotificationForm{
	Type:     "text",
	Title:    "Example Input 2",
	DbField:  "ApiKey",
	Required: true,
	IsHidden: false,
	IsList:   false,
	IsSwitch: false,
}

var example = &exampleNotif{&Notification{
	Method:    "example",
	Enabled:   null.NewNullBool(true),
	Limits:    3,
	Removable: false,
	Form:      []NotificationForm{form1, form2},
	Delay:     30,
}}

type exampleNotif struct {
	*Notification
}

func (e *exampleNotif) OnSave() error {
	return nil
}

func (e *exampleNotif) Select() *Notification {
	return e.Notification
}

func (e *exampleNotif) Send(data interface{}) error {
	return nil
}

func TestInit(t *testing.T) {
	db, err := database.OpenTester()
	require.Nil(t, err)
	db.CreateTable(&Notification{})
	db.Create(example.Select())
	SetDB(db)
}

func TestFind(t *testing.T) {
	Append(example)
	itemer, err := Find(example.Method)
	require.Nil(t, err)

	item := itemer.Select()
	require.NotNil(t, item)

	assert.Equal(t, "example", item.Method)
	assert.Len(t, allNotifiers, 1)
}

func TestAll(t *testing.T) {
	items := All()
	assert.Len(t, items, 1)
	assert.Len(t, allNotifiers, 1)
}

func TestCreate(t *testing.T) {
	assert.Len(t, allNotifiers, 1)

	example := &Notification{
		Method:      "anotherexample",
		Title:       "Example 2",
		Description: "New Message here",
	}
	err := example.Create()
	require.Nil(t, err)
	assert.NotZero(t, example.Id)
	assert.Equal(t, "anotherexample", example.Method)
	assert.Equal(t, "Example 2", example.Title)
	assert.NotZero(t, example.CreatedAt)

	items := All()
	assert.Len(t, items, 2)
	assert.Len(t, allNotifiers, 2)
}

func TestUpdate(t *testing.T) {
	itemer, err := Find("anotherexample")
	require.Nil(t, err)
	require.NotNil(t, itemer)

	item := itemer.Select()
	require.NotNil(t, item)

	item.Host = "Updated Host Var"
	err = item.Update()
	require.Nil(t, err)
	assert.Equal(t, "Updated Host Var", item.Host)
}

func TestDelete(t *testing.T) {
	all := All()
	assert.Len(t, all, 2)

	itemer, err := Find("example2")
	require.Nil(t, err)

	item := itemer.Select()
	require.NotNil(t, item)

	err = item.Delete()
	require.Nil(t, err)

	all = All()
	assert.Len(t, all, 2)
}

func TestClose(t *testing.T) {
	assert.Nil(t, db.Close())
}
