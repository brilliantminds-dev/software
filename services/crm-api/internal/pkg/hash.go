package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(data string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(data), 14)

	return string(bytes)
}

func CheckHash(data string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data))

	return err == nil
}
