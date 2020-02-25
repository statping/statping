package database

import "github.com/hunterlong/statping/types"

type UserObj struct {
	*types.User
	o *Object
}

func User(id int64) (*UserObj, error) {
	var user types.User
	query := database.Users().Where("id = ?", id)
	finder := query.First(&user)
	return &UserObj{User: &user, o: wrapObject(id, &user, query)}, finder.Error()
}

func UserByUsername(username string) (*UserObj, error) {
	var user types.User
	query := database.Users().Where("username = ?", username)
	finder := query.First(&user)
	return &UserObj{User: &user, o: wrapObject(user.Id, &user, query)}, finder.Error()
}

func AllUsers() []*types.User {
	var users []*types.User
	database.Users().Find(&users)
	return users
}

func (u *UserObj) object() *Object {
	return u.o
}
