package find

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user"
	"time"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

type IFindUseCase interface {
	Execute(InputFindUserDto) (OutputFindUserDto, status.InternalStatus, error)
}

type FindUseCase struct {
	repository user.IUserRepository
}

func NewFindUseCase(repository user.IUserRepository) IFindUseCase {
	return &FindUseCase{
		repository: repository,
	}
}

func (findUseCase *FindUseCase) Execute(input InputFindUserDto) (OutputFindUserDto, status.InternalStatus, error) {

	user, err := findUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputFindUserDto{}, status.InternalServerError, err
	}
	if user == nil {
		return OutputFindUserDto{}, status.InvalidResourceId, err
	}

	if user.GetId() == "" {
		return OutputFindUserDto{}, status.NoRecordsFound, err
	}

	outputFindUserDto := OutputFindUserDto{
		Id: user.GetId(),
		Tenant: tenantDto{
			Id: user.GetTenantId(),
		},
		Name:      user.GetName(),
		Phone:     user.GetPhone(),
		Email:     user.GetEmail(),
		CreatedAt: user.GetCreatedAt().Format(time.RFC3339),
	}

	if !user.GetUpdatedAt().IsZero() {
		outputFindUserDto.UpdatedAt = user.GetUpdatedAt().Format(time.RFC3339)
	}

	return outputFindUserDto, status.Success, nil
}
