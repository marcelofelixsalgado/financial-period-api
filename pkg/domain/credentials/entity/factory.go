package entity

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"
	"time"

	"github.com/google/uuid"
)

func Create(user entity.IUser, password string) (IUserCredentials, error) {

	hashedPassword, err := Hash(password)
	if err != nil {
		return nil, err
	}
	password = string(hashedPassword)

	userCredentials, err := NewUserCredentials(uuid.NewString(), user.GetId(), password, time.Now(), time.Time{})
	if err != nil {
		return nil, err
	}

	return userCredentials, nil
}
