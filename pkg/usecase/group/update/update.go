package update

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/group/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/group"
	repositoryStatus "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/status"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"time"
)

type IUpdateUseCase interface {
	Execute(InputUpdateGroupDto) (OutputUpdateGroupDto, status.InternalStatus, error)
}

type UpdateUseCase struct {
	repository group.IGroupRepository
}

func NewUpdateUseCase(repository group.IGroupRepository) IUpdateUseCase {
	return &UpdateUseCase{
		repository: repository,
	}
}

func (updateUseCase *UpdateUseCase) Execute(input InputUpdateGroupDto) (OutputUpdateGroupDto, status.InternalStatus, error) {

	// Find the entity before update
	currentEntity, err := updateUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputUpdateGroupDto{}, status.InternalServerError, err
	}
	if currentEntity == nil {
		return OutputUpdateGroupDto{}, status.InvalidResourceId, nil
	}

	groupType := entity.GroupType{
		Code: input.Type.Code,
	}

	entity, err := entity.NewGroup(input.Id, currentEntity.GetTenantId(), input.Code, input.Name, groupType, currentEntity.GetCreatedAt(), time.Now())
	if err != nil {
		return OutputUpdateGroupDto{}, status.InternalServerError, err
	}

	// Persists in dabatase
	repositoryInternalStatus, err := updateUseCase.repository.Update(entity)
	if repositoryInternalStatus == repositoryStatus.EntityWithSameKeyAlreadyExists {
		return OutputUpdateGroupDto{}, status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return OutputUpdateGroupDto{}, status.InternalServerError, err
	}

	outputUpdateGroupDto := OutputUpdateGroupDto{
		Id:   entity.GetId(),
		Code: entity.GetCode(),
		Name: entity.GetName(),
		Type: GroupType{
			Code: entity.GetGroupType().Code,
		},
		CreatedAt: entity.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: entity.GetUpdatedAt().Format(time.RFC3339),
	}

	return outputUpdateGroupDto, status.Success, nil
}
