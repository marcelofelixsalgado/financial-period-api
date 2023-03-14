package find

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/group"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"time"
)

type IFindUseCase interface {
	Execute(InputFindGroupDto) (OutputFindGroupDto, status.InternalStatus, error)
}

type FindUseCase struct {
	repository group.IGroupRepository
}

func NewFindUseCase(repository group.IGroupRepository) IFindUseCase {
	return &FindUseCase{
		repository: repository,
	}
}

func (findUseCase *FindUseCase) Execute(input InputFindGroupDto) (OutputFindGroupDto, status.InternalStatus, error) {

	group, err := findUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputFindGroupDto{}, status.InternalServerError, err
	}
	if group == nil {
		return OutputFindGroupDto{}, status.InvalidResourceId, err
	}

	if group.GetId() == "" {
		return OutputFindGroupDto{}, status.NoRecordsFound, err
	}

	outputFindGroupDto := OutputFindGroupDto{
		Id:   group.GetId(),
		Code: group.GetCode(),
		Name: group.GetName(),
		Type: GroupType{
			Code: group.GetGroupType().Code,
		},
		CreatedAt: group.GetCreatedAt().Format(time.RFC3339),
	}

	if !group.GetUpdatedAt().IsZero() {
		outputFindGroupDto.UpdatedAt = group.GetUpdatedAt().Format(time.RFC3339)
	}

	return outputFindGroupDto, status.Success, nil
}
