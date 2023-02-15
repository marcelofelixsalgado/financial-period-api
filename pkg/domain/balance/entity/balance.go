package entity

import (
	"errors"
	"time"
)

type IBalance interface {
	GetId() string
	GetTenantId() string
	GetPeriodId() string
	GetCategoryId() string
	GetActualAmount() float32
	GetLimitAmount() float32
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type Balance struct {
	id           string
	tenantId     string
	periodId     string
	categoryId   string
	actualAmount float32
	limitAmount  float32
	createdAt    time.Time
	updatedAt    time.Time
}

func (balance Balance) GetId() string {
	return balance.id
}

func (balance Balance) GetTenantId() string {
	return balance.tenantId
}

func (balance Balance) GetPeriodId() string {
	return balance.periodId
}

func (balance Balance) GetCategoryId() string {
	return balance.categoryId
}

func (balance Balance) GetActualAmount() float32 {
	return balance.actualAmount
}

func (balance Balance) GetLimitAmount() float32 {
	return balance.limitAmount
}

func (balance Balance) GetCreatedAt() time.Time {
	return balance.createdAt
}

func (balance Balance) GetUpdatedAt() time.Time {
	return balance.updatedAt
}

func NewBalance(id string, tenantId string, periodId string, categoryId string, actualAmount float32, limitAmount float32, createdAt time.Time, updatedAt time.Time) (IBalance, error) {
	balance := Balance{
		id:           id,
		tenantId:     tenantId,
		periodId:     periodId,
		categoryId:   categoryId,
		actualAmount: actualAmount,
		limitAmount:  limitAmount,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}

	if err := balance.validate(); err != nil {
		return nil, err
	}

	return balance, nil
}

func (balance *Balance) validate() error {
	if balance.id == "" {
		return errors.New("id is required")
	}

	if balance.tenantId == "" {
		return errors.New("tenant id is required")
	}

	if balance.periodId == "" {
		return errors.New("period is required")
	}

	if balance.categoryId == "" {
		return errors.New("category is required")
	}

	return nil
}
