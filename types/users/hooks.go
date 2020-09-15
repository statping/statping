package users

import (
	"github.com/statping/statping/types/errors"
	"github.com/statping/statping/utils"
	"gorm.io/gorm"
)

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username is empty")
	} else if u.Password == "" {
		return errors.New("password is empty")
	}
	return nil
}

func (u *User) BeforeDelete(*gorm.DB) error {
	if utils.Params.GetBool("DEMO_MODE") {
		if u.Username == "admin" {
			return errors.New("cannot delete admin in DEMO_MODE")
		}
	}
	return nil
}

func (u *User) BeforeCreate(*gorm.DB) error {
	if err := u.Validate(); err != nil {
		return err
	}
	u.Password = utils.HashPassword(u.Password)
	u.ApiKey = utils.NewSHA256Hash()
	return nil
}

func (u *User) BeforeUpdate(*gorm.DB) error {
	return u.Validate()
}
