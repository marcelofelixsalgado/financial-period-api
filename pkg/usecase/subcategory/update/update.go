package update

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/subcategory/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/category"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/subcategory"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"time"

	repositoryStatus "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/status"
)

type IUpdateUseCase interface {
	Execute(InputUpdateSubCategoryDto) (OutputUpdateSubCategoryDto, status.InternalStatus, error)
}

type UpdateUseCase struct {
	subCategoryRepository subcategory.ISubCategoryRepository
	categoryRepository    category.ICategoryRepository
}

func NewUpdateUseCase(subCategoryRepository subcategory.ISubCategoryRepository, categoryRepository category.ICategoryRepository) IUpdateUseCase {
	return &UpdateUseCase{
		subCategoryRepository: subCategoryRepository,
		categoryRepository:    categoryRepository,
	}
}

func (updateUseCase *UpdateUseCase) Execute(input InputUpdateSubCategoryDto) (OutputUpdateSubCategoryDto, status.InternalStatus, error) {

	// Checks if category exists
	category, err := updateUseCase.categoryRepository.FindById(input.Category.Id)
	if err != nil {
		return OutputUpdateSubCategoryDto{}, status.InternalServerError, err
	}
	if category == nil {
		return OutputUpdateSubCategoryDto{}, status.InvalidResourceId, err
	}
	if category.GetId() == "" {
		return OutputUpdateSubCategoryDto{}, status.NoRecordsFound, err
	}

	// Find the entity before update
	currentEntity, err := updateUseCase.subCategoryRepository.FindById(input.Id)
	if err != nil {
		return OutputUpdateSubCategoryDto{}, status.InternalServerError, err
	}
	if currentEntity == nil {
		return OutputUpdateSubCategoryDto{}, status.InvalidResourceId, nil
	}

	entity, err := entity.NewSubCategory(currentEntity.GetId(), currentEntity.GetTenantId(), input.Code, input.Name, category, currentEntity.GetCreatedAt(), time.Now())
	if err != nil {
		return OutputUpdateSubCategoryDto{}, status.InternalServerError, err
	}

	repositoryInternalStatus, err := updateUseCase.subCategoryRepository.Update(entity)
	if repositoryInternalStatus == repositoryStatus.EntityWithSameKeyAlreadyExists {
		return OutputUpdateSubCategoryDto{}, status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return OutputUpdateSubCategoryDto{}, status.InternalServerError, err
	}

	outputUpdateSubCategoryDto := OutputUpdateSubCategoryDto{
		Id:   entity.GetId(),
		Code: entity.GetCode(),
		Name: entity.GetName(),
		Category: CategoryOutput{
			Id:   entity.GetCategory().GetId(),
			Code: entity.GetCategory().GetCode(),
			Name: entity.GetCategory().GetName(),
			TransactionType: TransactionType{
				Code: entity.GetCategory().GetTransactionType().GetCode(),
				Name: entity.GetCategory().GetTransactionType().GetName(),
			},
		},
		CreatedAt: entity.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: entity.GetCreatedAt().Format(time.RFC3339),
	}

	return outputUpdateSubCategoryDto, status.Success, nil
}
