package list

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/category"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"time"
)

type IListCategoryUseCase interface {
	Execute(InputListCategoryDto, []filter.FilterParameter) (OutputListCategoryDto, status.InternalStatus, error)
}

type ListCategoryUseCase struct {
	repository category.ICategoryRepository
}

func NewListUseCase(repository category.ICategoryRepository) IListCategoryUseCase {
	return &ListCategoryUseCase{
		repository: repository,
	}
}

func (listCategoryUseCase ListCategoryUseCase) Execute(input InputListCategoryDto, filterParameters []filter.FilterParameter) (OutputListCategoryDto, status.InternalStatus, error) {

	categories, err := listCategoryUseCase.repository.List(filterParameters, input.TenantId)
	if err != nil {
		return OutputListCategoryDto{}, status.InternalServerError, err
	}

	outputListCategoryDto := OutputListCategoryDto{}

	if len(categories) == 0 {
		return outputListCategoryDto, status.NoRecordsFound, nil
	}

	for _, item := range categories {
		category := Category{
			Id:   item.GetId(),
			Code: item.GetCode(),
			Name: item.GetName(),
			TransactionType: TransactionTypeOutput{
				Code: item.GetTransactionType().GetCode(),
				Name: item.GetTransactionType().GetName(),
			},
			CreatedAt: item.GetCreatedAt().Format(time.RFC3339),
			UpdatedAt: item.GetCreatedAt().Format(time.RFC3339),
		}
		outputListCategoryDto.Categories = append(outputListCategoryDto.Categories, category)
	}

	return outputListCategoryDto, status.Success, nil
}
