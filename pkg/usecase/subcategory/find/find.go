package find

import (
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/subcategory"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

type IFindUseCase interface {
	Execute(InputFindSubCategoryDto) (OutputFindSubCategoryDto, status.InternalStatus, error)
}

type FindUseCase struct {
	repository subcategory.ISubCategoryRepository
}

func NewFindUseCase(repository subcategory.ISubCategoryRepository) IFindUseCase {
	return &FindUseCase{
		repository: repository,
	}
}

func (findUseCase FindUseCase) Execute(input InputFindSubCategoryDto) (OutputFindSubCategoryDto, status.InternalStatus, error) {

	subCategory, err := findUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputFindSubCategoryDto{}, status.InternalServerError, err
	}
	if subCategory == nil {
		return OutputFindSubCategoryDto{}, status.InvalidResourceId, err
	}
	if subCategory.GetId() == "" {
		return OutputFindSubCategoryDto{}, status.NoRecordsFound, err
	}

	outputFindSubCategoryDto := OutputFindSubCategoryDto{
		Id:   subCategory.GetId(),
		Code: subCategory.GetCode(),
		Name: subCategory.GetName(),
		Category: Category{
			Id:   subCategory.GetCategory().GetId(),
			Code: subCategory.GetCategory().GetCode(),
			Name: subCategory.GetCategory().GetName(),
			TransactionType: TransactionType{
				Code: subCategory.GetCategory().GetTransactionType().GetCode(),
				Name: subCategory.GetCategory().GetTransactionType().GetName(),
			},
		},
		CreatedAt: subCategory.GetCreatedAt().Format(time.RFC3339),
	}

	if !subCategory.GetUpdatedAt().IsZero() {
		outputFindSubCategoryDto.UpdatedAt = subCategory.GetUpdatedAt().Format(time.RFC3339)
	}

	return outputFindSubCategoryDto, status.Success, nil
}
