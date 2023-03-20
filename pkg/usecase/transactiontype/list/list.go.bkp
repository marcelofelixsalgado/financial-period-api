package list

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/group"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
)

type IListUseCase interface {
	Execute(InputListGroupDto, []filter.FilterParameter) (OutputListGroupDto, status.InternalStatus, error)
}

type ListUseCase struct {
	repository group.IGroupRepository
}

func NewListUseCase(repository group.IGroupRepository) IListUseCase {
	return &ListUseCase{
		repository: repository,
	}
}

func (listUseCase ListUseCase) Execute(input InputListGroupDto, filterParameters []filter.FilterParameter) (OutputListGroupDto, status.InternalStatus, error) {

	groups, err := listUseCase.repository.List(filterParameters, input.TenantId)
	if err != nil {
		return OutputListGroupDto{}, status.InternalServerError, err
	}

	outputListGroupDto := OutputListGroupDto{}

	if len(groups) == 0 {
		return OutputListGroupDto{}, status.NoRecordsFound, nil
	}

	for _, item := range groups {
		group := Group{
			Id:   item.GetId(),
			Code: item.GetCode(),
			Name: item.GetName(),
			Type: GroupType{
				Code: item.GetGroupType().Code,
			},
		}
		outputListGroupDto.Groups = append(outputListGroupDto.Groups, group)
	}
	return outputListGroupDto, status.Success, nil
}
