package database

import "github.com/hunterlong/statping/types"

type UserObj struct {
	*types.User
}

func (o *Object) AsUser() *UserObj {
	return &UserObj{
		User: o.model.(*types.User),
	}
}

func User(id int64) (*UserObj, error) {
	var user types.User
	query := database.Model(&types.User{}).Where("id = ?", id).Find(&user)
	return &UserObj{User: &user}, query.Error()
}
