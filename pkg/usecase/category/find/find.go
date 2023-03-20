package find

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/category"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"time"
)

type IFindUseCase interface {
	Execute(InputFindCategoryDto) (OutputFindCategoryDto, status.InternalStatus, error)
}

type FindUseCase struct {
	repository category.ICategoryRepository
}

func NewFindUseCase(repository category.ICategoryRepository) IFindUseCase {
	return &FindUseCase{
		repository: repository,
	}
}

func (findUseCase *FindUseCase) Execute(input InputFindCategoryDto) (OutputFindCategoryDto, status.InternalStatus, error) {

	category, err := findUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputFindCategoryDto{}, status.InternalServerError, err
	}
	if category == nil {
		return OutputFindCategoryDto{}, status.InvalidResourceId, err
	}
	if category.GetId() == "" {
		return OutputFindCategoryDto{}, status.NoRecordsFound, err
	}

	outputFindCategoryDto := OutputFindCategoryDto{
		Id:   category.GetId(),
		Code: category.GetCode(),
		Name: category.GetName(),
		TransactionType: TransactionTypeOutput{
			Code: category.GetTransactionType().GetCode(),
			Name: category.GetTransactionType().GetName(),
		},
		CreatedAt: category.GetCreatedAt().Format(time.RFC3339),
	}

	if !category.GetUpdatedAt().IsZero() {
		outputFindCategoryDto.UpdatedAt = category.GetUpdatedAt().Format(time.RFC3339)
	}

	return outputFindCategoryDto, status.Success, nil
}
