package update

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/credentials/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/credentials"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"time"
)

type IUpdateUseCase interface {
	Execute(InputUpdateUserCredentialsDto) (OutputUpdateUserCredentialsDto, status.InternalStatus, error)
}

type UpdateUseCase struct {
	repository credentials.IUserCredentialsRepository
}

func NewUpdateUseCase(repository credentials.IUserCredentialsRepository) IUpdateUseCase {
	return &UpdateUseCase{
		repository: repository,
	}
}

func (updateUseCase *UpdateUseCase) Execute(input InputUpdateUserCredentialsDto) (OutputUpdateUserCredentialsDto, status.InternalStatus, error) {

	// Find the entity before update
	currentEntity, err := updateUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputUpdateUserCredentialsDto{}, status.InternalServerError, err
	}
	if currentEntity == nil {
		return OutputUpdateUserCredentialsDto{}, status.InvalidResourceId, err
	}

	entity, err := entity.NewUserCredentials(input.Id, input.UserId, input.Password, currentEntity.GetCreatedAt(), time.Now())
	if err != nil {
		return OutputUpdateUserCredentialsDto{}, status.InternalServerError, err
	}

	// Persists in dabatase
	err = updateUseCase.repository.Update(entity)
	if err != nil {
		return OutputUpdateUserCredentialsDto{}, status.InternalServerError, err
	}

	outputUpdateUserDto := OutputUpdateUserCredentialsDto{
		Id:        entity.GetId(),
		UserId:    entity.GetUserId(),
		Password:  entity.GetPassword(),
		CreatedAt: entity.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: entity.GetUpdatedAt().Format(time.RFC3339),
	}

	return outputUpdateUserDto, status.Success, nil
}
