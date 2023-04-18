package find_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/user/find"
	"testing"
	"time"

	useCaseStatus "github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/assert"
)

func TestFindUserUseCase_Execute(t *testing.T) {
	m := &mocks.UserRepositoryMock{}

	user, _ := entity.NewUser("1234", "5678", "test", "1111-2222", "test@test.com", time.Time{}, time.Time{})

	m.On("FindById", user.GetId()).Return(user, nil)

	useCase := find.NewFindUseCase(m)

	input := find.InputFindUserDto{
		Id: user.GetId(),
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.Id)
	assert.NotEmpty(t, output.Tenant.Id)
	assert.NotEmpty(t, output.CreatedAt)
	assert.Equal(t, user.GetName(), output.Name)
	assert.Equal(t, user.GetPhone(), output.Phone)
	assert.Equal(t, user.GetEmail(), output.Email)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindById", 1)
}
