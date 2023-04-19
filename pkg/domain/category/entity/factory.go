package entity

import (
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"

	uuid "github.com/satori/go.uuid"
)

func Create(tenantId string, code string, name string, transactionType entity.ITransactionType) (ICategory, error) {
	return NewCategory(uuid.NewV4().String(), tenantId, code, name, transactionType, time.Now(), time.Time{})
}
