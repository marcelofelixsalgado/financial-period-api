package list_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/group/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/group/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/group/list"
	useCaseStatus "marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListGroupUseCase_Execute(t *testing.T) {
	m := &mocks.GroupRepositoryMock{}

	tenantId := "11"

	groupType := entity.GroupType{
		Code: "EXP",
	}

	group1, _ := entity.NewGroup("1", "11", "Group1", "Group 1", groupType, time.Time{}, time.Time{})
	group2, _ := entity.NewGroup("1", "11", "Group2", "Group 2", groupType, time.Time{}, time.Time{})

	groups := []entity.IGroup{group1, group2}

	m.On("List", []filter.FilterParameter{}, tenantId).Return(groups, nil)

	useCase := list.NewListUseCase(m)

	input := list.InputListGroupDto{
		TenantId: tenantId,
	}

	output, internalStatus, err := useCase.Execute(input, []filter.FilterParameter{})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Len(t, output.Groups, 2)

	assert.NotEmpty(t, output.Groups[0].Id)
	assert.Equal(t, groups[0].GetName(), output.Groups[0].Name)
	assert.Equal(t, groups[0].GetCode(), output.Groups[0].Code)
	assert.Equal(t, groups[0].GetGroupType().Code, output.Groups[0].Type.Code)

	assert.Equal(t, groups[1].GetName(), output.Groups[1].Name)
	assert.Equal(t, groups[1].GetCode(), output.Groups[1].Code)
	assert.Equal(t, groups[1].GetGroupType().Code, output.Groups[1].Type.Code)

	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "List", 1)
}
