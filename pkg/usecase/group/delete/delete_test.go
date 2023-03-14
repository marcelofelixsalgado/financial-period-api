package delete_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/group/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/group/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/group/delete"
	useCaseStatus "marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteGroupUseCase(t *testing.T) {
	m := &mocks.GroupRepositoryMock{}

	groupType := entity.GroupType{
		Code: "EXP",
	}

	group, _ := entity.NewGroup("1", "11", "Group1", "Group 1", groupType, time.Time{}, time.Time{})

	m.On("FindById", group.GetId()).Return(group, nil)
	m.On("Delete", mock.Anything).Return(nil)

	useCase := delete.NewDeleteUseCase(m)

	input := delete.InputDeleteGroupDto{
		Id: group.GetId(),
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindById", 1)
	m.AssertNumberOfCalls(t, "Delete", 1)
}
