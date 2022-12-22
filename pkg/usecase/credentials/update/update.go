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
	currentEntity, err := updateUseCase.repository.FindByUserId(input.UserId)
	if err != nil {
		return OutputUpdateUserCredentialsDto{}, status.InternalServerError, err
	}
	if currentEntity == nil {
		return OutputUpdateUserCredentialsDto{}, status.InvalidResourceId, err
	}

	// If current password persisted in database and the received current password don't match
	if err := entity.VerfifyPassword(currentEntity.GetPassword(), input.CurrentPassword); err != nil {
		return OutputUpdateUserCredentialsDto{}, status.PasswordsDontMatch, err
	}

	// Generates the hash from the new password
	hashedPassword, err := entity.Hash(input.NewPassword)
	if err != nil {
		return OutputUpdateUserCredentialsDto{}, status.InternalServerError, err
	}
	input.NewPassword = string(hashedPassword)

	entity, err := entity.NewUserCredentials(currentEntity.GetId(), input.UserId, input.NewPassword, currentEntity.GetCreatedAt(), time.Now())
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
