package update_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/group/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/group/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/group/update"
	"testing"
	"time"

	useCaseStatus "marcelofelixsalgado/financial-period-api/pkg/usecase/status"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateGroupUseCase_Execute(t *testing.T) {
	m := &mocks.GroupRepositoryMock{}

	groupType := entity.GroupType{
		Code: "EXP",
	}

	group, _ := entity.NewGroup("1", "11", "Group 1", "Group 1", groupType, time.Time{}, time.Time{})

	m.On("FindById", group.GetId()).Return(group, nil)
	m.On("Update", mock.Anything).Return(nil)

	useCase := update.NewUpdateUseCase(m)

	input := update.InputUpdateGroupDto{
		Id:   group.GetId(),
		Code: group.GetCode(),
		Name: group.GetName(),
		Type: update.GroupType{
			Code: groupType.Code,
		},
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
	m.AssertNumberOfCalls(t, "Update", 1)
}
