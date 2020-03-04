package users

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types/null"
)

func (u *User) Samples() []database.DbObject {

	var samples []database.DbObject

	u2 := &User{
		Username: "testadmin",
		Password: "password123",
		Email:    "info@betatude.com",
		Admin:    null.NewNullBool(true),
	}

	samples = append(samples, u2)

	u3 := &User{
		Username: "testadmin2",
		Password: "password123",
		Email:    "info@adminhere.com",
		Admin:    null.NewNullBool(true),
	}

	samples = append(samples, u3)

	return samples
}
