package hashing

import "golang.org/x/crypto/bcrypt"

func Hash(code string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(code), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckCode(code, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(code))
	return err == nil
}
