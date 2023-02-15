package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

func Create(tenantId string, periodId string, categoryId string, actualAmount float32, limitAmount float32) (IBalance, error) {
	return NewBalance(uuid.NewV4().String(), tenantId, periodId, categoryId, actualAmount, limitAmount, time.Now(), time.Time{})
}
