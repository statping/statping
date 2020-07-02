package users

import "github.com/statping/statping/utils"

func (u *User) BeforeCreate() error {
	u.Password = utils.HashPassword(u.Password)
	u.ApiKey = utils.NewSHA256Hash()
	return nil
}
