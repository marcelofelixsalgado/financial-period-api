package create

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/category"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/transactiontype"
	"time"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	repositoryStatus "github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"
)

type ICreateUseCase interface {
	Execute(InputCreateCategoryDto) (OutputCreateCategoryDto, status.InternalStatus, error)
}

type CreateUseCase struct {
	categoryRepository        category.ICategoryRepository
	transactionTypeRepository transactiontype.ITransactionTypeRepository
}

func NewCreateUseCase(categoryRepository category.ICategoryRepository, transactionTypeRepository transactiontype.ITransactionTypeRepository) ICreateUseCase {
	return &CreateUseCase{
		categoryRepository:        categoryRepository,
		transactionTypeRepository: transactionTypeRepository,
	}
}

func (createUseCase *CreateUseCase) Execute(input InputCreateCategoryDto) (OutputCreateCategoryDto, status.InternalStatus, error) {

	transactionType, err := createUseCase.transactionTypeRepository.FindByCode(input.TransactionType.Code)
	if err != nil {
		return OutputCreateCategoryDto{}, status.InternalServerError, err
	}
	if transactionType == nil {
		return OutputCreateCategoryDto{}, status.InvalidResourceId, err
	}
	if transactionType.GetCode() == "" {
		return OutputCreateCategoryDto{}, status.NoRecordsFound, err
	}

	category, err := entity.Create(input.TenantId, input.Code, input.Name, transactionType)
	if err != nil {
		return OutputCreateCategoryDto{}, status.InternalServerError, err
	}

	repositoryInternalStatus, err := createUseCase.categoryRepository.Create(category)
	if repositoryInternalStatus == repositoryStatus.EntityWithSameKeyAlreadyExists {
		return OutputCreateCategoryDto{}, status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return OutputCreateCategoryDto{}, status.InternalServerError, err
	}

	outputCreateCategoryDto := OutputCreateCategoryDto{
		Id:   category.GetId(),
		Code: category.GetCode(),
		Name: category.GetName(),
		TransactionType: TransactionTypeOutput{
			Code: category.GetTransactionType().GetCode(),
			Name: category.GetTransactionType().GetName(),
		},
		CreatedAt: category.GetCreatedAt().Format(time.RFC3339),
	}

	return outputCreateCategoryDto, status.Success, nil
}
