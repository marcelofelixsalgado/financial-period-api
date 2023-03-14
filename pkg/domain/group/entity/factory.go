package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

func Create(tenantId string, code string, name string, groupType GroupType) (IGroup, error) {
	return NewGroup(uuid.NewV4().String(), tenantId, code, name, groupType, time.Now(), time.Time{})
}
