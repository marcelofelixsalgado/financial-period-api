package entity

import (
	"errors"
	"strings"
	"time"
)

type IGroup interface {
	GetId() string
	GetTenantId() string
	GetCode() string
	GetName() string
	GetGroupType() GroupType
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type Group struct {
	id        string
	tenantId  string
	code      string
	name      string
	groupType GroupType
	createAt  time.Time
	updatedAt time.Time
}

type GroupType struct {
	Code string
}

func (group Group) GetId() string {
	return group.id
}

func (group Group) GetTenantId() string {
	return group.tenantId
}

func (group Group) GetCode() string {
	return group.code
}

func (group Group) GetName() string {
	return group.name
}

func (group Group) GetGroupType() GroupType {
	return group.groupType
}

func (group Group) GetCreatedAt() time.Time {
	return group.createAt
}

func (group Group) GetUpdatedAt() time.Time {
	return group.updatedAt
}

func (group *Group) SetUpdatedAt(updatedAt time.Time) {
	group.updatedAt = updatedAt
	group.validate()
}

func NewGroup(id string, tenantId string, code string, name string, groupType GroupType, createAt time.Time, updatedAt time.Time) (IGroup, error) {
	group := Group{
		id:        id,
		tenantId:  tenantId,
		code:      code,
		name:      name,
		groupType: groupType,
		createAt:  createAt,
		updatedAt: updatedAt,
	}
	group.format()
	if err := group.validate(); err != nil {
		return nil, err
	}
	return group, nil
}

func (group *Group) validate() error {
	if group.id == "" {
		return errors.New("id is required")
	}

	if group.tenantId == "" {
		return errors.New("tenant id is required")
	}

	if group.code == "" {
		return errors.New("code is required")
	}

	if group.name == "" {
		return errors.New("name is required")
	}

	if group.groupType.Code != "EAR" && group.groupType.Code != "EXP" {
		return errors.New("invalid group type")
	}

	return nil
}

func (group *Group) format() {
	group.code = strings.TrimSpace(group.code)
	group.name = strings.TrimSpace(group.name)
	group.groupType.Code = strings.TrimSpace(group.groupType.Code)
}
