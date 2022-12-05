package entity

import "golang.org/x/crypto/bcrypt"

// Receive a string and put a hash on it
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Compares a password with a hash and returs if they are equal
func VerfifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
