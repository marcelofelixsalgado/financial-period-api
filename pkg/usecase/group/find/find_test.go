package find_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/group/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/group/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/group/find"
	useCaseStatus "marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFindGroupUseCase_Execute(t *testing.T) {
	m := &mocks.GroupRepositoryMock{}

	groupType := entity.GroupType{
		Code: "EXP",
	}

	group, _ := entity.NewGroup("1", "11", "Group1", "Group 1", groupType, time.Time{}, time.Time{})

	m.On("FindById", group.GetId()).Return(group, nil)

	useCase := find.NewFindUseCase(m)

	input := find.InputFindGroupDto{
		Id: group.GetId(),
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.Id)
	assert.NotEmpty(t, output.CreatedAt)
	assert.Equal(t, group.GetName(), output.Name)
	assert.Equal(t, group.GetCode(), output.Code)
	assert.Equal(t, group.GetGroupType().Code, output.Type.Code)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindById", 1)
}
