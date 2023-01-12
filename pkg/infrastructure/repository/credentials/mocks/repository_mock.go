package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/credentials/entity"

	"github.com/stretchr/testify/mock"
)

type UserCredentialsMock struct {
	mock.Mock
}

func (m *UserCredentialsMock) Create(userCredencials entity.IUserCredentials) error {
	args := m.Called(userCredencials)
	return args.Error(0)
}

func (m *UserCredentialsMock) Update(userCredencials entity.IUserCredentials) error {
	args := m.Called(userCredencials)
	return args.Error(0)
}

func (m *UserCredentialsMock) FindById(id string) (entity.IUserCredentials, error) {
	args := m.Called(id)
	return args.Get(0).(entity.IUserCredentials), args.Error(1)
}

func (m *UserCredentialsMock) FindByUserId(userId string) (entity.IUserCredentials, error) {
	args := m.Called(userId)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(entity.IUserCredentials), args.Error(1)
}

func (m *UserCredentialsMock) FindByUserEmail(email string) (entity.IUserCredentials, error) {
	args := m.Called(email)
	return args.Get(0).(entity.IUserCredentials), args.Error(1)
}
