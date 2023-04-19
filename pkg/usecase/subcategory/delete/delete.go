package delete

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/subcategory"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

type IDeleteUseCase interface {
	Execute(InputDeleteSubCategoryDto) (OutputDeleteSubCategoryDto, status.InternalStatus, error)
}

type DeleteUseCase struct {
	repository subcategory.ISubCategoryRepository
}

func NewDeleteUseCase(repository subcategory.ISubCategoryRepository) IDeleteUseCase {
	return &DeleteUseCase{
		repository: repository,
	}
}

func (deleteUseCase *DeleteUseCase) Execute(input InputDeleteSubCategoryDto) (OutputDeleteSubCategoryDto, status.InternalStatus, error) {

	// Find the entity before update
	subCategory, err := deleteUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputDeleteSubCategoryDto{}, status.InternalServerError, err
	}
	if subCategory == nil {
		return OutputDeleteSubCategoryDto{}, status.InvalidResourceId, err
	}

	// Apply in dabatase
	err = deleteUseCase.repository.Delete(input.Id)
	if err != nil {
		return OutputDeleteSubCategoryDto{}, status.InternalServerError, err
	}

	outputDeleteUseCaseDto := OutputDeleteSubCategoryDto{}

	return outputDeleteUseCaseDto, status.Success, nil
}
