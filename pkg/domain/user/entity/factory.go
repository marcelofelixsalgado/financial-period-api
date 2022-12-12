package entity

import (
	"time"

	"github.com/google/uuid"
)

func Create(name string, phone string, email string) (IUser, error) {

	user, err := NewUser(uuid.NewString(), name, phone, email, time.Now(), time.Time{})
	if err != nil {
		return nil, err
	}

	return user, nil
}
