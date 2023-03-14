package create

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/group/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/group"
	repositoryStatus "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/status"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"time"
)

type ICreateUseCase interface {
	Execute(InputCreateGroupDto) (OutputCreateGroupDto, status.InternalStatus, error)
}

type CreateUseCase struct {
	repository group.IGroupRepository
}

func NewCreateUseCase(repository group.IGroupRepository) ICreateUseCase {
	return &CreateUseCase{
		repository: repository,
	}
}

func (createUseCase CreateUseCase) Execute(input InputCreateGroupDto) (OutputCreateGroupDto, status.InternalStatus, error) {

	// Creates an entity
	groupType := entity.GroupType{
		Code: input.Type.Code,
	}
	entity, err := entity.Create(input.TenantId, input.Code, input.Name, groupType)
	if err != nil {
		return OutputCreateGroupDto{}, status.InternalServerError, err
	}

	// Persists in database
	repositoryInternalStatus, err := createUseCase.repository.Create(entity)
	if repositoryInternalStatus == repositoryStatus.EntityWithSameKeyAlreadyExists {
		return OutputCreateGroupDto{}, status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return OutputCreateGroupDto{}, status.InternalServerError, err
	}

	groupTypeDto := GroupType{
		Code: entity.GetGroupType().Code,
	}

	outputCreateGroupDto := OutputCreateGroupDto{
		Id:        entity.GetId(),
		Code:      entity.GetCode(),
		Name:      entity.GetName(),
		Type:      groupTypeDto,
		CreatedAt: entity.GetCreatedAt().Format(time.RFC3339),
	}

	return outputCreateGroupDto, status.Success, nil
}
