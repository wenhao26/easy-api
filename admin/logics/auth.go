package logics

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

func CheckPassword(cipherText, plainText []byte) bool {
	err := bcrypt.CompareHashAndPassword(cipherText, plainText)
	if err != nil {
		return false
	}
	return true
}
