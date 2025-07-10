package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword genera un hash a partir de una contraseña
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compara una contraseña y un hash
func CheckPasswordHash(password, hash string) bool {
	passHashed := hash
	err := bcrypt.CompareHashAndPassword([]byte(passHashed), []byte(password))
	return err == nil
}
