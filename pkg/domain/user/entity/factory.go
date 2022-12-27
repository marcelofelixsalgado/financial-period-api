package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

func Create(name string, phone string, email string) (IUser, error) {

	user, err := NewUser(uuid.NewV4().String(), name, phone, email, time.Now(), time.Time{})
	if err != nil {
		return nil, err
	}

	return user, nil
}
