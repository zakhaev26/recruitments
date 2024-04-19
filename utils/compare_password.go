package utils

import "golang.org/x/crypto/bcrypt"

func ComparePassword(hashedPass string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
}
