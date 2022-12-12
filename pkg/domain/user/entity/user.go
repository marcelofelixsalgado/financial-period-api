package entity

import (
	"errors"
	"strings"
	"time"
)

type IUser interface {
	GetId() string
	GetName() string
	GetPhone() string
	GetEmail() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type User struct {
	id        string
	name      string
	phone     string
	email     string
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(id string, name string, phone string, email string, createdAt time.Time, updatedAt time.Time) (IUser, error) {
	user := User{
		id:        id,
		name:      name,
		phone:     phone,
		email:     email,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
	user.format()

	if err := user.validate(); err != nil {
		return nil, err
	}
	return user, nil

}

func (user User) GetId() string {
	return user.id
}

func (user User) GetName() string {
	return user.name
}

func (user User) GetPhone() string {
	return user.phone
}

func (user User) GetEmail() string {
	return user.email
}

func (user User) GetCreatedAt() time.Time {
	return user.createdAt
}

func (user User) GetUpdatedAt() time.Time {
	return user.updatedAt
}

func (user *User) SetUpdatedAt(updatedAt time.Time) {
	user.updatedAt = updatedAt
}

func (user *User) format() {
	user.name = strings.TrimSpace(user.name)
}

func (user *User) validate() error {
	if user.id == "" {
		return errors.New("id is required")
	}

	if user.name == "" {
		return errors.New("name is required")
	}

	if user.phone == "" {
		return errors.New("phone is required")
	}

	if user.email == "" {
		return errors.New("email is required")
	}

	return nil
}
