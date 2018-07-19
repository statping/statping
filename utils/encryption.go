package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func NewSHA1Hash(n ...int) string {
	noRandomCharacters := 32
	if len(n) > 0 {
		noRandomCharacters = n[0]
	}
	randString := RandomString(noRandomCharacters)
	hash := sha1.New()
	hash.Write([]byte(randString))
	bs := hash.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString generates a random string of n length
func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	return string(b)
}

func Sha256(data []byte) string {
	h := sha1.New()
	h.Write(data)
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}
