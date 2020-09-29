package utils

import (
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

// HashPassword returns the bcrypt hash of a password string
func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

// CheckHash returns true if the password matches with a hashed bcrypt password
func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// NewSHA1Hash returns a random SHA1 hash based on a specific length
func NewSHA256Hash() string {
	d := make([]byte, 10)
	rand.Seed(Now().UnixNano())
	rand.Read(d)
	return fmt.Sprintf("%x", sha256.Sum256(d))
}

// NewSHA1Hash returns a random SHA1 hash based on a specific length
func Sha256Hash(val string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(val)))
}

var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString generates a random string of n length
func RandomString(n int) string {
	b := make([]rune, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	return string(b)
}
