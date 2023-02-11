package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

func Create(tenantId string, name string, phone string, email string) (IUser, error) {

	user, err := NewUser(uuid.NewV4().String(), tenantId, name, phone, email, time.Now(), time.Time{})
	if err != nil {
		return nil, err
	}

	return user, nil
}
