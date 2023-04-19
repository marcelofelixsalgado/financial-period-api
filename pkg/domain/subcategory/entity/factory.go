package entity

import (
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"

	uuid "github.com/satori/go.uuid"
)

func Create(tenantId string, code string, name string, category entity.ICategory) (ISubCategory, error) {
	return NewSubCategory(uuid.NewV4().String(), tenantId, code, name, category, time.Now(), time.Time{})
}
