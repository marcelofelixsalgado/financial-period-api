package create

import (
	tenantEntity "marcelofelixsalgado/financial-period-api/pkg/domain/tenant/entity"
	userEntity "marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"
	repositoryStatus "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/status"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/tenant"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"

	"time"
)

type ICreateUseCase interface {
	Execute(InputCreateUserDto) (OutputCreateUserDto, status.InternalStatus, error)
}

type CreateUseCase struct {
	tenantRepository tenant.ITenantRepository
	userRepository   user.IUserRepository
}

func NewCreateUseCase(userRepository user.IUserRepository, tenantRepository tenant.ITenantRepository) ICreateUseCase {
	return &CreateUseCase{
		userRepository:   userRepository,
		tenantRepository: tenantRepository,
	}
}

func (createUseCase *CreateUseCase) Execute(input InputCreateUserDto) (OutputCreateUserDto, status.InternalStatus, error) {

	// Creates a tenant entity
	tenant, err := tenantEntity.Create()
	if err != nil {
		return OutputCreateUserDto{}, status.InternalServerError, err
	}

	// Persists the tenant
	repositoryInternalStatus, err := createUseCase.tenantRepository.Create(tenant)
	if err != nil || repositoryInternalStatus == repositoryStatus.InternalServerError {
		return OutputCreateUserDto{}, status.InternalServerError, err
	}

	// Creates a user entity
	user, err := userEntity.Create(tenant.GetId(), input.Name, input.Phone, input.Email)
	if err != nil {
		return OutputCreateUserDto{}, status.InternalServerError, err
	}

	// Persists the user
	repositoryInternalStatus, err = createUseCase.userRepository.Create(user)
	if repositoryInternalStatus == repositoryStatus.EntityWithSameKeyAlreadyExists {
		return OutputCreateUserDto{}, status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return OutputCreateUserDto{}, status.InternalServerError, err
	}

	outputCreateUserDto := OutputCreateUserDto{
		Id: user.GetId(),
		Tenant: tenantDto{
			Id: user.GetTenantId(),
		},
		Name:      user.GetName(),
		Phone:     user.GetPhone(),
		Email:     user.GetEmail(),
		CreatedAt: user.GetCreatedAt().Format(time.RFC3339),
	}

	return outputCreateUserDto, status.Success, nil
}
