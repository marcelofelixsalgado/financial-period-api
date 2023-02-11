package list_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user/mocks"
	useCaseStatus "marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/user/list"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListUserUseCase_Execute(t *testing.T) {
	m := &mocks.UserRepositoryMock{}

	users1, _ := entity.NewUser("1", "11", "user1", "1111-1111", "user1@test.com", time.Time{}, time.Time{})
	users2, _ := entity.NewUser("2", "22", "user2", "2222-2222", "user2@test.com", time.Time{}, time.Time{})

	users := []entity.IUser{users1, users2}

	m.On("List", []filter.FilterParameter{}).Return(users, nil)

	useCase := list.NewListUseCase(m)

	input := list.InputListUserDto{}

	output, internalStatus, err := useCase.Execute(input, []filter.FilterParameter{})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Len(t, output.Users, 2)

	assert.NotEmpty(t, output.Users[0].Id)
	assert.Equal(t, users[0].GetName(), output.Users[0].Name)
	assert.Equal(t, users[0].GetPhone(), output.Users[0].Phone)
	assert.Equal(t, users[0].GetEmail(), output.Users[0].Email)

	assert.NotEmpty(t, output.Users[1].Id)
	assert.Equal(t, users[1].GetName(), output.Users[1].Name)
	assert.Equal(t, users[1].GetPhone(), output.Users[1].Phone)
	assert.Equal(t, users[1].GetEmail(), output.Users[1].Email)

	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "List", 1)
}
