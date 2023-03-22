package list

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/subcategory"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
)

type IListUseCase interface {
	Execute(InputListSubCategoryDto, []filter.FilterParameter) (OutputListSubCategoryDto, status.InternalStatus, error)
}

type ListUseCase struct {
	repository subcategory.ISubCategoryRepository
}

func NewListUseCase(repository subcategory.ISubCategoryRepository) IListUseCase {
	return &ListUseCase{
		repository: repository,
	}
}

func (listUseCase *ListUseCase) Execute(input InputListSubCategoryDto, filterParameters []filter.FilterParameter) (OutputListSubCategoryDto, status.InternalStatus, error) {

	subCategories, err := listUseCase.repository.List(filterParameters, input.TenantId)
	if err != nil {
		return OutputListSubCategoryDto{}, status.InternalServerError, err
	}

	outputListSubCategoryDto := OutputListSubCategoryDto{}

	if len(subCategories) == 0 {
		return outputListSubCategoryDto, status.NoRecordsFound, nil
	}

	for _, item := range subCategories {
		subCategory := SubCategory{
			Id:   item.GetId(),
			Code: item.GetCode(),
			Name: item.GetName(),
			Category: Category{
				Id:   item.GetCategory().GetId(),
				Code: item.GetCategory().GetCode(),
				Name: item.GetCategory().GetName(),
				TransactionType: TransactionType{
					Code: item.GetCategory().GetTransactionType().GetCode(),
					Name: item.GetCategory().GetTransactionType().GetName(),
				},
			},
		}

		outputListSubCategoryDto.SubCategories = append(outputListSubCategoryDto.SubCategories, subCategory)
	}

	return outputListSubCategoryDto, status.Success, nil
}
