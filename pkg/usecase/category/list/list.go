package list

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/category"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

type IListUseCase interface {
	Execute(InputListCategoryDto, []filter.FilterParameter) (OutputListCategoryDto, status.InternalStatus, error)
}

type ListUseCase struct {
	repository category.ICategoryRepository
}

func NewListUseCase(repository category.ICategoryRepository) IListUseCase {
	return &ListUseCase{
		repository: repository,
	}
}

func (listUseCase ListUseCase) Execute(input InputListCategoryDto, filterParameters []filter.FilterParameter) (OutputListCategoryDto, status.InternalStatus, error) {

	categories, err := listUseCase.repository.List(filterParameters, input.TenantId)
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
		}
		outputListCategoryDto.Categories = append(outputListCategoryDto.Categories, category)
	}

	return outputListCategoryDto, status.Success, nil
}
