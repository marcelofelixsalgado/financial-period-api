package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
)

type ICategory interface {
	GetId() string
	GetTenantId() string
	GetCode() string
	GetName() string
	GetTransactionType() entity.ITransactionType
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type Category struct {
	id              string
	tenantId        string
	code            string
	name            string
	transactionType entity.ITransactionType
	createAt        time.Time
	updatedAt       time.Time
}

func (category Category) GetId() string {
	return category.id
}

func (category Category) GetTenantId() string {
	return category.tenantId
}

func (category Category) GetCode() string {
	return category.code
}

func (category Category) GetName() string {
	return category.name
}

func (category Category) GetTransactionType() entity.ITransactionType {
	return category.transactionType
}

func (category Category) GetCreatedAt() time.Time {
	return category.createAt
}

func (category Category) GetUpdatedAt() time.Time {
	return category.updatedAt
}

func (category *Category) SetUpdatedAt(updatedAt time.Time) {
	category.updatedAt = updatedAt
	category.validate()
}

func NewCategory(id string, tenantId string, code string, name string, transactionType entity.ITransactionType, createAt time.Time, updatedAt time.Time) (ICategory, error) {
	category := Category{
		id:              id,
		tenantId:        tenantId,
		code:            code,
		name:            name,
		transactionType: transactionType,
		createAt:        createAt,
		updatedAt:       updatedAt,
	}
	category.format()
	if err := category.validate(); err != nil {
		return nil, err
	}
	return category, nil
}

func (category *Category) validate() error {
	if category.id == "" {
		return errors.New("id is required")
	}

	if category.tenantId == "" {
		return errors.New("tenant id is required")
	}

	if category.code == "" {
		return errors.New("code is required")
	}

	if category.name == "" {
		return errors.New("name is required")
	}

	if category.transactionType == nil {
		return errors.New("transaction type is required")
	}

	return nil
}

func (category *Category) format() {
	category.code = strings.TrimSpace(category.code)
	category.name = strings.TrimSpace(category.name)
}
