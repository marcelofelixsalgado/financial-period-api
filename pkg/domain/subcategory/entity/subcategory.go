package entity

import (
	"errors"
	"marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	"strings"
	"time"
)

type ISubCategory interface {
	GetId() string
	GetTenantId() string
	GetCode() string
	GetName() string
	GetCategory() entity.ICategory
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type SubCategory struct {
	id        string
	tenantId  string
	code      string
	name      string
	category  entity.ICategory
	createdAt time.Time
	updatedAt time.Time
}

func NewSubCategory(id string, tenantId string, code string, name string, category entity.ICategory, createdAt time.Time, updatedAt time.Time) (ISubCategory, error) {
	subCategory := SubCategory{
		id:        id,
		tenantId:  tenantId,
		code:      code,
		name:      name,
		category:  category,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
	subCategory.format()
	if err := subCategory.validate(); err != nil {
		return nil, err
	}
	return subCategory, nil
}

func (subCategory SubCategory) GetId() string {
	return subCategory.id
}

func (subCategory SubCategory) GetTenantId() string {
	return subCategory.tenantId
}

func (subCategory SubCategory) GetCode() string {
	return subCategory.code
}

func (subCategory SubCategory) GetName() string {
	return subCategory.name
}

func (subCategory SubCategory) GetCategory() entity.ICategory {
	return subCategory.category
}

func (subCategory SubCategory) GetCreatedAt() time.Time {
	return subCategory.createdAt
}

func (subCategory SubCategory) GetUpdatedAt() time.Time {
	return subCategory.updatedAt
}

func (subCategory *SubCategory) format() {
	subCategory.code = strings.TrimSpace(subCategory.code)
	subCategory.name = strings.TrimSpace(subCategory.name)
}

func (subCategory *SubCategory) validate() error {
	if subCategory.id == "" {
		return errors.New("id is required")
	}

	if subCategory.tenantId == "" {
		return errors.New("tenant id is required")
	}

	if subCategory.code == "" {
		return errors.New("code is required")
	}

	if subCategory.name == "" {
		return errors.New("name is required")
	}

	if subCategory.category == nil {
		return errors.New("category is required")
	}

	return nil
}
