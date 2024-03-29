package delete_test

import (
	"testing"
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user/mocks"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/user/delete"

	useCaseStatus "github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteUserUseCase_Execute(t *testing.T) {
	m := &mocks.UserRepositoryMock{}

	user, _ := entity.NewUser("1234", "5678", "test", "1111-2222", "test@test.com", time.Time{}, time.Time{})

	m.On("FindById", user.GetId()).Return(user, nil)
	m.On("Delete", mock.Anything).Return(nil)

	useCase := delete.NewDeleteUseCase(m)

	input := delete.InputDeleteUserDto{
		Id: user.GetId(),
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindById", 1)
	m.AssertNumberOfCalls(t, "Delete", 1)
}
