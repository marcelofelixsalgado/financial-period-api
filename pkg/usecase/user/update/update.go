package update

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"time"
)

type IUpdateUseCase interface {
	Execute(InputUpdateUserDto) (OutputUpdateUserDto, status.InternalStatus, error)
}

type UpdateUseCase struct {
	repository user.IRepository
}

func NewUpdateUseCase(repository user.IRepository) IUpdateUseCase {
	return &UpdateUseCase{
		repository: repository,
	}
}

func (updateUseCase *UpdateUseCase) Execute(input InputUpdateUserDto) (OutputUpdateUserDto, status.InternalStatus, error) {

	// Find the entity before update
	currentEntity, err := updateUseCase.repository.Find(input.Id)
	if err != nil {
		return OutputUpdateUserDto{}, status.InternalServerError, err
	}
	if currentEntity == nil {
		return OutputUpdateUserDto{}, status.InvalidResourceId, err
	}

	entity, err := entity.NewUser(input.Id, input.Name, input.Password, input.Phone, input.Email, currentEntity.GetCreatedAt(), time.Now())
	if err != nil {
		return OutputUpdateUserDto{}, status.InternalServerError, err
	}

	// Persists in dabatase
	err = updateUseCase.repository.Update(entity)
	if err != nil {
		return OutputUpdateUserDto{}, status.InternalServerError, err
	}

	outputUpdateUserDto := OutputUpdateUserDto{
		Id:        entity.GetId(),
		Name:      entity.GetName(),
		Password:  entity.GetPassword(),
		Phone:     entity.GetPhone(),
		Email:     entity.GetEmail(),
		CreatedAt: entity.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: entity.GetUpdatedAt().Format(time.RFC3339),
	}

	return outputUpdateUserDto, status.Success, nil
}
