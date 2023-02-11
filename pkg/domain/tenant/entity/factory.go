package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

func Create() (ITenant, error) {

	tenant, err := NewTenant(uuid.NewV4().String(), time.Now())
	if err != nil {
		return nil, err
	}

	return tenant, nil
}
