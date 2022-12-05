package entity

import (
	"time"

	"github.com/google/uuid"
)

func Create(name string, password string, phone string, email string) (IUser, error) {

	hashedPassword, err := Hash(password)
	if err != nil {
		return nil, err
	}
	password = string(hashedPassword)

	return NewUser(uuid.NewString(), name, password, phone, email, time.Now(), time.Time{})
}
