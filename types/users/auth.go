package users

import (
	"fmt"
	"github.com/prometheus/common/log"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// AuthUser will return the User and a boolean if authentication was correct.
// AuthUser accepts username, and password as a string
func AuthUser(username, password string) (*User, bool) {
	user, err := FindByUsername(username)
	if err != nil {
		log.Warnln(fmt.Errorf("user %v not found", username))
		return nil, false
	}
	if checkHash(password, user.Password) {
		user.UpdatedAt = time.Now().UTC()
		user.Update()
		return user, true
	}
	return nil, false
}

// checkHash returns true if the password matches with a hashed bcrypt password
func checkHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
