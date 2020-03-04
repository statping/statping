package users

import (
	"github.com/hunterlong/statping/types/null"
)

func Samples() {
	u2 := &User{
		Username: "testadmin",
		Password: "password123",
		Email:    "info@betatude.com",
		Admin:    null.NewNullBool(true),
	}

	u2.Create()

	u3 := &User{
		Username: "testadmin2",
		Password: "password123",
		Email:    "info@adminhere.com",
		Admin:    null.NewNullBool(true),
	}

	u3.Create()
}
