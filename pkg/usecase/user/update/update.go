package update

import (
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

type IUpdateUseCase interface {
	Execute(InputUpdateUserDto) (OutputUpdateUserDto, status.InternalStatus, error)
}

type UpdateUseCase struct {
	repository user.IUserRepository
}

func NewUpdateUseCase(repository user.IUserRepository) IUpdateUseCase {
	return &UpdateUseCase{
		repository: repository,
	}
}

func (updateUseCase *UpdateUseCase) Execute(input InputUpdateUserDto) (OutputUpdateUserDto, status.InternalStatus, error) {

	// Find the entity before update
	currentEntity, err := updateUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputUpdateUserDto{}, status.InternalServerError, err
	}
	if currentEntity == nil {
		return OutputUpdateUserDto{}, status.InvalidResourceId, err
	}

	entity, err := entity.NewUser(input.Id, input.TenantId, input.Name, input.Phone, input.Email, currentEntity.GetCreatedAt(), time.Now())
	if err != nil {
		return OutputUpdateUserDto{}, status.InternalServerError, err
	}

	// Persists in dabatase
	err = updateUseCase.repository.Update(entity)
	if err != nil {
		return OutputUpdateUserDto{}, status.InternalServerError, err
	}

	outputUpdateUserDto := OutputUpdateUserDto{
		Id: entity.GetId(),
		Tenant: tenantDto{
			Id: entity.GetTenantId(),
		},
		Name:      entity.GetName(),
		Phone:     entity.GetPhone(),
		Email:     entity.GetEmail(),
		CreatedAt: entity.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: entity.GetUpdatedAt().Format(time.RFC3339),
	}

	return outputUpdateUserDto, status.Success, nil
}
