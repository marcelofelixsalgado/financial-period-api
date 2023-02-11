package entity

import (
	"errors"
	"time"
)

type ITenant interface {
	GetId() string
	GetCreatedAt() time.Time
}

type Tenant struct {
	id        string
	createdAt time.Time
}

func NewTenant(id string, createdAt time.Time) (ITenant, error) {
	tenant := Tenant{
		id:        id,
		createdAt: createdAt,
	}

	if err := tenant.validate(); err != nil {
		return nil, err
	}
	return tenant, nil

}

func (tenant Tenant) GetId() string {
	return tenant.id
}

func (tenant Tenant) GetCreatedAt() time.Time {
	return tenant.createdAt
}

func (tenant *Tenant) validate() error {
	if tenant.id == "" {
		return errors.New("id is required")
	}

	return nil
}
