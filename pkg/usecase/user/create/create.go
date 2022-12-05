package create

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"

	"time"
)

type ICreateUseCase interface {
	Execute(InputCreateUserDto) (OutputCreateUserDto, status.InternalStatus, error)
}

type CreateUseCase struct {
	repository user.IRepository
}

func NewCreateUseCase(repository user.IRepository) ICreateUseCase {
	return &CreateUseCase{
		repository: repository,
	}
}

func (createUseCase *CreateUseCase) Execute(input InputCreateUserDto) (OutputCreateUserDto, status.InternalStatus, error) {

	// Creates an entity
	entity, err := entity.Create(input.Name, input.Password, input.Phone, input.Email)
	if err != nil {
		return OutputCreateUserDto{}, status.InternalServerError, err
	}

	// Persists in dabatase
	err = createUseCase.repository.Create(entity)
	if err != nil {
		return OutputCreateUserDto{}, status.InternalServerError, err
	}

	outputCreateUserDto := OutputCreateUserDto{
		Id:        entity.GetId(),
		Name:      entity.GetName(),
		Password:  entity.GetPassword(),
		Phone:     entity.GetPhone(),
		Email:     entity.GetEmail(),
		CreatedAt: entity.GetCreatedAt().Format(time.RFC3339),
	}

	return outputCreateUserDto, status.Success, nil
}
