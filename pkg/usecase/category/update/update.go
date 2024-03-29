package update

import (
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	transactionType "github.com/marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/category"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	repositoryStatus "github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"
)

type IUpdateUseCase interface {
	Execute(InputUpdateCategoryDto) (OutputUpdateCategoryDto, status.InternalStatus, error)
}

type UpdateUseCase struct {
	repository category.ICategoryRepository
}

func NewUpdateUseCase(repository category.ICategoryRepository) IUpdateUseCase {
	return &UpdateUseCase{
		repository: repository,
	}
}

func (updateUseCase UpdateUseCase) Execute(input InputUpdateCategoryDto) (OutputUpdateCategoryDto, status.InternalStatus, error) {

	// Find the entity before update
	currentEntity, err := updateUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputUpdateCategoryDto{}, status.InternalServerError, err
	}
	if currentEntity == nil {
		return OutputUpdateCategoryDto{}, status.InvalidResourceId, nil
	}

	transactionType, err := transactionType.NewTransactionType(input.TransactionType.Code, "")
	if err != nil {
		return OutputUpdateCategoryDto{}, status.InternalServerError, err
	}

	entity, err := entity.NewCategory(input.Id, currentEntity.GetTenantId(), input.Code, input.Name, transactionType, currentEntity.GetCreatedAt(), time.Now())
	if err != nil {
		return OutputUpdateCategoryDto{}, status.InternalServerError, err
	}

	// Persists in dabatase
	repositoryInternalStatus, err := updateUseCase.repository.Update(entity)
	if repositoryInternalStatus == repositoryStatus.EntityWithSameKeyAlreadyExists {
		return OutputUpdateCategoryDto{}, status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return OutputUpdateCategoryDto{}, status.InternalServerError, err
	}

	outputUpdateCategoryDto := OutputUpdateCategoryDto{
		Id:   entity.GetId(),
		Code: entity.GetCode(),
		Name: entity.GetName(),
		TransactionType: TransactionTypeOutput{
			Code: entity.GetTransactionType().GetCode(),
			Name: currentEntity.GetTransactionType().GetName(),
		},
		CreatedAt: entity.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: entity.GetUpdatedAt().Format(time.RFC3339),
	}

	return outputUpdateCategoryDto, status.Success, nil
}
