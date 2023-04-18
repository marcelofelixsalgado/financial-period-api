package create

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/subcategory/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/category"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/subcategory"
	"time"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	repositoryStatus "github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"
)

type ICreateUseCase interface {
	Execute(InputCreateSubCategoryDto) (OutputCreateSubCategoryDto, status.InternalStatus, error)
}

type CreateUseCase struct {
	subCategoryRepository subcategory.ISubCategoryRepository
	categoryRepository    category.ICategoryRepository
}

func NewCreateUseCase(subCategoryRepository subcategory.ISubCategoryRepository, categoryRepository category.ICategoryRepository) ICreateUseCase {
	return &CreateUseCase{
		subCategoryRepository: subCategoryRepository,
		categoryRepository:    categoryRepository,
	}
}

func (createUseCase *CreateUseCase) Execute(input InputCreateSubCategoryDto) (OutputCreateSubCategoryDto, status.InternalStatus, error) {

	// Checks if category exists
	category, err := createUseCase.categoryRepository.FindById(input.Category.Id)
	if err != nil {
		return OutputCreateSubCategoryDto{}, status.InternalServerError, err
	}
	if category == nil {
		return OutputCreateSubCategoryDto{}, status.InvalidResourceId, err
	}
	if category.GetId() == "" {
		return OutputCreateSubCategoryDto{}, status.NoRecordsFound, err
	}

	// Creates the entity
	subcategory, err := entity.Create(input.TenantId, input.Code, input.Name, category)
	if err != nil {
		return OutputCreateSubCategoryDto{}, status.InternalServerError, err
	}

	// Persists Sub-Category
	repositoryInternalStatus, err := createUseCase.subCategoryRepository.Create(subcategory)
	if repositoryInternalStatus == repositoryStatus.EntityWithSameKeyAlreadyExists {
		return OutputCreateSubCategoryDto{}, status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return OutputCreateSubCategoryDto{}, status.InternalServerError, err
	}

	outputCreateSubCategoryDto := OutputCreateSubCategoryDto{
		Id:   subcategory.GetId(),
		Code: subcategory.GetCode(),
		Name: subcategory.GetName(),
		Category: CategoryOutput{
			Id:   subcategory.GetCategory().GetId(),
			Code: subcategory.GetCategory().GetCode(),
			Name: subcategory.GetCategory().GetName(),
			TransactionType: TransactionType{
				Code: subcategory.GetCategory().GetTransactionType().GetCode(),
				Name: subcategory.GetCategory().GetTransactionType().GetName(),
			},
		},
		CreatedAt: subcategory.GetCreatedAt().Format(time.RFC3339),
	}

	return outputCreateSubCategoryDto, status.Success, nil
}
