package entity

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type IUserCredentials interface {
	GetId() string
	GetUserId() string
	GetPassword() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type UserCredentials struct {
	id        string
	userId    string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

func NewUserCredentials(id string, userId string, password string, createdAt time.Time, updatedAt time.Time) (IUserCredentials, error) {
	userCredentials := UserCredentials{
		id:        id,
		userId:    userId,
		password:  password,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}

	if err := userCredentials.validate(); err != nil {
		return nil, err
	}
	return userCredentials, nil
}

func (userCredentials UserCredentials) GetId() string {
	return userCredentials.id
}

func (userCredentials UserCredentials) GetUserId() string {
	return userCredentials.userId
}

func (userCredentials UserCredentials) GetPassword() string {
	return userCredentials.password
}

func (userCredentials UserCredentials) GetCreatedAt() time.Time {
	return userCredentials.createdAt
}

func (userCredentials UserCredentials) GetUpdatedAt() time.Time {
	return userCredentials.updatedAt
}

func (userCredentials *UserCredentials) SetUpdatedAt(updatedAt time.Time) {
	userCredentials.updatedAt = updatedAt
}

func (userCredentials *UserCredentials) validate() error {
	if userCredentials.id == "" {
		return errors.New("id is required")
	}

	if userCredentials.userId == "" {
		return errors.New("user id is required")
	}

	if userCredentials.password == "" {
		return errors.New("password is required")
	}

	return nil
}

// Receive a string and put a hash on it
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Compares a password with a hash and returs if they are equal
func VerfifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
